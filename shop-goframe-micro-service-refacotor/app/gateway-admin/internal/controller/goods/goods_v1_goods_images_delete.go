package goods

import (
	"context"
	goods_images "shop-goframe-micro-service-refacotor/app/goods/api/goods_images/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) GoodsImagesDelete(ctx context.Context, req *v1.GoodsImagesDeleteReq) (res *v1.GoodsImagesDeleteRes, err error) {
	// 调用gRPC服务
	_, err = c.GoodsImagesClient.Delete(ctx, &goods_images.GoodsImagesDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &v1.GoodsImagesDeleteRes{}, nil
}
