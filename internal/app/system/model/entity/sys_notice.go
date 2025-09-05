// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-11-09 11:41:17
// 生成路径: internal/app/system/model/entity/sys_notice.go
// 生成人：gfast
// desc:通知公告
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysNotice is the golang structure for table sys_notice.
type SysNotice struct {
	gmeta.Meta `orm:"table:sys_notice"`
	Id         int64       `orm:"id,primary" json:"id"`        // ID
	Title      string      `orm:"title" json:"title"`          // 标题
	Type       int64       `orm:"type" json:"type"`            // 类型
	Tag        int         `orm:"tag" json:"tag"`              // 标签
	Content    string      `orm:"content" json:"content"`      // 内容
	Remark     string      `orm:"remark" json:"remark"`        // 备注
	Sort       int         `orm:"sort" json:"sort"`            // 排序
	Status     int         `orm:"status" json:"status"`        // 状态
	CreatedBy  int64       `orm:"created_by" json:"createdBy"` // 发送人
	UpdatedBy  int64       `orm:"updated_by" json:"updatedBy"` // 修改人
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 更新时间
	DeletedAt  *gtime.Time `orm:"deleted_at" json:"deletedAt"` // 删除时间
}
