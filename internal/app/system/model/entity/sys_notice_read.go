// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-11-09 17:37:25
// 生成路径: internal/app/system/model/entity/sys_notice_read.go
// 生成人：gfast
// desc:已读记录
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysNoticeRead is the golang structure for table sys_notice_read.
type SysNoticeRead struct {
	gmeta.Meta `orm:"table:sys_notice_read"`
	Id         int64       `orm:"id,primary" json:"id"`        // id
	NoticeId   int64       `orm:"notice_id" json:"noticeId"`   // 信息id
	UserId     int64       `orm:"user_id" json:"userId"`       // 用户id
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 更新时间
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 阅读时间
}
