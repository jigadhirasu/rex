package rex

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMap1(t *testing.T) {

	start := time.Now()

	pipe := Pipe2(
		Map1[int](func(ctx Context, a int) (int, error) {
			return a + 1, nil
		})(),
		Map1[int](func(ctx Context, a int) (int, error) {
			return a - 1, nil
		})(),
	)(
		Range(1, 1000),
	)(
		NewContext(context.TODO()),
	)

	_, err := pipe.ToSlice()

	assert.NoError(t, err)

	// fmt.Println(result)

	fmt.Println(time.Since(start))

}
