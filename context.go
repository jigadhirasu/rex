package rex

import (
	"context"
	"sync"
	"time"
)

// NewContext 產生一個新的 context
func NewContext(original context.Context) Context {

	ctx := &ContextImpl{
		Context:  original,
		ctxMutex: new(sync.RWMutex),
		RWMutex:  new(sync.RWMutex),
	}

	if _, ok := Maybe[func()](ctx, ctxCancelKey); !ok {
		original, cancel := context.WithCancel(original)
		ctx.WithContext(original)
		Overwrite(ctx, ctxCancelKey, cancel)
	}

	return ctx
}

// Context 封裝了 context.Context，並且提供了一些額外的功能
type Context interface {
	context.Context

	Lock()
	Unlock()
	RLock()
	RUnlock()
	Err() error

	Original() context.Context
	WithContext(context.Context)
	Cancel()
	WithDeadline(time.Time)
	WithTimeout(time.Duration)
}

type ContextImpl struct {
	context.Context
	ctxMutex *sync.RWMutex
	*sync.RWMutex
}

// 去得 context.Context
func (ctx *ContextImpl) Original() context.Context {
	return ctx.Context
}

func (ctx *ContextImpl) Err() error {
	if err, ok := Maybe[error](ctx, ctxErrorKey); ok {
		return err
	}

	return ctx.Context.Err()
}

// 用原生的 context.Context 傳遞值
func (ctx *ContextImpl) WithContext(nextCtx context.Context) {
	ctx.ctxMutex.Lock()
	ctx.Context = nextCtx
	ctx.ctxMutex.Unlock()
}

func (ctx *ContextImpl) Cancel() {
	Get[context.CancelFunc](ctx, ctxCancelKey)()
}

func (ctx *ContextImpl) WithDeadline(deadline time.Time) {
	newCtx, cancel := context.WithDeadline(ctx, deadline)
	ctx.WithContext(newCtx)
	Overwrite(ctx, ctxCancelKey, cancel)
}

func (ctx *ContextImpl) WithTimeout(timeout time.Duration) {
	newCtx, cancel := context.WithTimeout(ctx, timeout)
	ctx.WithContext(newCtx)
	Overwrite(ctx, ctxCancelKey, cancel)
}
