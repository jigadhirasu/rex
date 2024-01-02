package rex

import (
	"context"
	"testing"
	"time"
)

func TestZip1(t *testing.T) {

	iter2 := Interval(100 * time.Millisecond)

	pipe := Pipe2(
		ZipFromIterable[int, int, int](iter2, func(ctx Context, a, b int) (int, error) {
			return a + b, nil
		}, WithOnErrorStrategy(ContinueOnError)),
		Map[int](func(ctx Context, a int) (int, error) {
			if a > 5 {
				ctx.Cancel()
			}

			return a, nil
		}),
	)(
		Range(0, 5),
	)

	ctx := NewContext(context.TODO())

	result, err := pipe(ctx).ToSlice()

	if err != nil {
		t.Error(err)
	}

	if len(result) != 5 {
		t.Error("Expected 5 items")
	}
}
