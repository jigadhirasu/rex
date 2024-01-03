package rex

import (
	"context"
	"testing"
	"time"
)

func TestCombineLatest(t *testing.T) {

	pipe := Pipe2(
		Take[int](3),
		CombineLatestWith[int, int, int](Interval(time.Second), func(ctx Context, a, b int) (int, error) {
			return a + b, nil
		}),
	)(
		Interval(time.Millisecond * 200),
	)

	ctx := NewContext(context.TODO())

	Assert(t, pipe(ctx),
		FromItem[int](
			ItemOf(0),
			ItemOf(1),
			ItemOf(2),
		),
	)

}
