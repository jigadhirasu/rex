package rex

// Tap 用來在每個元素上執行指定的函數，並將元素原封不動地傳遞到下一個操作
func Tap[A any](f Func0[A]) PipeLine[A, A] {
	return _tap(f)
}

func _tap[A any](f Func0[A], opts ...applyOption) PipeLine[A, A] {
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
							if !SendItem(ctx, ch, ItemError[A](err)) {
								ch <- ItemError[A](ctx.Err())
							}

							return
						}

						if err := f(ctx, a); err != nil {
							if !SendItem(ctx, ch, ItemError[A](err)) {
								ch <- ItemError[A](ctx.Err())
							}

							return
						}

						if !SendItem(ctx, ch, ItemOf(a)) {
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
