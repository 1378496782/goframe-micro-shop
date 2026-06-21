package order

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	refund_info "shop-goframe-micro-service-refacotor/app/order/api/refund_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
)

func (c *ControllerV1) RefundInfoCreate(ctx context.Context, req *v1.RefundInfoCreateReq) (res *v1.RefundInfoCreateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &refund_info.RefundInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		// 处理类型不匹配的情况
		panic("用户ID类型错误或不存在")
	}
	grpcReq.UserId = userId
	fmt.Println(grpcReq.UserId)
	// 调用gRPC服务
	grpcRes, err := c.RefundInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.RefundInfoCreateRes{Id: grpcRes.Id}, nil
}
