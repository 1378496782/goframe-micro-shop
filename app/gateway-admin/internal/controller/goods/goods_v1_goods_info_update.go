package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	goods_info "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) GoodsInfoUpdate(ctx context.Context, req *v1.GoodsInfoUpdateReq) (res *v1.GoodsInfoUpdateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &goods_info.GoodsInfoUpdateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.GoodsInfoClient.Update(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &v1.GoodsInfoUpdateRes{
		Id: grpcRes.Id,
	}, nil
}
