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
		<-time.After(time.Second)
		Next(subject)(4, 5, 6)
		subject.Close()
	}()

	go func() {
		result, err := Subscribe(subject).ToSlice()

		assert.NoError(t, err)

		assert.Equal(t,
			result,
			[]int{1, 2, 3, 4, 5, 6},
		)
	}()

	<-time.After(time.Millisecond * 500)
	result, err := Subscribe(subject).ToSlice()
	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{2, 3, 4, 5, 6},
	)
}
