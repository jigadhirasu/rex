package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat1(t *testing.T) {

	p1 := Range(1, 5)
	p2 := Range(1, 5)

	ctx := NewContext(context.TODO())
	p3 := Concat(p1, p2)(ctx)

	result, err := p3.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, result)
}
