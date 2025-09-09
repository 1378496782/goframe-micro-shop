/*
* @desc:任务日志
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/1/13 21:42
 */

package sysJobLog

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysJobLog(New())
}

func New() service.ISysJobLog {
	return new(sSysJobLog)
}

type sSysJobLog struct {
}

func (s *sSysJobLog) Add(ctx context.Context, data *do.SysJobLog) (err error) {
	_, err = dao.SysJobLog.Ctx(ctx).Insert(data)
	return
}

func (s *sSysJobLog) List(ctx context.Context, req *system.SysJobLogListReq) (listRes *system.SysJobLogListRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		listRes = new(system.SysJobLogListRes)
		m := dao.SysJobLog.Ctx(ctx).Where(dao.SysJobLog.Columns().TargetName, req.TargetName)
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		err = m.Fields(system.SysJobLogListRes{}).Page(req.PageNum, req.PageSize).Order("id desc").Scan(&listRes.List)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *sSysJobLog) Delete(ctx context.Context, targetName string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysJobLog.Ctx(ctx).Delete(dao.SysJobLog.Columns().TargetName, targetName)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
