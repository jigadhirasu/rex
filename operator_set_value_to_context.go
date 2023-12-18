package rex

// 將一些值設定到 Context 中
func SetValueToContext[A any](f Func0) func(iterable Iterable[A]) Reader[A] {
	return func(iterable Iterable[A]) Reader[A] {
		return func(ctx Context) Iterable[A] {
			if err := f(ctx); err != nil {
				return FromItem[A](ItemError[A](err))
			}

			return iterable
		}
	}
}
