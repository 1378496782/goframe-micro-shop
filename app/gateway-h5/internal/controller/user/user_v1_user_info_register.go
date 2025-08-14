package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	user_info "shop-goframe-micro-service-refacotor/app/user/api/user_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) UserInfoRegister(ctx context.Context, req *v1.UserInfoRegisterReq) (res *v1.UserInfoRegisterRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &user_info.UserInfoRegisterReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC登录服务
	grpcRes, err := c.UserInfoClient.Register(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	return &v1.UserInfoRegisterRes{Id: grpcRes.Id}, nil
}
