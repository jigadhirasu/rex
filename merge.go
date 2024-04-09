package rex

import (
	"sync"
)

// MergeMap 用來將一個 Iterable[A] 轉換成 Iterable[B]，不保證順序
func MergeMap[A, B any](f HFunc1[A, B]) PipeLine[A, B] {
	return func(iterable Iterable[A]) Reader[B] {
		return func(ctx Context) Iterable[B] {
			hf := func(ctx Context, a A) (Iterable[B], error) {
				return f(ctx, a), nil
			}

			return Pipe2(
				_map(hf),
				MergeALL[B](),
			)(iterable)(ctx)
		}
	}
}

// MergeALL 會將所有的 Iterable 串接起來，不保證順序
func MergeALL[A any](opts ...applyOption) PipeLine[Iterable[A], A] {
	return func(iterable Iterable[Iterable[A]]) Reader[A] {
		return _merge(iterable)
	}
}

// Merge 會將所有的 Iterable 串接起來，不保證順序
func Merge[A any](iterables ...Iterable[A]) Reader[A] {
	return _merge(From(iterables...))
}

func _merge[A any](iterables Iterable[Iterable[A]]) Reader[A] {
	return func(ctx Context) Iterable[A] {

		return func() <-chan Item[A] {
			ch := make(chan Item[A])

			go func() {
				defer close(ch)

				wg := new(sync.WaitGroup)

				wf := func(iterable Iterable[A]) {
					defer Catcher[A](ch)
					defer wg.Done()

					source := iterable()

					for {
						item, ok := <-source
						if !ok {
							return
						}

						if !SendItem(ctx, ch, item) {
							ch <- ItemError[A](ctx.Err())
							return
						}
					}
				}

				source := iterables()

				for {
					item, ok := <-source
					if !ok {
						break
					}

					iterable, err := item()
					if err != nil {
						ch <- ItemError[A](err)
						return
					}

					wg.Add(1)

					go wf(iterable)
				}

				wg.Wait()

			}()

			return ch
		}

	}
}
