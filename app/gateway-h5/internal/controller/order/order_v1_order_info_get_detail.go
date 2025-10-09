package order

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"
)

func (c *ControllerV1) OrderInfoGetDetail(ctx context.Context, req *v1.OrderInfoGetDetailReq) (res *v1.OrderInfoGetDetailRes, err error) {
	// 1. Get userId from context
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "无法获取用户信息，请重新登录")
	}

	// 2. Call gRPC service
	grpcReq := &order_info.OrderInfoGetDetailReq{
		Id:     req.Id,
		UserId: userId,
	}
	detailRes, err := c.OrderInfoClient.GetDetail(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 3. Assemble response
	res = &v1.OrderInfoGetDetailRes{}
	if err := gconv.Struct(detailRes.OrderInfo, &res.OrderInfo); err != nil {
		return nil, err
	}
	if err := gconv.Structs(detailRes.OrderGoodsInfos, &res.OrderGoodsInfos); err != nil {
		return nil, err
	}

	return res, nil
}
