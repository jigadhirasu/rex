package rex

// 取得最後一個元素
func Last[A any](iterable Iterable[A]) func(ctx Context) Iterable[A] {
	return func(ctx Context) Iterable[A] {

		ch := make(chan Item[A])

		go func() {
			defer close(ch)

			source := iterable()

			var last Item[A]
			for i := 0; i < 4096; i++ {
				item, ok := <-source
				if !ok {
					break
				}

				last = item
			}

			if last != nil {
				sendItem[A](ctx, ch, last)
			}
		}()

		return FromChanItem[A](ch)
	}
}