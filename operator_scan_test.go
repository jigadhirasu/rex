package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan1(t *testing.T) {

	pipe := Pipe1(
		Scan[int, int](0, func(a, b int) int {
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

func TestScan2(t *testing.T) {

	pipe := Pipe1(
		Scan[int, []int]([]int{}, func(a []int, b int) []int {
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

func TestScan3(t *testing.T) {

	pipe := Pipe1(
		Scan[int, []int]([]int{}, func(a []int, b int) []int {
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
		[][]int{{}},
	)
}

func TestScan4(t *testing.T) {

	pipe := Pipe1(
		Scan[int, map[int]string](map[int]string{}, func(a map[int]string, b int) map[int]string {
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
