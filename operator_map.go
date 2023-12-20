package rex

import "sync"

// 有Side Effect的Step
func Map[A, B any](f Func1[A, B]) func(opts ...applyOption) PipeLine[A, B] {
	return func(opts ...applyOption) PipeLine[A, B] {

		op := newOptions(opts...)

		return func(iterable Iterable[A]) Reader[B] {
			return func(ctx Context) Iterable[B] {

				ch := make(chan Item[B], op.bufferSize)

				go func() {
					defer close(ch)

					source := iterable()

					var wg *sync.WaitGroup

					task := func(wg *sync.WaitGroup) {
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
								ch <- ItemError[B](err)
								break
							}

							b, err := f(ctx, a)

							if err != nil {
								if !sendItem(ctx, ch, ItemError[B](err)) {
									ch <- ItemError[B](ctx.Err())
									return
								}

								if op.OnErrorStrategy == ContinueOnError {
									continue
								}
								return
							}

							if !sendItem(ctx, ch, ItemOf(b)) {
								ch <- ItemError[B](ctx.Err())
								return
							}
						}
					}

					if op.poolSize > 1 {
						wg = new(sync.WaitGroup)
						wg.Add(int(op.poolSize))
						for i := uint32(0); i < op.poolSize; i++ {
							go task(wg)
						}
						wg.Wait()
					} else {
						task(nil)
					}
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
