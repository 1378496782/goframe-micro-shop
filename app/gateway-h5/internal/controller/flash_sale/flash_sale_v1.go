package flash_sale

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	flash_sale "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/flash_sale/v1"
)

func (c *ControllerV1) FlashSaleGoodsList(ctx context.Context, req *v1.FlashSaleGoodsListReq) (res *v1.FlashSaleGoodsListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &flash_sale.FlashSaleGoodsListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.FlashSaleClient.GetFlashSaleGoodsList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.FlashSaleGoodsListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	// 批量转换列表项
	if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ControllerV1) FlashSaleGoodsDetail(ctx context.Context, req *v1.FlashSaleGoodsDetailReq) (res *v1.FlashSaleGoodsDetailRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &flash_sale.FlashSaleGoodsDetailReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.FlashSaleClient.GetFlashSaleGoodsDetail(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.FlashSaleGoodsDetailRes{}
	if err := gconv.Struct(grpcRes.Data, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ControllerV1) CreateFlashSaleOrder(ctx context.Context, req *v1.CreateFlashSaleOrderReq) (res *v1.CreateFlashSaleOrderRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &flash_sale.CreateFlashSaleOrderReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.FlashSaleClient.CreateFlashSaleOrder(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.CreateFlashSaleOrderRes{}
	if err := gconv.Struct(grpcRes.Data, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ControllerV1) GetFlashSaleResult(ctx context.Context, req *v1.GetFlashSaleResultReq) (res *v1.GetFlashSaleResultRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &flash_sale.GetFlashSaleResultReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.FlashSaleClient.GetFlashSaleResult(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.GetFlashSaleResultRes{}
	if err := gconv.Struct(grpcRes.Data, res); err != nil {
		return nil, err
	}

	return res, nil
}