package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var (
	User = userController{}
)

type userController struct {
	BaseController
}

// GetUserMenus 获取用户菜单及按钮权限
func (c *userController) GetUserMenus(ctx context.Context, req *system.UserMenusReq) (res *system.UserMenusRes, err error) {
	var (
		permissions []string
		menuList    []*model.UserMenus
		loginUser   = service.Context().GetLoginUser(ctx)
	)
	userId := loginUser.Id
	menuList, permissions, err = service.SysUser().GetAdminRules(ctx, userId)
	res = &system.UserMenusRes{
		MenuList:    menuList,
		Permissions: permissions,
		UserInfo:    loginUser.LoginUserRes,
	}
	return
}

// List 用户列表
func (c *userController) List(ctx context.Context, req *system.UserSearchReq) (res *system.UserSearchRes, err error) {
	var (
		total    interface{}
		userList []*entity.SysUser
	)
	res = new(system.UserSearchRes)
	req.UserInfo = service.Context().GetLoginUser(ctx)
	total, userList, err = service.SysUser().List(ctx, req)
	if err != nil || total == 0 {
		return
	}
	res.Total = total
	res.UserList, err = service.SysUser().GetUsersRoleDept(ctx, userList)
	return
}

// GetUsersByRoleId 通过角色id获取用户数据
func (c *userController) GetUsersByRoleId(ctx context.Context, req *system.UsersRoleIdReq) (res *system.UsersRoleIdRes, err error) {
	res = new(system.UsersRoleIdRes)
	res.UserList, err = service.SysUser().GetUsersByRoleId(ctx, req.RoleId)
	return
}

// GetParams 获取用户维护相关参数
func (c *userController) GetParams(ctx context.Context, req *system.UserGetParamsReq) (res *system.UserGetParamsRes, err error) {
	res = new(system.UserGetParamsRes)
	res.RoleList, err = service.SysRole().GetRoleList(ctx)
	if err != nil {
		return
	}
	res.Posts, err = service.SysPost().GetUsedPost(ctx)
	if err != nil {
		return
	}
	userId := service.Context().GetUserId(ctx)
	//判断是否超管
	if service.SysUser().IsSupperAdmin(ctx, userId) {
		//自己创建的角色可以被授权
		for _, v := range res.RoleList {
			res.RoleAccess = append(res.RoleAccess, v.Id)
		}
	} else {
		res.RoleAccess, err = service.SysUser().GetAdminRoleIds(ctx, userId, true)
		if err != nil {
			return
		}
		//自己创建的角色可以被授权
		for _, v := range res.RoleList {
			if v.CreatedBy == userId {
				res.RoleAccess = append(res.RoleAccess, v.Id)
			}
		}
	}
	return
}

// Add 添加用户
func (c *userController) Add(ctx context.Context, req *system.UserAddReq) (res *system.UserAddRes, err error) {
	err = service.SysUser().Add(ctx, req)
	return
}

// GetEditUser 获取修改用户信息
func (c *userController) GetEditUser(ctx context.Context, req *system.UserGetEditReq) (res *system.UserGetEditRes, err error) {
	res, err = service.SysUser().GetEditUser(ctx, req.Id)
	return
}

// Edit 修改用户
func (c *userController) Edit(ctx context.Context, req *system.UserEditReq) (res *system.UserEditRes, err error) {
	err = service.SysUser().Edit(ctx, req)
	return
}

// ResetPwd 重置密码
func (c *userController) ResetPwd(ctx context.Context, req *system.UserResetPwdReq) (res *system.UserResetPwdRes, err error) {
	err = service.SysUser().ResetUserPwd(ctx, req)
	return
}

// SetStatus 修改用户状态
func (c *userController) SetStatus(ctx context.Context, req *system.UserStatusReq) (res *system.UserStatusRes, err error) {
	err = service.SysUser().ChangeUserStatus(ctx, req)
	return
}

// Delete 删除用户
func (c *userController) Delete(ctx context.Context, req *system.UserDeleteReq) (res *system.UserDeleteRes, err error) {
	err = service.SysUser().Delete(ctx, req.Ids)
	return
}

// GetUserSelector 获取用户选择器数据
func (c *userController) GetUserSelector(ctx context.Context, req *system.UserSelectorReq) (res *system.UserSelectorRes, err error) {
	res = new(system.UserSelectorRes)
	res.Total, res.UserList, err = service.SysUser().GetUserSelector(ctx, req)
	return
}

// GetByIds 根据id 获取用户信息
func (c *userController) GetByIds(ctx context.Context, req *system.UserByIdsReq) (res *system.UserByIdsRes, err error) {
	res = new(system.UserByIdsRes)
	res.UserList, err = service.SysUser().GetUsers(ctx, gconv.Interfaces(req.Ids))
	return
}
