package rex

func BufferCount[A any](count int, opts ...applyOption) PipeLine[A, []A] {
	return func(iterable Iterable[A]) Reader[[]A] {
		return func(ctx Context) Iterable[[]A] {

			op := newOptions(opts...)

			ch := make(chan Item[[]A], op.bufferSize)

			go func() {
				defer close(ch)
				defer Catcher[[]A](ch)

				source := iterable()

				buf := []A{}

				for {
					item, ok := <-source
					if !ok {
						if len(buf) > 0 {
							if !sendItem(ctx, ch, ItemOf(buf)) {
								ch <- ItemError[[]A](ctx.Err())
								return
							}
						}
						return
					}

					a, err := item()
					if err != nil {
						if !sendItem(ctx, ch, ItemError[[]A](err)) {
							ch <- ItemError[[]A](ctx.Err())
							return
						}
					}

					buf = append(buf, a)

					if len(buf) == count {
						if !sendItem(ctx, ch, ItemOf(buf)) {
							ch <- ItemError[[]A](ctx.Err())
							return
						}

						buf = []A{}
					}
				}

			}()

			return FromChanItem[[]A](ch)
		}
	}

}
