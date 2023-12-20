package rex

import "sync"

// 有Side Effect的Step
func Map[A, B any](f Func1[A, B]) func(opts ...applyOption) PipeLine[A, B] {
	return func(opts ...applyOption) PipeLine[A, B] {

		op := newOptions(opts...)

		return func(iterable Iterable[A]) Reader[B] {
			return func(ctx Context) Iterable[B] {

				ch := make(chan Item[B])

				go func() {
					defer close(ch)

					source := iterable()

					wg := new(sync.WaitGroup)
					wg.Add(op.poolSize)

					task := func() {
						defer Catcher[B](ch)
						defer wg.Done()

						for {
							item, ok := <-source
							if !ok {
								return
							}

							a, err := item()
							if err != nil {
								ch <- ItemError[B](err)
								break
							}

							if !sendItem(ctx, ch, ItemAError(f(ctx, a))) {
								ch <- ItemError[B](ctx.Err())
								return
							}
						}
					}

					doWithPool(ctx, op.poolSize, task)

					wg.Wait()
				}()

				return FromChanItem[B](ch)
			}
		}
	}
}

func Map1[A any](f Func1[A, A]) func(opts ...applyOption) PipeLine[A, A] {
	return func(opts ...applyOption) PipeLine[A, A] {
		return func(iterable Iterable[A]) Reader[A] {
			return Map[A, A](f)(opts...)(iterable)
		}
	}
}
