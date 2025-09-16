package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	user_info "shop-goframe-micro-service-refacotor/app/user/api/user_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) UserInfoUpdate(ctx context.Context, req *v1.UserInfoUpdateReq) (res *v1.UserInfoUpdateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &user_info.UserInfoUpdateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		// 处理类型不匹配的情况
		panic("用户ID类型错误或不存在")
	}
	grpcReq.Id = userId

	grpcRes, err := c.UserInfoClient.UpdateInfo(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoUpdateRes{Id: grpcRes.Id}, nil
}
