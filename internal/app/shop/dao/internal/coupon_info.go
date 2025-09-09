// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-09-09 15:10:49
// 生成路径: internal/app/shop/dao/internal/coupon_info.go
// 生成人：gfast
// desc:优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CouponInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
type CouponInfoDao struct {
	table   string            // Table is the underlying table name of the DAO.
	group   string            // Group is the database configuration group name of current DAO.
	columns CouponInfoColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// CouponInfoColumns defines and stores column names for table coupon_info.
type CouponInfoColumns struct {
	Id        string // ID
	GoodsId   string // 关联商品id（0表示全场通用）
	Name      string // 名称
	Type      string // 类型
	Amount    string // 优惠金额（元）
	Deadline  string // 过期时间
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间（软删除）
}

var couponInfoColumns = CouponInfoColumns{
	Id:        "id",
	GoodsId:   "goods_id",
	Name:      "name",
	Type:      "type",
	Amount:    "amount",
	Deadline:  "deadline",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewCouponInfoDao creates and returns a new DAO object for table data access.
func NewCouponInfoDao() *CouponInfoDao {
	return &CouponInfoDao{
		group:   "goods",
		table:   "coupon_info",
		columns: couponInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CouponInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CouponInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CouponInfoDao) Columns() CouponInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CouponInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CouponInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CouponInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
