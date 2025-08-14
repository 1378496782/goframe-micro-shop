package interaction

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	praise "shop-goframe-micro-service-refacotor/app/interaction/api/praise_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"
)

func (c *ControllerV1) PraiseInfoCreate(ctx context.Context, req *v1.PraiseInfoCreateReq) (res *v1.PraiseInfoCreateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &praise.PraiseInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	// 调用gRPC服务
	grpcRes, err := c.PraiseInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.PraiseInfoCreateRes{Id: grpcRes.Id}, nil
}
