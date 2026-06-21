package stock

import (
	"context"
)

// StockManager 库存管理器接口
type StockManager interface {
	// ReduceStock 扣减库存
	// goodsId: 商品ID
	// count: 扣减数量
	// 返回值: 是否成功，错误信息
	ReduceStock(ctx context.Context, goodsId uint32, count int) (bool, error)

	// ReturnStock 返还库存
	// goodsId: 商品ID
	// count: 返还数量
	// 返回值: 是否成功，错误信息
	ReturnStock(ctx context.Context, goodsId uint32, count int) (bool, error)

	// GetStock 获取当前库存
	// goodsId: 商品ID
	// 返回值: 库存数量，错误信息
	GetStock(ctx context.Context, goodsId uint32) (int, error)

	// InitStock 初始化库存
	// goodsId: 商品ID
	// count: 初始库存数量
	// 返回值: 是否成功，错误信息
	InitStock(ctx context.Context, goodsId uint32, count int) (bool, error)
}
