package utility

import (
	"context"
	"sync"
)

// IFlashSaleService 秒杀服务接口（定义在utility包中避免循环依赖）
type IFlashSaleService interface {
	// 获取秒杀商品列表
	GetFlashSaleGoodsList(ctx context.Context, req interface{}) (interface{}, error)

	// 获取秒杀商品详情
	GetFlashSaleGoodsDetail(ctx context.Context, req interface{}) (interface{}, error)

	// 创建秒杀订单
	CreateFlashSaleOrder(ctx context.Context, req interface{}) (interface{}, error)

	// 查询秒杀结果
	GetFlashSaleResult(ctx context.Context, req interface{}) (interface{}, error)
}

var (
	flashSaleService IFlashSaleService
	once             sync.Once
)

// RegisterFlashSale 注册秒杀服务
func RegisterFlashSale(svc IFlashSaleService) {
	flashSaleService = svc
}

// GetFlashSaleService 获取秒杀服务
func GetFlashSaleService() IFlashSaleService {
	return flashSaleService
}
