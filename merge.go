package rex

import "sync"

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

func MergeALL[A any](opts ...applyOption) PipeLine[Iterable[A], A] {
	return func(iterable Iterable[Iterable[A]]) Reader[A] {
		return _merge(iterable)
	}
}

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

					wg.Add(1)

					source := iterable()

					for {
						item, ok := <-source
						if !ok {
							return
						}

						if !sendItem(ctx, ch, item) {
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

					go wf(iterable)
				}

				wg.Wait()

			}()

			return ch
		}

	}
}
