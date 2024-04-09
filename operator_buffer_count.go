package rex

// BufferCount 用數量來將 Iterable 分組後在輸出, EX count = 2, [A,B,C,D,E] => [[A,B],[C,D],[E]]
func BufferCount[A any](count int, opts ...applyOption) PipeLine[A, []A] {
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
						item, ok := <-source
						if !ok {
							if len(buf) > 0 {
								if !SendItem(ctx, ch, ItemOf(buf)) {
									ch <- ItemError[[]A](ctx.Err())
									return
								}
							}
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

						if len(buf) == count {
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
