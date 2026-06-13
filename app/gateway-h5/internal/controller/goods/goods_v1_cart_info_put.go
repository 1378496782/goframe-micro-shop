package goods

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
	cart_info "shop-goframe-micro-service-refacotor/app/goods/api/cart_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"
)

func (c *ControllerV1) CartInfoPut(ctx context.Context, req *v1.CartInfoPutReq) (res *v1.CartInfoPutRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &cart_info.CartInfoPutReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		// 处理类型不匹配的情况
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "无法获取用户信息，请重新登录")
	}
	grpcReq.UserId = userId
	// 调用gRPC服务
	_, err = c.CartInfoClient.Put(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.CartInfoPutRes{}, nil
}
