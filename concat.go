package rex

func ConcatMap[A, B any](f HFunc1[A, B]) PipeLine[A, B] {
	return func(iterable Iterable[A]) Reader[B] {
		return func(ctx Context) Iterable[B] {
			hf := func(ctx Context, a A) (Iterable[B], error) {
				return f(ctx, a), nil
			}

			return Pipe2(
				_map(hf),
				ConcatALL[B](),
			)(iterable)(ctx)
		}
	}
}

func ConcatALL[A any](opts ...applyOption) PipeLine[Iterable[A], A] {
	return func(iterable Iterable[Iterable[A]]) Reader[A] {
		return _concat(iterable)
	}
}

func Concat[A any](iterables ...Iterable[A]) Reader[A] {
	return _concat(From(iterables...))
}

func _concat[A any](iterables Iterable[Iterable[A]]) Reader[A] {

	return func(ctx Context) Iterable[A] {

		return func() <-chan Item[A] {
			ch := make(chan Item[A])

			go func() {
				defer close(ch)

				wf := func(iterable Iterable[A]) {
					defer Catcher[A](ch)

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

					wf(iterable)
				}
			}()

			return ch
		}

	}
}
