package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartWith1(t *testing.T) {

	ctx := NewContext(context.TODO())

	pipe := Pipe1(
		StartWith[int](From[int](1, 2, 3)),
	)(
		Range(1, 5),
	)(ctx)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t, []int{1, 2, 3, 1, 2, 3, 4, 5}, result)

}
