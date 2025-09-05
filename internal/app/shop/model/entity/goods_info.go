// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-09-05 12:04:34
// 生成路径: internal/app/shop/model/entity/goods_info.go
// 生成人：王中阳
// desc:商品表
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// GoodsInfo is the golang structure for table goods_info.
type GoodsInfo struct {
	gmeta.Meta             `orm:"table:goods_info"`
	Id                     uint                         `orm:"id,primary" json:"id"`                       // ID
	Name                   string                       `orm:"name" json:"name"`                           // 名称
	Images                 string                       `orm:"images" json:"images"`                       // 支持单图,多图
	Price                  int                          `orm:"price" json:"price"`                         // 价格(分)
	Level1CategoryId       int                          `orm:"level1_category_id" json:"level1CategoryId"` // 一级分类
	LinkedLevel1CategoryId *LinkedGoodsInfoCategoryInfo `orm:"with:id=level1_category_id" json:"linkedLevel1CategoryId"`
	Level2CategoryId       int                          `orm:"level2_category_id" json:"level2CategoryId"` // 二级分类
	LinkedLevel2CategoryId *LinkedGoodsInfoCategoryInfo `orm:"with:id=level2_category_id" json:"linkedLevel2CategoryId"`
	Level3CategoryId       int                          `orm:"level3_category_id" json:"level3CategoryId"` // 三级分类
	LinkedLevel3CategoryId *LinkedGoodsInfoCategoryInfo `orm:"with:id=level3_category_id" json:"linkedLevel3CategoryId"`
	Brand                  string                       `orm:"brand" json:"brand"`            // 品牌
	Stock                  int                          `orm:"stock" json:"stock"`            // 库存
	Sale                   int                          `orm:"sale" json:"sale"`              // 销量
	Tags                   string                       `orm:"tags" json:"tags"`              // 标签
	DetailInfo             string                       `orm:"detail_info" json:"detailInfo"` // 商品详情
	CreatedAt              *gtime.Time                  `orm:"created_at" json:"createdAt"`   //
	UpdatedAt              *gtime.Time                  `orm:"updated_at" json:"updatedAt"`   //
	DeletedAt              *gtime.Time                  `orm:"deleted_at" json:"deletedAt"`   //
}

type LinkedGoodsInfoCategoryInfo struct {
	gmeta.Meta `orm:"table:category_info"`
	Id         int    `orm:"id" json:"id"`     //
	Name       string `orm:"name" json:"name"` //
}
