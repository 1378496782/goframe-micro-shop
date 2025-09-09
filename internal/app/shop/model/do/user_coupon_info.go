// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-09-09 15:39:41
// 生成路径: internal/app/shop/model/entity/user_coupon_info.go
// 生成人：gfast
// desc:用户优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// UserCouponInfo is the golang structure for table user_coupon_info.
type UserCouponInfo struct {
	gmeta.Meta `orm:"table:user_coupon_info, do:true"`
	Id         interface{} `orm:"id,primary" json:"id"`        //
	UserId     interface{} `orm:"user_id" json:"userId"`       // 用户id
	CouponId   interface{} `orm:"coupon_id" json:"couponId"`   // 优惠券id
	Status     interface{} `orm:"status" json:"status"`        // 状态
	Amount     interface{} `orm:"amount" json:"amount"`        // 优惠金额（元）
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 更新时间
	DeletedAt  *gtime.Time `orm:"deleted_at" json:"deletedAt"` // 删除时间（软删除）
}
