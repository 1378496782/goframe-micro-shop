package goodsRedis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

// SetEmptyGoodsDetail 添加设置空缓存的函数，防止缓存穿透
func SetEmptyGoodsDetail(ctx context.Context, productId uint32) error {
	key := fmt.Sprintf("goods:detail:%d", productId)
	// 设置一个短时间的空值，防止缓存穿透
	return goodsCache.Set(ctx, key, nil, 1*time.Minute)
}

// SetGoodsDetail 设置商品详情缓存
func SetGoodsDetail(ctx context.Context, productId uint32, data interface{}) error {
	key := fmt.Sprintf("goods:detail:%d", productId)

	// 使用JSON序列化确保数据类型一致性
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return goodsCache.Set(ctx, key, jsonData, time.Hour)
}

// GetGoodsDetail 获取商品详情缓存
func GetGoodsDetail(ctx context.Context, productId uint32) (*g.Var, error) {
	key := fmt.Sprintf("goods:detail:%d", productId)
	result, err := goodsCache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	// 检查是否是空值标记
	if result.IsEmpty() || result.String() == "null" {
		return g.NewVar(nil), nil
	}

	return result, nil
}
