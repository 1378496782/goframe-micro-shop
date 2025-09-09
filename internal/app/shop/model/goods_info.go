// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-09-08 11:37:29
// 生成路径: internal/app/shop/model/goods_info.go
// 生成人：王中阳
// desc:商品
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
)

// GoodsInfoInfoRes is the golang structure for table goods_info.
type GoodsInfoInfoRes struct {
	gmeta.Meta             `orm:"table:goods_info"`
	Id                     uint                         `orm:"id,primary" json:"id" dc:"ID"`                         // ID
	Name                   string                       `orm:"name" json:"name" dc:"名称"`                             // 名称
	Images                 string                       `orm:"images" json:"images" dc:"多图"`                         // 多图
	PicUrl                 string                       `orm:"pic_url" json:"picUrl" dc:"封面图"`                       // 封面图
	Price                  int                          `orm:"price" json:"price" dc:"价格(分)"`                        // 价格(分)
	Level1CategoryId       int                          `orm:"level1_category_id" json:"level1CategoryId" dc:"一级分类"` // 一级分类
	LinkedLevel1CategoryId *LinkedGoodsInfoCategoryInfo `orm:"with:id=level1_category_id" json:"linkedLevel1CategoryId"`
	Level2CategoryId       int                          `orm:"level2_category_id" json:"level2CategoryId" dc:"二级分类"` // 二级分类
	LinkedLevel2CategoryId *LinkedGoodsInfoCategoryInfo `orm:"with:id=level2_category_id" json:"linkedLevel2CategoryId"`
	Level3CategoryId       int                          `orm:"level3_category_id" json:"level3CategoryId" dc:"三级分类"` // 三级分类
	LinkedLevel3CategoryId *LinkedGoodsInfoCategoryInfo `orm:"with:id=level3_category_id" json:"linkedLevel3CategoryId"`
	Brand                  string                       `orm:"brand" json:"brand" dc:"品牌"`              // 品牌
	Stock                  int                          `orm:"stock" json:"stock" dc:"库存"`              // 库存
	Sale                   int                          `orm:"sale" json:"sale" dc:"销量"`                // 销量
	Tags                   string                       `orm:"tags" json:"tags" dc:"标签"`                // 标签
	DetailInfo             string                       `orm:"detail_info" json:"detailInfo" dc:"商品详情"` // 商品详情
	CreatedAt              *gtime.Time                  `orm:"created_at" json:"createdAt" dc:""`       //
	Sort                   int                          `orm:"sort" json:"sort" dc:"排序 倒叙"`             // 排序 倒叙
	UpdatedAt              *gtime.Time                  `orm:"updated_at" json:"updatedAt" dc:""`       //
	DeletedAt              *gtime.Time                  `orm:"deleted_at" json:"deletedAt" dc:""`       //
}
type LinkedGoodsInfoCategoryInfo struct {
	gmeta.Meta `orm:"table:category_info"`
	Id         int    `orm:"id" json:"id" dc:""`     //
	Name       string `orm:"name" json:"name" dc:""` //
}

type GoodsInfoListRes struct {
	Id                     uint                         `json:"id" dc:"ID"`
	Name                   string                       `json:"name" dc:"名称"`
	PicUrl                 string                       `json:"picUrl" dc:"封面图"`
	Price                  int                          `json:"price" dc:"价格(分)"`
	Level1CategoryId       int                          `json:"level1CategoryId" dc:"一级分类"`
	LinkedLevel1CategoryId *LinkedGoodsInfoCategoryInfo `orm:"with:id=level1_category_id" json:"linkedLevel1CategoryId" dc:"一级分类"`
	Level2CategoryId       int                          `json:"level2CategoryId" dc:"二级分类"`
	LinkedLevel2CategoryId *LinkedGoodsInfoCategoryInfo `orm:"with:id=level2_category_id" json:"linkedLevel2CategoryId" dc:"二级分类"`
	Level3CategoryId       int                          `json:"level3CategoryId" dc:"三级分类"`
	LinkedLevel3CategoryId *LinkedGoodsInfoCategoryInfo `orm:"with:id=level3_category_id" json:"linkedLevel3CategoryId" dc:"三级分类"`
	Brand                  string                       `json:"brand" dc:"品牌"`
	Stock                  int                          `json:"stock" dc:"库存"`
	Sale                   int                          `json:"sale" dc:"销量"`
	Tags                   string                       `json:"tags" dc:"标签"`
	CreatedAt              *gtime.Time                  `json:"createdAt" dc:""`
	Sort                   int                          `json:"sort" dc:"排序 倒叙"`
}

