package rex

import "sync"

// 有Side Effect的Step
func Map[A, B any](f Func1[A, B], opts ...applyOption) PipeLine[A, B] {
	return _map(f, opts...)
}

func _map[A, B any](f Func1[A, B], opts ...applyOption) PipeLine[A, B] {
	return func(iterable Iterable[A]) Reader[B] {
		return func(ctx Context) Iterable[B] {

			op := newOptions(opts...)

			ch := make(chan Item[B], op.bufferSize)

			go func() {
				defer close(ch)

				source := iterable()

				var wg *sync.WaitGroup

				mf := func() {
					defer Catcher[B](ch)
					if wg != nil {
						defer wg.Done()
					}

					for {

						item, ok := <-source
						if !ok {
							return
						}

						a, err := item()
						if err != nil {
							if !SendItem(ctx, ch, ItemError[B](err)) {
								ch <- ItemError[B](ctx.Err())
								return
							}

							if op.OnErrorStrategy == ContinueOnError {
								continue
							}
							return
						}

						b, err := f(ctx, a)

						if err != nil {
							if !SendItem(ctx, ch, ItemError[B](err)) {
								ch <- ItemError[B](ctx.Err())
								return
							}

							if op.OnErrorStrategy == ContinueOnError {
								continue
							}
							return
						}

						if !SendItem(ctx, ch, ItemOf(b)) {
							ch <- ItemError[B](ctx.Err())
							return
						}

					}
				}

				if op.poolSize == 1 {
					mf()
					return
				}

				wg = new(sync.WaitGroup)
				wg.Add(int(op.poolSize))
				for i := uint32(0); i < op.poolSize; i++ {
					go mf()
				}
				wg.Wait()

			}()

			return func() <-chan Item[B] {
				return ch
			}
		}
	}
}
