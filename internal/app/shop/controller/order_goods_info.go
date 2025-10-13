// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-10-11 23:55:29
// 生成路径: internal/app/shop/controller/order_goods_info.go
// 生成人：gfast
// desc:订单物品表
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/shop"
	"github.com/tiger1103/gfast/v3/internal/app/shop/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

type orderGoodsInfoController struct {
	systemController.BaseController
}

var OrderGoodsInfo = new(orderGoodsInfoController)

// List 列表
func (c *orderGoodsInfoController) List(ctx context.Context, req *shop.OrderGoodsInfoSearchReq) (res *shop.OrderGoodsInfoSearchRes, err error) {
	res = new(shop.OrderGoodsInfoSearchRes)
	res.OrderGoodsInfoSearchRes, err = service.OrderGoodsInfo().List(ctx, &req.OrderGoodsInfoSearchReq)
	return
}

// Get 获取订单物品表
func (c *orderGoodsInfoController) Get(ctx context.Context, req *shop.OrderGoodsInfoGetReq) (res *shop.OrderGoodsInfoGetRes, err error) {
	res = new(shop.OrderGoodsInfoGetRes)
	res.OrderGoodsInfoInfoRes, err = service.OrderGoodsInfo().GetById(ctx, req.Id)
	return
}

// Add 添加订单物品表
func (c *orderGoodsInfoController) Add(ctx context.Context, req *shop.OrderGoodsInfoAddReq) (res *shop.OrderGoodsInfoAddRes, err error) {
	err = service.OrderGoodsInfo().Add(ctx, req.OrderGoodsInfoAddReq)
	return
}

// Edit 修改订单物品表
func (c *orderGoodsInfoController) Edit(ctx context.Context, req *shop.OrderGoodsInfoEditReq) (res *shop.OrderGoodsInfoEditRes, err error) {
	err = service.OrderGoodsInfo().Edit(ctx, req.OrderGoodsInfoEditReq)
	return
}

// Delete 删除订单物品表
func (c *orderGoodsInfoController) Delete(ctx context.Context, req *shop.OrderGoodsInfoDeleteReq) (res *shop.OrderGoodsInfoDeleteRes, err error) {
	err = service.OrderGoodsInfo().Delete(ctx, req.Ids)
	return
}

// GetByOrderId 根据订单ID获取订单商品列表
func (c *orderGoodsInfoController) GetByOrderId(ctx context.Context, req *shop.OrderGoodsInfoGetByOrderIdReq) (res *shop.OrderGoodsInfoGetByOrderIdRes, err error) {
	res = new(shop.OrderGoodsInfoGetByOrderIdRes)
	res.List, err = service.OrderGoodsInfo().GetByOrderId(ctx, req.OrderId)
	return
}

// GetDetail 获取订单商品详细信息
func (c *orderGoodsInfoController) GetDetail(ctx context.Context, req *shop.OrderGoodsInfoGetDetailReq) (res *shop.OrderGoodsInfoGetDetailRes, err error) {
	res = new(shop.OrderGoodsInfoGetDetailRes)
	res.OrderGoodsDetailRes, err = service.OrderGoodsInfo().GetOrderGoodsDetail(ctx, req.Id)
	return
}

// AddOrderGoods 添加订单商品
func (c *orderGoodsInfoController) AddOrderGoods(ctx context.Context, req *shop.OrderGoodsInfoAddOrderGoodsReq) (res *shop.OrderGoodsInfoAddOrderGoodsRes, err error) {
	err = service.OrderGoodsInfo().AddOrderGoods(ctx, req.OrderGoodsAddReq)
	return
}
