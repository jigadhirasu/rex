package rex

import (
	"cmp"
)

// Distinct 會回傳一個新的 Iterable，此 Iterable 會將重複的元素過濾掉。
func Distinct[A any, B cmp.Ordered](f Transfer1[A, B]) func(iterable Iterable[A]) func(ctx Context) Iterable[A] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[A] {
		return func(ctx Context) Iterable[A] {
			ch := make(chan Item[A])

			go func() {
				defer close(ch)

				offset := map[B]bool{}

				source := iterable()
				for i := 0; i < 4096; i++ {
					item, ok := <-source
					if !ok {
						return
					}

					a, err := item()
					if err != nil {
						ch <- ItemError[A](err)
						return
					}

					if _, ok := offset[f(a)]; ok {
						continue
					}

					offset[f(a)] = true

					if !sendItem(ctx, ch, item) {
						ch <- ItemError[A](ctx.Err())
						return
					}
				}
			}()

			return FromChanItem[A](ch)
		}
	}
}

func Distinct1[A cmp.Ordered](f Transfer1[A, A]) func(iterable Iterable[A]) func(ctx Context) Iterable[A] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[A] {
		return Distinct[A, A](f)(iterable)
	}
}
