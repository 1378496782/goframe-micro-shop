package user_coupon_info

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	//v1.UnimplementedUserCouponInfoServer
}

func Register(s *grpcx.GrpcServer) {
	//v1.RegisterUserCouponInfoServer(s.Server, &Controller{})
}

//func (*Controller) Create(ctx context.Context, req *v1.UserCouponInfoCreateReq) (res *v1.UserCouponInfoCreateRes, err error) {
//	return nil, gerror.NewCode(gcode.CodeNotImplemented)
//}
