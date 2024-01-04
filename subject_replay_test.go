package rex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubjectReplay1(t *testing.T) {

	subject := NewSubjectReplay[int](2)

	Next(subject)(1, 2, 3, 4, 5)

	subject.Close()

	f := func() {
		result, err := Subscribe(subject).ToSlice()
		assert.NoError(t, err)
		assert.Equal(t, []int{4, 5}, result)
	}

	f()
	f()
	f()

}
