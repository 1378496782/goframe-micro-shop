// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2024-10-23 16:10:12
// 生成路径: internal/app/system/dao/internal/sys_attachment.go
// 生成人：gfast
// desc:附件管理
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAttachmentDao is the manager for logic model data accessing and custom defined data operations functions management.
type SysAttachmentDao struct {
	table   string               // Table is the underlying table name of the DAO.
	group   string               // Group is the database configuration group name of current DAO.
	columns SysAttachmentColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// SysAttachmentColumns defines and stores column names for table sys_attachment.
type SysAttachmentColumns struct {
	Id        string // 文件ID
	AppId     string // 应用ID
	Drive     string // 上传驱动
	Name      string // 文件原始名
	Kind      string // 上传类型
	MimeType  string // 扩展类型
	Path      string // 本地路径
	Size      string // 文件大小
	Ext       string // 扩展名
	Md5       string // md5校验码
	CreatedBy string // 上传人ID
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

var sysAttachmentColumns = SysAttachmentColumns{
	Id:        "id",
	AppId:     "app_id",
	Drive:     "drive",
	Name:      "name",
	Kind:      "kind",
	MimeType:  "mime_type",
	Path:      "path",
	Size:      "size",
	Ext:       "ext",
	Md5:       "md5",
	CreatedBy: "created_by",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysAttachmentDao creates and returns a new DAO object for table data access.
func NewSysAttachmentDao() *SysAttachmentDao {
	return &SysAttachmentDao{
		group:   "default",
		table:   "sys_attachment",
		columns: sysAttachmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysAttachmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysAttachmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysAttachmentDao) Columns() SysAttachmentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysAttachmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysAttachmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysAttachmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
