// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RefundInfo is the golang structure for table refund_info.
type RefundInfo struct {
	Id           int         `json:"id"           orm:"id"            description:"售后退款表"`                      // 售后退款表
	Number       string      `json:"number"       orm:"number"        description:"售后订单号"`                      // 售后订单号
	OrderId      int         `json:"orderId"      orm:"order_id"      description:"订单id"`                       // 订单id
	GoodsId      int         `json:"goodsId"      orm:"goods_id"      description:"要售后的商品id"`                   // 要售后的商品id
	Reason       string      `json:"reason"       orm:"reason"        description:"退款原因"`                       // 退款原因
	Status       int         `json:"status"       orm:"status"        description:"审核状态 1待处理 2同意退款 3拒绝退款"`      // 审核状态 1待处理 2同意退款 3拒绝退款
	RefundStatus int         `json:"refundStatus" orm:"refund_status" description:"退款状态 0未退款 1退款中 2退款成功 3退款失败"` // 退款状态 0未退款 1退款中 2退款成功 3退款失败
	RefundAmount int         `json:"refundAmount" orm:"refund_amount" description:"退款金额 单位分"`                   // 退款金额 单位分
	UserId       int         `json:"userId"       orm:"user_id"       description:"用户id"`                       // 用户id
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`                           //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`                           //
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""`                           //
}
