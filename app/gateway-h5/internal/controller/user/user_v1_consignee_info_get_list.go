package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	consignee "shop-goframe-micro-service-refacotor/app/user/api/consignee_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) ConsigneeInfoGetList(ctx context.Context, req *v1.ConsigneeInfoGetListReq) (res *v1.ConsigneeInfoGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &consignee.ConsigneeInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.ConsigneeInfoClient.GetList(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.ConsigneeInfoGetListRes{
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
