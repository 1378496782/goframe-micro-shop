package order

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
)

func (c *ControllerV1) CancelTimeoutPendingOrders(ctx context.Context, req *v1.CancelTimeoutPendingOrdersReq) (res *v1.CancelTimeoutPendingOrdersRes, err error) {
	grpcRes, err := c.OrderInfoClient.CancelTimeout(ctx, &order_info.CancelTimeoutPendingOrdersReq{
		TimeoutMinutes: req.TimeoutMinutes,
		Limit:          req.Limit,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CancelTimeoutPendingOrdersRes{
		Message:     grpcRes.Message,
		CancelCount: grpcRes.CancelCount,
	}, nil
}
