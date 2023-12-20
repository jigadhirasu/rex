package rex

type options struct {
	poolSize        uint32
	bufferSize      uint32
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

func WithPool(size uint32) applyOption {
	return func(o *options) {
		o.poolSize = size
	}
}

func WithBuffer(size uint32) applyOption {
	return func(o *options) {
		o.bufferSize = size
	}
}

func WithOnErrorStrategy(strategy OnErrorStrategy) applyOption {
	return func(o *options) {
		o.OnErrorStrategy = strategy
	}
}
