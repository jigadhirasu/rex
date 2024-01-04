package rex

func StartWith[A any](iterable Iterable[A]) PipeLine[A, A] {
	return func(source Iterable[A]) Reader[A] {
		return func(ctx Context) Iterable[A] {

			ch := make(chan Item[A])

			go func() {
				defer close(ch)
				defer Catcher[A](ch)

				for item := range iterable() {
					if !SendItem(ctx, ch, item) {
						ch <- ItemError[A](ctx.Err())
						return
					}
				}

				for item := range source() {
					if !SendItem(ctx, ch, item) {
						ch <- ItemError[A](ctx.Err())
						return
					}
				}
			}()

			return func() <-chan Item[A] {
				return ch
			}
		}
	}
}
