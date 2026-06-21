// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2024-10-23 16:10:12
// 生成路径: internal/app/system/controller/sys_attachment.go
// 生成人：gfast
// desc:附件管理
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
	"errors"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

type sysAttachmentController struct {
	BaseController
}

var SysAttachment = new(sysAttachmentController)

// List 列表
func (c *sysAttachmentController) List(ctx context.Context, req *system.SysAttachmentSearchReq) (res *system.SysAttachmentSearchRes, err error) {
	res = new(system.SysAttachmentSearchRes)
	res.SysAttachmentSearchRes, err = commonService.SysAttachment().List(ctx, &req.SysAttachmentSearchReq)
	return
}

// Get 获取附件管理
func (c *sysAttachmentController) Get(ctx context.Context, req *system.SysAttachmentGetReq) (res *system.SysAttachmentGetRes, err error) {
	res = new(system.SysAttachmentGetRes)
	res.SysAttachmentInfoRes, err = commonService.SysAttachment().GetById(ctx, req.Id)
	return
}

// Add 添加附件管理
func (c *sysAttachmentController) Add(ctx context.Context, req *system.SysAttachmentAddReq) (res *system.SysAttachmentAddRes, err error) {
	req.CreatedBy = service.Context().GetUserId(ctx)
	err = commonService.SysAttachment().Add(ctx, req.SysAttachmentAddReq)
	return
}

// Edit 修改附件管理
func (c *sysAttachmentController) Edit(ctx context.Context, req *system.SysAttachmentEditReq) (res *system.SysAttachmentEditRes, err error) {
	err = commonService.SysAttachment().Edit(ctx, req.SysAttachmentEditReq)
	return
}

// Delete 删除附件管理
func (c *sysAttachmentController) Delete(ctx context.Context, req *system.SysAttachmentDeleteReq) (res *system.SysAttachmentDeleteRes, err error) {
	err = commonService.SysAttachment().Delete(ctx, req.Ids)
	return
}

// 附件管理状态修改（状态）
func (c *sysAttachmentController) ChangeStatus(ctx context.Context, req *system.SysAttachmentStatusSwitchReq) (res *system.SysAttachmentStatusSwitchRes, err error) {
	if !service.SysUser().AccessRule(ctx, service.Context().GetUserId(ctx), "api/v1/system/sysAttachment/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = commonService.SysAttachment().ChangeStatus(ctx, req.Id, req.Status)
	return
}
