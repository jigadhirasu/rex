package rex

func Share[A any]() PipeLine[A, A] {
	return func(iterable Iterable[A]) Reader[A] {

		subject := NewSubject[A]()

		source := iterable()

		go func() {
			NextChanItem[A](subject)(source)
			subject.Close()
		}()

		return func(ctx Context) Iterable[A] {
			return Subscribe(subject)
		}
	}
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
			return Subscribe(subject)
		}
	}
}
