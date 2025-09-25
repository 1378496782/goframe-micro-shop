package order

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
)

func (c *ControllerV1) Payment(ctx context.Context, req *v1.PaymentReq) (res *v1.PaymentRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &order_info.PaymentReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.OrderInfoClient.Payment(ctx, grpcReq)
	if err != nil {
		g.Log().Warningf(ctx, "支付失败, err:%v", err.Error())
		return nil, err
	}

	// 返回响应
	res = &v1.PaymentRes{}
	if err := gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}

	g.Log().Info(ctx, "支付成功")
	return res, nil
}
