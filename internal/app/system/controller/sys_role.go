/*
* @desc:角色管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/30 9:08
 */

package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var Role = roleController{}

type roleController struct {
	BaseController
}

// List 角色列表
func (c *roleController) List(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error) {
	res, err = service.SysRole().GetRoleListSearch(ctx, req)
	return
}

// GetParams 获取角色表单参数
func (c *roleController) GetParams(ctx context.Context, req *system.RoleGetParamsReq) (res *system.RoleGetParamsRes, err error) {
	res = new(system.RoleGetParamsRes)
	res.Menu, err = service.SysAuthRule().GetMenuList(ctx)
	if err != nil {
		return
	}
	roleIds, err := service.SysUser().GetAdminRoleIds(ctx, service.Context().GetUserId(ctx))
	if err != nil {
		return
	}
	res.AccessMenus, err = service.SysUser().GetAdminMenusIdsByRoleIds(ctx, roleIds)
	return
}

// Add 添加角色信息
func (c *roleController) Add(ctx context.Context, req *system.RoleAddReq) (res *system.RoleAddRes, err error) {
	err = service.SysRole().AddRole(ctx, req)
	return
}

// Get 获取角色信息
func (c *roleController) Get(ctx context.Context, req *system.RoleGetReq) (res *system.RoleGetRes, err error) {
	res = new(system.RoleGetRes)
	res.Role, err = service.SysRole().Get(ctx, req.Id)
	if err != nil {
		return
	}
	res.MenuIds, err = service.SysRole().GetFilteredNamedPolicy(ctx, req.Id)
	return
}

// Edit 修改角色信息
func (c *roleController) Edit(ctx context.Context, req *system.RoleEditReq) (res *system.RoleEditRes, err error) {
	err = service.SysRole().EditRole(ctx, req)
	return
}

// Delete 删除角色
func (c *roleController) Delete(ctx context.Context, req *system.RoleDeleteReq) (res *system.RoleDeleteRes, err error) {
	err = service.SysRole().DeleteByIds(ctx, req.Ids)
	return
}

// DeptTreeSelect 获取角色授权部门数据
func (c *roleController) DeptTreeSelect(ctx context.Context, req *system.RoleDeptTreeSelectReq) (res *system.RoleDeptTreeSelectRes, err error) {
	res, err = service.SysRole().RoleDeptTreeSelect(ctx)
	return
}

// MenuTreeSelect 获取角色授权接口数据
func (c *roleController) MenuTreeSelect(ctx context.Context, req *system.RoleMenuTreeSelectReq) (res *system.RoleMenuTreeSelectRes, err error) {
	var list []*model.SysAuthRuleInfoRes
	res = &system.RoleMenuTreeSelectRes{
		Rules: make([]*model.SysAuthRuleTreeRes, 0),
	}
	list, err = service.SysAuthRule().GetMenuListSearch(ctx, &system.RuleSearchReq{})
	if err != nil {
		return
	}
	res.Rules = service.SysAuthRule().GetMenuListTree(0, list)
	res.DataScope, err = service.SysRole().GetRoleDataScope(ctx, req.RoleId)
	return
}

// RoleDataScope 设置角色数据权限
func (c *roleController) RoleDataScope(ctx context.Context, req *system.DataScopeReq) (res *system.DataScopeRes, err error) {
	err = service.SysRole().RoleDataScope(ctx, req)
	return
}

func (s *roleController) SetRoleUsers(ctx context.Context, req *system.SetRoleUserReq) (res *system.SetRoleUserRes, err error) {
	err = service.SysUser().SetUserRole(ctx, req.RoleId, req.UserIds)
	return
}
