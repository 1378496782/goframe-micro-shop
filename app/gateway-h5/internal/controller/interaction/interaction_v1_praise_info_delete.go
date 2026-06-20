package interaction

import (
	"context"
	"fmt"
	praise "shop-goframe-micro-service-refacotor/app/interaction/api/praise_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) PraiseInfoDelete(ctx context.Context, req *v1.PraiseInfoDeleteReq) (res *v1.PraiseInfoDeleteRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &praise.PraiseInfoDeleteReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	Value := ctx.Value(middleware.CtxUserId)
	userId, ok := Value.(uint32)
	if !ok {
		return nil, fmt.Errorf("用户ID类型错误或不存在")
	}
	grpcReq.UserId = userId
	// 调用gRPC服务
	_, err = c.PraiseInfoClient.Delete(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.PraiseInfoDeleteRes{Id: req.Id}, nil
}
