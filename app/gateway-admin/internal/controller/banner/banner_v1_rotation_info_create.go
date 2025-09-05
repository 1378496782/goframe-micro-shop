package banner

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	rotation_info "shop-goframe-micro-service-refacotor/app/banner/api/rotation_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/banner/v1"
)

func (c *ControllerV1) RotationInfoCreate(ctx context.Context, req *v1.RotationInfoCreateReq) (res *v1.RotationInfoCreateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &rotation_info.RotationInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	// 调用gRPC服务
	grpcRes, err := c.RotationInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.RotationInfoCreateRes{Id: grpcRes.Id}, nil
}
