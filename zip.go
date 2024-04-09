package rex

// Zip 用來將兩個 Iterable 合併成一個，並透過指定的函數對應元素
func Zip[A, B, C any](f Func2[A, B, C], opts ...applyOption) func(iterableA Iterable[A], iterableB Iterable[B]) Reader[C] {
	return _zip(f, opts...)
}

func _zip[A, B, C any](f Func2[A, B, C], opts ...applyOption) func(iterableA Iterable[A], iterableB Iterable[B]) Reader[C] {
	return func(iterableA Iterable[A], iterableB Iterable[B]) Reader[C] {
		return func(ctx Context) Iterable[C] {

			op := newOptions(opts...)

			ch := make(chan Item[C])

			go func() {
				defer close(ch)

				sourceA := iterableA()
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
