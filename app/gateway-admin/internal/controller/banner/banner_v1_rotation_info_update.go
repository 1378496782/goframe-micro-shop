package banner

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	rotation_info "shop-goframe-micro-service-refacotor/app/banner/api/rotation_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/banner/v1"
)

func (c *ControllerV1) RotationInfoUpdate(ctx context.Context, req *v1.RotationInfoUpdateReq) (res *v1.RotationInfoUpdateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &rotation_info.RotationInfoUpdateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.RotationInfoClient.Update(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &v1.RotationInfoUpdateRes{
		Id: grpcRes.Id,
	}, nil
}
