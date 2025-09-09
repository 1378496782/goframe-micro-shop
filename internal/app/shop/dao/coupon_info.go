// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-09-09 15:10:49
// 生成路径: internal/app/shop/dao/coupon_info.go
// 生成人：gfast
// desc:优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/shop/dao/internal"
)

// couponInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type couponInfoDao struct {
	*internal.CouponInfoDao
}

var (
	// CouponInfo is globally public accessible object for table tools_gen_table operations.
	CouponInfo = couponInfoDao{
		internal.NewCouponInfoDao(),
	}
)

// Fill with you ideas below.
