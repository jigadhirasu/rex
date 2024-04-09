package rex

// Default 來源如果為空，則傳回預設值
func Default[A any](a A) PipeLine[A, A] {
	return func(iterable Iterable[A]) Reader[A] {
		return func(ctx Context) Iterable[A] {
			return func() <-chan Item[A] {
				ch := make(chan Item[A])

				go func() {
					defer close(ch)
					defer Catcher[A](ch)

					source := iterable()

					count := 0
					for {
						item, ok := <-source
						if !ok {
							break
						}
						if !SendItem(ctx, ch, item) {
							ch <- ItemError[A](ctx.Err())
							return
						}
						count++
					}

					if count < 1 {
						ch <- ItemOf[A](a)
					}
				}()

				return ch
			}
		}
	}
}
