package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {

	data := []int{9, 2, 3, 1, 3, 5, 2, 4, 6}

	pipe := Pipe1(
		First[int],
	)(
		From[int](data...),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{9},
	)
}
