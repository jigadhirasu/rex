package rex

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestShare(t *testing.T) {

	ctx := NewContext(context.TODO())

	pipe := Pipe1(
		Share[int](),
	)(
		Range(1, 10),
	)

	go func() {
		for item := range pipe(ctx)() {
			fmt.Println(item())
		}
	}()

	<-time.After(time.Second * 3)
	for item := range pipe(ctx)() {
		fmt.Println(item())
	}

}
