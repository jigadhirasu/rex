package rex

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSubject1(t *testing.T) {

	subject := NewSubject[int]()

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
		[]int{4, 5, 6},
	)
}

func TestSubject2(t *testing.T) {

	subject := NewSubject[int]()

	go func() {
		<-time.After(time.Second)
		Next(subject)(4, 5, 6)
		subject.Close()
	}()

	Next(subject)(1, 2, 3)

	result, err := Subscribe(subject).ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{4, 5, 6},
	)
}

func TestSubject3(t *testing.T) {

	ctx := NewContext(context.TODO())

	pipe := Pipe3[int](
		Filter(func(a int) bool {
			return a%2 == 0
		}),
		Map(func(ctx Context, a int) (int, error) {
			return a * 2, nil
		}),
		MergeMap(func(ctx Context, a int) Iterable[int] {
			return From[int](a*2, a*3)
		}),
	)(
		Range(1, 2),
	)(ctx)

	subject := NewSubject[int]()

	go func() {
		pipe2 := Pipe1[int](
			Map(func(ctx Context, a int) (int, error) {
				return a + 100, nil
			}),
		)(
			pipe,
		)(ctx)

		<-time.After(time.Second)
		NextChanItem(subject)(pipe2())

		Next(subject)(3)

		subject.Close()
	}()

	result, err := Subscribe(subject).ToSlice()

	assert.NoError(t, err)

	assert.Equal(t,
		result,
		[]int{108, 112, 3},
	)
}
