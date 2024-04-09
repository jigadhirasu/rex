package rex

// Skip 取得前 n 個元素
func Skip[A any](n int) PipeLine[A, A] {
	return _skip[A](n)
}

func _skip[A any](n int) PipeLine[A, A] {
	return func(iterable Iterable[A]) Reader[A] {
		return func(ctx Context) Iterable[A] {
			return func() <-chan Item[A] {
				ch := make(chan Item[A])

				go func() {
					defer close(ch)
					defer Catcher[A](ch)

					source := iterable()

					for i := 0; i < n; i++ {
						_, ok := <-source
						if !ok {
							return
						}
						continue
					}

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
				}()

				return ch
			}
		}
	}
}
