package rex

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type A struct {
	Name string
}

func TestPipe(t *testing.T) {

	pipe := Pipe3(
		Map[float64, int](func(ctx Context, a float64) (int, error) {
			return int(a * 100), nil
		}),
		Map[int, string](func(ctx Context, a int) (string, error) {
			return fmt.Sprintf("%d", a), nil
		}),
		MergeMap[string, A](func(ctx Context, a string) Iterable[A] {
			return From[A](
				A{Name: a + "1"},
				A{Name: a + "2"},
				A{Name: a + "3"},
			)
		}),
	)(
		From[float64](1, 2, 3),
	)

	ctx := NewContext(context.TODO())

	result, err := pipe(ctx).ToSlice()

	assert.NoError(t, err)

	assert.Equal(t, 9, len(result))

}
