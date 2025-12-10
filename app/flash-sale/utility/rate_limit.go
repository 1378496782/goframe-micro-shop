package utility

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

// RateLimiter 简单限流器
type RateLimiter struct {
	cache *gcache.Cache
}

// NewRateLimiter 创建限流器
func NewRateLimiter(cache *gcache.Cache) *RateLimiter {
	return &RateLimiter{
		cache: cache,
	}
}

// CheckLimit 检查是否超过限流阈值
func (r *RateLimiter) CheckLimit(ctx context.Context, key string, limit int, duration time.Duration) (bool, error) {
	// 获取当前计数
	count, err := r.cache.Get(ctx, key)
	if err != nil {
		return false, err
	}

	currentCount := gconv.Int(count)
	if currentCount >= limit {
		return false, nil // 超过限流
	}

	// 增加计数
	newCount := currentCount + 1
	if currentCount == 0 {
		// 第一次设置，设置过期时间
		err = r.cache.Set(ctx, key, newCount, duration)
	} else {
		// 更新计数，保持原有过期时间
		err = r.cache.Set(ctx, key, newCount, 0)
	}

	return err == nil, err
}

// UserRateLimit 用户限流检查
func UserRateLimit(ctx context.Context, userId uint32, cache *gcache.Cache) error {
	limiter := NewRateLimiter(cache)
	key := fmt.Sprintf("flash_sale:rate_limit:user:%d", userId)

	allowed, err := limiter.CheckLimit(ctx, key, 5, time.Second) // 每秒5次
	if err != nil {
		return fmt.Errorf("限流检查失败: %v", err)
	}

	if !allowed {
		return fmt.Errorf("请求过于频繁，请稍后再试")
	}

	return nil
}

// IPRateLimit IP限流检查
func IPRateLimit(ctx context.Context, ip string, cache *gcache.Cache) error {
	limiter := NewRateLimiter(cache)
	key := fmt.Sprintf("flash_sale:rate_limit:ip:%s", ip)

	allowed, err := limiter.CheckLimit(ctx, key, 20, time.Second) // 每秒20次
	if err != nil {
		return fmt.Errorf("IP限流检查失败: %v", err)
	}

	if !allowed {
		return fmt.Errorf("当前网络请求过多，请稍后再试")
	}

	return nil
}

// GetClientIP 获取客户端IP
func GetClientIP(ctx context.Context) string {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return ""
	}

	// 优先获取X-Forwarded-For
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip = r.GetClientIp()
	}

	return ip
}

// CheckGlobalLimit 检查全局限流
func (r *RateLimiter) CheckGlobalLimit(ctx context.Context) error {
	key := "flash_sale:rate_limit:global"
	allowed, err := r.CheckLimit(ctx, key, 1000, time.Second) // 每秒1000次全局请求
	if err != nil {
		return fmt.Errorf("全局限流检查失败: %v", err)
	}
	if !allowed {
		return fmt.Errorf("系统繁忙，请稍后再试")
	}
	return nil
}

// CheckUserLimit 检查用户限流
func (r *RateLimiter) CheckUserLimit(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("flash_sale:rate_limit:user:%d", userId)
	allowed, err := r.CheckLimit(ctx, key, 10, time.Second) // 每秒10次用户请求
	if err != nil {
		return fmt.Errorf("用户限流检查失败: %v", err)
	}
	if !allowed {
		return fmt.Errorf("请求过于频繁，请稍后再试")
	}
	return nil
}

// CheckIPLimit 检查IP限流
func (r *RateLimiter) CheckIPLimit(ctx context.Context, ip string) error {
	key := fmt.Sprintf("flash_sale:rate_limit:ip:%s", ip)
	allowed, err := r.CheckLimit(ctx, key, 50, time.Second) // 每秒50次IP请求
	if err != nil {
		return fmt.Errorf("IP限流检查失败: %v", err)
	}
	if !allowed {
		return fmt.Errorf("当前网络请求过多，请稍后再试")
	}
	return nil
}

// CheckPurchaseLimit 检查购买限制
func (r *RateLimiter) CheckPurchaseLimit(ctx context.Context, userId uint32, goodsId uint32) error {
	key := fmt.Sprintf("flash_sale:purchase_limit:%d:%d", userId, goodsId)
	allowed, err := r.CheckLimit(ctx, key, 1, time.Hour) // 每小时只能购买一次
	if err != nil {
		return fmt.Errorf("购买限制检查失败: %v", err)
	}
	if !allowed {
		return fmt.Errorf("您已购买过该商品，请稍后再试")
	}
	return nil
}

// RecordPurchase 记录购买记录
func (r *RateLimiter) RecordPurchase(ctx context.Context, userId uint32, goodsId uint32) error {
	key := fmt.Sprintf("flash_sale:purchase_record:%d:%d", userId, goodsId)
	return r.cache.Set(ctx, key, 1, time.Hour*24) // 记录24小时
}
