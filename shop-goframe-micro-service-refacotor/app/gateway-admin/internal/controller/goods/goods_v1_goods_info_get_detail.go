package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/goods/v1"
	goods_info "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
)

func (c *ControllerV1) GoodsInfoGetDetail(ctx context.Context, req *v1.GoodsInfoGetDetailReq) (res *v1.GoodsInfoGetDetailRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &goods_info.GoodsInfoGetDetailReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.GoodsInfoClient.GetDetail(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	// 批量转换列表项
	res = &v1.GoodsInfoGetDetailRes{}
	if err := gconv.Struct(grpcRes.Data, res); err != nil {
		return nil, err
	}

	return res, nil
}
