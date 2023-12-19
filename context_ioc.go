package rex

import (
	"context"
	"fmt"
)

// 取得值，若不存在則 panic
func Get[A any](ctx Context, key string) A {
	ctx.RLock()
	defer ctx.RUnlock()

	value := ctx.Value(key)
	v, ok := value.(A)
	if !ok {
		panic(fmt.Sprintln("key:", key, " not exists"))
	}

	return v
}

// 取得值，若不存在則回傳預設值
func Maybe[A any](ctx Context, key string) (A, bool) {
	ctx.RLock()
	defer ctx.RUnlock()

	value := ctx.Value(key)
	v, ok := value.(A)
	return v, ok
}

// 設定錯誤
func SetError(ctx Context, err error) {
	Set(ctx, ctxErrorKey, err)
	cancel, ok := Maybe[context.CancelFunc](ctx, ctxCancelKey)
	if ok {
		cancel()
	}
}

// 設定
func Set(ctx Context, key string, value any) {
	ctx.Lock()
	if ctx.Value(key) != nil {
		panic(fmt.Sprintln("key:", key, " already exists"))
	}

	ctx.WithContext(context.WithValue(ctx.Original(), String2Any(key), value))
	ctx.Unlock()
}

// 複寫
func Overwrite(ctx Context, key string, value any) {
	ctx.Lock()
	if ctx.Value(key) != nil {
		fmt.Println("overwrite key:", key, " value:", value)
	}
	ctx.WithContext(context.WithValue(ctx.Original(), String2Any(key), value))
	ctx.Unlock()
}

// 刪除
func Delete(ctx Context, key string) {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.Value(key) != nil {
		ctx.WithContext(context.WithValue(ctx.Original(), String2Any(key), nil))
		return
	}
	fmt.Println("key:", key, " not exists")
}

func String2Any(value string) any {
	return value
}
