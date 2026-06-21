// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-09-09 15:39:40
// 生成路径: internal/app/shop/dao/user_coupon_info.go
// 生成人：gfast
// desc:用户优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/shop/dao/internal"
)

// userCouponInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type userCouponInfoDao struct {
	*internal.UserCouponInfoDao
}

var (
	// UserCouponInfo is globally public accessible object for table tools_gen_table operations.
	UserCouponInfo = userCouponInfoDao{
		internal.NewUserCouponInfoDao(),
	}
)

// Fill with you ideas below.
