package rex

func Reduce[A, B any](f Transfer2[B, A, B]) func(iterable Iterable[A]) Reader[B] {
	return func(iterable Iterable[A]) Reader[B] {
		return func(ctx Context) Iterable[B] {
			ch := make(chan Item[B])

			go func() {
				defer close(ch)
				defer Catcher[B](ch)

				source := iterable()

				var initial B
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

				if !SendItem(ctx, ch, ItemOf(initial)) {
					ch <- ItemError[B](ctx.Err())
					return
				}
			}()

			return FromChanItem[B](ch)
		}
	}
}
