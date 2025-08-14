package interaction

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	praise "shop-goframe-micro-service-refacotor/app/interaction/api/praise_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"
)

func (c *ControllerV1) PraiseInfoGetList(ctx context.Context, req *v1.PraiseInfoGetListReq) (res *v1.PraiseInfoGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &praise.PraiseInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.PraiseInfoClient.GetList(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.PraiseInfoGetListRes{
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
