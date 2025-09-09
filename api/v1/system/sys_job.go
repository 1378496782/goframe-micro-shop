// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2023-01-12 17:43:50
// 生成路径: api/v1/system/sys_job.go
// 生成人：gfast
// desc:定时任务相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"

	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

// SysJobSearchReq 分页请求参数
type SysJobSearchReq struct {
	g.Meta   `path:"/list" tags:"系统后台/定时任务" method:"get" summary:"定时任务列表"`
	JobName  string `p:"jobName"`                          //任务名称
	JobGroup string `p:"jobGroup"`                         //任务组名
	Status   string `p:"status" v:"status@integer#状态需为整数"` //状态
	commonApi.PageReq
	commonApi.Author
}

// SysJobSearchRes 列表返回结果
type SysJobSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.SysJobListRes `json:"list"`
}

// SysJobAddReq 添加操作请求参数
type SysJobAddReq struct {
	g.Meta `path:"/add" tags:"系统后台/定时任务" method:"post" summary:"定时任务添加"`
	commonApi.Author
	JobName        string `p:"jobName" v:"required#任务名称不能为空"`
	JobParams      string `p:"jobParams" `
	JobGroup       string `p:"jobGroup" `
	InvokeTarget   string `p:"invokeTarget" v:"required#任务方法不能为空"`
	CronExpression string `p:"cronExpression" v:"required#cron执行表达式不能为空"`
	MisfirePolicy  int    `p:"misfirePolicy" `
	Status         int    `p:"status" v:"required#状态不能为空"`
	Remark         string `p:"remark" `
	CreatedBy      uint64
}

// SysJobAddRes 添加操作返回结果
type SysJobAddRes struct {
	commonApi.EmptyRes
}

// SysJobEditReq 修改操作请求参数
type SysJobEditReq struct {
	g.Meta `path:"/edit" tags:"系统后台/定时任务" method:"put" summary:"定时任务修改"`
	commonApi.Author
	JobId          uint64 `p:"jobId" v:"required#主键ID不能为空"`
	JobName        string `p:"jobName" v:"required#任务名称不能为空"`
	JobParams      string `p:"jobParams" `
	JobGroup       string `p:"jobGroup" `
	InvokeTarget   string `p:"invokeTarget" v:"required#任务方法不能为空"`
	CronExpression string `p:"cronExpression" v:"required#cron执行表达式不能为空"`
	MisfirePolicy  int    `p:"misfirePolicy" `
	Status         int    `p:"status" v:"required#状态不能为空"`
	Remark         string `p:"remark" `
	UpdatedBy      uint64
}

// SysJobEditRes 修改操作返回结果
type SysJobEditRes struct {
	commonApi.EmptyRes
}

// SysJobGetReq 获取一条数据请求
type SysJobGetReq struct {
	g.Meta `path:"/get" tags:"系统后台/定时任务" method:"get" summary:"获取定时任务信息"`
	commonApi.Author
	JobId uint64 `p:"jobId" v:"required#主键必须"` //通过主键获取
}

// SysJobGetRes 获取一条数据结果
type SysJobGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SysJobInfoRes
}

// SysJobDeleteReq 删除数据请求
type SysJobDeleteReq struct {
	g.Meta `path:"/delete" tags:"系统后台/定时任务" method:"delete" summary:"删除定时任务"`
	commonApi.Author
	JobIds []uint64 `p:"jobIds" v:"required#主键必须"` //通过主键删除
}

// SysJobDeleteRes 删除数据返回
type SysJobDeleteRes struct {
	commonApi.EmptyRes
}

type SysJobStartReq struct {
	g.Meta `path:"/start" tags:"系统后台/定时任务" method:"put" summary:"启动任务"`
	commonApi.Author
	JobId uint64 `p:"jobId" v:"required#jobId必须"`
}

type SysJobStartRes struct {
	commonApi.EmptyRes
}

type SysJobStopReq struct {
	g.Meta `path:"/stop" tags:"系统后台/定时任务" method:"put" summary:"停止任务"`
	commonApi.Author
	JobId uint64 `p:"jobId" v:"required#jobId必须"`
}

type SysJobStopRes struct {
	commonApi.EmptyRes
}

type SysJobRunReq struct {
	g.Meta `path:"/run" tags:"系统后台/定时任务" method:"put" summary:"运行任务"`
	commonApi.Author
	JobId uint64 `p:"jobId" v:"required#jobId必须"`
}

type SysJobRunRes struct {
	commonApi.EmptyRes
}

type SysJobLogListReq struct {
	g.Meta `path:"/logs" tags:"系统后台/定时任务" method:"get" summary:"执行日志"`
	commonApi.Author
	commonApi.PageReq
	TargetName string `p:"targetName" v:"required#targetName必须"`
}

type SysJobLogListRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysJobLog `json:"list"`
}

type SysJobLogDeleteReq struct {
	g.Meta `path:"/deleteLogs" tags:"系统后台/定时任务" method:"delete" summary:"删除执行日志"`
	commonApi.Author
	TargetName string `p:"targetName" v:"required#日志标签名不能为空"`
}

type SysJobLogDeleteRes struct {
	commonApi.EmptyRes
}
