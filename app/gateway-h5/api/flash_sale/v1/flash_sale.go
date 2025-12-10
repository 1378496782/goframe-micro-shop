package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 秒杀商品列表
type FlashSaleGoodsListReq struct {
	g.Meta `path:"/flash-sale/goods" method:"get" tags:"秒杀管理" sm:"秒杀商品列表"`
	Page   uint32 `json:"page" d:"1" v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type FlashSaleGoodsListRes struct {
	List  []*FlashSaleGoodsItem `json:"list" dc:"秒杀商品列表"`
	Page  uint32                `json:"page" dc:"当前页码"`
	Size  uint32                `json:"size" dc:"每页数量"`
	Total uint32                `json:"total" dc:"总数"`
}

type FlashSaleGoodsItem struct {
	Id          uint32 `json:"id" dc:"秒杀商品ID"`
	GoodsId     uint32 `json:"goods_id" dc:"商品ID"`
	Title       string `json:"title" dc:"秒杀标题"`
	Price       uint64 `json:"price" dc:"秒杀价格"`
	Stock       uint32 `json:"stock" dc:"秒杀库存"`
	Sold        uint32 `json:"sold" dc:"已售数量"`
	StartTime   int64  `json:"start_time" dc:"开始时间戳"`
	EndTime     int64  `json:"end_time" dc:"结束时间戳"`
	Status      uint32 `json:"status" dc:"状态：1-未开始，2-进行中，3-已结束"`
	PicUrl      string `json:"pic_url" dc:"商品主图"`
	OriginalPrice uint64 `json:"original_price" dc:"原价"`
}

// 秒杀商品详情
type FlashSaleGoodsDetailReq struct {
	g.Meta `path:"/flash-sale/goods/detail" method:"get" tags:"秒杀管理" sm:"秒杀商品详情"`
	Id     uint32 `json:"id" v:"required" dc:"秒杀商品ID"`
}

type FlashSaleGoodsDetailRes struct {
	*FlashSaleGoodsItem
	DetailInfo string `json:"detail_info" dc:"商品详情"`
	Images     string `json:"images" dc:"商品图片"`
}

// 创建秒杀订单
type CreateFlashSaleOrderReq struct {
	g.Meta  `path:"/flash-sale/order" method:"post" tags:"秒杀管理" sm:"创建秒杀订单"`
	GoodsId uint32 `json:"goods_id" v:"required" dc:"秒杀商品ID"`
	Count   uint32 `json:"count" v:"required|min:1" dc:"购买数量"`
}

type CreateFlashSaleOrderRes struct {
	OrderId string `json:"order_id" dc:"订单ID"`
	Status  uint32 `json:"status" dc:"订单状态：1-处理中，2-成功，3-失败"`
	Message string `json:"message" dc:"状态描述"`
}

// 获取秒杀结果
type GetFlashSaleResultReq struct {
	g.Meta  `path:"/flash-sale/result" method:"get" tags:"秒杀管理" sm:"获取秒杀结果"`
	OrderId string `json:"order_id" v:"required" dc:"订单ID"`
}

type GetFlashSaleResultRes struct {
	OrderId   string `json:"order_id" dc:"订单ID"`
	Status    uint32 `json:"status" dc:"订单状态：1-处理中，2-成功，3-失败"`
	Message   string `json:"message" dc:"状态描述"`
	GoodsName string `json:"goods_name" dc:"商品名称"`
	Price     uint64 `json:"price" dc:"实际支付价格"`
}