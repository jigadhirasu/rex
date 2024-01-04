package rex

// 取得前 n 個元素
func Take[A any](n int) PipeLine[A, A] {
	return _take[A](n)
}

func _take[A any](n int) PipeLine[A, A] {
	return func(iterable Iterable[A]) Reader[A] {
		return func(ctx Context) Iterable[A] {
			ch := make(chan Item[A])

			go func() {
				defer close(ch)
				defer Catcher[A](ch)

				source := iterable()

				for i := 0; i < n; i++ {
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

			return func() <-chan Item[A] {
				return ch
			}
		}
	}
}
