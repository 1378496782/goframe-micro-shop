package order

import (
	"context"
	"fmt"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) CancelOrder(ctx context.Context, req *v1.CancelOrderReq) (res *v1.CancelOrderRes, err error) {
	// 从上下文中获取userId
	userId := ctx.Value(middleware.CtxUserId)
	fmt.Println(userId)

	grpcReq := &order_info.CancelOrderReq{
		Id: req.Id,
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
