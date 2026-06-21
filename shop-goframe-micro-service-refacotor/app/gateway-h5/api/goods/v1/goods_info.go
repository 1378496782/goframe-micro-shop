package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 商品详情
type GoodsInfoGetDetailReq struct {
	g.Meta `path:"/goods/detail" method:"get" tags:"商品管理" sm:"商品详情"`
	Id     uint32 `json:"id" v:"required" dc:"商品ID"`
}

type GoodsInfoGetDetailRes struct {
	*GoodsInfoItem // 复用列表项结构
}

// 商品分页查询
type GoodsInfoGetListReq struct {
	g.Meta      `path:"/goods" method:"get" tags:"商品管理" sm:"商品分页列表"`
	Page        uint32 `json:"page" d:"1"  v:"min:1" dc:"页码"`
	Size        uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
	IsHot       uint32 `json:"is_hot" d:"0"  dc:"热门推荐(1 开启)"`
	Keyword     string `json:"keyword" dc:"搜索关键词"`
	CategoryId  uint32 `json:"category_id" dc:"分类ID"`
	PriceMin    uint64 `json:"price_min" dc:"最低价格(分)"`
	PriceMax    uint64 `json:"price_max" dc:"最高价格(分)"`
	SortType    uint32 `json:"sort_type" d:"0" v:"in:0,1,2,3" dc:"排序类型：0默认排序 1价格升序 2价格降序 3销量降序"`
	OnlyInStock uint32 `json:"only_in_stock" d:"0" v:"in:0,1" dc:"是否只看有库存：0否 1是"`
}

type GoodsInfoGetListRes struct {
	List  []*GoodsInfoItem `json:"list" dc:"商品列表"`
	Page  uint32           `json:"page" dc:"当前页码"`
	Size  uint32           `json:"size" dc:"每页数量"`
	Total uint32           `json:"total" dc:"总数"`
}

type GoodsInfoItem struct {
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
