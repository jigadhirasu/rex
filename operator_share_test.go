package rex

import (
	"context"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShare(t *testing.T) {

	ctx := NewContext(context.TODO())

	pipe := Pipe2(
		Take[int](2),
		Share[int],
	)(
		Interval(time.Millisecond * 200),
	)(ctx)

	okf := func(result []int) {
		fmt.Println(result)
		assert.LessOrEqual(t, len(result), 2)
	}

	f := func() {
		result, err := pipe.ToSlice()
		assert.NoError(t, err)
		okf(result)
	}

	go f()
	go f()
	go f()
	f()
	f()
	f()

	// wait all complete
	<-time.After(time.Millisecond * 500)

}

func TestShareReplay(t *testing.T) {

	ctx := NewContext(context.TODO())

	pipe := Pipe1(
		ShareReplay[int](3),
	)(
		Range(1, 5),
	)(ctx)

	okf := func(result []int) {
		for _, v := range []int{3, 4, 5} {
			assert.True(t, slices.Contains(result, v))
		}
	}

	f := func() {
		result, err := pipe.ToSlice()
		assert.NoError(t, err)
		okf(result)
	}

	go f()
	go f()
	go f()
	f()
	f()
	f()

}
