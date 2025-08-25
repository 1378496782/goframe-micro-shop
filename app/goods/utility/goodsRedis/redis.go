package goodsRedis

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

var goodsCache *gcache.Cache

// InitGoodsRedis 初始化商品服务Redis
func InitGoodsRedis(ctx context.Context) error {
	// 获取Redis配置
	redisConfig, err := g.Cfg().Get(ctx, "redis.goods")
	if err != nil {
		return fmt.Errorf("获取Redis配置失败: %v", err)
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
	goodsCache = gcache.New()
	goodsCache.SetAdapter(gcache.NewAdapterRedis(redis))

	// 测试连接
	if _, err := redis.Do(ctx, "PING"); err != nil {
		return fmt.Errorf("Redis连接测试失败: %v", err)
	}

	g.Log().Info(ctx, "商品服务Redis初始化成功")
	return nil
}

// GetGoodsCache 获取商品缓存实例
func GetGoodsCache() *gcache.Cache {
	return goodsCache
}
