package rex

import "time"

func BufferTime[A any](duration time.Duration, opts ...applyOption) PipeLine[A, []A] {
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

					select {
					case item, ok := <-source:
						if !ok {
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

					case <-time.After(duration):
						if len(buf) == 0 {
							continue
						}

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
