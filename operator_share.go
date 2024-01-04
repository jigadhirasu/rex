package rex

func Share[A any](iterable Iterable[A]) Reader[A] {
	return ShareReplay[A](1)(iterable)
}

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
