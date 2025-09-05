// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2023-11-09 11:41:17
// 生成路径: internal/app/system/dao/internal/sys_notice.go
// 生成人：gfast
// desc:通知公告
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysNoticeDao is the manager for logic model data accessing and custom defined data operations functions management.
type SysNoticeDao struct {
	table   string           // Table is the underlying table name of the DAO.
	group   string           // Group is the database configuration group name of current DAO.
	columns SysNoticeColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// SysNoticeColumns defines and stores column names for table sys_notice.
type SysNoticeColumns struct {
	Id        string // ID
	Title     string // 标题
	Type      string // 类型
	Tag       string // 标签
	Content   string // 内容
	Remark    string // 备注
	Sort      string // 排序
	Status    string // 状态
	Receiver  string //接收者
	CreatedBy string // 发送人
	UpdatedBy string // 修改人
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

var sysNoticeColumns = SysNoticeColumns{
	Id:        "id",
	Title:     "title",
	Type:      "type",
	Tag:       "tag",
	Content:   "content",
	Remark:    "remark",
	Sort:      "sort",
	Status:    "status",
	Receiver:  "receiver",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewSysNoticeDao creates and returns a new DAO object for table data access.
func NewSysNoticeDao() *SysNoticeDao {
	return &SysNoticeDao{
		group:   "default",
		table:   "sys_notice",
		columns: sysNoticeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysNoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysNoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysNoticeDao) Columns() SysNoticeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysNoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysNoticeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysNoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
