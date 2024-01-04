package rex

import (
	"context"
	"testing"
	"time"
)

func TestShare(t *testing.T) {

	pipe := Pipe1(
		Share[int](),
	)(
		Range(1, 5),
	)

	ctx := NewContext(context.TODO())

	go func() {
		Assert(t, pipe(ctx), FromItem[int](
			ItemOf(1),
			ItemOf(2),
			ItemOf(3),
			ItemOf(4),
			ItemOf(5),
		))
	}()

	<-time.After(time.Second)

	Assert(t, pipe(ctx), FromItem[int](
		ItemOf(1),
		ItemOf(2),
		ItemOf(3),
		ItemOf(4),
		ItemOf(5),
	))
}
