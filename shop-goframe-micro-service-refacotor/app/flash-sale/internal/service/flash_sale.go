package service

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
)

// IFlashSale 秒杀服务接口
type IFlashSale interface {
	// 获取秒杀商品列表
	GetFlashSaleGoodsList(ctx context.Context, req *v1.FlashSaleGoodsListReq) (*v1.FlashSaleGoodsListRes, error)

	// 获取秒杀商品详情
	GetFlashSaleGoodsDetail(ctx context.Context, req *v1.FlashSaleGoodsDetailReq) (*v1.FlashSaleGoodsDetailRes, error)

	// 创建秒杀订单
	CreateFlashSaleOrder(ctx context.Context, req *v1.CreateFlashSaleOrderReq) (*v1.CreateFlashSaleOrderRes, error)

	// 查询秒杀结果
	GetFlashSaleResult(ctx context.Context, req *v1.GetFlashSaleResultReq) (*v1.GetFlashSaleResultRes, error)

	// 初始化秒杀商品库存
	InitFlashSaleStock(ctx context.Context, goodsId uint32, stock int) error

	// 处理秒杀订单（异步）
	ProcessFlashSaleOrder(ctx context.Context, orderNo string) error
}
