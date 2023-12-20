package rex

import "sync"

// 有Side Effect的Step
func Map[A, B any](f Func1[A, B]) func(iterable Iterable[A]) Reader[B] {
	return func(iterable Iterable[A]) Reader[B] {
		return func(ctx Context) Iterable[B] {
			ch := make(chan Item[B])

			go func() {
				defer close(ch)

				source := iterable()

				poolSize := 1
				wg := new(sync.WaitGroup)
				wg.Add(poolSize)

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

				WithPool(ctx, poolSize, task)

				wg.Wait()
			}()

			return FromChanItem[B](ch)
		}
	}
}

func Map1[A any](f Func1[A, A]) func(iterable Iterable[A]) Reader[A] {
	return func(iterable Iterable[A]) Reader[A] {
		return Map[A, A](f)(iterable)
	}
}
