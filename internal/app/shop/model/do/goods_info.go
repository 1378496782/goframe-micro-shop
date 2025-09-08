// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-09-08 11:37:29
// 生成路径: internal/app/shop/model/entity/goods_info.go
// 生成人：王中阳
// desc:商品
// company:云南奇讯科技有限公司
// ==========================================================================

package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// GoodsInfo is the golang structure for table goods_info.
type GoodsInfo struct {
	gmeta.Meta       `orm:"table:goods_info, do:true"`
	Id               interface{} `orm:"id,primary" json:"id"`                       // ID
	Name             interface{} `orm:"name" json:"name"`                           // 名称
	Images           interface{} `orm:"images" json:"images"`                       // 多图
	PicUrl           interface{} `orm:"pic_url" json:"picUrl"`                      // 封面图
	Price            interface{} `orm:"price" json:"price"`                         // 价格(分)
	Level1CategoryId interface{} `orm:"level1_category_id" json:"level1CategoryId"` // 一级分类
	Level2CategoryId interface{} `orm:"level2_category_id" json:"level2CategoryId"` // 二级分类
	Level3CategoryId interface{} `orm:"level3_category_id" json:"level3CategoryId"` // 三级分类
	Brand            interface{} `orm:"brand" json:"brand"`                         // 品牌
	Stock            interface{} `orm:"stock" json:"stock"`                         // 库存
	Sale             interface{} `orm:"sale" json:"sale"`                           // 销量
	Tags             interface{} `orm:"tags" json:"tags"`                           // 标签
	DetailInfo       interface{} `orm:"detail_info" json:"detailInfo"`              // 商品详情
	CreatedAt        *gtime.Time `orm:"created_at" json:"createdAt"`                //
	Sort             interface{} `orm:"sort" json:"sort"`                           // 排序 倒叙
	UpdatedAt        *gtime.Time `orm:"updated_at" json:"updatedAt"`                //
	DeletedAt        *gtime.Time `orm:"deleted_at" json:"deletedAt"`                //
}
