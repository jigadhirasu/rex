package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap1(t *testing.T) {

	pipe := Pipe2(
		Map1[int](func(ctx Context, a int) (int, error) {
			return a + 1, nil
		})(WithOnErrorStrategy(ContinueOnError)),
		Map1[int](func(ctx Context, a int) (int, error) {
			return a - 1, nil
		})(),
	)(
		Range(1, 10),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, result)

}
