package rex

// 過濾掉不符合條件的元素
func Filter[A any](f Predicate[A]) func(iterable Iterable[A]) Reader[A] {
	return func(iterable Iterable[A]) Reader[A] {
		return func(ctx Context) Iterable[A] {
			ch := make(chan Item[A])

			go func() {
				defer close(ch)
				defer Catcher[A](ch)

				source := iterable()
				for i := 0; i < 4096; i++ {
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

					if !sendItem(ctx, ch, item) {
						ch <- ItemError[A](ctx.Err())
						return
					}
				}
			}()

			return FromChanItem[A](ch)
		}
	}
}
