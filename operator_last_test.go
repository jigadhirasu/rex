package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLast1(t *testing.T) {

	data := []int{9, 2, 3, 1, 3, 5, 2, 4, 6}

	pipe := Pipe1(
		Last[int],
	)(
		From[int](data...),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{6},
	)
}

func TestLast2(t *testing.T) {

	data := []int{}

	pipe := Pipe1(
		Last[int],
	)(
		From[int](data...),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{},
	)
}

func TestLast3(t *testing.T) {

	data := []int{}

	pipe := Pipe2(
		Default[int](6),
		Last[int],
	)(
		From[int](data...),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{6},
	)
}
