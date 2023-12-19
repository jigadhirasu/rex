package rex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSubjectReplay1(t *testing.T) {

	subject := NewSubjectReplay[int](2)

	go func() {
		Next(subject)(1, 2, 3)
		Next(subject)(4, 5, 6)
		<-time.After(time.Second * 2)
		subject.Close()
	}()

	<-time.After(time.Second)
	result, err := Subscribe(subject).ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{5, 6},
	)
}
