// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CouponInfoDao is the data access object for the table coupon_info.
type CouponInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CouponInfoColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CouponInfoColumns defines and stores column names for the table coupon_info.
type CouponInfoColumns struct {
	Id        string //
	GoodsId   string // 关联商品id（0表示全场通用）
	Name      string // 优惠券名称
	Type      string // 优惠券类型：0-新人券，1-活动券，2-其他
	Amount    string // 优惠金额（单位：分）
	Deadline  string // 过期时间
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间（软删除）
}

// couponInfoColumns holds the columns for the table coupon_info.
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
func NewCouponInfoDao(handlers ...gdb.ModelHandler) *CouponInfoDao {
	return &CouponInfoDao{
		group:    "default",
		table:    "coupon_info",
		columns:  couponInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CouponInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CouponInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CouponInfoDao) Columns() CouponInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CouponInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CouponInfoDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CouponInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
