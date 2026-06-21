package logic

import (
	"context"
	"fmt"

	v1 "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"shop-goframe-micro-service-refacotor/app/flash-sale/utility"
)

func init() {
	// 注册服务
	utility.RegisterFlashSale(New())
}

// FlashSaleService 秒杀服务实现
type FlashSaleService struct{}

// New 创建秒杀服务实例（实现utility.IFlashSaleService接口）
func New() utility.IFlashSaleService {
	return &FlashSaleService{}
}

// GetFlashSaleGoodsList 获取秒杀商品列表
func (s *FlashSaleService) GetFlashSaleGoodsList(ctx context.Context, req interface{}) (interface{}, error) {
	request, ok := req.(*v1.FlashSaleGoodsListReq)
	if !ok {
		return nil, fmt.Errorf("invalid request type: %T", req)
	}
	return NewFlashSale().GetFlashSaleGoodsList(ctx, request)
}

// GetFlashSaleGoodsDetail 获取秒杀商品详情
func (s *FlashSaleService) GetFlashSaleGoodsDetail(ctx context.Context, req interface{}) (interface{}, error) {
	request, ok := req.(*v1.FlashSaleGoodsDetailReq)
	if !ok {
		return nil, fmt.Errorf("invalid request type: %T", req)
	}
	return NewFlashSale().GetFlashSaleGoodsDetail(ctx, request)
}

// CreateFlashSaleOrder 创建秒杀订单
func (s *FlashSaleService) CreateFlashSaleOrder(ctx context.Context, req interface{}) (interface{}, error) {
	request, ok := req.(*v1.CreateFlashSaleOrderReq)
	if !ok {
		return nil, fmt.Errorf("invalid request type: %T", req)
	}
	return NewFlashSale().CreateFlashSaleOrder(ctx, request)
}

// GetFlashSaleResult 查询秒杀结果
func (s *FlashSaleService) GetFlashSaleResult(ctx context.Context, req interface{}) (interface{}, error) {
	request, ok := req.(*v1.GetFlashSaleResultReq)
	if !ok {
		return nil, fmt.Errorf("invalid request type: %T", req)
	}
	return NewFlashSale().GetFlashSaleResult(ctx, request)
}
