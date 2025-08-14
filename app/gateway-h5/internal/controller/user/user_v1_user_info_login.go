package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	user_info "shop-goframe-micro-service-refacotor/app/user/api/user_info/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) UserInfoLogin(ctx context.Context, req *v1.UserInfoLoginReq) (res *v1.UserInfoLoginRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &user_info.UserInfoLoginReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC登录服务
	grpcRes, err := c.UserInfoClient.Login(ctx, grpcReq)

	if err != nil {
		// 这里可以根据gRPC返回的错误码转换成本地错误码
		// 例如，如果gRPC返回的是用户不存在，可以转换为CodeNotFound
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "登录失败")
	}

	// 使用gconv转换响应
	res = &v1.UserInfoLoginRes{}
	if err := gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}

	return res, nil
}
