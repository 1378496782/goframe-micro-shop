package interaction

import (
	"context"
	praise "shop-goframe-micro-service-refacotor/app/interaction/api/praise_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"
)

func (c *ControllerV1) PraiseInfoCreate(ctx context.Context, req *v1.PraiseInfoCreateReq) (res *v1.PraiseInfoCreateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &praise.PraiseInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "无法获取用户信息，请重新登录")
	}
	grpcReq.UserId = userId
	// 调用gRPC服务
	grpcRes, err := c.PraiseInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.PraiseInfoCreateRes{Id: grpcRes.Id}, nil
}
