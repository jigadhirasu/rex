package rex

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type A struct {
	Name string
}

func TestPipe(t *testing.T) {

	ctx := NewContext(context.TODO())
	go func() {
		<-time.After(time.Millisecond * 2000)
		ctx.Cancel()
	}()

	// r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := Pipe4(
		Map[float64, int](func(ctx Context, a float64) (int, error) {
			return int(a * 100), nil
		}),
		Map[int, string](func(ctx Context, a int) (string, error) {
			return fmt.Sprintf("%d", a), nil
		}),
		Tap[string](func(ctx Context, a string) {
			fmt.Println(a)
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
	)(ctx)

	for item := range result() {
		fmt.Println(item())
	}

}
