// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderGoodsInfo is the golang structure of table order_goods_info for DAO operations like Where/Data.
type OrderGoodsInfo struct {
	g.Meta         `orm:"table:order_goods_info, do:true"`
	Id             any         // 商品维度的订单表
	OrderId        any         // 关联的主订单表
	GoodsId        any         // 商品id
	GoodsOptionsId any         // 商品规格id sku id
	Count          any         // 商品数量
	Remark         any         // 备注
	Price          any         // 订单金额 单位分
	CouponPrice    any         // 优惠券金额 单位分
	ActualPrice    any         // 实际支付金额 单位分
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
