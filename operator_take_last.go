package rex

// 取最後 n 個元素
// 需等待 iterable 結束後才能知道最後的是誰，會阻塞
func TakeLast[A any](n int) PipeLine[A, A] {
	return takeLast[A](n)
}

func takeLast[A any](n int) PipeLine[A, A] {
	return func(iterable Iterable[A]) Reader[A] {
		return func(ctx Context) Iterable[A] {

			ch := make(chan Item[A])

			go func() {
				defer close(ch)
				defer Catcher[A](ch)

				buf := []A{}
				defer func() {
					for _, a := range buf {
						ch <- ItemOf[A](a)
					}
				}()

				source := iterable()

				for {
					item, ok := <-source
					if !ok {
						break
					}

					a, err := item()
					if err != nil {
						ch <- ItemError[A](err)
						return
					}

					buf = append(buf, a)
					if len(buf) > n {
						buf = buf[1:]
					}
				}

			}()

			return FromChanItem[A](ch)
		}
	}
}
