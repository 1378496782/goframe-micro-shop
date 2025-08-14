package interaction

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	collection "shop-goframe-micro-service-refacotor/app/interaction/api/collection_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"
)

func (c *ControllerV1) CollectionInfoCreate(ctx context.Context, req *v1.CollectionInfoCreateReq) (res *v1.CollectionInfoCreateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &collection.CollectionInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	// 调用gRPC服务
	grpcRes, err := c.CollectionInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.CollectionInfoCreateRes{Id: grpcRes.Id}, nil
}
