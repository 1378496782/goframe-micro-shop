// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-09-05 11:52:54
// 生成路径: internal/app/shop/model/entity/category_info.go
// 生成人：王中阳
// desc:商品分类
// company:云南奇讯科技有限公司
// ==========================================================================

package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// CategoryInfo is the golang structure for table category_info.
type CategoryInfo struct {
	gmeta.Meta `orm:"table:category_info, do:true"`
	Id         interface{} `orm:"id,primary" json:"id"`        // ID
	ParentId   interface{} `orm:"parent_id" json:"parentId"`   // 父级id
	Name       interface{} `orm:"name" json:"name"`            // 名称
	PicUrl     interface{} `orm:"pic_url" json:"picUrl"`       // 图片
	Level      interface{} `orm:"level" json:"level"`          // 等级
	Sort       interface{} `orm:"sort" json:"sort"`            // 排序
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` //
	DeletedAt  *gtime.Time `orm:"deleted_at" json:"deletedAt"` //
}
