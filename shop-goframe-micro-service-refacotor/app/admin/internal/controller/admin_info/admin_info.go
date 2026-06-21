package admin_info

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
	"shop-goframe-micro-service-refacotor/app/admin/api/admin_info/v1"
	"shop-goframe-micro-service-refacotor/app/admin/internal/logic/admin_info"
	"shop-goframe-micro-service-refacotor/utility/consts"
)

type Controller struct {
	v1.UnimplementedAdminInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAdminInfoServer(s.Server, &Controller{})
}

func (c *Controller) Login(ctx context.Context, req *v1.AdminInfoLoginReq) (*v1.AdminInfoLoginRes, error) {
	// 调用logic层
	token, expire, err := admin_info.Login(ctx, req.Name, req.Password)
	// 错误类型
	infoError := consts.InfoError(consts.AdminInfo, consts.LoginFail)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 转换时间格式
	expireProto := timestamppb.New(expire)
	if err := expireProto.CheckValid(); err != nil {
		return nil, err
	}

	return &v1.AdminInfoLoginRes{
		Token:  token,
		Expire: expireProto,
	}, nil
}

// Register 注册接口
func (c *Controller) Register(ctx context.Context, req *v1.AdminInfoRegisterReq) (*v1.AdminInfoRegisterRes, error) {
	// 调用logic层注册
	admin, err := admin_info.Register(ctx, req.Name, req.Password)
	// 错误类型
	infoError := consts.InfoError(consts.AdminInfo, consts.LoginFail)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 转换时间格式
	createdAtProto := timestamppb.New(admin.CreatedAt.Time)
	if err := createdAtProto.CheckValid(); err != nil {
		return nil, err
	}

	// 返回响应
	return &v1.AdminInfoRegisterRes{
		Id:        uint32(admin.Id),
		Name:      admin.Name,
		RoleIds:   admin.RoleIds,
		IsAdmin:   uint32(admin.IsAdmin),
		CreatedAt: createdAtProto,
	}, nil
}
