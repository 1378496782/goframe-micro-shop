// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleScopeDao is the data access object for table sys_role_scope.
type SysRoleScopeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SysRoleScopeColumns // columns contains all the column names of Table for convenient usage.
}

// SysRoleScopeColumns defines and stores column names for table sys_role_scope.
type SysRoleScopeColumns struct {
	Id        string // ID
	RoleId    string // 角色id
	MenuId    string // api接口id
	DataScope string // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	DeptIds   string // 扩展数据
}

// sysRoleScopeColumns holds the columns for table sys_role_scope.
var sysRoleScopeColumns = SysRoleScopeColumns{
	Id:        "id",
	RoleId:    "role_id",
	MenuId:    "menu_id",
	DataScope: "data_scope",
	DeptIds:   "dept_ids",
}

// NewSysRoleScopeDao creates and returns a new DAO object for table data access.
func NewSysRoleScopeDao() *SysRoleScopeDao {
	return &SysRoleScopeDao{
		group:   "default",
		table:   "sys_role_scope",
		columns: sysRoleScopeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysRoleScopeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysRoleScopeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysRoleScopeDao) Columns() SysRoleScopeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysRoleScopeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysRoleScopeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysRoleScopeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
