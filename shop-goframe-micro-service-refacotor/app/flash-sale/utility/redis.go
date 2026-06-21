package utility

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

var (
	flashSaleCache *gcache.Cache
)

// InitFlashSaleRedis 初始化秒杀服务Redis
func InitFlashSaleRedis(ctx context.Context) error {
	// 获取Redis配置
	redisConfig, err := g.Cfg().Get(ctx, "redis.flash_sale")
	if err != nil {
		// 如果没有配置，使用默认商品Redis
		redisConfig, err = g.Cfg().Get(ctx, "redis.goods")
		if err != nil {
			return fmt.Errorf("获取Redis配置失败: %v", err)
		}
	}

	// 创建Redis实例
	config := &gredis.Config{}
	if err := redisConfig.Scan(config); err != nil {
		return fmt.Errorf("解析Redis配置失败: %v", err)
	}

	redis, err := gredis.New(config)
	if err != nil {
		return fmt.Errorf("创建Redis连接失败: %v", err)
	}

	// 创建缓存适配器
	flashSaleCache = gcache.New()
	flashSaleCache.SetAdapter(gcache.NewAdapterRedis(redis))

	// 测试连接
	if _, err := redis.Do(ctx, "PING"); err != nil {
		return fmt.Errorf("Redis连接测试失败: %v", err)
	}

	g.Log().Info(ctx, "秒杀服务Redis初始化成功")
	return nil
}

// GetFlashSaleCache 获取秒杀缓存实例
func GetFlashSaleCache() *gcache.Cache {
	return flashSaleCache
}
