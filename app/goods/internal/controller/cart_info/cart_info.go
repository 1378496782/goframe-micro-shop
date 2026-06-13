package cart_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/cart_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/logic/cart_info"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"strings"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
	if req.Count == 0 || req.UserId == 0 || req.GoodsId == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "参数错误")
	}

	// 错误类型
	infoError := consts.InfoError(consts.CartInfo, consts.CreateFail)

	// 加购前校验商品是否存在
	goodsRecord, err := dao.GoodsInfo.Ctx(ctx).Where(dao.GoodsInfo.Columns().Id, req.GoodsId).One()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	if goodsRecord.IsEmpty() {
		return nil, gerror.NewCode(gcode.CodeNotFound, "商品不存在")
	}
	var goodsInfo entity.GoodsInfo
	if err := goodsRecord.Struct(&goodsInfo); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "数据转换失败")
	}

	// 	先根据 user_id + goods_id 查 cart_info
	cartRecord, err := dao.CartInfo.Ctx(ctx).Where(g.Map{
		dao.CartInfo.Columns().UserId:  req.UserId,
		dao.CartInfo.Columns().GoodsId: req.GoodsId,
	}).One()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 如果不存在：正常插入一条新购物车记录
	// Insert 成功 -> 返回新 id
	// Insert 失败，并且是 Duplicate entry -> 说明并发请求已经插入了，重新查购物车，继续走更新逻辑
	// Insert 失败，但不是 Duplicate entry -> 真正的数据库错误
	if cartRecord.IsEmpty() {
		// 先判断库存是否足够
		if int(req.Count) > int(goodsInfo.Stock) {
			return nil, gerror.NewCode(gcode.CodeInvalidParameter, "商品库存不足")
		}
		id, err := dao.CartInfo.Ctx(ctx).InsertAndGetId(req)
		if err == nil {
			return &v1.CartInfoCreateRes{Id: uint32(id)}, nil
		}
		if !isDuplicateKeyError(err) {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}

		// 走到这里，说明另一个并发请求已经插入了同一条购物车记录。
		// 重新查询，然后继续走下面“已存在购物车记录”的更新逻辑。
		cartRecord, err = dao.CartInfo.Ctx(ctx).Where(g.Map{
			dao.CartInfo.Columns().UserId:  req.UserId,
			dao.CartInfo.Columns().GoodsId: req.GoodsId,
		}).One()
		if err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}
		if cartRecord.IsEmpty() {
			return nil, gerror.NewCode(gcode.CodeInternalError, "购物车记录创建冲突，请重试")
		}
	}

	// 如果已存在：更新 count = old_count + req.Count
	var existingCartInfo entity.CartInfo
	err = cartRecord.Struct(&existingCartInfo)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	newCount := existingCartInfo.Count + int(req.Count)
	// 先校验库存是否足够
	// TODO: 这里可能存在并发问题, 需要添加锁
	if newCount > int(goodsInfo.Stock) {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "商品库存不足")
	}
	_, err = dao.CartInfo.Ctx(ctx).Where(dao.CartInfo.Columns().Id, existingCartInfo.Id).
		Update(g.Map{
			dao.CartInfo.Columns().Count: newCount,
		})
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.CartInfoCreateRes{Id: uint32(existingCartInfo.Id)}, nil
}

func isDuplicateKeyError(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	return strings.Contains(errMsg, "Duplicate entry") || strings.Contains(errMsg, "Error 1062")
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
