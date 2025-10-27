//package api
//
//import (
//	"context"
//	v1 "shop-goframe-micro-service-refacotor/app/gateway-resource/api/file/v1"
//
//	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
//	"github.com/gogf/gf/v2/errors/gcode"
//	"github.com/gogf/gf/v2/errors/gerror"
//)
//
//type Controller struct {
//	v1.UnimplementedFileServer
//}
//
//func Register(s *grpcx.GrpcServer) {
//	v1.RegisterFileServer(s.Server, &Controller{})
//}
//
//func (*Controller) GetSignUrl(ctx context.Context, req *v1.GetSignUrlReq) (res *v1.GetSignUrlRes, err error) {
//	return nil, gerror.NewCode(gcode.CodeNotImplemented)
//}
