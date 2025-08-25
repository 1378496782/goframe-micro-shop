package category_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/category_info/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedCategoryInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCategoryInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.CategoryInfoGetListReq) (res *v1.CategoryInfoGetListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetTree(ctx context.Context, req *v1.CategoryInfoGetTreeReq) (res *v1.CategoryInfoGetTreeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Create(ctx context.Context, req *v1.CategoryInfoCreateReq) (res *v1.CategoryInfoCreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Update(ctx context.Context, req *v1.CategoryInfoUpdateReq) (res *v1.CategoryInfoUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Delete(ctx context.Context, req *v1.CategoryInfoDeleteReq) (res *v1.CategoryInfoDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
