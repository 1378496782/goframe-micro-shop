package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	category_info "shop-goframe-micro-service-refacotor/app/goods/api/category_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) CategoryInfoCreate(ctx context.Context, req *v1.CategoryInfoCreateReq) (res *v1.CategoryInfoCreateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &category_info.CategoryInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	// 调用gRPC服务
	grpcRes, err := c.CategoryInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.CategoryInfoCreateRes{Id: grpcRes.Id}, nil
}
