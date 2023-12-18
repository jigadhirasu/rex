package rex

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetValueToContext(t *testing.T) {

	ctx := NewContext(context.TODO())

	pipe := Pipe1(
		SetValueToContext[int](func(ctx Context) error {
			panic("ggc")
		}),
	)(
		From(1),
	)(ctx)

	_, err := pipe.ToSlice()

	assert.EqualError(t,
		err,
		fmt.Sprintf("panic error: %v", "ggc"),
	)

}
