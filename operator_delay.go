package rex

import "time"

// Delay 用來延遲發送元素
func Delay[A any](duration time.Duration, opts ...applyOption) PipeLine[A, A] {
	return _map[A](func(ctx Context, a A) (A, error) {
		<-time.After(duration)
		return a, nil
	}, opts...)
}
