package rex

import (
	"fmt"
	"reflect"
)

func ItemOf[A any](a A) Item[A] {
	return func() (A, error) {
		return a, nil
	}
}

func ItemError[A any](err error) Item[A] {
	return func() (A, error) {
		var a A
		return a, err
	}
}

func ItemAError[A any](a A, err error) Item[A] {
	return func() (A, error) {
		return a, err
	}
}

type Item[A any] func() (A, error)

func (i Item[A]) Error() error {
	_, err := i()
	return err
}

func (i Item[A]) Equal(itemB Item[A]) bool {
	a, errA := i()
	b, errB := itemB()

	fmt.Println("A:", a, errA, "B:", b, errB)

	equal := reflect.DeepEqual(a, b)

	if errA != nil {
		if errB != nil {
			return errA.Error() == errB.Error() && equal
		}
		return false
	}

	return equal
}

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
