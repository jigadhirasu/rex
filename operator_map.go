package rex

import "sync"

// 有Side Effect的Step
func Map[A, B any](f Func1[A, B], opts ...applyOption) PipeLine[A, B] {

	return func(iterable Iterable[A]) Reader[B] {

		return func(ctx Context) Iterable[B] {

			op := newOptions(opts...)

			ch := make(chan Item[B], op.bufferSize)

			go func() {
				defer close(ch)

				source := iterable()

				var wg *sync.WaitGroup

				if op.poolSize == 1 {
					mapDo[A, B](ctx, source, ch, op, wg, f)
					return
				}

				wg = new(sync.WaitGroup)
				wg.Add(int(op.poolSize))
				for i := uint32(0); i < op.poolSize; i++ {
					go mapDo[A, B](ctx, source, ch, op, wg, f)
				}
				wg.Wait()
			}()

			return FromChanItem[B](ch)
		}
	}
}

func Map1[A any](f Func1[A, A], opts ...applyOption) PipeLine[A, A] {
	return func(iterable Iterable[A]) Reader[A] {
		return Map[A, A](f, opts...)(iterable)
	}
}

func mapDo[A, B any](ctx Context, source <-chan Item[A], expose chan<- Item[B], opt options, wg *sync.WaitGroup, f Func1[A, B]) {
	defer Catcher[B](expose)
	if wg != nil {
		defer wg.Done()
	}

	for {
		select {
		case <-ctx.Done():
			return
		case item, ok := <-source:
			if !ok {
				return
			}

			a, err := item()
			if err != nil {
				if !sendItem(ctx, expose, ItemError[B](err)) {
					expose <- ItemError[B](ctx.Err())
					return
				}
			}

			b, err := f(ctx, a)

			if err != nil {
				if !sendItem(ctx, expose, ItemError[B](err)) {
					expose <- ItemError[B](ctx.Err())
					return
				}

				if opt.OnErrorStrategy == ContinueOnError {
					continue
				}
				return
			}

			if !sendItem(ctx, expose, ItemOf(b)) {
				expose <- ItemError[B](ctx.Err())
				return
			}
		}
	}
}
