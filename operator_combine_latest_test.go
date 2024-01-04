package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombineLatest(t *testing.T) {

	ctx := NewContext(context.Background())

	pipe := Pipe1(
		ShareReplay[int](1),
	)(
		From[int](10),
	)(ctx)

	pipe1 := Pipe1(
		CombineLatestWith(pipe, func(ctx Context, a, b int) (int, error) {
			return a + b, nil
		}),
	)(
		Range(1, 5),
	)

	f := func() {
		result, err := pipe1(ctx).ToSlice()
		assert.NoError(t, err)
		assert.Equal(t, []int{11, 12, 13, 14, 15}, result)
	}

	go f()
	go f()
	f()

}
