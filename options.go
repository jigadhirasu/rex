package rex

type options struct {
	poolSize        int
	bufferSize      int
	OnErrorStrategy OnErrorStrategy
}

type applyOption func(*options)

func newOptions(opts ...applyOption) options {
	op := &options{
		poolSize:        1,
		bufferSize:      0,
		OnErrorStrategy: StopOnError,
	}

	for _, opt := range opts {
		opt(op)
	}

	return *op
}

func WithPool(size int) applyOption {
	return func(o *options) {
		o.poolSize = size
	}
}
