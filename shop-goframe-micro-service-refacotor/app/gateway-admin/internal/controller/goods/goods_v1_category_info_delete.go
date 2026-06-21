package goods

import (
	"context"
	category_info "shop-goframe-micro-service-refacotor/app/goods/api/category_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) CategoryInfoDelete(ctx context.Context, req *v1.CategoryInfoDeleteReq) (res *v1.CategoryInfoDeleteRes, err error) {
	// 调用gRPC服务
	_, err = c.CategoryInfoClient.Delete(ctx, &category_info.CategoryInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &v1.CategoryInfoDeleteRes{}, nil
}
