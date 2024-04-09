package rex

// Share 用來共享來源，讓多個讀取者可以共享相同的來源
func Share[A any](iterable Iterable[A]) Reader[A] {
	return ShareReplay[A](1)(iterable)
}

// ShareReplay 用來共享來源，讓多個讀取者可以共享相同的來源，並且可以回放指定數量的元素
func ShareReplay[A any](bufferSize int) PipeLine[A, A] {
	return func(iterable Iterable[A]) Reader[A] {

		subject := NewSubjectReplay[A](bufferSize)

		source := iterable()

		go func() {
			NextChanItem[A](subject)(source)
			subject.Close()
		}()

		return func(ctx Context) Iterable[A] {
			return func() <-chan Item[A] {
				return Subscribe(subject)()
			}
		}
	}
}
