package rex

// Last 取得最後一個元素，會阻塞直到來源結束
func Last[A any](iterable Iterable[A]) Reader[A] {
	return takeLast[A](1)(iterable)
}
