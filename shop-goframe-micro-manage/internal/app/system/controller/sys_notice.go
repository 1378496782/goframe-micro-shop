// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2023-11-09 11:41:17
// 生成路径: internal/app/system/controller/sys_notice.go
// 生成人：gfast
// desc:通知公告
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
	"fmt"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

type sysNoticeController struct {
	BaseController
}

var SysNotice = new(sysNoticeController)

// List 列表（管理）
func (c *sysNoticeController) List(ctx context.Context, req *system.SysNoticeSearchReq) (res *system.SysNoticeSearchRes, err error) {
	res = new(system.SysNoticeSearchRes)
	res.SysNoticeSearchRes, err = service.SysNotice().List(ctx, &req.SysNoticeSearchReq)
	return
}

// ShoWList(展示)
func (c *sysNoticeController) ShowList(ctx context.Context, req *system.SysNoticeShowSearchReq) (res *system.SysNoticeShowSearchRes, err error) {
	res = new(system.SysNoticeShowSearchRes)
	res.SysNoticeSearchRes, err = service.SysNotice().ListShow(ctx, &req.SysNoticeSearchReq)
	return
}

// Get 获取通知公告
func (c *sysNoticeController) Get(ctx context.Context, req *system.SysNoticeGetReq) (res *system.SysNoticeGetRes, err error) {
	res = new(system.SysNoticeGetRes)
	res.SysNoticeInfoRes, err = service.SysNotice().GetById(ctx, req.Id)
	return
}

// Add 添加通知公告
func (c *sysNoticeController) Add(ctx context.Context, req *system.SysNoticeAddReq) (res *system.SysNoticeAddRes, err error) {
	req.CreatedBy = service.Context().GetUserId(ctx)
	err = service.SysNotice().Add(ctx, req.SysNoticeAddReq)
	return
}

// Edit 修改通知公告
func (c *sysNoticeController) Edit(ctx context.Context, req *system.SysNoticeEditReq) (res *system.SysNoticeEditRes, err error) {
	req.UpdatedBy = service.Context().GetUserId(ctx)
	err = service.SysNotice().Edit(ctx, req.SysNoticeEditReq)
	return
}

// Delete 删除通知公告
func (c *sysNoticeController) Delete(ctx context.Context, req *system.SysNoticeDeleteReq) (res *system.SysNoticeDeleteRes, err error) {
	err = service.SysNotice().Delete(ctx, req.Ids)
	return
}

// IndexData 获取首页数据
func (c *sysNoticeController) IndexData(ctx context.Context, req *system.SysNoticeIndexDataReq) (res *system.SysNoticeIndexDataRes, err error) {
	//res.SysNoticeIndexRes, err = service.SysNotice().IndexData(ctx)
	return
}

// UnReadCount
func (c *sysNoticeController) UnReadCount(ctx context.Context, req *system.SysNoticeUnReadCountReq) (res *system.SysNoticeUnReadCountRes, err error) {
	res = new(system.SysNoticeUnReadCountRes)
	currentUser := service.Context().Get(ctx).User.Id
	//noticeids, err := service.SysNotice().CurrentUseWithIds(ctx, currentUser)
	res.SysNoticeUnreadCount, err = service.SysNotice().UnReadCount(ctx, currentUser)
	fmt.Println(res)
	return
}

// GetUserNameList 得到供选择用户列表
func (c *sysNoticeController) GetUserNameList(ctx context.Context, req *system.SysNoticeUserSearchReq) (res *system.SysNoticeUserSearchRes, err error) {
	res = new(system.SysNoticeUserSearchRes)
	list, err := service.SysNotice().GetUserNameList(ctx, req.UserNickName)
	res.UserList = list
	return
}
