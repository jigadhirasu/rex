package rex

func Tap[A any](f TapFunc[A]) PipeLine[A, A] {
	return _map[A, A](func(ctx Context, a A) (A, error) {
		f(ctx, a)
		return a, nil
	})
}
