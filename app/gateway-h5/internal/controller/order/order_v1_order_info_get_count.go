package order

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"
)

func (c *ControllerV1) OrderInfoGetCount(ctx context.Context, req *v1.OrderInfoGetCountReq) (res *v1.OrderInfoGetCountRes, err error) {
	// 1. Get userId from context
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "无法获取用户信息，请重新登录")
	}

	// 2. Call gRPC service
	grpcReq := &order_info.OrderInfoGetCountReq{
		UserId: userId,
	}
	grpcRes, err := c.OrderInfoClient.GetCount(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 3. Assemble and return response
	res = &v1.OrderInfoGetCountRes{
		Pending:   grpcRes.Pending,
		Shipping:  grpcRes.Shipping,
		Delivered: grpcRes.Delivered,
		Completed: grpcRes.Completed,
		AfterSale: grpcRes.AfterSale,
	}

	return res, nil
}