// GoodsInfoSearchReq 分页请求参数
type GoodsInfoSearchReq struct {
	comModel.PageReq
	Id               string `p:"id" dc:"ID"`                                                       //ID
	Name             string `p:"name" dc:"名称"`                                                     //名称
	PicUrl           string `p:"picUrl" dc:"封面图"`                                                  //封面图
	Price            string `p:"price" v:"price@integer#价格(分)需为整数" dc:"价格(分)"`                     //价格(分)
	Level1CategoryId string `p:"level1CategoryId" v:"level1CategoryId@integer#一级分类需为整数" dc:"一级分类"` //一级分类
	Level2CategoryId string `p:"level2CategoryId" v:"level2CategoryId@integer#二级分类需为整数" dc:"二级分类"` //二级分类
	Level3CategoryId string `p:"level3CategoryId" v:"level3CategoryId@integer#三级分类需为整数" dc:"三级分类"` //三级分类
	Brand            string `p:"brand" dc:"品牌"`                                                    //品牌
	Stock            string `p:"stock" v:"stock@integer#库存需为整数" dc:"库存"`                           //库存
	Sale             string `p:"sale" v:"sale@integer#销量需为整数" dc:"销量"`                             //销量
	Tags             string `p:"tags" dc:"标签"`                                                     //标签
	CreatedAt        string `p:"createdAt" v:"createdAt@datetime#需为YYYY-MM-DD hh:mm:ss格式" dc:""`   //
	Sort             string `p:"sort" v:"sort@integer#排序 倒叙需为整数" dc:"排序 倒叙"`                       //排序 倒叙
}

// GoodsInfoSearchRes 列表返回结果
type GoodsInfoSearchRes struct {
	comModel.ListRes
	List []*GoodsInfoListRes `json:"list"`
}

// 相关连表查询数据
type LinkedGoodsInfoDataSearchRes struct {
	LinkedGoodsInfoCategoryInfo []*LinkedGoodsInfoCategoryInfo `json:"linkedGoodsInfoCategoryInfo"`
}

// GoodsInfoAddReq 添加操作请求参数
type GoodsInfoAddReq struct {
	Name             string `p:"name" v:"required#名称不能为空" dc:"名称"`
	PicUrl           string `p:"picUrl"  dc:"封面图"`
	Price            int    `p:"price" v:"required#价格(分)不能为空" dc:"价格(分)"`
	Level1CategoryId int    `p:"level1CategoryId" v:"required#一级分类不能为空" dc:"一级分类"`
	Level2CategoryId int    `p:"level2CategoryId"  dc:"二级分类"`
	Level3CategoryId int    `p:"level3CategoryId"  dc:"三级分类"`
	Brand            string `p:"brand"  dc:"品牌"`
	Stock            int    `p:"stock"  dc:"库存"`
	Sale             int    `p:"sale"  dc:"销量"`
	Tags             string `p:"tags"  dc:"标签"`
	DetailInfo       string `p:"detailInfo"  dc:"商品详情"`
	Sort             int    `p:"sort"  dc:"排序 倒叙"`
}

// GoodsInfoEditReq 修改操作请求参数
type GoodsInfoEditReq struct {
	Id               uint   `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	Name             string `p:"name" v:"required#名称不能为空" dc:"名称"`
	PicUrl           string `p:"picUrl"  dc:"封面图"`
	Price            int    `p:"price" v:"required#价格(分)不能为空" dc:"价格(分)"`
	Level1CategoryId int    `p:"level1CategoryId" v:"required#一级分类不能为空" dc:"一级分类"`
	Level2CategoryId int    `p:"level2CategoryId"  dc:"二级分类"`
	Level3CategoryId int    `p:"level3CategoryId"  dc:"三级分类"`
	Brand            string `p:"brand"  dc:"品牌"`
	Stock            int    `p:"stock"  dc:"库存"`
	Sale             int    `p:"sale"  dc:"销量"`
	Tags             string `p:"tags"  dc:"标签"`
	DetailInfo       string `p:"detailInfo"  dc:"商品详情"`
	Sort             int    `p:"sort"  dc:"排序 倒叙"`
}
