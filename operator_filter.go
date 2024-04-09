package rex

// Filter 過濾掉不符合條件的元素
func Filter[A any](f Predicate[A]) func(iterable Iterable[A]) Reader[A] {
	return func(iterable Iterable[A]) Reader[A] {
		return func(ctx Context) Iterable[A] {
			return func() <-chan Item[A] {
				ch := make(chan Item[A])

				go func() {
					defer close(ch)
					defer Catcher[A](ch)

					source := iterable()
					for {
						item, ok := <-source
						if !ok {
							return
						}

						a, err := item()
						if err != nil {
							ch <- ItemError[A](err)
							return
						}

						if !f(a) {
							continue
						}

						if !SendItem(ctx, ch, item) {
							ch <- ItemError[A](ctx.Err())
							return
						}
					}
				}()

				return ch
			}
		}
	}
}
