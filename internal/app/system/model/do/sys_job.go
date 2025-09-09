// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-01-12 17:43:50
// 生成路径: internal/app/system/model/entity/sys_job.go
// 生成人：gfast
// desc:定时任务
// company:云南奇讯科技有限公司
// ==========================================================================

package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysJob is the golang structure for table sys_job.
type SysJob struct {
	gmeta.Meta     `orm:"table:sys_job, do:true"`
	JobId          interface{} `orm:"job_id,primary" json:"jobId"`           // 任务ID
	JobName        interface{} `orm:"job_name" json:"jobName"`               // 任务名称
	JobParams      interface{} `orm:"job_params" json:"jobParams"`           // 参数
	JobGroup       interface{} `orm:"job_group" json:"jobGroup"`             // 任务组名
	InvokeTarget   interface{} `orm:"invoke_target" json:"invokeTarget"`     // 任务方法
	CronExpression interface{} `orm:"cron_expression" json:"cronExpression"` // cron执行表达式
	MisfirePolicy  interface{} `orm:"misfire_policy" json:"misfirePolicy"`   // 计划执行策略
	Concurrent     interface{} `orm:"concurrent" json:"concurrent"`          // 是否并发执行
	Status         interface{} `orm:"status" json:"status"`                  // 状态
	CreatedBy      interface{} `orm:"created_by" json:"createdBy"`           // 创建者
	UpdatedBy      interface{} `orm:"updated_by" json:"updatedBy"`           // 更新者
	Remark         interface{} `orm:"remark" json:"remark"`                  // 备注信息
	CreatedAt      *gtime.Time `orm:"created_at" json:"createdAt"`           // 创建时间
	UpdatedAt      *gtime.Time `orm:"updated_at" json:"updatedAt"`           // 更新时间
}
