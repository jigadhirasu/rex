package rex

// 有Side Effect的Step
func Map[A, B any](f Func1[A, B]) func(iterable Iterable[A]) Reader[B] {
	return func(iterable Iterable[A]) Reader[B] {
		return func(ctx Context) Iterable[B] {
			ch := make(chan Item[B])

			go func() {
				defer close(ch)
				defer Catcher[B](ch)

				source := iterable()
				for {
					item, ok := <-source
					if !ok {
						return
					}

					a, err := item()
					if err != nil {
						ch <- ItemError[B](err)
						break
					}

					if !sendItem(ctx, ch, ItemAError(f(ctx, a))) {
						ch <- ItemError[B](ctx.Err())
						return
					}
				}
			}()

			return FromChanItem[B](ch)
		}
	}
}

func Map1[A any](f Func1[A, A]) func(iterable Iterable[A]) Reader[A] {
	return func(iterable Iterable[A]) Reader[A] {
		return Map[A, A](f)(iterable)
	}
}
