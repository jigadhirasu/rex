package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnce(t *testing.T) {

	ctx := NewContext(context.TODO())

	pipe := Pipe1(
		Once(func(ctx Context, a int) error {

			Set(ctx, "a", a)

			return nil
		}),
	)(Range(1, 3))(ctx)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t, []int{1, 2, 3}, result)

	assert.Equal(t, 1, Get[int](ctx, "a"))

}
