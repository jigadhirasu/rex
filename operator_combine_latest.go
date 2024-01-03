package rex

import "fmt"

func CombineLatestWith[A, B, C any](iterableB Iterable[B], f Func2[A, B, C], opts ...applyOption) PipeLine[A, C] {
	return _combineLatest[A, B, C](iterableB, f, opts...)
}

func _combineLatest[A, B, C any](iterableB Iterable[B], f Func2[A, B, C], opts ...applyOption) PipeLine[A, C] {
	return func(iterableA Iterable[A]) Reader[C] {
		return func(ctx Context) Iterable[C] {

			ch := make(chan Item[C])

			go func() {
				defer close(ch)

				sourceA := iterableA()
				sourceB := iterableB()

				var Va A
				var Vb B

				// ------所有來源都至少有一個才會開始------
				next := func(ctx Context, a A, b B) {
					c, err := f(ctx, Va, Vb)
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

				itemA, okA := <-sourceA
				if !okA {
					return
				}
				a, err := itemA()
				if err != nil {
					if !SendItem(ctx, ch, ItemError[C](err)) {
						ch <- ItemError[C](ctx.Err())
					}

					return
				}
				Va = a

				itemB, okB := <-sourceB
				if !okB {
					return
				}
				b, err := itemB()
				if err != nil {
					if !SendItem(ctx, ch, ItemError[C](err)) {
						ch <- ItemError[C](ctx.Err())
					}

					return
				}
				Vb = b

				next(ctx, Va, Vb)
				// ------所有來源都至少有一個才會開始------

				for {
					fmt.Println("GGG", Va, Vb)
					select {
					case itemA, okA := <-sourceA:
						if !okA {
							return
						}
						a, err := itemA()
						if err != nil {
							if !SendItem(ctx, ch, ItemError[C](err)) {
								ch <- ItemError[C](ctx.Err())
							}

							return
						}
						Va = a
					case itemB, okB := <-sourceB:
						if !okB {
							return
						}
						b, err := itemB()
						if err != nil {
							if !SendItem(ctx, ch, ItemError[C](err)) {
								ch <- ItemError[C](ctx.Err())
							}

							return
						}
						Vb = b
					}

					next(ctx, Va, Vb)
				}
			}()

			return FromChanItem[C](ch)

		}
	}
}
