package rex

import (
	"fmt"
	"os"
)

// 只有在 PROJECT_MODE=main 時才會捕捉 panic
func Catcher[A any](ch chan<- Item[A]) {

	projectMode := os.Getenv("PROJECT_MODE")

	if projectMode == "main" {
		if err := recover(); err != nil {
			ch <- ItemError[A](fmt.Errorf("panic error: %v", err))
		}
	}
}
