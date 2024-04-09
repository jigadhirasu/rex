package rex

// Once 用來執行一次指定的函數
func Once[A any](f Func0[A]) PipeLine[A, A] {
	return _once(f)
}

func _once[A any](f Func0[A], opts ...applyOption) PipeLine[A, A] {
	return func(iterable Iterable[A]) Reader[A] {
		return func(ctx Context) Iterable[A] {
			return func() <-chan Item[A] {
				ch := make(chan Item[A])

				go func() {
					defer close(ch)
					defer Catcher[A](ch)

					source := iterable()

					once := true

					for {
						item, ok := <-source
						if !ok {
							return
						}

						if once {
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

							once = false
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
