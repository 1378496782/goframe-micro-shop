/*
* @desc:角色管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/26 15:54
 */

package sysRole

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libWebsocket"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysRole(New())
}

func New() service.ISysRole {
	return &sSysRole{}
}

type sSysRole struct {
}

func (s *sSysRole) GetRoleListSearch(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error) {
	res = new(system.RoleListRes)
	err = g.Try(ctx, func(ctx context.Context) {
		model := dao.SysRole.Ctx(ctx)
		if req.RoleName != "" {
			model = model.Where("a.name like ?", "%"+req.RoleName+"%")
		}
		if req.Status != "" {
			model = model.Where("a.status", gconv.Int(req.Status))
		}
		userId := service.Context().GetUserId(ctx)
		//获取当前用户所属角色ids
		if !service.SysUser().IsSupperAdmin(ctx, userId) {
			var roleIds []uint
			roleIds, err = service.SysUser().GetAdminRoleIds(ctx, userId, true)
			liberr.ErrIsNil(ctx, err)
			if len(roleIds) == 0 {
				return
			}
			model = model.Where("a."+dao.SysRole.Columns().Id+" in(?) OR a.created_by = ?", roleIds, userId)
		}
		model = model.As("a")
		fields := "a.*, count(u.id) user_cnt"
		if service.ToolsGenTable().IsMysql(dao.SysRole.Group()) {
			model = model.LeftJoin("casbin_rule", "b", "b.v1  = a.id ")
			model = model.LeftJoin("sys_user", "u", "CONCAT('u_',u.id) = b.v0 ")
		} else if service.ToolsGenTable().IsDM(dao.SysRole.Group()) {
			fields = "a.id,a.pid,a.status,a.list_order,a.name,a.remark, a.created_at,COUNT(u.id) AS user_cnt"
			model = model.LeftJoin("casbin_rule", "b", "b.v1  = a.id ")
			model = model.LeftJoin("sys_user", "u", "('u_' || u.id) = b.v0 ")
		} else {
			model = model.LeftJoin("casbin_rule", "b", "b.v1  = cast(a.id AS VARCHAR) ")
			model = model.LeftJoin("sys_user", "u", "CONCAT('u_',u.id)  =  b.v0")
		}
		model = model.Group("a.id")
		err = model.Order("list_order asc,id asc").Fields(fields).Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

// GetRoleList 获取角色列表
func (s *sSysRole) GetRoleList(ctx context.Context) (list []*entity.SysRole, err error) {
	cache := commonService.Cache()
	//从缓存获取
	iList := cache.GetOrSetFuncLock(ctx, consts.CacheSysRole, s.getRoleListFromDb, 0, consts.CacheSysAuthTag)
	if !iList.IsEmpty() {
		err = gconv.Struct(iList, &list)
	}
	return
}

// 从数据库获取所有角色
func (s *sSysRole) getRoleListFromDb(ctx context.Context) (value interface{}, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var v []*entity.SysRole
		//从数据库获取
		err = dao.SysRole.Ctx(ctx).
			Order(dao.SysRole.Columns().ListOrder + " asc," + dao.SysRole.Columns().Id + " asc").
			Scan(&v)
		liberr.ErrIsNil(ctx, err, "获取角色数据失败")
		value = v
	})
	return
}

// AddRoleRule 添加角色权限
func (s *sSysRole) AddRoleRule(ctx context.Context, ruleIds []uint, roleId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		ruleIdsStr := gconv.Strings(ruleIds)
		rules := make([][]string, len(ruleIdsStr))
		for k, v := range ruleIdsStr {
			rules[k] = []string{gconv.String(roleId), v, "All"}
		}
		if len(rules) > 0 {
			_, err = enforcer.AddPolicies(rules)
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// DelRoleRule 删除角色权限
func (s *sSysRole) DelRoleRule(ctx context.Context, roleId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		_, err = enforcer.RemoveFilteredPolicy(0, gconv.String(roleId))
		liberr.ErrIsNil(ctx, e)
	})
	return
}

func (s *sSysRole) AddRole(ctx context.Context, req *system.RoleAddReq) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			req.CreatedBy = service.Context().GetUserId(ctx)
			roleId, e := dao.SysRole.Ctx(ctx).TX(tx).InsertAndGetId(do.SysRole{
				Pid:           req.Pid,
				Status:        req.Status,
				ListOrder:     req.ListOrder,
				Name:          req.Name,
				Remark:        req.Remark,
				CreatedBy:     req.CreatedBy,
				EffectiveTime: req.EffectiveTimeInfo,
			})
			liberr.ErrIsNil(ctx, e, "添加角色失败")
			//过滤ruleIds 把没有权限的过滤掉
			req.MenuIds, err = s.filterAccessRuleIds(ctx, req.MenuIds)
			liberr.ErrIsNil(ctx, err)
			//添加角色权限
			e = s.AddRoleRule(ctx, req.MenuIds, roleId)
			liberr.ErrIsNil(ctx, e)
			//清除缓存
			commonService.Cache().Remove(ctx, consts.CacheSysRole)
		})
		return err
	})
	return
}

