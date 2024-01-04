package rex

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZip1(t *testing.T) {

	iter2 := Interval(100 * time.Millisecond)

	pipe := Pipe2(
		ZipWith[int, int, int](iter2, func(ctx Context, a, b int) (int, error) {
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

	assert.NoError(t, Equal(pipe(ctx),
		FromItem[int](
			ItemOf(0),
			ItemOf(2),
			ItemOf(4),
			ItemError[int](errors.New("context canceled")),
		),
	))

}

func TestZip2(t *testing.T) {

	db := NewSubjectReplay[int](1)
	go func() {
		Next(db)(10)

		<-time.After(time.Second)
		db.Close()
	}()

	ctx := NewContext(context.TODO())

	go func() {
		pipe1 := Pipe2(
			Take[int](3),
			ZipWith[int, int, int](Subscribe(db), func(ctx Context, a, b int) (int, error) {
				return a + b, nil
			}),
		)(
			Range(0, 10),
		)

		assert.NoError(t, Equal(pipe1(ctx),
			FromItem[int](
				ItemOf(10),
			),
		))

	}()

	pipe2 := Pipe1(
		ZipWith[int, int, int](Subscribe(db), func(ctx Context, a, b int) (int, error) {
			return a + b, nil
		}),
	)(
		Range(10, 10),
	)

	assert.NoError(t, Equal(
		pipe2(ctx),
		FromItem[int](
			ItemOf(20),
		),
	))

}
