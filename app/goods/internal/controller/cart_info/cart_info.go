package cart_info

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/cart_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/logic/cart_info"
	"shop-goframe-micro-service-refacotor/utility/consts"
)

type Controller struct {
	v1.UnimplementedCartInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCartInfoServer(s.Server, &Controller{})
}

func (c *Controller) GetList(ctx context.Context, req *v1.CartInfoGetListReq) (res *v1.CartInfoGetListRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.CartInfo, consts.GetListFail)
	// 调用逻辑层方法
	response, err := cart_info.GetList(ctx, req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.CartInfoGetListRes{Data: response}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.CartInfoCreateReq) (res *v1.CartInfoCreateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.CartInfo, consts.CreateFail)
	// 向数据库中插入数据并获取自动生成的ID
	result, err := dao.CartInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回创建成功响应，包含新创建的ID
	return &v1.CartInfoCreateRes{Id: uint32(result)}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.CartInfoDeleteReq) (res *v1.CartInfoDeleteRes, err error) {
	// 根据ID从数据库中删除对应信息
	_, err = dao.CartInfo.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.CartInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回删除成功的空响应
	return &v1.CartInfoDeleteRes{}, nil // 返回空结构体
}
