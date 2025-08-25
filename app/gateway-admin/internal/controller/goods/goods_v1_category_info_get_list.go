package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	category_info "shop-goframe-micro-service-refacotor/app/goods/api/category_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) CategoryInfoGetList(ctx context.Context, req *v1.CategoryInfoGetListReq) (res *v1.CategoryInfoGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &category_info.CategoryInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.CategoryInfoClient.GetList(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.CategoryInfoGetListRes{
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
