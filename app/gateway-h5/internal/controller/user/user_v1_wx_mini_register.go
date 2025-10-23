package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	user_info "shop-goframe-micro-service-refacotor/app/user/api/user_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) WxMiniRegister(ctx context.Context, req *v1.WxMiniRegisterReq) (res *v1.WxMiniRegisterRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &user_info.WxMiniRegisterReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC登录服务
	grpcRes, err := c.UserInfoClient.WxMiniRegister(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 使用gconv转换响应
	res = &v1.WxMiniRegisterRes{}
	if err := gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}
	return res, nil
}
