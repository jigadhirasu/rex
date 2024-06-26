package rex

import "time"

// BufferTime 用時間來將 Iterable 分組後在輸出, EX duration = 2ms, [A,B,C,D,E] => [[A],[B,C,D],[E]]
func BufferTime[A any](duration time.Duration, opts ...applyOption) PipeLine[A, []A] {
	return func(iterable Iterable[A]) Reader[[]A] {
		return func(ctx Context) Iterable[[]A] {
			return func() <-chan Item[[]A] {
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
								if !SendItem(ctx, ch, ItemError[[]A](err)) {
									ch <- ItemError[[]A](ctx.Err())
									return
								}
							}

							buf = append(buf, a)

						case <-time.After(duration):
							if len(buf) == 0 {
								continue
							}

							if !SendItem(ctx, ch, ItemOf(buf)) {
								ch <- ItemError[[]A](ctx.Err())
								return
							}

							buf = []A{}
						}
					}

				}()

				return ch
			}
		}
	}

}
