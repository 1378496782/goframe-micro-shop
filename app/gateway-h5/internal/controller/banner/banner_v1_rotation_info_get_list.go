package banner

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	rotation_info "shop-goframe-micro-service-refacotor/app/banner/api/rotation_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/banner/v1"
)

func (c *ControllerV1) RotationInfoGetList(ctx context.Context, req *v1.RotationInfoGetListReq) (res *v1.RotationInfoGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &rotation_info.RotationInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.RotationInfoClient.GetList(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.RotationInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	// 批量转换列表项
	if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}

	return res, nil
}
