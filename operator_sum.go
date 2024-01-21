package rex

func Sum[A Number](iterable Iterable[A]) Reader[A] {
	return _reduce[A, A](func(a, b A) A {
		return a + b
	})(iterable)
}
