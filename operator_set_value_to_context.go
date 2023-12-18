package rex

import "fmt"

// 將一些值設定到 Context 中
func SetValueToContext[A any](f Func0) func(iterable Iterable[A]) Reader[A] {
	return func(iterable Iterable[A]) Reader[A] {
		return func(ctx Context) (next Iterable[A]) {

			next = iterable

			defer func() {
				if r := recover(); r != nil {
					next = FromItem[A](ItemError[A](fmt.Errorf("panic error: %v", r)))
				}
			}()

			if err := f(ctx); err != nil {
				next = FromItem[A](ItemError[A](err))
				return
			}

			return
		}
	}
}
