package rex

// Transfer1 is a function that takes an A and returns a B.
type Transfer1[A, B any] func(a A) B

// Transfer2 is a function that takes an A and a B and returns a C.
type Transfer2[A, B, C any] func(a A, b B) C

type Func0[A any] func(ctx Context, a A) error
type Func1[A, B any] func(ctx Context, a A) (B, error)
type Func2[A, B, C any] func(ctx Context, a A, b B) (C, error)
type Func3[A, B, C, D any] func(ctx Context, a A, b B, c C) (D, error)
type Func4[A, B, C, D, E any] func(ctx Context, a A, b B, c C, d D) (E, error)
type Func5[A, B, C, D, E, F any] func(ctx Context, a A, b B, c C, d D, e E) (F, error)
type Func6[A, B, C, D, E, F, G any] func(ctx Context, a A, b B, c C, d D, e E, f F) (G, error)
type Func7[A, B, C, D, E, F, G, H any] func(ctx Context, a A, b B, c C, d D, e E, f F, g G) (H, error)
type Func8[A, B, C, D, E, F, G, H, I any] func(ctx Context, a A, b B, c C, d D, e E, f F, g G, h H) (I, error)
type Func9[A, B, C, D, E, F, G, H, I, J any] func(ctx Context, a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, error)
type Func10[A, B, C, D, E, F, G, H, I, J, K any] func(ctx Context, a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) (K, error)

// HFunc1 is a function that takes an A and returns an Iterable of B.
type HFunc1[A, B any] func(ctx Context, a A) Iterable[B]

// HFunc2 is a function that takes an A and a B and returns an Iterable of C.
type HFunc2[A, B, C any] func(ctx Context, a A, b B) Iterable[C]

// Predicate is a function that takes an A and returns a bool.
type Predicate[A any] func(a A) bool

// Subject is a function that returns a channel to send items and an iterable to receive items.
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
