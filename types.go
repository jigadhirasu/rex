package rex

type Transfer1[A, B any] func(a A) B
type Transfer2[A, B, C any] func(a A, b B) C

type TapFunc[A any] func(ctx Context, a A)

type Func0 func(ctx Context) error
type Func1[A, B any] func(ctx Context, a A) (B, error)
type Func2[A, B, C any] func(ctx Context, a A, b B) (C, error)

type FlatFunc1[A, B any] func(ctx Context, a A) Iterable[B]
type FlatFunc2[A, B, C any] func(ctx Context, a A, b B) Iterable[C]

type Predicate[A any] func(a A) bool

type Subject[A any] func() (chan<- Item[A], Iterable[A])

func (subject Subject[A]) Close() {
	next, _ := subject()
	close(next)
}

// OnErrorStrategy is the Observable error strategy.
type OnErrorStrategy uint32

const (
	// StopOnError is the default error strategy.
	// An operator will stop processing items on error.
	StopOnError OnErrorStrategy = iota
	// ContinueOnError means an operator will continue processing items after an error.
	ContinueOnError
)
