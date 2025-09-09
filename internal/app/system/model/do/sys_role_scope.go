// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleScope is the golang structure of table sys_role_scope for DAO operations like Where/Data.
type SysRoleScope struct {
	g.Meta    `orm:"table:sys_role_scope, do:true"`
	Id        interface{} // ID
	RoleId    interface{} // 角色id
	MenuId    interface{} // api接口id
	DataScope interface{} // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	DeptIds   interface{} // 扩展数据
}
