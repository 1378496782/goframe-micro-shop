// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-10-11 23:55:29
// 生成路径: internal/app/shop/dao/order_goods_info.go
// 生成人：gfast
// desc:订单物品表
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/shop/dao/internal"
)

// orderGoodsInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type orderGoodsInfoDao struct {
	*internal.OrderGoodsInfoDao
}

var (
	// OrderGoodsInfo is globally public accessible object for table tools_gen_table operations.
	OrderGoodsInfo = orderGoodsInfoDao{
		internal.NewOrderGoodsInfoDao(),
	}
)

// Fill with you ideas below.
