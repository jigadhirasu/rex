package rex

// 將slice 轉成 Iterable
func From[A any](aa ...A) Iterable[A] {
	return func() <-chan Item[A] {
		ch := make(chan Item[A])

		go func() {
			for _, a := range aa {
				ch <- ItemOf[A](a)
			}
			close(ch)
		}()

		return ch
	}
}

// 將Item 轉成 Iterable
func FromItem[A any](aa ...Item[A]) Iterable[A] {
	return func() <-chan Item[A] {
		ch := make(chan Item[A])

		go func() {
			for _, itemA := range aa {
				ch <- itemA
			}
			close(ch)
		}()

		return ch
	}
}

// 將chan 轉成 Iterable
func FromChan[A any](source <-chan A) Iterable[A] {
	return func() <-chan Item[A] {
		ch := make(chan Item[A])

		go func() {
			for a := range source {
				ch <- ItemOf[A](a)
			}
			close(ch)
		}()

		return ch
	}
}

// 將chan item 轉成 Iterable
func FromChanItem[A any](ch <-chan Item[A]) Iterable[A] {
	return func() <-chan Item[A] {
		return ch
	}
}

// 產生一個數字序列
func Range(start, count int) Iterable[int] {
	return func() <-chan Item[int] {
		ch := make(chan Item[int])

		go func() {
			for i := start; i < start+count; i++ {
				ch <- ItemOf[int](i)
			}
			close(ch)
		}()

		return ch
	}
}

type Iterable[A any] func() <-chan Item[A]

func (iter Iterable[A]) ToSlice() ([]A, error) {
	aa := []A{}

	source := iter()

	for {

		item, ok := <-source
		if !ok {
			break
		}

		a, err := item()
		if err != nil {
			return nil, err
		}

		aa = append(aa, a)
	}

	return aa, nil
}
