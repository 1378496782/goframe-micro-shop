// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-09-09 17:28:51
// 生成路径: internal/app/shop/controller/rotation_info.go
// 生成人：gfast
// desc:轮播图
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/shop"
	"github.com/tiger1103/gfast/v3/internal/app/shop/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

type rotationInfoController struct {
	systemController.BaseController
}

var RotationInfo = new(rotationInfoController)

// List 列表
func (c *rotationInfoController) List(ctx context.Context, req *shop.RotationInfoSearchReq) (res *shop.RotationInfoSearchRes, err error) {
	res = new(shop.RotationInfoSearchRes)
	res.RotationInfoSearchRes, err = service.RotationInfo().List(ctx, &req.RotationInfoSearchReq)
	return
}

// Get 获取轮播图
func (c *rotationInfoController) Get(ctx context.Context, req *shop.RotationInfoGetReq) (res *shop.RotationInfoGetRes, err error) {
	res = new(shop.RotationInfoGetRes)
	res.RotationInfoInfoRes, err = service.RotationInfo().GetById(ctx, req.Id)
	return
}

// Add 添加轮播图
func (c *rotationInfoController) Add(ctx context.Context, req *shop.RotationInfoAddReq) (res *shop.RotationInfoAddRes, err error) {
	err = service.RotationInfo().Add(ctx, req.RotationInfoAddReq)
	return
}

// Edit 修改轮播图
func (c *rotationInfoController) Edit(ctx context.Context, req *shop.RotationInfoEditReq) (res *shop.RotationInfoEditRes, err error) {
	err = service.RotationInfo().Edit(ctx, req.RotationInfoEditReq)
	return
}

// Delete 删除轮播图
func (c *rotationInfoController) Delete(ctx context.Context, req *shop.RotationInfoDeleteReq) (res *shop.RotationInfoDeleteRes, err error) {
	err = service.RotationInfo().Delete(ctx, req.Ids)
	return
}
