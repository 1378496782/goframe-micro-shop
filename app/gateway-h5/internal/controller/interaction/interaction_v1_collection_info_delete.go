package interaction

import (
	"context"
	collection "shop-goframe-micro-service-refacotor/app/interaction/api/collection_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"
)

func (c *ControllerV1) CollectionInfoDelete(ctx context.Context, req *v1.CollectionInfoDeleteReq) (res *v1.CollectionInfoDeleteRes, err error) {
	// 调用gRPC服务
	_, err = c.CollectionInfoClient.Delete(ctx, &collection.CollectionInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &v1.CollectionInfoDeleteRes{}, nil
}
