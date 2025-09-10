package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CartInfoGetListReq 购物车选项分页查询
type CartInfoGetListReq struct {
	g.Meta `path:"/cart" method:"get" tags:"购物车管理" sm:"购物车选项列表"`
	Page   uint32 `json:"page" d:"1"  v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type CartInfoGetListRes struct {
	List  []*CartInfoItem `json:"list" dc:"购物车选项列表"`
	Page  uint32          `json:"page" dc:"当前页码"`
	Size  uint32          `json:"size" dc:"每页数量"`
	Total uint32          `json:"total" dc:"总数"`
}

type CartInfoItem struct {
	// 购物车字段
	Id     uint32 `json:"id" dc:"购物车ID"`
	UserId uint32 `json:"user_id" dc:"用户ID"`
	Count  uint32 `json:"count" dc:"商品数量"`

	// 商品字段
	GoodsId        uint32                 `json:"goods_id" dc:"商品ID"`
	GoodsName      string                 `json:"goods_name" dc:"商品名称"`
	GoodsPicUrl    string                 `json:"goods_pic_url" dc:"商品图片"`
	GoodsPrice     uint64                 `json:"goods_price" dc:"商品价格(分)"`
	GoodsBrand     string                 `json:"goods_brand" dc:"商品品牌"`
	GoodsStock     uint32                 `json:"goods_stock" dc:"商品库存"`
	GoodsSale      uint32                 `json:"goods_sale" dc:"商品销量"`
	GoodsTags      string                 `json:"goods_tags" dc:"商品标签"`
	GoodsSort      uint32                 `json:"goods_sort" dc:"商品排序"`
	GoodsCreatedAt *timestamppb.Timestamp `json:"goods_created_at" dc:"商品创建时间"`
	GoodsUpdatedAt *timestamppb.Timestamp `json:"goods_updated_at" dc:"商品更新时间"`
}

// CartInfoCreateReq 创建购物车选项
type CartInfoCreateReq struct {
	g.Meta  `path:"/cart" method:"post" tags:"购物车管理" sm:"创建购物车选项"`
	GoodsId uint32 `json:"goods_id" dc:"商品id"`
	Count   uint32 `json:"count" dc:"商品数量"`
}

type CartInfoCreateRes struct {
	Id uint32 `json:"id" dc:"购物车ID"`
}

// CartInfoDeleteReq 删除购物车选项
type CartInfoDeleteReq struct {
	g.Meta `path:"/cart" method:"delete" tags:"购物车管理" sm:"删除购物车选项"`
	Id     uint32 `json:"id" v:"required" dc:"购物车ID"`
}

type CartInfoDeleteRes struct {
	Id uint32 `json:"id" v:"required" dc:"购物车ID"`
}
