package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 商品分页查询
type AddGoodsInfoGetListReq struct {
	g.Meta `path:"/add/goods" method:"get" tags:"加购商品" sm:"加购商品列表"`
}

type AddGoodsInfoGetListRes struct {
	List  []*AddGoodsInfoItem `json:"list" dc:"商品列表"`
	Total uint32              `json:"total" dc:"总数"`
}

type AddGoodsInfoItem struct {
	Id               uint32                 `json:"id" dc:"商品ID"`
	PicUrl           string                 `json:"pic_url" dc:"主图"`
	Images           string                 `json:"images" dc:"支持单图,多图"`
	Name             string                 `json:"name" dc:"商品名称"`
	Price            uint64                 `json:"price" dc:"价格"`
	Level1CategoryId uint32                 `json:"level1_category_id" dc:"一级分类ID"`
	Level2CategoryId uint32                 `json:"level2_category_id" dc:"二级分类ID"`
	Level3CategoryId uint32                 `json:"level3_category_id" dc:"三级分类ID"`
	Brand            string                 `json:"brand" dc:"品牌"`
	Stock            uint32                 `json:"stock" dc:"库存"`
	Sale             uint32                 `json:"sale" dc:"销量"`
	Tags             string                 `json:"tags" dc:"标签"`
	DetailInfo       string                 `json:"detail_info" dc:"详情"`
	Sort             uint32                 `json:"sort" dc:"排序 倒序"`
	CreatedAt        *timestamppb.Timestamp `json:"created_at" dc:"创建时间"`
	UpdatedAt        *timestamppb.Timestamp `json:"updated_at" dc:"更新时间"`
}
