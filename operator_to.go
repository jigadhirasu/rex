package rex

func MapTo[A, B any](b B) PipeLine[A, B] {
	return _to[A, B](b)
}

func _to[A, B any](b B) PipeLine[A, B] {
	return func(iterable Iterable[A]) Reader[B] {
		return func(ctx Context) Iterable[B] {

			ch := make(chan Item[B])

			go func() {
				defer close(ch)
				defer Catcher[B](ch)

				source := iterable()

				for {
					item, ok := <-source
					if !ok {
						return
					}

					_, err := item()
					if err != nil {
						if !SendItem(ctx, ch, ItemError[B](err)) {
							ch <- ItemError[B](ctx.Err())
						}

						return
					}

					if !SendItem(ctx, ch, ItemOf(b)) {
						ch <- ItemError[B](ctx.Err())
						return
					}
				}
			}()

			return func() <-chan Item[B] {
				return ch
			}
		}
	}
}
