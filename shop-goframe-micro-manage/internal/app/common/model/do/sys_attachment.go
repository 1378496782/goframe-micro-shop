// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2024-10-23 16:10:12
// 生成路径: internal/app/system/model/entity/sys_attachment.go
// 生成人：gfast
// desc:附件管理
// company:云南奇讯科技有限公司
// ==========================================================================

package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysAttachment is the golang structure for table sys_attachment.
type SysAttachment struct {
	gmeta.Meta `orm:"table:sys_attachment, do:true"`
	Id         interface{} `orm:"id,primary" json:"id"`        // 文件ID
	AppId      interface{} `orm:"app_id" json:"appId"`         // 应用ID
	Drive      interface{} `orm:"drive" json:"drive"`          // 上传驱动
	Name       interface{} `orm:"name" json:"name"`            // 文件原始名
	Kind       interface{} `orm:"kind" json:"kind"`            // 上传类型
	MimeType   interface{} `orm:"mime_type" json:"mimeType"`   // 扩展类型
	Path       interface{} `orm:"path" json:"path"`            // 本地路径
	Size       interface{} `orm:"size" json:"size"`            // 文件大小
	Ext        interface{} `orm:"ext" json:"ext"`              // 扩展名
	Md5        interface{} `orm:"md5" json:"md5"`              // md5校验码
	CreatedBy  interface{} `orm:"created_by" json:"createdBy"` // 上传人ID
	Status     interface{} `orm:"status" json:"status"`        // 状态
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 修改时间
}
