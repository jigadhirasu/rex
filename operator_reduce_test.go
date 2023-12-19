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
		Reduce[int, []int](func(a []int, b int) []int {
			if a == nil {
				a = []int{}
			}
			return append(a, b)
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
		[][]int{{1, 2, 3, 4, 5}},
	)
}

func TestReduce3(t *testing.T) {

	pipe := Pipe1(
		Reduce[int, []int](func(a []int, b int) []int {
			if a == nil {
				a = []int{}
			}
			return append(a, b)
		}),
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

func TestReduce4(t *testing.T) {

	pipe := Pipe1(
		Reduce[int, map[int]string](func(a map[int]string, b int) map[int]string {
			if a == nil {
				a = map[int]string{}
			}
			a[b] = "a"
			return a
		}),
	)(
		Range(1, 3),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]map[int]string{
			{
				1: "a",
				2: "a",
				3: "a",
			},
		},
	)
}
