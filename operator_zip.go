package rex

func ZipFromIterable[A, B, C any](iterableB Iterable[B], f Func2[A, B, C], opts ...applyOption) PipeLine[A, C] {
	return _zip[A, B, C](iterableB, f, opts...)
}

func _zip[A, B, C any](iterableB Iterable[B], f Func2[A, B, C], opts ...applyOption) PipeLine[A, C] {
	return func(iterable Iterable[A]) Reader[C] {
		return func(ctx Context) Iterable[C] {

			op := newOptions(opts...)

			ch := make(chan Item[C])

			go func() {
				defer close(ch)

				sourceA := iterable()
				sourceB := iterableB()

				for {
					itemA, okA := <-sourceA
					itemB, okB := <-sourceB

					if !okA || !okB {
						return
					}

					a, err := itemA()
					if err != nil {
						if !SendItem(ctx, ch, ItemError[C](err)) {
							ch <- ItemError[C](ctx.Err())
							return
						}

						if op.OnErrorStrategy == ContinueOnError {
							continue
						}
						return
					}

					b, err := itemB()
					if err != nil {
						if !SendItem(ctx, ch, ItemError[C](err)) {
							ch <- ItemError[C](ctx.Err())
							return
						}

						if op.OnErrorStrategy == ContinueOnError {
							continue
						}
						return
					}

					c, err := f(ctx, a, b)
					if err != nil {
						if !SendItem(ctx, ch, ItemError[C](err)) {
							ch <- ItemError[C](ctx.Err())
							return
						}

						if op.OnErrorStrategy == ContinueOnError {
							continue
						}
						return
					}

					if !SendItem(ctx, ch, ItemOf[C](c)) {
						ch <- ItemError[C](ctx.Err())
						return
					}
				}
			}()

			return FromChanItem[C](ch)
		}
	}
}
