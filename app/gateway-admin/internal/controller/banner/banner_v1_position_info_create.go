package banner

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	position_info "shop-goframe-micro-service-refacotor/app/banner/api/position_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/banner/v1"
)

func (c *ControllerV1) PositionInfoCreate(ctx context.Context, req *v1.PositionInfoCreateReq) (res *v1.PositionInfoCreateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &position_info.PositionInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	// 调用gRPC服务
	grpcRes, err := c.PositionInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.PositionInfoCreateRes{Id: grpcRes.Id}, nil
}
