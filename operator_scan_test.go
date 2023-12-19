// package rex

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestScan1(t *testing.T) {

// 	pipe := Pipe1(
// 		Scan[int, int](func(a, b int) int {
// 			return a + b
// 		}),
// 	)(
// 		Range(1, 5),
// 	)(
// 		NewContext(context.TODO()),
// 	)

// 	result, err := pipe.ToSlice()

// 	assert.NoError(t, err)

// 	assert.Equal(t,
// 		result,
// 		[]int{1, 3, 6, 10, 15},
// 	)
// }

// func TestScan2(t *testing.T) {

// 	pipe := Pipe1(
// 		Scan[int, []int](func(a []int, b int) []int {
// 			if a == nil {
// 				a = []int{}
// 			}
// 			return append(a, b)
// 		}),
// 	)(
// 		Range(1, 5),
// 	)(
// 		NewContext(context.TODO()),
// 	)

// 	result, err := pipe.ToSlice()

// 	assert.NoError(t, err)

// 	assert.Equal(t,
// 		result,
// 		[][]int{
// 			{1},
// 			{1, 2},
// 			{1, 2, 3},
// 			{1, 2, 3, 4},
// 			{1, 2, 3, 4, 5},
// 		},
// 	)
// }

// func TestScan3(t *testing.T) {

// 	pipe := Pipe1(
// 		Scan[int, []int](func(a []int, b int) []int {
// 			if a == nil {
// 				a = []int{}
// 			}
// 			return append(a, b)
// 		}),
// 	)(
// 		From[int](),
// 	)(
// 		NewContext(context.TODO()),
// 	)

// 	result, err := pipe.ToSlice()

// 	assert.NoError(t, err)

// 	assert.Equal(t,
// 		result,
// 		[][]int{nil},
// 	)
// }

// func TestScan4(t *testing.T) {

// 	pipe := Pipe1(
// 		Scan[int, map[int]string](func(a map[int]string, b int) map[int]string {
// 			if a == nil {
// 				a = map[int]string{}
// 			}
// 			a[b] = "a"
// 			return a
// 		}),
// 	)(
// 		Range(1, 3),
// 	)(
// 		NewContext(context.TODO()),
// 	)

// 	result, err := pipe.ToSlice()

// 	assert.NoError(t, err)

// 	assert.Equal(t,
// 		result,
// 		[]map[int]string{
// 			{1: "a"},
// 			{1: "a", 2: "a"},
// 			{1: "a", 2: "a", 3: "a"},
// 		},
// 	)
// }

// func TestScan5(t *testing.T) {

// 	pipe := Pipe1(
// 		Scan[int, A](func(s A, a int) A {
// 			s.Name += fmt.Sprintf("%d", a)
// 			return s
// 		}),
// 	)(
// 		Range(1, 3),
// 	)(
// 		NewContext(context.TODO()),
// 	)

// 	result, err := pipe.ToSlice()

// 	assert.NoError(t, err)

// 	assert.Equal(t,
// 		result,
// 		[]A{
// 			{Name: "1"},
// 			{Name: "12"},
// 			{Name: "123"},
// 		},
// 	)
// }
