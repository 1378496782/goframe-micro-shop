package coupon

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/user/api/coupon/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedCouponServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCouponServiceServer(s.Server, &Controller{})
}

func (*Controller) CreateUserCoupon(ctx context.Context, req *v1.CouponCreateReq) (res *v1.CouponCreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
