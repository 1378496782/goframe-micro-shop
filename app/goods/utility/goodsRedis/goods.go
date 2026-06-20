package goodsRedis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"math/rand"
	"time"
)

const (
	categoryAllKey = "category:all:data"
	GoodsDetailKey = "goods:detail:"
	EmptyValue     = "__EMPTY__"
)

func ttlWithJitter(base time.Duration, maxJitter time.Duration) time.Duration {
	if maxJitter <= 0 {
		return base
	}
	return base + time.Duration(rand.Int63n(int64(maxJitter)))
}

// SetEmptyGoodsDetail 添加设置空缓存的函数，防止缓存穿透
func SetEmptyGoodsDetail(ctx context.Context, productId uint32) error {
	key := fmt.Sprintf("%s%d", GoodsDetailKey, productId)
	// 设置一个短时间的空值，防止缓存穿透
	return goodsCache.Set(ctx, key, EmptyValue, ttlWithJitter(1*time.Minute, 30*time.Second))
}

// SetGoodsDetail 设置商品详情缓存
func SetGoodsDetail(ctx context.Context, productId uint32, data interface{}) error {
	key := fmt.Sprintf("%s%d", GoodsDetailKey, productId)

	// 使用JSON序列化确保数据类型一致性
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return goodsCache.Set(ctx, key, jsonData, ttlWithJitter(time.Hour, 10*time.Minute))
}

// GetGoodsDetail 获取商品详情缓存
func GetGoodsDetail(ctx context.Context, productId uint32) (*g.Var, error) {
	key := fmt.Sprintf("%s%d", GoodsDetailKey, productId)
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

// DeleteGoodsDetail 删除商品详情数据缓存
func DeleteGoodsDetail(ctx context.Context, productId uint32) error {
	cache := GetGoodsCache()
	if cache == nil {
		return fmt.Errorf("goodsCache 未初始化")
	}

	key := fmt.Sprintf("%s%d", GoodsDetailKey, productId)
	_, err := cache.Remove(ctx, key)
	return err
}

// SetCategoryAll 设置分类全量数据缓存
func SetCategoryAll(ctx context.Context, data interface{}) error {
	// 使用JSON序列化确保数据类型一致性
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// 设置一周的缓存时间
	return goodsCache.Set(ctx, categoryAllKey, jsonData, 7*24*time.Hour)
}

// GetCategoryAll 获取分类全量数据缓存
func GetCategoryAll(ctx context.Context) (*gvar.Var, error) {
	result, err := goodsCache.Get(ctx, categoryAllKey)
	if err != nil {
		return nil, err
	}

	if result.IsEmpty() || result.String() == "null" {
		return gvar.New(nil), nil
	}

	return result, nil
}

// DeleteCategoryAll 删除分类全量数据缓存
func DeleteCategoryAll(ctx context.Context) error {
	_, err := goodsCache.Remove(ctx, categoryAllKey)
	return err
}

// DeleteKeys 批量删除多个缓存
func DeleteKeys(ctx context.Context, keys []uint32) error {
	cache := GetGoodsCache()
	if cache == nil {
		return fmt.Errorf("goodsCache 未初始化")
	}

	// 构建缓存 key 切片
	cacheKeys := make([]any, len(keys))
	for i, key := range keys {
		cacheKeys[i] = fmt.Sprintf("%s%d", GoodsDetailKey, key)
	}

	// 批量删除
	if _, err := cache.Remove(ctx, cacheKeys...); err != nil {
		g.Log().Warningf(ctx, "批量删除缓存失败: %v, keys=%v", err, cacheKeys)
		// 不 return，尝试延迟双删
	}

	// 延迟双删
	time.AfterFunc(300*time.Millisecond, func() {
		if _, err := cache.Remove(ctx, cacheKeys...); err != nil {
			g.Log().Warningf(ctx, "延迟双删缓存失败: %v, keys=%v", err, cacheKeys)
		}
	})

	return nil
}
