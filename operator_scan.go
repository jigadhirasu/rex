package rex

func Scan[A, B any](initial B, f Transfer2[B, A, B]) func(iterable Iterable[A]) func(ctx Context) Iterable[B] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[B] {
		return func(ctx Context) Iterable[B] {
			ch := make(chan Item[B])

			go func() {
				defer close(ch)

				source := iterable()

				for {
					item, ok := <-source
					if !ok {
						break
					}

					a, err := item()
					if err != nil {
						ch <- ItemError[B](err)
						break
					}

					initial = f(initial, a)
				}

				if !sendItem(ctx, ch, ItemOf(initial)) {
					ch <- ItemError[B](ctx.Err())
					return
				}
			}()

			return FromChanItem[B](ch)
		}
	}
}
