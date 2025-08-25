package goodsRedis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

const (
	categoryAllKey = "category:all:data"
	GoodsDetailKey = "goods:detail:"
	EmptyValue     = "__EMPTY__"
)

// SetEmptyGoodsDetail 添加设置空缓存的函数，防止缓存穿透
func SetEmptyGoodsDetail(ctx context.Context, productId uint32) error {
	key := fmt.Sprintf("%s%d", GoodsDetailKey, productId)
	// 设置一个短时间的空值，防止缓存穿透
	return goodsCache.Set(ctx, key, EmptyValue, 1*time.Minute)
}

// SetGoodsDetail 设置商品详情缓存
func SetGoodsDetail(ctx context.Context, productId uint32, data interface{}) error {
	key := fmt.Sprintf("%s%d", GoodsDetailKey, productId)

	// 使用JSON序列化确保数据类型一致性
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return goodsCache.Set(ctx, key, jsonData, time.Hour)
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
	key := fmt.Sprintf("%s%d", GoodsDetailKey, productId)
	_, err := goodsCache.Remove(ctx, key)
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
