// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysRoleScope is the golang structure for table sys_role_scope.
type SysRoleScope struct {
	Id        uint64 `json:"id"        description:"ID"`
	RoleId    uint64 `json:"roleId"    description:"角色id"`
	MenuId    int    `json:"menuId"    description:"api接口id"`
	DataScope int    `json:"dataScope" description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	DeptIds   string `json:"deptIds"   description:"扩展数据"`
}
