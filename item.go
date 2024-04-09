package rex

import (
	"reflect"
)

// ItemOf 用來產生一個 Item
func ItemOf[A any](a A) Item[A] {
	return func() (A, error) {
		return a, nil
	}
}

// ItemError 用來產生一個錯誤的 Item
func ItemError[A any](err error) Item[A] {
	return func() (A, error) {
		var a A
		return a, err
	}
}

// ItemAError 用來產生一個錯誤的 Item
func ItemAError[A any](a A, err error) Item[A] {
	return func() (A, error) {
		return a, err
	}
}

// Item 是一個Either型別
type Item[A any] func() (A, error)

// Error 用來取得錯誤
func (i Item[A]) Error() error {
	_, err := i()
	return err
}

// Equal 用來比較兩個 Item 是否相等
func (i Item[A]) Equal(itemB Item[A]) bool {
	a, errA := i()
	b, errB := itemB()

	equal := reflect.DeepEqual(a, b)

	if errA != nil {
		if errB != nil {
			return errA.Error() == errB.Error() && equal
		}
		return false
	}

	return equal
}

// SendItem 用來將Item 輸入 Iterable
func SendItem[A any](ctx Context, ch chan<- Item[A], item Item[A]) bool {
	select {
	case <-ctx.Done():
		return false
	default:
		select {
		case <-ctx.Done():
			return false
		case ch <- item:
			return true
		}
	}
}
