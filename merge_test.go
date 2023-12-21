package rex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge1(t *testing.T) {

	p1 := Range(1, 10)
	p2 := Range(1, 10)

	ctx := NewContext(context.TODO())
	p3 := Merge(p1, p2)(ctx)

	result, err := p3.ToSlice()

	assert.NoError(t, err)

	assert.Equal(t, 20, len(result))

}
