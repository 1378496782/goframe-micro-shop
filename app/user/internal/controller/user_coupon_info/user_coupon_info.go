package user_coupon_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/user/api/user_coupon_info/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedUserCouponInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserCouponInfoServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.UserCouponInfoCreateReq) (res *v1.UserCouponInfoCreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
