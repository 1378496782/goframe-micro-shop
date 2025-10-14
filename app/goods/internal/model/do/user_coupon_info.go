// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCouponInfo is the golang structure of table user_coupon_info for DAO operations like Where/Data.
type UserCouponInfo struct {
	g.Meta    `orm:"table:user_coupon_info, do:true"`
	Id        any         //
	UserId    any         // 用户id
	CouponId  any         // 优惠券id
	Status    any         // 状态：0-待使用，1-已使用，2-已过期
	Amount    any         // 优惠金额（单位：分）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间（软删除）
}
