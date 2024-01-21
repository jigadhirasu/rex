package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce1(t *testing.T) {

	pipe := Pipe1(
		Reduce[int, int](func(a, b int) int {
			return a + b
		}),
	)(
		Range(1, 5),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{15},
	)
}

func TestReduce2(t *testing.T) {

	pipe := Pipe1(
		ReduceSlice[int](),
	)(
		Range(1, 5),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		[][]int{{1, 2, 3, 4, 5}},
		result,
	)
}

func TestReduce3(t *testing.T) {

	pipe := Pipe1(
		ReduceSlice[int](),
	)(
		From[int](),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[][]int{nil},
	)
}

func TestReduceMap(t *testing.T) {

	type Data struct {
		A int
		B string
	}

	pipe := Pipe1(
		ReduceMap[Data, int](func(a Data) int {
			return a.A
		}),
	)(
		From(
			Data{A: 1, B: "a"},
			Data{A: 2, B: "b"},
			Data{A: 1, B: "d"},
		),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		[]map[int]Data{
			{
				1: {A: 1, B: "d"},
				2: {A: 2, B: "b"},
			},
		},
		result,
	)
}
