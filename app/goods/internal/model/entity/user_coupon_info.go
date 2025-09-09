// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCouponInfo is the golang structure for table user_coupon_info.
type UserCouponInfo struct {
	Id        int         `json:"id"        orm:"id"         description:""`                     //
	UserId    int         `json:"userId"    orm:"user_id"    description:"用户id"`                 // 用户id
	CouponId  int         `json:"couponId"  orm:"coupon_id"  description:"优惠券id"`                // 优惠券id
	Status    int         `json:"status"    orm:"status"     description:"状态：0-待使用，1-已使用，2-已过期"` // 状态：0-待使用，1-已使用，2-已过期
	Amount    int         `json:"amount"    orm:"amount"     description:"优惠金额（单位：分）"`           // 优惠金额（单位：分）
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                 // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                 // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间（软删除）"`            // 删除时间（软删除）
}
