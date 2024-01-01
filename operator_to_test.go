package rex

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapTo1(t *testing.T) {

	ctx := NewContext(context.TODO())

	pipe := Pipe1(
		MapTo[int, string]("hello"),
	)(
		Range(1, 3),
	)

	result, err := pipe(ctx).ToSlice()

	assert.NoError(t, err)

	assert.Equal(t, []string{"hello", "hello", "hello"}, result)
}

func TestMapTo2(t *testing.T) {

	ctx := NewContext(context.TODO())

	pipe := Pipe1(
		MapTo[int, string]("hello"),
	)(FromItem[int](
		ItemOf(1),
		ItemError[int](errors.New("error6996")),
		ItemOf(3),
	))

	_, err := pipe(ctx).ToSlice()

	assert.EqualError(t, err, "error6996")
}
