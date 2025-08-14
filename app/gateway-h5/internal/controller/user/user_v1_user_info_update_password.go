package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	user_info "shop-goframe-micro-service-refacotor/app/user/api/user_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) UserInfoUpdatePassword(ctx context.Context, req *v1.UserInfoUpdatePasswordReq) (res *v1.UserInfoUpdatePasswordRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &user_info.UserInfoUpdatePasswordReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.UserInfoClient.UpdatePassword(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoUpdatePasswordRes{Id: grpcRes.Id}, nil
}
