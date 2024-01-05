package rex

func CombineLatestWith[A, B, C any](iterableB Iterable[B], f Func2[A, B, C], opts ...applyOption) PipeLine[A, C] {
	return func(iterable Iterable[A]) Reader[C] {
		return _combineLatest[A, B, C](f, opts...)(iterable, iterableB)
	}
}
