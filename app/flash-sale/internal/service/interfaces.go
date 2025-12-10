package service

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/model"
)

// IFlashSaleService 秒杀服务接口
type IFlashSaleService interface {
	// 获取秒杀商品列表
	GetFlashSaleGoodsList(ctx context.Context, req *v1.FlashSaleGoodsListReq) (*v1.FlashSaleGoodsListRes, error)

	// 获取秒杀商品详情
	GetFlashSaleGoodsDetail(ctx context.Context, req *v1.FlashSaleGoodsDetailReq) (*v1.FlashSaleGoodsDetailRes, error)

	// 创建秒杀订单
	CreateFlashSaleOrder(ctx context.Context, req *v1.CreateFlashSaleOrderReq) (*v1.CreateFlashSaleOrderRes, error)

	// 查询秒杀结果
	GetFlashSaleResult(ctx context.Context, req *v1.GetFlashSaleResultReq) (*v1.GetFlashSaleResultRes, error)

	// 处理秒杀订单（异步）
	ProcessFlashSaleOrder(ctx context.Context, orderMsg *model.FlashSaleOrderMessage) error
}
