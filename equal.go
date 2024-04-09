package rex

import (
	"errors"
)

// Equal 用來比較兩個 Iterable 是否相等
func Equal[A any](iterableA, iterableB Iterable[A]) error {
	chA := iterableA()
	chB := iterableB()

	for {
		itemA, okA := <-chA
		itemB, okB := <-chB

		if !okA && !okB {
			return nil
		}

		if okA && !okB {
			return errors.New("iterableA has more items")
		}

		if !okA && okB {
			return errors.New("iterableB has more items")
		}

		if !itemA.Equal(itemB) {
			return errors.New("items not equal")
		}
	}
}
