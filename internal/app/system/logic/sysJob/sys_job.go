// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2023-01-12 17:43:50
// 生成路径: internal/app/system/logic/sys_job.go
// 生成人：gfast
// desc:定时任务
// company:云南奇讯科技有限公司
// ==========================================================================

package logic

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	systemService "github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"strings"
)

func init() {
	service.RegisterSysJob(New())
}

func New() service.ISysJob {
	return &sSysJob{}
}

type sSysJob struct{}

func (s *sSysJob) List(ctx context.Context, req *system.SysJobSearchReq) (listRes *system.SysJobSearchRes, err error) {
	listRes = new(system.SysJobSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysJob.Ctx(ctx).WithAll()
		if req.JobName != "" {
			m = m.Where(dao.SysJob.Columns().JobName+" like ?", "%"+req.JobName+"%")
		}
		if req.JobGroup != "" {
			m = m.Where(dao.SysJob.Columns().JobGroup+" = ?", req.JobGroup)
		}
		if req.Status != "" {
			m = m.Where(dao.SysJob.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "job_id asc"
		var res []*model.SysJobInfoRes
		err = m.Fields(system.SysJobSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SysJobListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SysJobListRes{
				JobId:          v.JobId,
				JobName:        v.JobName,
				JobGroup:       v.JobGroup,
				InvokeTarget:   v.InvokeTarget,
				CronExpression: v.CronExpression,
				MisfirePolicy:  v.MisfirePolicy,
				Status:         v.Status,
			}
		}
	})
	return
}

func (s *sSysJob) GetByJobId(ctx context.Context, jobId uint64) (res *model.SysJobInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysJob.Ctx(ctx).WithAll().Where(dao.SysJob.Columns().JobId, jobId).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSysJob) Add(ctx context.Context, req *system.SysJobAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysJob.Ctx(ctx).Insert(do.SysJob{
			JobName:        req.JobName,
			JobParams:      req.JobParams,
			JobGroup:       req.JobGroup,
			InvokeTarget:   req.InvokeTarget,
			CronExpression: req.CronExpression,
			MisfirePolicy:  req.MisfirePolicy,
			Status:         req.Status,
			Remark:         req.Remark,
			CreatedBy:      systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sSysJob) Edit(ctx context.Context, req *system.SysJobEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysJob.Ctx(ctx).WherePri(req.JobId).Update(do.SysJob{
			JobName:        req.JobName,
			JobParams:      req.JobParams,
			JobGroup:       req.JobGroup,
			InvokeTarget:   req.InvokeTarget,
			CronExpression: req.CronExpression,
			MisfirePolicy:  req.MisfirePolicy,
			Status:         req.Status,
			Remark:         req.Remark,
			UpdatedBy:      systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSysJob) Delete(ctx context.Context, jobIds []uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysJob.Ctx(ctx).Delete(dao.SysJob.Columns().JobId+" in (?)", jobIds)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

func (s *sSysJob) Start(ctx context.Context, jobId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var job *model.SysJobInfoRes
		job, err = s.GetByJobId(ctx, jobId)
		liberr.ErrIsNil(ctx, err)
		if job != nil {
			s.JobStart(ctx, job)
		}
	})
	return
}

func (s *sSysJob) Stop(ctx context.Context, jobId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var job *model.SysJobInfoRes
		job, err = s.GetByJobId(ctx, jobId)
		liberr.ErrIsNil(ctx, err)
		if job != nil {
			s.jobStop(ctx, job)
		}
	})
	return
}

func (s *sSysJob) Run(ctx context.Context, jobId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var job *model.SysJobInfoRes
		job, err = s.GetByJobId(ctx, jobId)
		liberr.ErrIsNil(ctx, err)
		if job != nil {
			s.jobRun(ctx, job)
		}
	})
	return
}

// JobRun 执行任务
func (s *sSysJob) jobRun(ctx context.Context, job *model.SysJobInfoRes) error {
	//可以task目录下是否绑定对应的方法
	f := service.TaskList().GetByName(job.InvokeTarget)
	if f == nil {
		return gerror.New("当前task目录下没有绑定这个方法")
	}
	//传参
	paramArr := strings.Split(job.JobParams, "|")
	service.TaskList().EditParams(f.FuncName, paramArr)
	task, err := gcron.AddOnce(gctx.New(), "@every 1s", f.Run)
	if err != nil || task == nil {
		return gerror.New("启动执行失败")
	}
	return nil
}

// GetJobs 获取已开启执行的任务
func (s *sSysJob) GetJobs(ctx context.Context) (jobs []*model.SysJobInfoRes, err error) {
	err = dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().Status, 0).Scan(&jobs)
	return
}

// JobStart 启动任务
func (s *sSysJob) JobStart(ctx context.Context, job *model.SysJobInfoRes) error {
	//获取task目录下是否绑定对应的方法
	f := service.TaskList().GetByName(job.InvokeTarget)
	if f == nil {
		return gerror.New("没有绑定对应的方法")
	}
	//传参
	paramArr := strings.Split(job.JobParams, "|")
	service.TaskList().EditParams(f.FuncName, paramArr)
	rs := gcron.Search(job.InvokeTarget)
	if rs == nil {
		if job.MisfirePolicy == 1 {
			t, err := gcron.AddSingleton(gctx.New(), job.CronExpression, f.Run, job.InvokeTarget)
			if err != nil {
				return err
			}
			if t == nil {
				return gerror.New("启动任务失败")
			}
		} else {
			t, err := gcron.AddOnce(gctx.New(), job.CronExpression, f.Run, job.InvokeTarget)
			if err != nil {
				return err
			}
			if t == nil {
				return gerror.New("启动任务失败")
			}
		}
	}
	gcron.Start(job.InvokeTarget)
	if job.MisfirePolicy == 1 {
		job.Status = 0
		_, err := dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().JobId, job.JobId).Unscoped().Update(g.Map{
			dao.SysJob.Columns().Status: job.Status,
		})
		return err
	}
	return nil
}

// JobStop 停止任务
func (s *sSysJob) jobStop(ctx context.Context, job *model.SysJobInfoRes) error {
	//获取task目录下是否绑定对应的方法
	f := service.TaskList().GetByName(job.InvokeTarget)
	if f == nil {
		return gerror.New("没有绑定对应的方法")
	}
	rs := gcron.Search(job.InvokeTarget)
	if rs != nil {
		gcron.Remove(job.InvokeTarget)
	}
	job.Status = 1
	_, err := dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().JobId, job.JobId).Unscoped().Update(g.Map{
		dao.SysJob.Columns().Status: job.Status,
	})
	return err
}
