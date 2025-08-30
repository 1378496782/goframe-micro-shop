package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	goods_images "shop-goframe-micro-service-refacotor/app/goods/api/goods_images/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
)

func (c *ControllerV1) GoodsImagesGetList(ctx context.Context, req *v1.GoodsImagesGetListReq) (res *v1.GoodsImagesGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &goods_images.GoodsImagesGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.GoodsImagesClient.GetList(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.GoodsImagesGetListRes{
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
