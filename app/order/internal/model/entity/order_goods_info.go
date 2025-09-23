// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderGoodsInfo is the golang structure for table order_goods_info.
type OrderGoodsInfo struct {
	Id             int         `json:"id"             orm:"id"               ` // 商品维度的订单表
	OrderId        int         `json:"orderId"        orm:"order_id"         ` // 关联的主订单表
	GoodsId        int         `json:"goodsId"        orm:"goods_id"         ` // 商品id
	GoodsOptionsId int         `json:"goodsOptionsId" orm:"goods_options_id" ` // 商品规格id sku id
	Count          int         `json:"count"          orm:"count"            ` // 商品数量
	Remark         string      `json:"remark"         orm:"remark"           ` // 备注
	Price          int         `json:"price"          orm:"price"            ` // 订单金额 单位分
	CouponPrice    int         `json:"couponPrice"    orm:"coupon_price"     ` // 优惠券金额 单位分
	ActualPrice    int         `json:"actualPrice"    orm:"actual_price"     ` // 实际支付金额 单位分
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       ` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       ` //
}
