package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault1(t *testing.T) {

	data := []int{}

	pipe := Pipe1(
		Default[int](6),
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

func TestDefault2(t *testing.T) {

	data := []int{5}

	pipe := Pipe1(
		Default[int](6),
	)(
		From[int](data...),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{5},
	)

}
