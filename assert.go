package rex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Assert[A any](t *testing.T, iterableA, iterableB Iterable[A]) {

	chA := iterableA()
	chB := iterableB()

	for {
		itemA, okA := <-chA
		itemB, okB := <-chB

		if !okA && !okB {
			return
		}

		if okA && !okB {
			assert.Fail(t, "iterableA has more items")
			return
		}

		if !okA && okB {
			assert.Fail(t, "iterableB has more items")
			return
		}

		assert.True(t, itemA.Equal(itemB))
	}
}
