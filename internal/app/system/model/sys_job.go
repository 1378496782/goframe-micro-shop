// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2023-01-12 17:43:50
// 生成路径: internal/app/system/model/sys_job.go
// 生成人：gfast
// desc:定时任务
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysJobInfoRes is the golang structure for table sys_job.
type SysJobInfoRes struct {
	gmeta.Meta     `orm:"table:sys_job"`
	JobId          uint64       `orm:"job_id,primary" json:"jobId"`           // 任务ID
	JobName        string       `orm:"job_name" json:"jobName"`               // 任务名称
	JobParams      string       `orm:"job_params" json:"jobParams"`           // 参数
	JobGroup       string       `orm:"job_group" json:"jobGroup"`             // 任务组名
	InvokeTarget   string       `orm:"invoke_target" json:"invokeTarget"`     // 任务方法
	CronExpression string       `orm:"cron_expression" json:"cronExpression"` // cron执行表达式
	MisfirePolicy  int          `orm:"misfire_policy" json:"misfirePolicy"`   // 计划执行策略
	Concurrent     int          `orm:"concurrent" json:"concurrent"`          // 是否并发执行
	Status         int          `orm:"status" json:"status"`                  // 状态
	CreatedBy      uint64       `orm:"created_by" json:"createdBy"`           // 创建者
	UpdatedBy      uint64       `orm:"updated_by" json:"updatedBy"`           // 更新者
	Remark         string       `orm:"remark" json:"remark"`                  // 备注信息
	CreatedAt      *gtime.Time  `orm:"created_at" json:"createdAt"`           // 创建时间
	UpdatedAt      *gtime.Time  `orm:"updated_at" json:"updatedAt"`           // 更新时间
	CreatedUser    *LinkUserRes `orm:"with:id=created_by" json:"createdUser"`
	UpdatedUser    *LinkUserRes `orm:"with:id=updated_by" json:"updatedUser"`
}

type SysJobListRes struct {
	JobId          uint64 `json:"jobId"`
	JobName        string `json:"jobName"`
	JobGroup       string `json:"jobGroup"`
	InvokeTarget   string `json:"invokeTarget"`
	CronExpression string `json:"cronExpression"`
	MisfirePolicy  int    `json:"misfirePolicy"`
	Status         int    `json:"status"`
}
