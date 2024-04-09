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

// WithPool 設定 goroutine pool 的大小
func WithPool(size uint32) applyOption {
	return func(o *options) {
		o.poolSize = size
	}
}

// WithBuffer 設定 buffer 的大小
func WithBuffer(size uint32) applyOption {
	return func(o *options) {
		o.bufferSize = size
	}
}

// WithOnErrorStrategy 設定錯誤處理策略
func WithOnErrorStrategy(strategy OnErrorStrategy) applyOption {
	return func(o *options) {
		o.OnErrorStrategy = strategy
	}
}