func (s *sSysRole) Get(ctx context.Context, id uint) (res *model.RoleInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//判断是否具有此角色的权限
		if !s.hasManageAccess(ctx, id, true) {
			liberr.ErrIsNil(ctx, errors.New("没有查看这个角色的权限"))
		}
		res = new(model.RoleInfoRes)
		err = dao.SysRole.Ctx(ctx).WherePri(id).Scan(&res.SysRole)
		liberr.ErrIsNil(ctx, err, "获取角色信息失败")
		err = gconv.Struct(res.SysRole.EffectiveTime, &res.EffectiveTimeInfo)
		if err != nil {
			res.EffectiveTimeInfo = new(model.EffectiveTimeInfo)
		}
	})
	return
}

// GetFilteredNamedPolicy 获取角色关联的菜单规则
func (s *sSysRole) GetFilteredNamedPolicy(ctx context.Context, id uint) (gpSlice []int, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		gp := enforcer.GetFilteredNamedPolicy("p", 0, gconv.String(id))
		gpSlice = make([]int, len(gp))
		for k, v := range gp {
			gpSlice[k] = gconv.Int(v[1])
		}
	})
	return
}

func (s *sSysRole) hasManageAccess(ctx context.Context, roleId uint, includeChildren ...bool) bool {
	currentUserId := service.Context().GetUserId(ctx)
	if !service.SysUser().IsSupperAdmin(ctx, currentUserId) {
		var (
			roleIds   []uint
			hasAccess bool
			err       error
			list      []*entity.SysRole
		)
		list, err = s.GetRoleList(ctx)
		if err != nil {
			g.Log().Error(ctx, err)
			return false
		}
		for _, v := range list {
			//判断是否当前用户所建角色
			if roleId == v.Id && v.CreatedBy == currentUserId {
				return true
			}
		}
		roleIds, err = service.SysUser().GetAdminRoleIds(ctx, service.Context().GetUserId(ctx), includeChildren...)
		if err != nil {
			g.Log().Error(ctx, err)
			return false
		}
		if len(roleIds) > 0 {
			for _, v := range roleIds {
				if v == roleId {
					hasAccess = true
					break
				}
			}
		}
		return hasAccess
	}
	return true
}

// EditRole 修改角色
func (s *sSysRole) EditRole(ctx context.Context, req *system.RoleEditReq) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//判断是否具有修改此角色的权限
			if !s.hasManageAccess(ctx, gconv.Uint(req.Id), true) {
				liberr.ErrIsNil(ctx, errors.New("没有修改这个角色的权限"))
			}
			_, e := dao.SysRole.Ctx(ctx).TX(tx).WherePri(req.Id).Data(&do.SysRole{
				Pid:           req.Pid,
				Status:        req.Status,
				ListOrder:     req.ListOrder,
				Name:          req.Name,
				Remark:        req.Remark,
				EffectiveTime: gconv.String(req.EffectiveTimeInfo),
			}).Update()
			liberr.ErrIsNil(ctx, e, "修改角色失败")
			//过滤ruleIds 把没有权限的过滤掉
			req.MenuIds, err = s.filterAccessRuleIds(ctx, req.MenuIds)
			liberr.ErrIsNil(ctx, err)
			//删除角色权限
			e = s.DelRoleRule(ctx, req.Id)
			liberr.ErrIsNil(ctx, e)
			//添加角色权限
			e = s.AddRoleRule(ctx, req.MenuIds, req.Id)
			liberr.ErrIsNil(ctx, e)
			//清除缓存
			commonService.Cache().Remove(ctx, consts.CacheSysRole)
			//通知刷新token
			s.refreshToken(ctx, req.Id)
		})
		return err
	})
	return
}

// 从给定的menuIds中过滤掉用户没有操作权限的菜单id
func (s *sSysRole) filterAccessRuleIds(ctx context.Context, menuIds []uint) (newRuleIds []uint, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//若不是超管，过滤ruleIds 把没有权限的过滤掉
		if !service.SysUser().IsSupperAdmin(ctx, service.Context().GetUserId(ctx)) {
			var (
				userRoleIds []uint
				accessMenus *garray.Array
			)
			userRoleIds, err = service.SysUser().GetAdminRoleIds(ctx, service.Context().GetUserId(ctx))
			liberr.ErrIsNil(ctx, err)
			accessMenus, err = service.SysUser().GetAdminMenusIdsByRoleIds(ctx, userRoleIds)
			for _, v := range menuIds {
				if accessMenus.Contains(v) {
					newRuleIds = append(newRuleIds, v)
				}
			}
		} else {
			newRuleIds = menuIds
		}
	})
	return
}

