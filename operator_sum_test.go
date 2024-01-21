package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum1(t *testing.T) {

	ctx := NewContext(context.TODO())

	pipe := Pipe2(
		Map[int, float64](func(ctx Context, a int) (float64, error) {
			return float64(a) * 3.33, nil
		}),
		Sum[float64],
	)(Range(1, 3))(NewContext(ctx))

	result, err := pipe.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t, []float64{6 * 3.33}, result)
}
