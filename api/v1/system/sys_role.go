/*
* @desc:角色api
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/30 9:16
 */

package system

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type RoleListReq struct {
	g.Meta   `path:"/role/list" tags:"系统后台/角色管理" method:"get" summary:"角色列表"`
	RoleName string `p:"roleName"`   //参数名称
	Status   string `p:"roleStatus"` //状态
	commonApi.PageReq
}

type RoleListRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysRole `json:"list"`
}

type RoleGetParamsReq struct {
	g.Meta `path:"/role/getParams" tags:"系统后台/角色管理" method:"get" summary:"角色编辑参数"`
}

type RoleGetParamsRes struct {
	g.Meta      `mime:"application/json"`
	Menu        []*model.SysAuthRuleInfoRes `json:"menu"`
	AccessMenus *garray.Array               `json:"accessMenus"`
}

type RoleAddReq struct {
	g.Meta    `path:"/role/add" tags:"系统后台/角色管理" method:"post" summary:"添加角色"`
	Pid       uint   `p:"pid"`
	Name      string `p:"name" v:"required#角色名称不能为空"`
	Status    uint   `p:"status"    `
	ListOrder uint   `p:"listOrder" `
	Remark    string `p:"remark"    `
	MenuIds   []uint `p:"menuIds"`
	CreatedBy uint64
	model.EffectiveTimeInfo
}

type RoleAddRes struct {
}

type RoleGetReq struct {
	g.Meta `path:"/role/get" tags:"系统后台/角色管理" method:"get" summary:"获取角色信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#角色id不能为空"`
}

type RoleGetRes struct {
	g.Meta  `mime:"application/json"`
	Role    *model.RoleInfoRes `json:"role"`
	MenuIds []int              `json:"menuIds"`
}

type RoleEditReq struct {
	g.Meta `path:"/role/edit" tags:"系统后台/角色管理" method:"put" summary:"修改角色"`
	commonApi.Author
	Id        int64  `p:"id" v:"required#角色id必须"`
	Pid       uint   `p:"pid"`
	Name      string `p:"name" v:"required#角色名称不能为空"`
	Status    uint   `p:"status"    `
	ListOrder uint   `p:"listOrder" `
	Remark    string `p:"remark"    `
	MenuIds   []uint `p:"menuIds"`
	model.EffectiveTimeInfo
}

type RoleEditRes struct {
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" tags:"系统后台/角色管理" method:"delete" summary:"删除角色"`
	Ids    []int64 `p:"ids" v:"required#角色id不能为空"`
}

type RoleDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type RoleDeptTreeSelectReq struct {
	g.Meta `path:"/role/deptTreeSelect" tags:"系统后台/角色管理" method:"get" summary:"获取部门树"`
	commonApi.Author
}

type RoleDeptTreeSelectRes struct {
	g.Meta `mime:"application/json"`
	Depts  []*model.SysDeptTreeRes `json:"depts"`
}

type RoleMenuTreeSelectReq struct {
	g.Meta `path:"/role/menuTreeSelect" tags:"系统后台/角色管理" method:"get" summary:"获取菜单树"`
	commonApi.Author
	RoleId uint `p:"roleId" v:"required#角色id必须"`
}

type RoleMenuTreeSelectRes struct {
	g.Meta    `mime:"application/json"`
	Rules     []*model.SysAuthRuleTreeRes `json:"rules"`
	DataScope []*model.ScopeAuthData      `json:"dataScope"`
}

// DataScopeReq 角色数据授权参数
type DataScopeReq struct {
	g.Meta   `path:"/role/dataScope" tags:"系统后台/角色管理" method:"put" summary:"角色数据授权"`
	RoleId   uint                      `p:"roleId" v:"required#角色ID不能为空"`
	AuthData []*model.ScopeAuthDataReq `p:"authData"`
}

type DataScopeRes struct {
	commonApi.EmptyRes
}

type SetRoleUserReq struct {
	g.Meta  `path:"/role/setRoleUser" tags:"系统后台/角色管理" method:"put" summary:"角色用户授权"`
	RoleId  uint     `p:"roleId" v:"required#角色ID不能为空"`
	UserIds []uint64 `p:"userIds"`
}

type SetRoleUserRes struct {
	commonApi.EmptyRes
}
