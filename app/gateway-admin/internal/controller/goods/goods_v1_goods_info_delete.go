package goods

import (
	"context"
	goods_info "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) GoodsInfoDelete(ctx context.Context, req *v1.GoodsInfoDeleteReq) (res *v1.GoodsInfoDeleteRes, err error) {
	// 调用gRPC服务
	_, err = c.GoodsInfoClient.Delete(ctx, &goods_info.GoodsInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &v1.GoodsInfoDeleteRes{}, nil
}
