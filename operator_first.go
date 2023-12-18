package rex

// 取得第一個元素
func First[A any](iterable Iterable[A]) Reader[A] {
	return func(ctx Context) Iterable[A] {

		ch := make(chan Item[A])

		go func() {
			defer close(ch)
			defer Catcher[A](ch)

			source := iterable()

			item, ok := <-source
			if !ok {
				return
			}
			if !sendItem(ctx, ch, item) {
				ch <- ItemError[A](ctx.Err())
				return
			}
		}()

		return FromChanItem[A](ch)
	}
}
