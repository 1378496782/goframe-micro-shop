package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	user_info "shop-goframe-micro-service-refacotor/app/user/api/user_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) FillPhone(ctx context.Context, req *v1.FillPhoneReq) (res *v1.FillPhoneRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &user_info.FillPhoneReq{}
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
	// 调用gRPC服务
	gerpRes, err := c.UserInfoClient.FillPhone(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.FillPhoneRes{gerpRes.Id}, nil
}
