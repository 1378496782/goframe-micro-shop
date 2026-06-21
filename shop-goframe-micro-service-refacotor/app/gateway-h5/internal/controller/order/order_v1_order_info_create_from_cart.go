package order

import (
	"context"

	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
)

func (c *ControllerV1) OrderInfoCreateFromCart(ctx context.Context, req *v1.OrderInfoCreateFromCartReq) (res *v1.OrderInfoCreateFromCartRes, err error) {
	if len(req.CartIds) == 0 {
		return nil, gerror.New("订单必须包含购物车ID列表")
	}

	grpcReq := &order_info.OrderInfoCreateFromCartReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "无法获取用户信息，请重新登录")
	}
	grpcReq.UserId = userId

	grpcRes, err := c.OrderInfoClient.CreateFromCart(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.OrderInfoCreateFromCartRes{}
	if err := gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}

	return res, nil
}
