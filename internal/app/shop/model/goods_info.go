// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-09-22 16:30:35
// 生成路径: internal/app/shop/model/goods_info.go
// 生成人：gfast
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
	gmeta.Meta       `orm:"table:goods_info"`
	Id               uint        `orm:"id,primary" json:"id" dc:"ID"`                           // ID
	Name             string      `orm:"name" json:"name" dc:"名字"`                               // 名字
	PicUrl           string      `orm:"pic_url" json:"picUrl" dc:"主图"`                          // 主图
	Images           string      `orm:"images" json:"images" dc:"详情配图"`                         // 详情配图
	Price            int         `orm:"price" json:"price" dc:"价格(分)"`                          // 价格(分)
	Level1CategoryId int         `orm:"level1_category_id" json:"level1CategoryId" dc:"1级分类id"` // 1级分类id
	Level2CategoryId int         `orm:"level2_category_id" json:"level2CategoryId" dc:"2级分类id"` // 2级分类id
	Level3CategoryId int         `orm:"level3_category_id" json:"level3CategoryId" dc:"3级分类id"` // 3级分类id
	Brand            string      `orm:"brand" json:"brand" dc:"品牌"`                             // 品牌
	Stock            int         `orm:"stock" json:"stock" dc:"库存"`                             // 库存
	Sale             int         `orm:"sale" json:"sale" dc:"销量"`                               // 销量
	Tags             string      `orm:"tags" json:"tags" dc:"标签"`                               // 标签
	Sort             int         `orm:"sort" json:"sort" dc:"排序 倒叙"`                            // 排序 倒叙
	DetailInfo       string      `orm:"detail_info" json:"detailInfo" dc:"商品详情"`                // 商品详情
	EnableBargain    int         `orm:"enable_bargain" json:"enableBargain" dc:"允许砍价"`          // 允许砍价
	CreatedAt        *gtime.Time `orm:"created_at" json:"createdAt" dc:""`                      //
	UpdatedAt        *gtime.Time `orm:"updated_at" json:"updatedAt" dc:""`                      //
	DeletedAt        *gtime.Time `orm:"deleted_at" json:"deletedAt" dc:""`                      //
}

type GoodsInfoListRes struct {
	Id            uint        `json:"id" dc:"ID"`
	Name          string      `json:"name" dc:"名字"`
	PicUrl        string      `json:"picUrl" dc:"主图"`
	Images        string      `json:"images" dc:"详情配图"`
	Price         int         `json:"price" dc:"价格(分)"`
	Stock         int         `json:"stock" dc:"库存"`
	Sale          int         `json:"sale" dc:"销量"`
	Tags          string      `json:"tags" dc:"标签"`
	Sort          int         `json:"sort" dc:"排序 倒叙"`
	DetailInfo    string      `json:"detailInfo" dc:"商品详情"`
	EnableBargain int         `json:"enableBargain" dc:"允许砍价"`
	CreatedAt     *gtime.Time `json:"createdAt" dc:""`
}

// GoodsInfoSearchReq 分页请求参数
type GoodsInfoSearchReq struct {
	comModel.PageReq
	Id            string `p:"id" dc:"ID"`                                                     //ID
	Name          string `p:"name" dc:"名字"`                                                   //名字
	PicUrl        string `p:"picUrl" dc:"主图"`                                                 //主图
	Images        string `p:"images" dc:"详情配图"`                                               //详情配图
	Price         string `p:"price" v:"price@integer#价格(分)需为整数" dc:"价格(分)"`                   //价格(分)
	Stock         string `p:"stock" v:"stock@integer#库存需为整数" dc:"库存"`                         //库存
	Sale          string `p:"sale" v:"sale@integer#销量需为整数" dc:"销量"`                           //销量
	Tags          string `p:"tags" dc:"标签"`                                                   //标签
	Sort          string `p:"sort" v:"sort@integer#排序 倒叙需为整数" dc:"排序 倒叙"`                     //排序 倒叙
	DetailInfo    string `p:"detailInfo" dc:"商品详情"`                                           //商品详情
	EnableBargain string `p:"enableBargain" v:"enableBargain@integer#允许砍价需为整数" dc:"允许砍价"`     //允许砍价
	CreatedAt     string `p:"createdAt" v:"createdAt@datetime#需为YYYY-MM-DD hh:mm:ss格式" dc:""` //
}

// GoodsInfoSearchRes 列表返回结果
type GoodsInfoSearchRes struct {
	comModel.ListRes
	List []*GoodsInfoListRes `json:"list"`
}

// GoodsInfoAddReq 添加操作请求参数
type GoodsInfoAddReq struct {
	Name          string             `p:"name" v:"required#名字不能为空" dc:"名字"`
	PicUrl        string             `p:"picUrl"  dc:"主图"`
	Images        []*comModel.UpFile `p:"images"  dc:"详情配图"`
	Price         int                `p:"price" v:"required#价格(分)不能为空" dc:"价格(分)"`
	Stock         int                `p:"stock"  dc:"库存"`
	Sale          int                `p:"sale"  dc:"销量"`
	Tags          string             `p:"tags"  dc:"标签"`
	Sort          int                `p:"sort"  dc:"排序 倒叙"`
	DetailInfo    string             `p:"detailInfo"  dc:"商品详情"`
	EnableBargain int                `p:"enableBargain"  dc:"允许砍价"`
}

// GoodsInfoEditReq 修改操作请求参数
type GoodsInfoEditReq struct {
	Id            uint               `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	Name          string             `p:"name" v:"required#名字不能为空" dc:"名字"`
	PicUrl        string             `p:"picUrl"  dc:"主图"`
	Images        []*comModel.UpFile `p:"images"  dc:"详情配图"`
	Price         int                `p:"price" v:"required#价格(分)不能为空" dc:"价格(分)"`
	Stock         int                `p:"stock"  dc:"库存"`
	Sale          int                `p:"sale"  dc:"销量"`
	Tags          string             `p:"tags"  dc:"标签"`
	Sort          int                `p:"sort"  dc:"排序 倒叙"`
	DetailInfo    string             `p:"detailInfo"  dc:"商品详情"`
	EnableBargain int                `p:"enableBargain"  dc:"允许砍价"`
}
