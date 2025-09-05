package banner

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	position_info "shop-goframe-micro-service-refacotor/app/banner/api/position_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/banner/v1"
)

func (c *ControllerV1) PositionInfoUpdate(ctx context.Context, req *v1.PositionInfoUpdateReq) (res *v1.PositionInfoUpdateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &position_info.PositionInfoUpdateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.PositionInfoClient.Update(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &v1.PositionInfoUpdateRes{
		Id: grpcRes.Id,
	}, nil
}
