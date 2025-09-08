package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/user/v1"
	consignee "shop-goframe-micro-service-refacotor/app/user/api/consignee_info/v1"
)

func (c *ControllerV1) ConsigneeInfoCreate(ctx context.Context, req *v1.ConsigneeInfoCreateReq) (res *v1.ConsigneeInfoCreateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &consignee.ConsigneeInfoCreateReq{}
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
	// 调用gRPC服务
	grpcRes, err := c.ConsigneeInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.ConsigneeInfoCreateRes{Id: grpcRes.Id}, nil
}
