// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2024-10-23 16:10:12
// 生成路径: internal/app/system/model/entity/sys_attachment.go
// 生成人：gfast
// desc:附件管理
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysAttachment is the golang structure for table sys_attachment.
type SysAttachment struct {
	gmeta.Meta `orm:"table:sys_attachment"`
	Id         int64       `orm:"id,primary" json:"id"`        // 文件ID
	AppId      string      `orm:"app_id" json:"appId"`         // 应用ID
	Drive      string      `orm:"drive" json:"drive"`          // 上传驱动
	Name       string      `orm:"name" json:"name"`            // 文件原始名
	Kind       string      `orm:"kind" json:"kind"`            // 上传类型
	MimeType   string      `orm:"mime_type" json:"mimeType"`   // 扩展类型
	Path       string      `orm:"path" json:"path"`            // 本地路径
	Size       int64       `orm:"size" json:"size"`            // 文件大小
	Ext        string      `orm:"ext" json:"ext"`              // 扩展名
	Md5        string      `orm:"md5" json:"md5"`              // md5校验码
	CreatedBy  int64       `orm:"created_by" json:"createdBy"` // 上传人ID
	Status     bool        `orm:"status" json:"status"`        // 状态
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 修改时间
}
