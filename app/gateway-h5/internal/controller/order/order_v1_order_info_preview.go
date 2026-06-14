package order

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) OrderInfoPreview(ctx context.Context, req *v1.OrderInfoPreviewReq) (res *v1.OrderInfoPreviewRes, err error) {
	// 通过 token 获取用户的id
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		// 处理类型不匹配的情况
		return nil, gerror.New("用户ID类型错误或不存在")
	}

	// 使用 gconv 自动转换结构体
	grpcReq := &order_info.OrderInfoPreviewReq{
		UserId: userId,
	}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	// 调用gRPC服务
	grpcRes, err := c.OrderInfoClient.Preview(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	items := []*v1.OrderInfoPreviewItem{}
	if err := gconv.Structs(grpcRes.Items, &items); err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.OrderInfoPreviewRes{
		Items:      items,
		TotalPrice: grpcRes.TotalPrice,
		TotalCount: grpcRes.TotalCount,
	}
	return res, nil
}
