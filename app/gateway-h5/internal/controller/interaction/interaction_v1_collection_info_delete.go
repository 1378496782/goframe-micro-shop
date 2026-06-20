package interaction

import (
	"context"
	collection "shop-goframe-micro-service-refacotor/app/interaction/api/collection_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) CollectionInfoDelete(ctx context.Context, req *v1.CollectionInfoDeleteReq) (res *v1.CollectionInfoDeleteRes, err error) {

	// 从token获取用户的id
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "用户ID类型错误或不存在")
	}
	// 调用gRPC服务
	collectionInfoDeleteRes, err := c.CollectionInfoClient.Delete(ctx, &collection.CollectionInfoDeleteReq{Id: req.Id, UserId: userId})
	if err != nil {
		return nil, err
	}

	return &v1.CollectionInfoDeleteRes{Id: collectionInfoDeleteRes.Id}, nil
}
