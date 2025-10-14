// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BargainHistoryDao is the data access object for the table bargain_history.
type BargainHistoryDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  BargainHistoryColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// BargainHistoryColumns defines and stores column names for the table bargain_history.
type BargainHistoryColumns struct {
	Id          string //
	BargainId   string //
	Amount      string //
	UserId      string //
	CreatedTime string //
	DeletedTime string //
}

// bargainHistoryColumns holds the columns for the table bargain_history.
var bargainHistoryColumns = BargainHistoryColumns{
	Id:          "id",
	BargainId:   "bargain_id",
	Amount:      "amount",
	UserId:      "user_id",
	CreatedTime: "created_time",
	DeletedTime: "deleted_time",
}

// NewBargainHistoryDao creates and returns a new DAO object for table data access.
func NewBargainHistoryDao(handlers ...gdb.ModelHandler) *BargainHistoryDao {
	return &BargainHistoryDao{
		group:    "default",
		table:    "bargain_history",
		columns:  bargainHistoryColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BargainHistoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BargainHistoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BargainHistoryDao) Columns() BargainHistoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BargainHistoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BargainHistoryDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BargainHistoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
