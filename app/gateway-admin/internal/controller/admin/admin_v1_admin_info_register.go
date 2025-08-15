package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	admin_info "shop-goframe-micro-service-refacotor/app/admin/api/admin_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/admin/v1"
)

func (c *ControllerV1) AdminInfoRegister(ctx context.Context, req *v1.AdminInfoRegisterReq) (res *v1.AdminInfoRegisterRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &admin_info.AdminInfoRegisterReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC登录服务
	grpcRes, err := c.AdminInfoClient.Register(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	return &v1.AdminInfoRegisterRes{Id: grpcRes.Id}, nil
}
