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
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
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
	if req.GoodsId == 0 || req.Count == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "商品ID和商品数量不能为空")
	}

	// 错误类型
	infoError := consts.InfoError(consts.CartInfo, consts.CreateFail)

	record, err := dao.CartInfo.Ctx(ctx).
		Where("user_id", req.UserId).
		Where("goods_id", req.GoodsId).
		One()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	if !record.IsEmpty() {
		var existingCart entity.CartInfo
		if err := record.Struct(&existingCart); err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, gerror.WrapCode(gcode.CodeInternalError, err, "购物车数据转换失败")
		}

		newCount := existingCart.Count + int(req.Count)
		if _, err := dao.CartInfo.Ctx(ctx).
			Where("id", existingCart.Id).
			Data(g.Map{"count": newCount}).
			Update(); err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}

		return &v1.CartInfoCreateRes{Id: uint32(existingCart.Id)}, nil
	}

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
	// 根据ID和用户ID从数据库中删除对应信息
	result, err := dao.CartInfo.Ctx(ctx).Where("id", req.Id).Where("user_id", req.UserId).Delete()
	infoError := consts.InfoError(consts.CartInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to get rows affected: %v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "删除失败")
	}
	if rowsAffected == 0 {
		return nil, gerror.NewCode(gcode.CodeNotFound, "购物车中没有该商品或无权删除")
	}

	// 返回删除成功的空响应
	return &v1.CartInfoDeleteRes{}, nil
}
