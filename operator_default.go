package rex

// 來源如果為空，則傳回預設值
func Default[A any](a A) func(iterable Iterable[A]) Reader[A] {
	return func(iterable Iterable[A]) Reader[A] {
		return func(ctx Context) Iterable[A] {
			ch := make(chan Item[A])

			go func() {
				defer close(ch)

				source := iterable()

				count := 0
				for i := 0; i < 4096; i++ {
					item, ok := <-source
					if !ok {
						break
					}
					if !sendItem(ctx, ch, item) {
						ch <- ItemError[A](ctx.Err())
						return
					}
					count++
				}

				if count < 1 {
					ch <- ItemOf[A](a)
				}
			}()

			return FromChanItem[A](ch)
		}
	}
}
