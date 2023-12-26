package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufferCount1(t *testing.T) {

	ctx := NewContext(context.TODO())

	pipe := Pipe1(
		BufferCount[int](3),
	)(
		Range(1, 20),
	)

	result, err := pipe(ctx).ToSlice()

	assert.NoError(t, err)

	assert.Equal(t, [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
		{13, 14, 15},
		{16, 17, 18},
		{19, 20},
	}, result)

}
