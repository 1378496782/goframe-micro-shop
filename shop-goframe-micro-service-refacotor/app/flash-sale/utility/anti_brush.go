package utility

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

// AntiBrushChecker 防刷检查器
type AntiBrushChecker struct {
	cache *gcache.Cache
}

// NewAntiBrushChecker 创建防刷检查器
func NewAntiBrushChecker(cache *gcache.Cache) *AntiBrushChecker {
	return &AntiBrushChecker{
		cache: cache,
	}
}

// CheckUserBehavior 检查用户行为
func (a *AntiBrushChecker) CheckUserBehavior(ctx context.Context, userId uint32, ip string) error {
	// 检查用户行为频率
	userKey := fmt.Sprintf("flash_sale:anti_brush:user:%d", userId)
	ipKey := fmt.Sprintf("flash_sale:anti_brush:ip:%s", ip)

	// 检查用户行为频率 - 每分钟最多10次请求
	userCount, err := a.cache.Get(ctx, userKey)
	if err != nil {
		return fmt.Errorf("获取用户行为数据失败: %v", err)
	}

	if userCount != nil {
		count := gconv.Int(userCount)
		if count > 10 {
			return fmt.Errorf("用户行为异常，请稍后再试")
		}
	}

	// 检查IP行为频率 - 每分钟最多50次请求
	ipCount, err := a.cache.Get(ctx, ipKey)
	if err != nil {
		return fmt.Errorf("获取IP行为数据失败: %v", err)
	}

	if ipCount != nil {
		count := gconv.Int(ipCount)
		if count > 50 {
			return fmt.Errorf("网络行为异常，请稍后再试")
		}
	}

	// 增加计数
	if userCount == nil {
		err = a.cache.Set(ctx, userKey, 1, time.Minute)
	} else {
		count := gconv.Int(userCount)
		err = a.cache.Set(ctx, userKey, count+1, time.Minute)
	}

	if err != nil {
		return fmt.Errorf("更新用户行为数据失败: %v", err)
	}

	if ipCount == nil {
		err = a.cache.Set(ctx, ipKey, 1, time.Minute)
	} else {
		count := gconv.Int(ipCount)
		err = a.cache.Set(ctx, ipKey, count+1, time.Minute)
	}

	if err != nil {
		return fmt.Errorf("更新IP行为数据失败: %v", err)
	}

	return nil
}
