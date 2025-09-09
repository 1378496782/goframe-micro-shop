// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2023-11-09 17:37:25
// 生成路径: internal/app/system/controller/sys_notice_read.go
// 生成人：gfast
// desc:已读记录
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

type sysNoticeReadController struct {
	BaseController
}

var SysNoticeRead = new(sysNoticeReadController)

// List 列表
func (c *sysNoticeReadController) List(ctx context.Context, req *system.SysNoticeReadSearchReq) (res *system.SysNoticeReadSearchRes, err error) {
	res = new(system.SysNoticeReadSearchRes)
	res.SysNoticeReadSearchRes, err = service.SysNoticeRead().List(ctx, &req.SysNoticeReadSearchReq)
	return
}

// Get 获取已读记录
func (c *sysNoticeReadController) Get(ctx context.Context, req *system.SysNoticeReadGetReq) (res *system.SysNoticeReadGetRes, err error) {
	res = new(system.SysNoticeReadGetRes)
	res.SysNoticeReadInfoRes, err = service.SysNoticeRead().GetById(ctx, req.Id)
	return
}

// Add 添加已读记录
func (c *sysNoticeReadController) Add(ctx context.Context, req *system.SysNoticeReadAddReq) (res *system.SysNoticeReadAddRes, err error) {
	err = service.SysNoticeRead().Add(ctx, req.SysNoticeReadAddReq)
	return
}

// Add 添加已读记录
func (c *sysNoticeReadController) ReadNotice(ctx context.Context, req *system.SysNoticeReadNoticeAddReq) (res *system.SysNoticeReadNoticeAddRes, err error) {
	err = service.SysNoticeRead().ReadNotice(ctx, req.SysNoticeReadNoticeReq)
	return
}

// Edit 修改已读记录
func (c *sysNoticeReadController) Edit(ctx context.Context, req *system.SysNoticeReadEditReq) (res *system.SysNoticeReadEditRes, err error) {
	err = service.SysNoticeRead().Edit(ctx, req.SysNoticeReadEditReq)
	return
}

// Delete 删除已读记录
func (c *sysNoticeReadController) Delete(ctx context.Context, req *system.SysNoticeReadDeleteReq) (res *system.SysNoticeReadDeleteRes, err error) {
	err = service.SysNoticeRead().Delete(ctx, req.Ids)
	return
}
