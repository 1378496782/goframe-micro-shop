// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-09-22 16:48:52
// 生成路径: internal/app/shop/dao/goods_info.go
// 生成人：gfast
// desc:商品
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/shop/dao/internal"
)

// goodsInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type goodsInfoDao struct {
	*internal.GoodsInfoDao
}

var (
	// GoodsInfo is globally public accessible object for table tools_gen_table operations.
	GoodsInfo = goodsInfoDao{
		internal.NewGoodsInfoDao(),
	}
)

// Fill with you ideas below.
