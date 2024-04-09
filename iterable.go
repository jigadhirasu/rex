package rex

import "time"

// From 將 Any 轉成 Iterable
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

// FromItem 將 Item 轉成 Iterable
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

// FromChan 將chan 轉成 Iterable
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

// FromChanItem 將chan item 轉成 Iterable
func FromChanItem[A any](ch <-chan Item[A]) Iterable[A] {
	return func() <-chan Item[A] {
		return ch
	}
}

// Range 產生一個數字序列
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

// Interval 產生一個間隔序列
func Interval(duration time.Duration) Iterable[int] {
	return func() <-chan Item[int] {
		ch := make(chan Item[int])

		go func() {
			i := 0
			for {
				ch <- ItemOf[int](i)
				i++

				<-time.After(duration)
			}
		}()

		return ch
	}
}

// Iterable 是一個迭代器
type Iterable[A any] func() <-chan Item[A]

// ToSlice 將 Iterable 轉成 Slice
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
