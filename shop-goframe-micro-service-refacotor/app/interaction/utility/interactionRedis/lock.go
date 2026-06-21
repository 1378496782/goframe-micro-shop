package interactionRedis

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

const (
	commentLikeCountCalibrateLockKey = "interaction:comment:like_count:calibrate_lock"
	commentLikeCountCalibrateLockTTL = 30 * time.Minute
)

func TryAcquireCommentLikeCountCalibrateLock(ctx context.Context) (lockValue string, acquired bool, err error) {
	lockValue = newLockValue()
	expireSeconds := int64(commentLikeCountCalibrateLockTTL / time.Second)
	if expireSeconds <= 0 {
		expireSeconds = 1
	}

	v, err := g.Redis().GroupString().Set(ctx, commentLikeCountCalibrateLockKey, lockValue, gredis.SetOption{
		TTLOption: gredis.TTLOption{EX: &expireSeconds},
		NX:        true,
	})
	if err != nil {
		return "", false, err
	}

	return lockValue, v != nil && !v.IsNil(), nil
}

func ReleaseCommentLikeCountCalibrateLock(ctx context.Context, lockValue string) error {
	if lockValue == "" {
		return nil
	}

	luaScript := `
if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
else
	return 0
end
`
	_, err := g.Redis().Do(ctx, "EVAL", luaScript, 1, commentLikeCountCalibrateLockKey, lockValue)
	return err
}

func newLockValue() string {
	hostname, _ := os.Hostname()
	return fmt.Sprintf("%s:%d:%d", hostname, os.Getpid(), time.Now().UnixNano())
}
