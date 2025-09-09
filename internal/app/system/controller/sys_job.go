// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2023-01-12 17:43:50
// 生成路径: internal/app/system/controller/sys_job.go
// 生成人：gfast
// desc:定时任务
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

type sysJobController struct {
	BaseController
}

var SysJob = new(sysJobController)

// List 列表
func (c *sysJobController) List(ctx context.Context, req *system.SysJobSearchReq) (res *system.SysJobSearchRes, err error) {
	res, err = service.SysJob().List(ctx, req)
	return
}

// Get 获取定时任务
func (c *sysJobController) Get(ctx context.Context, req *system.SysJobGetReq) (res *system.SysJobGetRes, err error) {
	res = new(system.SysJobGetRes)
	res.SysJobInfoRes, err = service.SysJob().GetByJobId(ctx, req.JobId)
	return
}

// Add 添加定时任务
func (c *sysJobController) Add(ctx context.Context, req *system.SysJobAddReq) (res *system.SysJobAddRes, err error) {
	err = service.SysJob().Add(ctx, req)
	return
}

// Edit 修改定时任务
func (c *sysJobController) Edit(ctx context.Context, req *system.SysJobEditReq) (res *system.SysJobEditRes, err error) {
	err = service.SysJob().Edit(ctx, req)
	return
}

// Delete 删除定时任务
func (c *sysJobController) Delete(ctx context.Context, req *system.SysJobDeleteReq) (res *system.SysJobDeleteRes, err error) {
	err = service.SysJob().Delete(ctx, req.JobIds)
	return
}

// Start 启动任务
func (c *sysJobController) Start(ctx context.Context, req *system.SysJobStartReq) (res *system.SysJobStartRes, err error) {
	err = service.SysJob().Start(ctx, req.JobId)
	return
}

// Stop 停止任务
func (c *sysJobController) Stop(ctx context.Context, req *system.SysJobStopReq) (res *system.SysJobStopRes, err error) {
	err = service.SysJob().Stop(ctx, req.JobId)
	return
}

// Run 运行任务
func (c *sysJobController) Run(ctx context.Context, req *system.SysJobRunReq) (res *system.SysJobRunRes, err error) {
	err = service.SysJob().Run(ctx, req.JobId)
	return
}

// Log 获取执行日志列表
func (c *sysJobController) Log(ctx context.Context, req *system.SysJobLogListReq) (res *system.SysJobLogListRes, err error) {
	res, err = service.SysJobLog().List(ctx, req)
	return
}

func (c *sysJobController) LogDelete(ctx context.Context, req *system.SysJobLogDeleteReq) (res *system.SysJobLogDeleteRes, err error) {
	err = service.SysJobLog().Delete(ctx, req.TargetName)
	return
}
