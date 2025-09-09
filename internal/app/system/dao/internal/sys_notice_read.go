// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2023-11-09 17:37:25
// 生成路径: internal/app/system/dao/internal/sys_notice_read.go
// 生成人：gfast
// desc:已读记录
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysNoticeReadDao is the manager for logic model data accessing and custom defined data operations functions management.
type SysNoticeReadDao struct {
	table   string               // Table is the underlying table name of the DAO.
	group   string               // Group is the database configuration group name of current DAO.
	columns SysNoticeReadColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// SysNoticeReadColumns defines and stores column names for table sys_notice_read.
type SysNoticeReadColumns struct {
	Id        string // id
	NoticeId  string // 信息id
	UserId    string // 用户id
	UpdatedAt string // 更新时间
	CreatedAt string // 阅读时间
}

var sysNoticeReadColumns = SysNoticeReadColumns{
	Id:        "id",
	NoticeId:  "notice_id",
	UserId:    "user_id",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

// NewSysNoticeReadDao creates and returns a new DAO object for table data access.
func NewSysNoticeReadDao() *SysNoticeReadDao {
	return &SysNoticeReadDao{
		group:   "default",
		table:   "sys_notice_read",
		columns: sysNoticeReadColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysNoticeReadDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysNoticeReadDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysNoticeReadDao) Columns() SysNoticeReadColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysNoticeReadDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysNoticeReadDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysNoticeReadDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
