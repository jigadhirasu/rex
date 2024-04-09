package rex

// Zip 用來將兩個 Iterable 合併成一個，並透過指定的函數對應元素
func ZipWith[A, B, C any](iterableB Iterable[B], f Func2[A, B, C], opts ...applyOption) PipeLine[A, C] {
	return func(iterable Iterable[A]) Reader[C] {
		return _zip[A, B, C](f, opts...)(iterable, iterableB)
	}
}
