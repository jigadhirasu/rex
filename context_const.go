package rex

import (
	"fmt"
	"strings"
	"time"
)

// GenKey 用來產生一個新的 key
func GenKey(prefixKey ...string) string {
	var prefix string
	if len(prefixKey) > 0 {
		prefix = "_" + strings.Join(prefixKey, "_")
	}

	return fmt.Sprintf("ctx%s_%d", prefix, time.Now().UnixNano())
}

const ctxErrorKey = "ctx_error"

const ctxCancelKey = "ctx_cancel"

const CtxRequestKey = "ctx_request"
