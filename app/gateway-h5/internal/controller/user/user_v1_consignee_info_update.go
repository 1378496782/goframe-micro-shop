package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	consignee "shop-goframe-micro-service-refacotor/app/user/api/consignee_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) ConsigneeInfoUpdate(ctx context.Context, req *v1.ConsigneeInfoUpdateReq) (res *v1.ConsigneeInfoUpdateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &consignee.ConsigneeInfoUpdateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.ConsigneeInfoClient.Update(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &v1.ConsigneeInfoUpdateRes{
		Id: grpcRes.Id,
	}, nil
}
