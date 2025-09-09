// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-09-09 15:39:40
// 生成路径: internal/app/shop/model/entity/user_coupon_info.go
// 生成人：gfast
// desc:用户优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// UserCouponInfo is the golang structure for table user_coupon_info.
type UserCouponInfo struct {
	gmeta.Meta     `orm:"table:user_coupon_info"`
	Id             int                             `orm:"id,primary" json:"id"`      //
	UserId         int                             `orm:"user_id" json:"userId"`     // 用户id
	CouponId       int                             `orm:"coupon_id" json:"couponId"` // 优惠券id
	LinkedCouponId *LinkedUserCouponInfoCouponInfo `orm:"with:id=coupon_id" json:"linkedCouponId"`
	Status         int                             `orm:"status" json:"status"`        // 状态
	Amount         int                             `orm:"amount" json:"amount"`        // 优惠金额（元）
	CreatedAt      *gtime.Time                     `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt      *gtime.Time                     `orm:"updated_at" json:"updatedAt"` // 更新时间
	DeletedAt      *gtime.Time                     `orm:"deleted_at" json:"deletedAt"` // 删除时间（软删除）
}

type LinkedUserCouponInfoCouponInfo struct {
	gmeta.Meta `orm:"table:coupon_info"`
	Id         int    `orm:"id" json:"id"`     //
	Name       string `orm:"name" json:"name"` // 优惠券名称
}