// DeleteByIds 删除角色
func (s *sSysRole) DeleteByIds(ctx context.Context, ids []int64) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			for _, id := range ids {
				//判断是否有删除该角色的权限
				if !s.hasManageAccess(ctx, gconv.Uint(id)) {
					liberr.ErrIsNil(ctx, errors.New("没有删除这个角色的权限"))
				}
			}
			_, err = dao.SysRole.Ctx(ctx).TX(tx).Where(dao.SysRole.Columns().Id+" in(?)", ids).Delete()
			liberr.ErrIsNil(ctx, err, "删除角色失败")
			//删除角色权限
			for _, v := range ids {
				err = s.DelRoleRule(ctx, v)
				liberr.ErrIsNil(ctx, err)
			}
			//清除缓存
			commonService.Cache().Remove(ctx, consts.CacheSysRole)
		})
		return err
	})
	return
}

func (s *sSysRole) RoleDeptTreeSelect(ctx context.Context) (res *system.RoleDeptTreeSelectRes, err error) {
	res = new(system.RoleDeptTreeSelectRes)
	err = g.Try(ctx, func(ctx context.Context) {
		list, err := service.SysDept().GetList(ctx, &system.DeptSearchReq{
			Status:  "1",
			ShowAll: true,
		})
		liberr.ErrIsNil(ctx, err)
		dList := service.SysDept().GetListTree(0, list)
		res.Depts = dList
	})
	return
}

func (s *sSysRole) GetRoleDataScope(ctx context.Context, roleId uint) (data []*model.ScopeAuthData, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysRoleScope.Ctx(ctx).Where("role_id", roleId).Scan(&data)
		liberr.ErrIsNil(ctx, err, "获取角色数据权限失败")
	})
	return
}

func (s *sSysRole) GetRoleMenuScope(ctx context.Context, roleIds []uint, menuId uint) (data []*model.ScopeAuthData, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysRoleScope.Ctx(ctx).WhereIn("role_id", roleIds).
			Where("menu_id", menuId).
			Scan(&data)
		liberr.ErrIsNil(ctx, err, "获取角色数据权限失败")
	})
	return
}

// RoleDataScope 设置角色数据权限
func (s *sSysRole) RoleDataScope(ctx context.Context, req *system.DataScopeReq) error {
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err := g.Try(ctx, func(ctx context.Context) {
			data := make([]do.SysRoleScope, 0, len(req.AuthData))
			for _, v := range req.AuthData {
				if v.MenuId != 0 && v.Scope != 0 {
					data = append(data, do.SysRoleScope{
						RoleId:    req.RoleId,
						MenuId:    v.MenuId,
						DataScope: v.Scope,
						DeptIds:   gconv.String(v.DeptIds),
					})
				}
			}
			//清除旧权限
			_, err := dao.SysRoleScope.Ctx(ctx).Where(dao.SysRoleScope.Columns().RoleId, req.RoleId).
				Delete()
			liberr.ErrIsNil(ctx, err, "清除旧权限信息失败")
			if len(data) > 0 {
				_, err = dao.SysRoleScope.Ctx(ctx).Data(data).Insert()
				liberr.ErrIsNil(ctx, err, "设置权限信息失败")
			}
			//通知刷新token
			s.refreshToken(ctx, gconv.Int64(req.RoleId))
		})
		return err
	})
	return err
}

func (s *sSysRole) FindSonByParentId(roleList []*entity.SysRole, id uint) []*entity.SysRole {
	children := make([]*entity.SysRole, 0, len(roleList))
	for _, v := range roleList {
		if v.Pid == id {
			children = append(children, v)
			fChildren := s.FindSonByParentId(roleList, v.Id)
			children = append(children, fChildren...)
		}
	}
	return children
}

func (s *sSysRole) FindSonIdsByParentId(roleList []*entity.SysRole, id uint) []uint {
	children := make([]uint, 0, len(roleList))
	for _, v := range roleList {
		if v.Pid == id {
			children = append(children, v.Id)
			fChildren := s.FindSonIdsByParentId(roleList, v.Id)
			children = append(children, fChildren...)
		}
	}
	return children
}

// 刷新角色下用户token
func (s *sSysRole) refreshToken(ctx context.Context, roleId int64) {
	_ = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		userRoleIds := enforcer.GetFilteredGroupingPolicy(1, gconv.String(roleId))
		for _, v := range userRoleIds {
			userId := gstr.Split(v[0], "_")[1]
			//通知用户更新token
			libWebsocket.SendToUser(gconv.Uint64(userId), &libWebsocket.WResponse{
				Event: consts.WebsocketTypeTokenUpdated,
				Data:  nil,
			})
		}
	})
}
