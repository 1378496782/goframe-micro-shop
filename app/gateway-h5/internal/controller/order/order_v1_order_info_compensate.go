package order

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
)

func (c *ControllerV1) OrderInfoCompensate(ctx context.Context, req *v1.OrderInfoCompensateReq) (res *v1.OrderInfoCompensateRes, err error) {
	grpcRes, err := c.OrderInfoClient.Compensate(ctx, &order_info.OrderInfoCompensateReq{
		Limit: req.Limit,
	})
	if err != nil {
		return nil, err
	}

	return &v1.OrderInfoCompensateRes{
		Message: grpcRes.Message,
	}, nil
}
