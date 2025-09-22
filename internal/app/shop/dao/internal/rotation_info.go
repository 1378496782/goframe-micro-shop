// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-09-22 16:50:03
// 生成路径: internal/app/shop/dao/internal/rotation_info.go
// 生成人：gfast
// desc:轮播图
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RotationInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
type RotationInfoDao struct {
	table   string              // Table is the underlying table name of the DAO.
	group   string              // Group is the database configuration group name of current DAO.
	columns RotationInfoColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// RotationInfoColumns defines and stores column names for table rotation_info.
type RotationInfoColumns struct {
	Id        string // ID
	PicUrl    string // 图片
	Link      string // 跳转链接
	Sort      string // 排序字段
	CreatedAt string // 创建时间
	UpdatedAt string //
	DeletedAt string //
}

var rotationInfoColumns = RotationInfoColumns{
	Id:        "id",
	PicUrl:    "pic_url",
	Link:      "link",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewRotationInfoDao creates and returns a new DAO object for table data access.
func NewRotationInfoDao() *RotationInfoDao {
	return &RotationInfoDao{
		group:   "banner",
		table:   "rotation_info",
		columns: rotationInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RotationInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RotationInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RotationInfoDao) Columns() RotationInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RotationInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RotationInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RotationInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
