// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-09-05 11:52:54
// 生成路径: internal/app/shop/dao/category_info.go
// 生成人：王中阳
// desc:商品分类
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/shop/dao/internal"
)

// categoryInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type categoryInfoDao struct {
	*internal.CategoryInfoDao
}

var (
	// CategoryInfo is globally public accessible object for table tools_gen_table operations.
	CategoryInfo = categoryInfoDao{
		internal.NewCategoryInfoDao(),
	}
)

// Fill with you ideas below.
