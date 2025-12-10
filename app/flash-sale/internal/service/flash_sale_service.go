package service

import (
	"context"
	"fmt"
	"time"

	v1 "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// FlashSaleService 秒杀服务实现
type FlashSaleService struct{}

// NewFlashSaleService 创建秒杀服务
func NewFlashSaleService() *FlashSaleService {
	return &FlashSaleService{}
}

// GetFlashSaleGoodsList 获取秒杀商品列表
func (s *FlashSaleService) GetFlashSaleGoodsList(ctx context.Context, req interface{}) (interface{}, error) {
	// 类型转换
	flashReq, ok := req.(*v1.FlashSaleGoodsListReq)
	if !ok {
		return nil, fmt.Errorf("invalid request type")
	}

	g.Log().Info(ctx, "获取秒杀商品列表, ActivityId:", flashReq.ActivityId, "PageNum:", flashReq.PageNum, "PageSize:", flashReq.PageSize)
	return &v1.FlashSaleGoodsListRes{
		Total: 0,
		List:  []*v1.FlashSaleGoodsInfo{},
	}, nil
}

// GetFlashSaleGoodsDetail 获取秒杀商品详情
func (s *FlashSaleService) GetFlashSaleGoodsDetail(ctx context.Context, req interface{}) (interface{}, error) {
	// 类型转换
	flashReq, ok := req.(*v1.FlashSaleGoodsDetailReq)
	if !ok {
		return nil, fmt.Errorf("invalid request type")
	}

	g.Log().Info(ctx, "获取秒杀商品详情, GoodsId:", flashReq.GoodsId)
	return &v1.FlashSaleGoodsDetailRes{
		GoodsInfo:     &v1.FlashSaleGoodsInfo{},
		RemainSeconds: 0,
		CanBuy:        false,
	}, nil
}

// CreateFlashSaleOrder 创建秒杀订单
func (s *FlashSaleService) CreateFlashSaleOrder(ctx context.Context, req interface{}) (interface{}, error) {
	// 类型转换
	flashReq, ok := req.(*v1.CreateFlashSaleOrderReq)
	if !ok {
		return nil, fmt.Errorf("invalid request type")
	}

	g.Log().Info(ctx, "创建秒杀订单, UserId:", flashReq.UserId, "GoodsId:", flashReq.GoodsId)
	return &v1.CreateFlashSaleOrderRes{
		Success:  true,
		OrderNo:  fmt.Sprintf("FS%d", time.Now().Unix()),
		Message:  "秒杀订单创建成功",
		ResultId: fmt.Sprintf("RESULT_%d", time.Now().Unix()),
		Status:   1, // 成功
	}, nil
}

// GetFlashSaleResult 查询秒杀结果
func (s *FlashSaleService) GetFlashSaleResult(ctx context.Context, req interface{}) (interface{}, error) {
	// 类型转换
	flashReq, ok := req.(*v1.GetFlashSaleResultReq)
	if !ok {
		return nil, fmt.Errorf("invalid request type")
	}

	g.Log().Info(ctx, "查询秒杀结果, ResultId:", flashReq.ResultId)
	return &v1.GetFlashSaleResultRes{
		Status:    1, // 成功
		Message:   "秒杀成功",
		OrderNo:   fmt.Sprintf("FS%d", time.Now().Unix()),
		GoodsId:   1,
		PayAmount: 100,
	}, nil
}

// ProcessFlashSaleOrder 处理秒杀订单（异步）- 兼容旧接口
func (s *FlashSaleService) ProcessFlashSaleOrder(ctx context.Context, orderMsg *model.FlashSaleOrderMessage) error {
	// TODO: 实现处理秒杀订单逻辑
	g.Log().Info(ctx, "处理秒杀订单, OrderId:", orderMsg.OrderId, "UserId:", orderMsg.UserId, "GoodsId:", orderMsg.GoodsId)

	// 模拟处理逻辑
	time.Sleep(100 * time.Millisecond)

	g.Log().Info(ctx, "秒杀订单处理完成, OrderId:", orderMsg.OrderId)
	return nil
}
