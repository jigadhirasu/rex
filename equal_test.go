package rex

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual1(t *testing.T) {

	var err error
	err = Equal(
		Range(0, 5),
		From(0, 1, 2, 3, 4),
	)

	assert.NoError(t, err)

	ctx := NewContext(context.TODO())
	pipe := Pipe1(
		Map[int, int](func(ctx Context, a int) (int, error) {
			if a > 2 {
				return 0, errors.New("error a > 2")
			}
			return a, nil
		}),
	)(Range(0, 10))(ctx)

	err = Equal(
		pipe,
		FromItem(
			ItemOf(0),
			ItemOf(1),
			ItemOf(2),
			ItemError[int](errors.New("error a > 2")),
		),
	)

	assert.NoError(t, err)
}
