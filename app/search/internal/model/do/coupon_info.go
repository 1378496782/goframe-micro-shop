// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CouponInfo is the golang structure of table coupon_info for DAO operations like Where/Data.
type CouponInfo struct {
	g.Meta    `orm:"table:coupon_info, do:true"`
	Id        any         //
	GoodsId   any         // 关联商品id（0表示全场通用）
	Name      any         // 优惠券名称
	Type      any         // 优惠券类型：0-新人券，1-活动券，2-其他
	Amount    any         // 优惠金额（单位：分）
	Deadline  *gtime.Time // 过期时间
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间（软删除）
}
