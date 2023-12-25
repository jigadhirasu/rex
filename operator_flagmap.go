package rex

import (
	"sync"
)

// Iterable 會是一個新的 goroutine 不保證順序
func FlatMap[A, B any](f HFunc1[A, B]) func(iterable Iterable[A]) Reader[B] {
	return func(iterable Iterable[A]) Reader[B] {
		return func(ctx Context) Iterable[B] {

			ch := make(chan Item[B])

			go func() {
				source := iterable()

				wg := new(sync.WaitGroup)

				defer func() {
					wg.Wait()
					close(ch)
				}()

				for {
					item, ok := <-source
					if !ok {
						return
					}

					a, err := item()
					if err != nil {
						ch <- ItemError[B](err)
						return
					}

					wg.Add(1)

					go func() {
						next := f(ctx, a)()

						for {
							item, ok := <-next
							if !ok {
								wg.Done()
								return
							}

							if !sendItem(ctx, ch, item) {
								ch <- ItemError[B](ctx.Err())
								wg.Done()
								return
							}
						}
					}()
				}
			}()

			return FromChanItem[B](ch)
		}
	}
}

func FlatMap1[A any](f HFunc1[A, A]) func(iterable Iterable[A]) Reader[A] {
	return func(iterable Iterable[A]) Reader[A] {
		return FlatMap[A, A](f)(iterable)
	}
}
