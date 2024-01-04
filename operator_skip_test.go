package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkip1(t *testing.T) {

	ctx := NewContext(context.Background())

	pipe := Pipe1(
		Skip[int](2),
	)(
		Range(0, 5),
	)(ctx)

	result, err := pipe.ToSlice()
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 3, 4}, result)

}

func TestSkip2(t *testing.T) {

	ctx := NewContext(context.Background())

	pipe := Pipe1(
		Skip[int](10),
	)(
		Range(0, 5),
	)(ctx)

	result, err := pipe.ToSlice()
	assert.NoError(t, err)
	assert.Equal(t, []int{}, result)

}
