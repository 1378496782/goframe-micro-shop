// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-11-09 11:41:17
// 生成路径: internal/app/system/model/entity/sys_notice.go
// 生成人：gfast
// desc:通知公告
// company:云南奇讯科技有限公司
// ==========================================================================

package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysNotice is the golang structure for table sys_notice.
type SysNotice struct {
	gmeta.Meta `orm:"table:sys_notice, do:true"`
	Id         interface{} `orm:"id,primary" json:"id"`        // ID
	Title      interface{} `orm:"title" json:"title"`          // 标题
	Type       interface{} `orm:"type" json:"type"`            // 类型
	Tag        interface{} `orm:"tag" json:"tag"`              // 标签
	Content    interface{} `orm:"content" json:"content"`      // 内容
	Remark     interface{} `orm:"remark" json:"remark"`        // 备注
	Sort       interface{} `orm:"sort" json:"sort"`            // 排序
	Status     interface{} `orm:"status" json:"status"`        // 状态
	Receiver   interface{} `orm:"receiver" json:"receiver"`    //接收者
	CreatedBy  interface{} `orm:"created_by" json:"createdBy"` // 发送人
	UpdatedBy  interface{} `orm:"updated_by" json:"updatedBy"` // 修改人
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 更新时间
	DeletedAt  *gtime.Time `orm:"deleted_at" json:"deletedAt"` // 删除时间
}
