// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-09-22 16:30:35
// 生成路径: internal/app/shop/model/entity/goods_info.go
// 生成人：gfast
// desc:商品
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// GoodsInfo is the golang structure for table goods_info.
type GoodsInfo struct {
	gmeta.Meta       `orm:"table:goods_info"`
	Id               uint        `orm:"id,primary" json:"id"`                       // ID
	Name             string      `orm:"name" json:"name"`                           // 名字
	PicUrl           string      `orm:"pic_url" json:"picUrl"`                      // 主图
	Images           string      `orm:"images" json:"images"`                       // 详情配图
	Price            int         `orm:"price" json:"price"`                         // 价格(分)
	Level1CategoryId int         `orm:"level1_category_id" json:"level1CategoryId"` // 1级分类id
	Level2CategoryId int         `orm:"level2_category_id" json:"level2CategoryId"` // 2级分类id
	Level3CategoryId int         `orm:"level3_category_id" json:"level3CategoryId"` // 3级分类id
	Brand            string      `orm:"brand" json:"brand"`                         // 品牌
	Stock            int         `orm:"stock" json:"stock"`                         // 库存
	Sale             int         `orm:"sale" json:"sale"`                           // 销量
	Tags             string      `orm:"tags" json:"tags"`                           // 标签
	Sort             int         `orm:"sort" json:"sort"`                           // 排序 倒叙
	DetailInfo       string      `orm:"detail_info" json:"detailInfo"`              // 商品详情
	EnableBargain    int         `orm:"enable_bargain" json:"enableBargain"`        // 允许砍价
	CreatedAt        *gtime.Time `orm:"created_at" json:"createdAt"`                //
	UpdatedAt        *gtime.Time `orm:"updated_at" json:"updatedAt"`                //
	DeletedAt        *gtime.Time `orm:"deleted_at" json:"deletedAt"`                //
}
