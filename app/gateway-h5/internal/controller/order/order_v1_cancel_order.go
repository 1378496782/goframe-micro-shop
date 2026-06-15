package order

import (
	"context"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) CancelOrder(ctx context.Context, req *v1.CancelOrderReq) (res *v1.CancelOrderRes, err error) {
	// 从上下文中获取userId
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "无法获取用户信息，请重新登录")
	}

	grpcReq := &order_info.CancelOrderReq{
		Id:     req.Id,
		UserId: userId,
	}
	grpcRes, err := c.OrderInfoClient.CancelOrder(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CancelOrderRes{}

	err = gconv.Struct(grpcRes, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
