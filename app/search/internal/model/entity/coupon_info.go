// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CouponInfo is the golang structure for table coupon_info.
type CouponInfo struct {
	Id        int         `json:"id"        orm:"id"         description:""`                       //
	GoodsId   int         `json:"goodsId"   orm:"goods_id"   description:"关联商品id（0表示全场通用）"`        // 关联商品id（0表示全场通用）
	Name      string      `json:"name"      orm:"name"       description:"优惠券名称"`                  // 优惠券名称
	Type      int         `json:"type"      orm:"type"       description:"优惠券类型：0-新人券，1-活动券，2-其他"` // 优惠券类型：0-新人券，1-活动券，2-其他
	Amount    int         `json:"amount"    orm:"amount"     description:"优惠金额（单位：分）"`             // 优惠金额（单位：分）
	Deadline  *gtime.Time `json:"deadline"  orm:"deadline"   description:"过期时间"`                   // 过期时间
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                   // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                   // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间（软删除）"`              // 删除时间（软删除）
}
