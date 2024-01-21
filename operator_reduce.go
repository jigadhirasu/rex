package rex

func Reduce[A, B any](f Transfer2[B, A, B]) PipeLine[A, B] {
	return _reduce(f)
}

func ReduceSlice[A any]() PipeLine[A, []A] {
	return _reduce[A, []A](func(b []A, a A) []A {
		if b == nil {
			b = []A{}
		}

		return append(b, a)
	})
}

func ReduceMap[A any, B comparable](f Transfer1[A, B]) PipeLine[A, map[B]A] {
	return _reduce[A, map[B]A](func(b map[B]A, a A) map[B]A {
		if b == nil {
			b = map[B]A{}
		}

		b[f(a)] = a

		return b
	})
}

func _reduce[A, B any](f Transfer2[B, A, B]) func(iterable Iterable[A]) Reader[B] {
	return func(iterable Iterable[A]) Reader[B] {
		return func(ctx Context) Iterable[B] {
			return func() <-chan Item[B] {
				ch := make(chan Item[B])

				go func() {
					defer close(ch)
					defer Catcher[B](ch)

					source := iterable()

					var initial B
					for {
						item, ok := <-source
						if !ok {
							break
						}

						a, err := item()
						if err != nil {
							ch <- ItemError[B](err)
							break
						}

						initial = f(initial, a)
					}

					if !SendItem(ctx, ch, ItemOf(initial)) {
						ch <- ItemError[B](ctx.Err())
						return
					}
				}()

				return ch
			}
		}
	}
}
