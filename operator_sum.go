package rex

// Sum 用來計算 Iterable 中所有元素的總和
func Sum[A Number](iterable Iterable[A]) Reader[A] {
	return _reduce[A, A](func(a, b A) A {
		return a + b
	})(iterable)
}
