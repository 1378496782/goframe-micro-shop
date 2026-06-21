// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-10-10 23:08:01
// 生成路径: internal/app/shop/controller/order_info.go
// 生成人：gfast
// desc:订单表
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/shop"
	"github.com/tiger1103/gfast/v3/internal/app/shop/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

type orderInfoController struct {
	systemController.BaseController
}

var OrderInfo = new(orderInfoController)

// List 列表
func (c *orderInfoController) List(ctx context.Context, req *shop.OrderInfoSearchReq) (res *shop.OrderInfoSearchRes, err error) {
	res = new(shop.OrderInfoSearchRes)
	res.OrderInfoSearchRes, err = service.OrderInfo().List(ctx, &req.OrderInfoSearchReq)
	return
}

// Get 获取订单表
func (c *orderInfoController) Get(ctx context.Context, req *shop.OrderInfoGetReq) (res *shop.OrderInfoGetRes, err error) {
	res = new(shop.OrderInfoGetRes)
	res.OrderInfoInfoRes, err = service.OrderInfo().GetById(ctx, req.Id)
	return
}

// Add 添加订单表
func (c *orderInfoController) Add(ctx context.Context, req *shop.OrderInfoAddReq) (res *shop.OrderInfoAddRes, err error) {
	err = service.OrderInfo().Add(ctx, req.OrderInfoAddReq)
	return
}

// Edit 修改订单表
func (c *orderInfoController) Edit(ctx context.Context, req *shop.OrderInfoEditReq) (res *shop.OrderInfoEditRes, err error) {
	err = service.OrderInfo().Edit(ctx, req.OrderInfoEditReq)
	return
}

// Delete 删除订单表
func (c *orderInfoController) Delete(ctx context.Context, req *shop.OrderInfoDeleteReq) (res *shop.OrderInfoDeleteRes, err error) {
	err = service.OrderInfo().Delete(ctx, req.Ids)
	return
}

// Ship 订单发货
func (c *orderInfoController) Ship(ctx context.Context, req *shop.OrderInfoShipReq) (res *shop.OrderInfoShipRes, err error) {
	err = service.OrderInfo().Ship(ctx, req.Id)
	return
}

// Refund 订单退款
func (c *orderInfoController) Refund(ctx context.Context, req *shop.OrderInfoRefundReq) (res *shop.OrderInfoRefundRes, err error) {
	err = service.OrderInfo().Refund(ctx, req.Id)
	return
}

// GetProducts 获取订单商品列表
func (c *orderInfoController) GetProducts(ctx context.Context, req *shop.OrderInfoGetProductsReq) (res *shop.OrderInfoGetProductsRes, err error) {
	res = new(shop.OrderInfoGetProductsRes)
	res.List, err = service.OrderInfo().GetOrderProducts(ctx, req.OrderId)
	return
}
