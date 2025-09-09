// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-09-09 17:28:51
// 生成路径: internal/app/shop/model/entity/rotation_info.go
// 生成人：gfast
// desc:轮播图
// company:云南奇讯科技有限公司
// ==========================================================================

package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// RotationInfo is the golang structure for table rotation_info.
type RotationInfo struct {
	gmeta.Meta `orm:"table:rotation_info, do:true"`
	Id         interface{} `orm:"id,primary" json:"id"`        // ID
	PicUrl     interface{} `orm:"pic_url" json:"picUrl"`       // 图片
	Link       interface{} `orm:"link" json:"link"`            // 跳转链接
	Sort       interface{} `orm:"sort" json:"sort"`            // 排序字段
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` //
	DeletedAt  *gtime.Time `orm:"deleted_at" json:"deletedAt"` //
}
