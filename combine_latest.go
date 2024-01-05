package rex

func CombineLatest[A, B, C any](f Func2[A, B, C], opts ...applyOption) func(iterableA Iterable[A], iterableB Iterable[B]) Reader[C] {
	return _combineLatest[A, B, C](f, opts...)
}

func _combineLatest[A, B, C any](f Func2[A, B, C], opts ...applyOption) func(Iterable[A], Iterable[B]) Reader[C] {
	return func(iterableA Iterable[A], iterableB Iterable[B]) Reader[C] {
		return func(ctx Context) Iterable[C] {
			return func() <-chan Item[C] {

				ch := make(chan Item[C])

				go func() {
					defer close(ch)

					sourceA := iterableA()
					sourceB := iterableB()

					// ------所有來源都至少有一個才會開始------
					next := func(ctx Context, itemA Item[A], itemB Item[B]) {
						a, err := itemA()
						if err != nil {
							if !SendItem(ctx, ch, ItemError[C](err)) {
								ch <- ItemError[C](ctx.Err())
							}

							return
						}

						b, err := itemB()
						if err != nil {
							if !SendItem(ctx, ch, ItemError[C](err)) {
								ch <- ItemError[C](ctx.Err())
							}

							return
						}

						c, err := f(ctx, a, b)

						if err != nil {
							if !SendItem(ctx, ch, ItemError[C](err)) {
								ch <- ItemError[C](ctx.Err())
							}

							return
						}
						if !SendItem(ctx, ch, ItemOf[C](c)) {
							ch <- ItemError[C](ctx.Err())
							return
						}
					}

					var lastA Item[A]
					var lastB Item[B]

					itemA, okA := <-sourceA
					if okA {
						lastA = itemA
					}

					itemB, okB := <-sourceB
					if okB {
						lastB = itemB
					}

					if !okA || !okB {
						return
					}

					next(ctx, lastA, lastB)
					// ------所有來源都至少有一個才會開始------

					for {
						select {
						case itemA, okA = <-sourceA:
							if okA {
								lastA = itemA
								next(ctx, lastA, lastB)
							}
						case itemB, okB = <-sourceB:
							if okB {
								lastB = itemB
								next(ctx, lastA, lastB)
							}
						}

						if !okA && !okB {
							return
						}
					}
				}()

				return ch
			}
		}
	}
}
