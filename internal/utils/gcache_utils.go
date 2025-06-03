package utils

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gcache"
)

// GcacheSet 如果存在缓存更新，如果不存在缓存创建
func GcacheSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (err error) {
	gcacheVal, err := gcache.Get(ctx, key)
	if err != nil {
		return
	}

	if gcacheVal.IsNil() {
		err = gcache.Set(ctx, key, value, duration)
		if err != nil {
			return
		}
	} else {
		_, _, err = gcache.Update(ctx, key, value)
		if err != nil {
			return
		}
	}

	return

}
