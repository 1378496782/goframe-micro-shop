// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BargainInfoDao is the data access object for the table bargain_info.
type BargainInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  BargainInfoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// BargainInfoColumns defines and stores column names for the table bargain_info.
type BargainInfoColumns struct {
	Id          string //
	UserId      string //
	GoodsId     string //
	Counts      string //
	CreatedTime string //
	UpdatedTime string //
	DeletedTime string //
	ExpiredTime string //
}

// bargainInfoColumns holds the columns for the table bargain_info.
var bargainInfoColumns = BargainInfoColumns{
	Id:          "id",
	UserId:      "user_id",
	GoodsId:     "goods_id",
	Counts:      "counts",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
	DeletedTime: "deleted_time",
	ExpiredTime: "expired_time",
}

// NewBargainInfoDao creates and returns a new DAO object for table data access.
func NewBargainInfoDao(handlers ...gdb.ModelHandler) *BargainInfoDao {
	return &BargainInfoDao{
		group:    "default",
		table:    "bargain_info",
		columns:  bargainInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BargainInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BargainInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BargainInfoDao) Columns() BargainInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BargainInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BargainInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BargainInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
