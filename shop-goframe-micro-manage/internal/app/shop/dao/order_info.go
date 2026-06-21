// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-10-10 23:08:01
// 生成路径: internal/app/shop/dao/order_info.go
// 生成人：gfast
// desc:订单表
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/shop/dao/internal"
)

// orderInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type orderInfoDao struct {
	*internal.OrderInfoDao
}

var (
	// OrderInfo is globally public accessible object for table tools_gen_table operations.
	OrderInfo = orderInfoDao{
		internal.NewOrderInfoDao(),
	}
)

// Fill with you ideas below.
