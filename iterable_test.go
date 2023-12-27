package rex

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {

	result, err := Range(0, 10).ToSlice()

	assert.NoError(t, err)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
}

func TestInterval(t *testing.T) {

	pipe := Pipe1(
		Take[int](3),
	)

	i1 := Interval(time.Second)

	ctx := NewContext(context.TODO())

	result, err := pipe(i1)(ctx).ToSlice()

	assert.NoError(t, err)

	assert.Equal(t, []int{0, 1, 2}, result)

}
