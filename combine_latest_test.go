package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAAAA(t *testing.T) {

	ctx := NewContext(context.Background())

	cb_ := _combineLatest1(
		func(ctx Context, a, b int) (int, error) {
			return a + b, nil
		},
	)(
		From[int](10),
		Range(1, 5),
	)(ctx)

	f := func() {
		result, err := cb_.ToSlice()
		assert.NoError(t, err)
		assert.Equal(t, []int{11, 12, 13, 14, 15}, result)
	}

	go f()
	go f()
	f()
}
