// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-09-09 15:39:40
// 生成路径: internal/app/shop/dao/internal/user_coupon_info.go
// 生成人：gfast
// desc:用户优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserCouponInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
type UserCouponInfoDao struct {
	table   string                // Table is the underlying table name of the DAO.
	group   string                // Group is the database configuration group name of current DAO.
	columns UserCouponInfoColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// UserCouponInfoColumns defines and stores column names for table user_coupon_info.
type UserCouponInfoColumns struct {
	Id        string //
	UserId    string // 用户id
	CouponId  string // 优惠券id
	Status    string // 状态
	Amount    string // 优惠金额（元）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间（软删除）
}

var userCouponInfoColumns = UserCouponInfoColumns{
	Id:        "id",
	UserId:    "user_id",
	CouponId:  "coupon_id",
	Status:    "status",
	Amount:    "amount",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewUserCouponInfoDao creates and returns a new DAO object for table data access.
func NewUserCouponInfoDao() *UserCouponInfoDao {
	return &UserCouponInfoDao{
		group:   "goods",
		table:   "user_coupon_info",
		columns: userCouponInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserCouponInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserCouponInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserCouponInfoDao) Columns() UserCouponInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserCouponInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserCouponInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserCouponInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
