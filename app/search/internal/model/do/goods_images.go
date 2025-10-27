// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsImages is the golang structure of table goods_images for DAO operations like Where/Data.
type GoodsImages struct {
	g.Meta  `orm:"table:goods_images, do:true"`
	Id      any //
	Url     any // 七牛云url
	GoodsId any // 商品ID
	FileId  any // 文件ID（关联file_info）
	Sort    any // 排序
}
