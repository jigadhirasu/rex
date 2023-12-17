package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {

	data := []int{1, 2, 3, 1, 3, 5, 2, 4, 6}

	pipe := Pipe1(
		Distinct1[int](func(a int) int {
			return a
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
		[]int{1, 2, 3, 5, 4, 6},
	)
}
