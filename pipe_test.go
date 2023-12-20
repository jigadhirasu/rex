package rex

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type A struct {
	Name string
}

func TestPipe(t *testing.T) {

	start := From[float64](1, 2, 3)

	ctx := NewContext(context.TODO())
	go func() {
		<-time.After(time.Millisecond * 500)
		ctx.Cancel()
	}()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := Pipe3(
		Map[float64, int](func(ctx Context, a float64) (int, error) {
			return int(a * 100), nil
		})(),
		Map[int, string](func(ctx Context, a int) (string, error) {
			return fmt.Sprintf("%d", a), nil
		})(),
		FlatMap[string, A](func(ctx Context, a string) Iterable[A] {
			<-time.After(time.Millisecond * time.Duration(r.Intn(1000)))
			return From[A](
				A{Name: a + "3"},
				A{Name: a + "1"},
				A{Name: a + "2"},
			)
		}),
	)(
		start,
	)(ctx)

	for item := range result() {
		fmt.Println(item())
	}
}
