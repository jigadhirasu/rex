package rex

// Number 定義數字型態
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Average 用來計算平均值
func Average[A Number](iterable Iterable[A]) Reader[A] {
	return func(ctx Context) Iterable[A] {
		return func() <-chan Item[A] {
			ch := make(chan Item[A])

			go func() {
				defer close(ch)
				defer Catcher[A](ch)

				source := iterable()

				var sum A
				count := 0
				for {
					item, ok := <-source
					if !ok {
						break
					}

					a, err := item()
					if err != nil {
						ch <- ItemError[A](err)
						break
					}

					count++
					sum += a
				}

				if count > 0 {
					ch <- ItemOf[A](sum / A(count))
				}
			}()

			return ch
		}
	}
}
