package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	category_info "shop-goframe-micro-service-refacotor/app/goods/api/category_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
)

func (c *ControllerV1) CategoryInfoGetAll(ctx context.Context, req *v1.CategoryInfoGetAllReq) (res *v1.CategoryInfoGetAllRes, err error) {
	// 调用gRPC服务（不需要传递任何参数）
	grpcRes, err := c.CategoryInfoClient.GetAll(ctx, &category_info.CategoryInfoGetAllReq{})
	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.CategoryInfoGetAllRes{
		Total: grpcRes.Total,
		List:  make([]*v1.CategoryInfoItem, len(grpcRes.List)),
	}
	// 批量转换列表项
	if err := gconv.Structs(grpcRes.List, &res.List); err != nil {
		return nil, err
	}

	return res, nil
}
