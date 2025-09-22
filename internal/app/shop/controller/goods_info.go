// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-09-22 16:30:35
// 生成路径: internal/app/shop/controller/goods_info.go
// 生成人：gfast
// desc:商品
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/shop"
	"github.com/tiger1103/gfast/v3/internal/app/shop/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

type goodsInfoController struct {
	systemController.BaseController
}

var GoodsInfo = new(goodsInfoController)

// List 列表
func (c *goodsInfoController) List(ctx context.Context, req *shop.GoodsInfoSearchReq) (res *shop.GoodsInfoSearchRes, err error) {
	res = new(shop.GoodsInfoSearchRes)
	res.GoodsInfoSearchRes, err = service.GoodsInfo().List(ctx, &req.GoodsInfoSearchReq)
	return
}

// Get 获取商品
func (c *goodsInfoController) Get(ctx context.Context, req *shop.GoodsInfoGetReq) (res *shop.GoodsInfoGetRes, err error) {
	res = new(shop.GoodsInfoGetRes)
	res.GoodsInfoInfoRes, err = service.GoodsInfo().GetById(ctx, req.Id)
	return
}

// Add 添加商品
func (c *goodsInfoController) Add(ctx context.Context, req *shop.GoodsInfoAddReq) (res *shop.GoodsInfoAddRes, err error) {
	err = service.GoodsInfo().Add(ctx, req.GoodsInfoAddReq)
	return
}

// Edit 修改商品
func (c *goodsInfoController) Edit(ctx context.Context, req *shop.GoodsInfoEditReq) (res *shop.GoodsInfoEditRes, err error) {
	err = service.GoodsInfo().Edit(ctx, req.GoodsInfoEditReq)
	return
}

// Delete 删除商品
func (c *goodsInfoController) Delete(ctx context.Context, req *shop.GoodsInfoDeleteReq) (res *shop.GoodsInfoDeleteRes, err error) {
	err = service.GoodsInfo().Delete(ctx, req.Ids)
	return
}
