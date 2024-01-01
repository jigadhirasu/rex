package rex

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTap1(t *testing.T) {

	os.Setenv("PROJECT_MODE", "main")

	ctx := NewContext(context.TODO())

	pipe := Pipe1(
		Tap[int](func(ctx Context, a int) error {
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
