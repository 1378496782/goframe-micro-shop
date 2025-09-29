package goods

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
	cart_info "shop-goframe-micro-service-refacotor/app/goods/api/cart_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"
)

func (c *ControllerV1) CartInfoDelete(ctx context.Context, req *v1.CartInfoDeleteReq) (res *v1.CartInfoDeleteRes, err error) {
	// 1. Get userId from context
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "无法获取用户信息，请重新登录")
	}

	// 2. Call gRPC service
	grpcReq := &cart_info.CartInfoDeleteReq{
		Id:     req.Id,
		UserId: userId,
	}
	_, err = c.CartInfoClient.Delete(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 3. Return empty response
	return &v1.CartInfoDeleteRes{}, nil
}