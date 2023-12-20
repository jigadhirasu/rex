package rex

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap1(t *testing.T) {

	pipe := Pipe2(
		Map1[int](func(ctx Context, a int) (int, error) {
			return a + 1, nil
		}),
		Map1[int](func(ctx Context, a int) (int, error) {
			return a - 1, nil
		}),
	)(
		Range(1, 100),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	fmt.Println(result)

}
