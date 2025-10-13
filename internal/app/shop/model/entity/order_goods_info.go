// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-10-11 23:55:29
// 生成路径: internal/app/shop/model/entity/order_goods_info.go
// 生成人：gfast
// desc:订单物品表
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// OrderGoodsInfo is the golang structure for table order_goods_info.
type OrderGoodsInfo struct {
	gmeta.Meta     `orm:"table:order_goods_info"`
	Id             int         `orm:"id,primary" json:"id"`                   // 商品维度的订单表
	OrderId        int         `orm:"order_id" json:"orderId"`                // 关联的主订单表
	GoodsId        int         `orm:"goods_id" json:"goodsId"`                // 商品id
	GoodsOptionsId int         `orm:"goods_options_id" json:"goodsOptionsId"` // 商品规格id sku id
	Count          int         `orm:"count" json:"count"`                     // 商品数量
	Remark         string      `orm:"remark" json:"remark"`                   // 备注
	Price          int         `orm:"price" json:"price"`                     // 订单金额 单位分
	CouponPrice    int         `orm:"coupon_price" json:"couponPrice"`        // 优惠券金额 单位分
	ActualPrice    int         `orm:"actual_price" json:"actualPrice"`        // 实际支付金额 单位分
	CreatedAt      *gtime.Time `orm:"created_at" json:"createdAt"`            //
	UpdatedAt      *gtime.Time `orm:"updated_at" json:"updatedAt"`            //
}
