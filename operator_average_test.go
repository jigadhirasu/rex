package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage1(t *testing.T) {

	pipe := Pipe1(
		Average[int],
	)(
		Range(1, 5),
	)(
		NewContext(context.TODO()),
	)

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{3},
	)
}
