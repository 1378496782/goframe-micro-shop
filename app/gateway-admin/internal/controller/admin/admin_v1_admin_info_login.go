package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	admin_info "shop-goframe-micro-service-refacotor/app/admin/api/admin_info/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/admin/v1"
)

func (c *ControllerV1) AdminInfoLogin(ctx context.Context, req *v1.AdminInfoLoginReq) (res *v1.AdminInfoLoginRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &admin_info.AdminInfoLoginReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC登录服务
	grpcRes, err := c.AdminInfoClient.Login(ctx, grpcReq)

	if err != nil {
		// 这里可以根据gRPC返回的错误码转换成本地错误码
		// 例如，如果gRPC返回的是用户不存在，可以转换为CodeNotFound
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "登录失败")
	}

	// 使用gconv转换响应
	res = &v1.AdminInfoLoginRes{}
	if err := gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}

	return res, nil
}
