package rex

// 取得第一個元素
func First[A any](iterable Iterable[A]) Reader[A] {
	return _take[A](1)(iterable)
}
