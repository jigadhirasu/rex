package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	pipe := Pipe1(
		Filter[int](func(a int) bool {
			return a%2 == 0
		}),
	)(
		From[int](data...),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{2, 4, 6, 8, 0},
	)
}
