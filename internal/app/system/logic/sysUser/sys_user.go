/*
* @desc:用户处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/23 15:08
 */

package sysUser

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/library/libWebsocket"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/mssola/user_agent"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysUser(New())
}

type sSysUser struct {
	casBinUserPrefix string //CasBin 用户id前缀
}

func New() service.ISysUser {
	return &sSysUser{
		casBinUserPrefix: "u_",
	}
}

func (s *sSysUser) GetCasBinUserPrefix() string {
	return s.casBinUserPrefix
}

// IsSupperAdmin 判断用户是否超管
func (s *sSysUser) IsSupperAdmin(ctx context.Context, userId uint64) bool {
	superAdminIds := s.NotCheckAuthAdminIds(ctx)
	if superAdminIds.Contains(userId) {
		return true
	}
	return false
}

func (s *sSysUser) NotCheckAuthAdminIds(ctx context.Context) *gset.Set {
	ids := g.Cfg().MustGet(ctx, "system.notCheckAuthAdminIds").Uint64s()
	if !g.IsNil(ids) {
		return gset.NewFrom(ids)
	}
	return gset.New()
}

func (s *sSysUser) GetAdminUserByUsernamePassword(ctx context.Context, req *system.UserLoginReq) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user, err = s.GetUserByUsername(ctx, req.Username)
		liberr.ErrIsNil(ctx, err)
		liberr.ValueIsNil(user, "账号密码错误")
		//验证密码
		if libUtils.EncryptPassword(req.Password, user.UserSalt) != user.UserPassword {
			liberr.ErrIsNil(ctx, gerror.New("账号密码错误"))
		}
		//账号状态
		if user.UserStatus == 0 {
			liberr.ErrIsNil(ctx, gerror.New("账号已被冻结"))
		}
	})
	return
}

// GetUserByUsername 通过用户名获取用户信息
func (s *sSysUser) GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user = &model.LoginUserRes{}
		err = dao.SysUser.Ctx(ctx).Fields(user).Where(dao.SysUser.Columns().UserName, userName).Scan(user)
		liberr.ErrIsNil(ctx, err, "账号密码错误")
	})
	return
}

func (s *sSysUser) GetUserByPhone(ctx context.Context, phone string) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysUser.Ctx(ctx).Fields(user).Where(dao.SysUser.Columns().Mobile, phone).
			Scan(&user)
		liberr.ErrIsNil(ctx, err, "登录失败，用户信息不存在")
	})
	return
}

// GetUserById 通过用户名获取用户信息
func (s *sSysUser) GetUserById(ctx context.Context, id uint64) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysUser.Ctx(ctx).Fields(user).WherePri(id).Scan(&user)
		liberr.ErrIsNil(ctx, err, "获取用户信息失败")
	})
	return
}

// LoginLog 记录登录日志
func (s *sSysUser) LoginLog(ctx context.Context, params *model.LoginLogParams) {
	ua := user_agent.New(params.UserAgent)
	browser, _ := ua.Browser()
	loginData := &do.SysLoginLog{
		LoginName:     params.Username,
		Ipaddr:        params.Ip,
		LoginLocation: libUtils.GetCityByIp(params.Ip),
		Browser:       browser,
		Os:            ua.OS(),
		Status:        params.Status,
		Msg:           params.Msg,
		LoginTime:     gtime.Now(),
		Module:        params.Module,
	}
	_, err := dao.SysLoginLog.Ctx(ctx).Insert(loginData)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func (s *sSysUser) UpdateLoginInfo(ctx context.Context, id uint64, ip string, openId ...string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		data := g.Map{
			dao.SysUser.Columns().LastLoginIp:   ip,
			dao.SysUser.Columns().LastLoginTime: gtime.Now(),
		}
		if len(openId) > 0 && openId[0] != "" {
			data[dao.SysUser.Columns().OpenId] = openId[0]
		}
		_, err = dao.SysUser.Ctx(ctx).WherePri(id).Unscoped().Update(data)
		liberr.ErrIsNil(ctx, err, "更新用户登录信息失败")
	})
	return
}

// GetAdminRules 获取用户菜单数据
func (s *sSysUser) GetAdminRules(ctx context.Context, userId uint64) (menuList []*model.UserMenus, permissions []string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//是否超管
		isSuperAdmin := s.IsSupperAdmin(ctx, userId)
		//获取用户菜单数据
		allRoles, err := service.SysRole().GetRoleList(ctx)
		liberr.ErrIsNil(ctx, err)
		roles, err := s.GetAdminRole(ctx, userId, allRoles)
		liberr.ErrIsNil(ctx, err)
		name := make([]string, len(roles))
		roleIds := make([]uint, len(roles))
		for k, v := range roles {
			name[k] = v.Name
			roleIds[k] = v.Id
		}
		//获取菜单信息
		if isSuperAdmin {
			//超管获取所有菜单
			permissions = []string{"*/*/*"}
			menuList, err = s.GetAllMenus(ctx)
			liberr.ErrIsNil(ctx, err)
		} else {
			menuList, err = s.GetAdminMenusByRoleIds(ctx, roleIds)
			liberr.ErrIsNil(ctx, err)
			permissions, err = s.GetPermissions(ctx, roleIds)
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// GetAdminRole 获取用户角色
func (s *sSysUser) GetAdminRole(ctx context.Context, userId uint64, allRoleList []*entity.SysRole) (roles []*entity.SysRole, err error) {
	var roleIds []uint
	roleIds, err = s.GetAdminRoleIds(ctx, userId)
	if err != nil {
		return
	}
	roles = make([]*entity.SysRole, 0, len(allRoleList))
	for _, v := range allRoleList {
		for _, id := range roleIds {
			if id == v.Id {
				roles = append(roles, v)
			}
		}
		if len(roles) == len(roleIds) {
			break
		}
	}
	return
}

// GetAdminRoleIds 获取用户角色ids
func (s *sSysUser) GetAdminRoleIds(ctx context.Context, userId uint64, includeChildren ...bool) (roleIds []uint, err error) {
	enforcer, e := commonService.CasbinEnforcer(ctx)
	if e != nil {
		err = e
		return
	}
	//查询关联角色规则
	groupPolicy := enforcer.GetFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", s.casBinUserPrefix, userId))
	if len(groupPolicy) > 0 {
		roleIds = make([]uint, len(groupPolicy))
		//得到角色id的切片
		for k, v := range groupPolicy {
			roleIds[k] = gconv.Uint(v[1])
		}
	}
	if len(includeChildren) > 0 && includeChildren[0] {
		//获取子级
		var allRoles []*entity.SysRole
		allRoles, err = service.SysRole().GetRoleList(ctx)
		if err != nil {
			return
		}
		childrenIds := make([]uint, 0, 100)
		for _, roleId := range roleIds {
			childrenIds = append(childrenIds, service.SysRole().FindSonIdsByParentId(allRoles, roleId)...)
		}
		//合并去重
		roleIds = append(roleIds, childrenIds...)
	}
	roleIds = libUtils.SliceUnique(roleIds)
	return
}

func (s *sSysUser) GetAllMenus(ctx context.Context) (menus []*model.UserMenus, err error) {
	//获取所有开启的菜单
	var allMenus []*model.SysAuthRuleInfoRes
	allMenus, err = service.SysAuthRule().GetIsMenuList(ctx)
	if err != nil {
		return
	}
	menus = make([]*model.UserMenus, len(allMenus))
	for k, v := range allMenus {
		var menu *model.UserMenu
		menu = s.setMenuData(menu, v)
		menus[k] = &model.UserMenus{UserMenu: menu}
	}
	menus = s.GetMenusTree(menus, 0)
	return
}

func (s *sSysUser) GetAdminMenusIdsByRoleIds(ctx context.Context, roleIds []uint) (menuIds *garray.Array, err error) {
	//获取角色对应的菜单id
	menuIds = garray.New()
	err = g.Try(ctx, func(ctx context.Context) {
		if s.IsSupperAdmin(ctx, service.Context().GetUserId(ctx)) {
			var menus []*model.SysAuthRuleInfoRes
			menus, err = service.SysAuthRule().GetMenuList(ctx)
			liberr.ErrIsNil(ctx, err)
			for _, m := range menus {
				menuIds.Append(m.Id)
			}
			return
		}
		enforcer, e := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		for _, roleId := range roleIds {
			//查询当前权限
			gp := enforcer.GetFilteredPolicy(0, gconv.String(roleId))
			for _, p := range gp {
				menuIds.Append(gconv.Uint(p[1]))
			}
		}
	})
	return
}

func (s *sSysUser) GetAdminMenusByRoleIds(ctx context.Context, roleIds []uint) (menus []*model.UserMenus, err error) {
	//获取角色对应的菜单id
	err = g.Try(ctx, func(ctx context.Context) {
		var (
			menuArr    *garray.Array
			menuIds    = map[uint]uint{}
			menuNameId = map[string]*model.SysAuthRuleInfoRes{}
		)
		menuArr, err = s.GetAdminMenusIdsByRoleIds(ctx, roleIds)
		liberr.ErrIsNil(ctx, err)
		menuArr.Iterator(func(k int, v interface{}) bool {
			mp := gconv.Uint(v)
			menuIds[mp] = mp
			return true
		})
		//获取所有开启的菜单
		allMenus, err := service.SysAuthRule().GetIsMenuList(ctx)
		liberr.ErrIsNil(ctx, err)
		menus = make([]*model.UserMenus, 0, len(allMenus))
		for _, v := range allMenus {
			urlMap, _ := gurl.ParseURL(v.Name, -1)
			if query, ok := urlMap["query"]; ok && query != "" {
				menuNameId[urlMap["path"]+"#query"] = v
			}
		}
		for _, v := range allMenus {
			var roleMenu *model.UserMenu
			if _, ok := menuIds[v.Id]; gstr.Equal(v.Condition, "nocheck") || ok {
				roleMenu = s.setMenuData(roleMenu, v)
				menus = append(menus, &model.UserMenus{UserMenu: roleMenu})
			}
			if info, ok := menuNameId[v.Name+"#query"]; ok && roleMenu != nil {
				roleMenu = s.setMenuData(roleMenu, info)
				menus = append(menus, &model.UserMenus{UserMenu: roleMenu})
			}
		}
		menus = s.GetMenusTree(menus, 0)
	})
	return
}

func (s *sSysUser) GetMenusTree(menus []*model.UserMenus, pid uint) []*model.UserMenus {
	returnList := make([]*model.UserMenus, 0, len(menus))
	for _, menu := range menus {
		if menu.Pid == pid {
			menu.Children = s.GetMenusTree(menus, menu.Id)
			returnList = append(returnList, menu)
		}
	}
	return returnList
}

func (s *sSysUser) setMenuData(menu *model.UserMenu, entity *model.SysAuthRuleInfoRes) *model.UserMenu {
	menu = &model.UserMenu{
		Id:        entity.Id,
		Pid:       entity.Pid,
		Name:      gstr.CaseCamelLower(gstr.Replace(entity.Name, "/", "_")),
		Component: entity.Component,
		Path:      entity.Path,
		MenuMeta: &model.MenuMeta{
			Icon:        entity.Icon,
			Title:       entity.Title,
			IsLink:      "",
			IsHide:      entity.IsHide == 1,
			IsKeepAlive: entity.IsCached == 1,
			IsAffix:     entity.IsAffix == 1,
			IsIframe:    entity.IsIframe == 1,
		},
	}
	if menu.MenuMeta.IsIframe || entity.IsLink == 1 {
		menu.MenuMeta.IsLink = entity.LinkUrl
	}
	return menu
}

func (s *sSysUser) GetPermissions(ctx context.Context, roleIds []uint) (userButtons []string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//获取角色对应的菜单id
		enforcer, err := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, err)
		menuIds := map[int64]int64{}
		for _, roleId := range roleIds {
			//查询当前权限
			gp := enforcer.GetFilteredPolicy(0, gconv.String(roleId))
			for _, p := range gp {
				mid := gconv.Int64(p[1])
				menuIds[mid] = mid
			}
		}
		//获取所有开启的按钮
		allButtons, err := service.SysAuthRule().GetIsButtonList(ctx)
		liberr.ErrIsNil(ctx, err)
		userButtons = make([]string, 0, len(allButtons))
		for _, button := range allButtons {
			if _, ok := menuIds[gconv.Int64(button.Id)]; gstr.Equal(button.Condition, "nocheck") || ok {
				userButtons = append(userButtons, button.Name)
			}
		}
	})
	return
}

func fieldsEx() []string {
	s := dao.SysUser.Columns()
	return []string{
		s.Id,
		s.UserName,
		s.Mobile,
		s.UserNickname,
		s.Birthday,
		s.UserStatus,
		s.UserEmail,
		s.Sex,
		s.Avatar,
		s.DeptId,
		s.Remark,
		s.IsAdmin,
		s.Address,
		s.Describe,
		s.LastLoginIp,
		s.LastLoginTime,
		s.CreatedAt,
		s.UpdatedAt,
		s.DeletedAt,
	}
}

func inSlice(field string, excludes []string) bool {
	for _, exclude := range excludes {
		if exclude == field {
			return true
		}
	}
	return false
}

// 根据用户ID 获取用户
func (s *sSysUser) GetByIdsUser(ctx context.Context, req *system.UserByIdsReq) (total interface{}, userList []*entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysUser.Ctx(ctx)
		if req.Ids != nil {
			m = m.Where("id in (?)", req.Ids)
		}
		//dao.SysUserPost.Columns().UserId+" in (?)", ids
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取用户数据失败")

		// 排除password字段
		m = m.Fields(fieldsEx())

		err = m.Page(req.PageNum, req.PageSize).Order("id asc").Scan(&userList)
		liberr.ErrIsNil(ctx, err, "获取用户列表失败")
	})
	return
}

// List 用户列表
func (s *sSysUser) List(ctx context.Context, req *system.UserSearchReq) (total interface{}, userList []*entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysUser.Ctx(ctx)
		if req.KeyWords != "" {
			keyWords := "%" + req.KeyWords + "%"
			m = m.Where("user_name like ? or  user_nickname like ?", keyWords, keyWords)
		}
		if req.DeptId != "" {
			deptIds, e := s.getSearchDeptIds(ctx, gconv.Uint64(req.DeptId))
			liberr.ErrIsNil(ctx, e)
			m = m.Where("dept_id in (?)", deptIds)
		}
		if req.Status != "" {
			m = m.Where("user_status", gconv.Int(req.Status))
		}
		if req.Mobile != "" {
			m = m.Where("mobile like ?", "%"+req.Mobile+"%")
		}
		if len(req.DateRange) > 0 {
			m = m.Where("created_at >=? AND created_at <=?", req.DateRange[0], req.DateRange[1])
		}

		if req.RoleId > 0 {
			m = m.LeftJoin("casbin_rule", "b", "b.v0 = CONCAT('u_',sys_user.id )")
			m = m.Where("v1 = ? and SUBSTR(v0,1,2) = 'u_'", req.RoleId)
		}
		//判断权限，普通管理只能按数据权限查看
		if !s.AccessRule(ctx, req.UserInfo.Id, "api/v1/system/user/all") {
			where := s.GetAuthDeptWhere(
				ctx,
				req.UserInfo,
				"sys_user", "dept_id", "id",
			)
			if len(where) > 0 {
				m = m.Where(where)
			}
			m = m.WhereNotIn(dao.SysUser.Columns().Id, s.NotCheckAuthAdminIds(ctx).Slice())
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取用户数据失败")

		// 排除password字段
		if req.RoleId > 0 {
			m = m.Fields(fieldsEx())
		} else {
			m = m.FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt)
		}

		err = m.Page(req.PageNum, req.PageSize).Order("id asc").Scan(&userList)
		liberr.ErrIsNil(ctx, err, "获取用户列表失败")
	})
	return
}

func (s *sSysUser) GetUsersByRoleId(ctx context.Context, roleId uint) (users []*model.SysUserRoleDeptRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysUser.Ctx(ctx)
		m = m.As("a").LeftJoin("casbin_rule", "b", "b.v0 = CONCAT('u_',a.id )").
			Where("v1 = ? and SUBSTR(v0,1,2) = 'u_'", roleId)
		err = m.Order("id asc").Scan(&users)
		liberr.ErrIsNil(ctx, err, "获取用户数据失败")
	})
	return
}

// GetUserSelector 获取用户选择器数据
func (s *sSysUser) GetUserSelector(ctx context.Context, req *system.UserSelectorReq) (total interface{}, userList []*model.SysUserSimpleRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysUser.Ctx(ctx)
		if req.KeyWords != "" {
			keyWords := "%" + req.KeyWords + "%"
			m = m.Where("user_name like ? or  user_nickname like ?", keyWords, keyWords)
		}
		if req.DeptId != "" {
			deptIds, e := s.getSearchDeptIds(ctx, gconv.Uint64(req.DeptId))
			liberr.ErrIsNil(ctx, e)
			m = m.Where("dept_id in (?)", deptIds)
		}
		if req.Status != "" {
			m = m.Where("user_status", gconv.Int(req.Status))
		}
		if req.Mobile != "" {
			m = m.Where("mobile like ?", "%"+req.Mobile+"%")
		}
		if len(req.DateRange) > 0 {
			m = m.Where("created_at >=? AND created_at <=?", req.DateRange[0], req.DateRange[1])
		}

		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取用户数据失败")

		// 排除password字段
		if req.RoleId > 0 {
			m = m.Fields(fieldsEx())
		} else {
			m = m.FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt)
		}

		err = m.Page(req.PageNum, req.PageSize).Order("id asc").Scan(&userList)
		liberr.ErrIsNil(ctx, err, "获取用户列表失败")
	})
	return
}

// GetUsersRoleDept 获取多个用户角色 部门信息
func (s *sSysUser) GetUsersRoleDept(ctx context.Context, userList []*entity.SysUser) (users []*model.SysUserRoleDeptRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		allRoles, e := service.SysRole().GetRoleList(ctx)
		liberr.ErrIsNil(ctx, e)
		depts, e := service.SysDept().GetFromCache(ctx)
		liberr.ErrIsNil(ctx, e)
		users = make([]*model.SysUserRoleDeptRes, len(userList))
		for k, u := range userList {
			var dept *entity.SysDept
			users[k] = &model.SysUserRoleDeptRes{
				SysUser: u,
			}
			for _, d := range depts {
				if u.DeptId == uint64(d.DeptId) {
					dept = d
				}
			}
			users[k].Dept = dept
			roles, e := s.GetAdminRole(ctx, u.Id, allRoles)
			liberr.ErrIsNil(ctx, e)
			for _, r := range roles {
				users[k].RoleInfo = append(users[k].RoleInfo, &model.SysUserRoleInfoRes{RoleId: r.Id, Name: r.Name})
			}
		}
	})
	return
}

func (s *sSysUser) getSearchDeptIds(ctx context.Context, deptId uint64) (deptIds []uint64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		deptAll, e := service.SysDept().GetFromCache(ctx)
		liberr.ErrIsNil(ctx, e)
		deptWithChildren := service.SysDept().FindSonByParentId(deptAll, deptId)
		deptIds = make([]uint64, len(deptWithChildren))
		for k, v := range deptWithChildren {
			deptIds[k] = v.DeptId
		}
		deptIds = append(deptIds, deptId)
	})
	return
}

// 过滤用户可操作的角色
func (s *sSysUser) filterRoleIds(ctx context.Context, roleIds []uint, userId uint64, includeChildren ...bool) (newRoleIds []uint, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var (
			accessRoleList []uint
			roleList       []*entity.SysRole
		)
		accessRoleList, err = service.SysUser().GetAdminRoleIds(ctx, userId, includeChildren...)
		liberr.ErrIsNil(ctx, err)
		roleList, err = service.SysRole().GetRoleList(ctx)
		liberr.ErrIsNil(ctx, err)
		//子角色也要能够被授权
		sonIds := make([]uint, 0, 10)
		for _, v := range accessRoleList {
			sonIds = append(sonIds, service.SysRole().FindSonIdsByParentId(roleList, v)...)
		}
		accessRoleList = append(accessRoleList, sonIds...)
		//自己创建的角色可以被授权
		for _, v := range roleList {
			if v.CreatedBy == userId {
				accessRoleList = append(accessRoleList, v.Id)
			}
		}
		//去重accessRoleList
		accessRoleList = gconv.Uints(garray.NewArrayFrom(gconv.Interfaces(accessRoleList)).Unique().Slice())
		for _, r := range roleIds {
			for _, a := range accessRoleList {
				if r == a {
					newRoleIds = append(newRoleIds, r)
					break
				}
			}
		}
	})
	return
}

func (s *sSysUser) Add(ctx context.Context, req *system.UserAddReq) (err error) {
	err = s.UserNameOrMobileExists(ctx, req.UserName, req.Mobile)
	if err != nil {
		return
	}
	req.UserSalt = grand.S(10)
	req.Password = libUtils.EncryptPassword(req.Password, req.UserSalt)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			userId, e := dao.SysUser.Ctx(ctx).TX(tx).InsertAndGetId(do.SysUser{
				UserName:     req.UserName,
				Mobile:       req.Mobile,
				UserNickname: req.NickName,
				UserPassword: req.Password,
				UserSalt:     req.UserSalt,
				UserStatus:   req.Status,
				UserEmail:    req.Email,
				Sex:          req.Sex,
				DeptId:       req.DeptId,
				Remark:       req.Remark,
				IsAdmin:      req.IsAdmin,
			})
			liberr.ErrIsNil(ctx, e, "添加用户失败")
			//不是超管过滤提交角色数据
			if !service.SysUser().IsSupperAdmin(ctx, service.Context().GetUserId(ctx)) {
				req.RoleIds, err = s.filterRoleIds(ctx, req.RoleIds, service.Context().GetUserId(ctx))
				liberr.ErrIsNil(ctx, err)
			}
			e = s.addUserRole(ctx, req.RoleIds, userId)
			liberr.ErrIsNil(ctx, e, "设置用户权限失败")
			e = s.AddUserPost(ctx, tx, req.PostIds, userId)
			liberr.ErrIsNil(ctx, e)
		})
		return err
	})
	return
}

func (s *sSysUser) Edit(ctx context.Context, req *system.UserEditReq) (err error) {
	err = s.UserNameOrMobileExists(ctx, "", req.Mobile, req.UserId)
	if err != nil {
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysUser.Ctx(ctx).TX(tx).WherePri(req.UserId).Update(do.SysUser{
				Mobile:       req.Mobile,
				UserNickname: req.NickName,
				UserStatus:   req.Status,
				UserEmail:    req.Email,
				Sex:          req.Sex,
				DeptId:       req.DeptId,
				Remark:       req.Remark,
				IsAdmin:      req.IsAdmin,
			})
			liberr.ErrIsNil(ctx, err, "修改用户信息失败")
			if !service.SysUser().IsSupperAdmin(ctx, service.Context().GetUserId(ctx)) {
				req.RoleIds, err = s.filterRoleIds(ctx, req.RoleIds, service.Context().GetUserId(ctx))
				liberr.ErrIsNil(ctx, err)
			}
			//设置用户所属角色信息
			err = s.EditUserRole(ctx, req.RoleIds, req.UserId)
			liberr.ErrIsNil(ctx, err, "设置用户权限失败")
			err = s.AddUserPost(ctx, tx, req.PostIds, req.UserId)
			liberr.ErrIsNil(ctx, err)
			//通知用户更新token
			libWebsocket.SendToUser(gconv.Uint64(req.UserId), &libWebsocket.WResponse{
				Event: consts.WebsocketTypeTokenUpdated,
				Data:  nil,
			})
		})
		return err
	})
	return
}

// AddUserPost 添加用户岗位信息
func (s *sSysUser) AddUserPost(ctx context.Context, tx gdb.TX, postIds []int64, userId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//删除旧岗位信息
		_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Where(dao.SysUserPost.Columns().UserId, userId).Delete()
		liberr.ErrIsNil(ctx, err, "设置用户岗位失败")
		if len(postIds) == 0 {
			return
		}
		//添加用户岗位信息
		data := g.List{}
		for _, v := range postIds {
			data = append(data, g.Map{
				dao.SysUserPost.Columns().UserId: userId,
				dao.SysUserPost.Columns().PostId: v,
			})
		}
		_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Data(data).Insert()
		liberr.ErrIsNil(ctx, err, "设置用户岗位失败")
	})
	return
}

// AddUserRole 添加用户角色信息
func (s *sSysUser) addUserRole(ctx context.Context, roleIds []uint, userId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)
		for _, v := range roleIds {
			_, e = enforcer.AddGroupingPolicy(fmt.Sprintf("%s%d", s.casBinUserPrefix, userId), gconv.String(v))
			liberr.ErrIsNil(ctx, e)
		}
	})
	return
}

// EditUserRole 修改用户角色信息
func (s *sSysUser) EditUserRole(ctx context.Context, roleIds []uint, userId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)

		//删除用户旧角色信息
		_, err = enforcer.RemoveFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", s.casBinUserPrefix, userId))
		liberr.ErrIsNil(ctx, err)
		for _, v := range roleIds {
			_, err = enforcer.AddGroupingPolicy(fmt.Sprintf("%s%d", s.casBinUserPrefix, userId), gconv.String(v))
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// SetUserRole 设置用户角色
func (s *sSysUser) SetUserRole(ctx context.Context, roleId uint, userIds []uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, e)

		//删除用户旧角色信息
		_, err = enforcer.RemoveFilteredGroupingPolicy(1, fmt.Sprintf("%d", roleId))
		liberr.ErrIsNil(ctx, err)
		for _, v := range userIds {
			_, err = enforcer.AddGroupingPolicy(fmt.Sprintf("%s%d", s.casBinUserPrefix, v), gconv.String(roleId))
			liberr.ErrIsNil(ctx, err)
			//通知用户更新token
			libWebsocket.SendToUser(v, &libWebsocket.WResponse{
				Event: consts.WebsocketTypeTokenUpdated,
				Data:  nil,
			})
		}
	})
	return
}

func (s *sSysUser) UserNameOrMobileExists(ctx context.Context, userName, mobile string, id ...int64) error {
	user := (*entity.SysUser)(nil)
	err := g.Try(ctx, func(ctx context.Context) {
		m := dao.SysUser.Ctx(ctx)
		if len(id) > 0 {
			m = m.Where(dao.SysUser.Columns().Id+" != ", id)
		}
		m = m.Where(fmt.Sprintf("%s='%s' OR %s='%s'",
			dao.SysUser.Columns().UserName,
			userName,
			dao.SysUser.Columns().Mobile,
			mobile))
		err := m.Limit(1).Scan(&user)
		liberr.ErrIsNil(ctx, err, "获取用户信息失败")
		if user == nil {
			return
		}
		if user.UserName == userName {
			liberr.ErrIsNil(ctx, gerror.New("用户名已存在"))
		}
		if user.Mobile == mobile {
			liberr.ErrIsNil(ctx, gerror.New("手机号已存在"))
		}
	})
	return err
}

// GetEditUser 获取编辑用户信息
func (s *sSysUser) GetEditUser(ctx context.Context, id uint64) (res *system.UserGetEditRes, err error) {
	res = new(system.UserGetEditRes)
	err = g.Try(ctx, func(ctx context.Context) {
		//获取用户信息
		res.User, err = s.GetUserInfoById(ctx, id)
		liberr.ErrIsNil(ctx, err)
		//获取已选择的角色信息
		res.CheckedRoleIds, err = s.GetAdminRoleIds(ctx, id)
		liberr.ErrIsNil(ctx, err)
		res.CheckedPosts, err = s.GetUserPostIds(ctx, id)
		liberr.ErrIsNil(ctx, err)
	})
	return
}

// GetUserInfoById 通过Id获取用户信息
func (s *sSysUser) GetUserInfoById(ctx context.Context, id uint64, withPwd ...bool) (user *entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		if len(withPwd) > 0 && withPwd[0] {
			//用户用户信息
			err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, id).Scan(&user)
		} else {
			//用户用户信息
			err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, id).
				FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt).Scan(&user)
		}
		liberr.ErrIsNil(ctx, err, "获取用户数据失败")
	})
	return
}

// GetUserPostIds 获取用户岗位
func (s *sSysUser) GetUserPostIds(ctx context.Context, userId uint64) (postIds []int64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var list []*entity.SysUserPost
		err = dao.SysUserPost.Ctx(ctx).Where(dao.SysUserPost.Columns().UserId, userId).Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取用户岗位信息失败")
		postIds = make([]int64, 0)
		for _, entity := range list {
			postIds = append(postIds, entity.PostId)
		}
	})
	return
}

// ResetUserPwd 重置用户密码
func (s *sSysUser) ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (err error) {
	salt := grand.S(10)
	password := libUtils.EncryptPassword(req.Password, salt)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(req.Id).Update(g.Map{
			dao.SysUser.Columns().UserSalt:     salt,
			dao.SysUser.Columns().UserPassword: password,
		})
		liberr.ErrIsNil(ctx, err, "重置用户密码失败")
	})
	return
}

func (s *sSysUser) ChangeUserStatus(ctx context.Context, req *system.UserStatusReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(req.Id).Update(do.SysUser{UserStatus: req.UserStatus})
		liberr.ErrIsNil(ctx, err, "设置用户状态失败")
	})
	return
}

// Delete 删除用户
func (s *sSysUser) Delete(ctx context.Context, ids []int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysUser.Ctx(ctx).TX(tx).Where(dao.SysUser.Columns().Id+" in(?)", ids).Delete()
			liberr.ErrIsNil(ctx, err, "删除用户失败")
			//删除对应权限
			enforcer, e := commonService.CasbinEnforcer(ctx)
			liberr.ErrIsNil(ctx, e)
			for _, v := range ids {
				enforcer.RemoveFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", s.casBinUserPrefix, v))
			}
			//删除用户对应的岗位
			_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Delete(dao.SysUserPost.Columns().UserId+" in (?)", ids)
			liberr.ErrIsNil(ctx, err, "删除用户的岗位失败")
		})
		return err
	})
	return
}

// GetUsers 通过用户ids查询多个用户信息
func (s *sSysUser) GetUsers(ctx context.Context, ids []interface{}) (users []*model.SysUserSimpleRes, err error) {
	if len(ids) == 0 {
		return
	}
	idsSet := gset.NewFrom(ids).Slice()
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, idsSet).
			Order(dao.SysUser.Columns().Id + " ASC").Scan(&users)
	})
	return
}

// GetDataWhere 获取数据权限判断条件
// Deprecated : 此方法已废弃，请使用更简单的GetAuthWhere方法或GetAuthDeptWhere方法
func (s *sSysUser) GetDataWhere(ctx context.Context, userInfo *model.ContextUser, entityData interface{}, menuId uint) (where g.Map, err error) {
	whereJustMe := g.Map{} //本人数据权限
	t := reflect.TypeOf(entityData)
	err = g.Try(ctx, func(ctx context.Context) {
		for i := 0; i < t.Elem().NumField(); i++ {
			if t.Elem().Field(i).Name == "CreatedBy" {
				//若存在用户id的字段，则生成判断数据权限的条件
				//1、获取当前用户所属角色Ids
				var (
					roleIds   []uint
					scope     []*model.ScopeAuthData
					deptIdArr = gset.New()
					allScope  = false
				)
				roleIds, err = s.GetAdminRoleIds(ctx, userInfo.Id)
				liberr.ErrIsNil(ctx, err)
				scope, err = service.SysRole().GetRoleMenuScope(ctx, roleIds, menuId)
				liberr.ErrIsNil(ctx, err)
				if scope == nil {
					//角色未设置数据权限，默认仅本人数据权限
					whereJustMe = g.Map{"user.id": userInfo.Id}
				} else {
					//2获取角色对应数据权限
					for _, sv := range scope {
						switch sv.DataScope {
						case 1: //全部数据权限
							allScope = true
							goto endLoop
						case 2: //自定数据权限
							deptIdArr.Add(gconv.Interfaces(sv.DeptIds)...)
						case 3: //本部门数据权限
							deptIdArr.Add(gconv.Int64(userInfo.DeptId))
						case 4: //本部门及以下数据权限
							deptIdArr.Add(gconv.Int64(userInfo.DeptId))
							//获取正常状态部门数据
							deptList := ([]*entity.SysDept)(nil)
							deptList, err = service.SysDept().GetList(ctx, &system.DeptSearchReq{Status: "1", ShowAll: true})
							liberr.ErrIsNil(ctx, err)
							var dList g.List
							for _, d := range deptList {
								m := g.Map{
									"id":    d.DeptId,
									"pid":   d.ParentId,
									"label": d.DeptName,
								}
								dList = append(dList, m)
							}
							l := libUtils.FindSonByParentId(dList, userInfo.DeptId, "pid", "id")
							for _, li := range l {
								deptIdArr.Add(gconv.Int64(li["id"]))
							}
						case 5: //仅本人数据权限
							whereJustMe = g.Map{"user.id": userInfo.Id}
						}
					}
				}
			endLoop:
				if !allScope && deptIdArr.Size() > 0 {
					where = g.Map{"user.dept_id": deptIdArr.Slice()}
				} else if !allScope && len(whereJustMe) > 0 {
					where = whereJustMe
				}
				break
			}
		}
	})
	return
}

// GetAuthWhere 获取数据权限判断条件-按创建人id获取(created_by),数据权限会跟随创建人转移
func (s *sSysUser) GetAuthWhere(ctx context.Context, userInfo *model.ContextUser, field ...string) g.Map {
	var (
		//当前请求api接口对应的菜单
		url    = gstr.TrimLeft(ghttp.RequestFromCtx(ctx).Request.URL.Path, "/")
		menuId uint
		err    error
		nm     g.Map
	)
	//获取菜单ID
	menuId, err = service.SysAuthRule().GetIdByName(ctx, url)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}
	nm, err = s.GetAuthDataWhere(ctx, userInfo, menuId, field...)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}
	return nm
}

// GetAuthDataWhere 获取数据权限判断条件-按创建用户
func (s *sSysUser) GetAuthDataWhere(ctx context.Context, userInfo *model.ContextUser, menuId uint, field ...string) (where g.Map, err error) {
	whereJustMe := g.Map{} //本人数据权限
	createdUserField := "created_by"
	//表别名
	tableAlias := ""
	if len(field) > 0 && field[0] != "" {
		tableAlias = field[0]
	}
	if len(field) > 1 && field[1] != "" {
		createdUserField = field[1]
	}

	if tableAlias != "" {
		createdUserField = tableAlias + "." + createdUserField
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//若存在用户id的字段，则生成判断数据权限的条件
		//1、获取当前用户所属角色Ids
		var (
			roleIds   []uint
			scope     []*model.ScopeAuthData
			deptIdArr = gset.New()
		)
		roleIds, err = s.GetAdminRoleIds(ctx, userInfo.Id)
		liberr.ErrIsNil(ctx, err)
		scope, err = service.SysRole().GetRoleMenuScope(ctx, roleIds, menuId)
		liberr.ErrIsNil(ctx, err)
		if scope == nil {
			//角色未设置数据权限，默认仅本人数据权限
			whereJustMe = g.Map{createdUserField: userInfo.Id}
		} else {
			//2获取角色对应数据权限
			for _, sv := range scope {
				switch sv.DataScope {
				case 1: //全部数据权限
					goto endLoop
				case 2: //自定数据权限
					deptIdArr.Add(gconv.Interfaces(sv.DeptIds)...)
				case 3: //本部门数据权限
					deptIdArr.Add(gconv.Int64(userInfo.DeptId))
				case 4: //本部门及以下数据权限
					deptIdArr.Add(gconv.Int64(userInfo.DeptId))
					//获取正常状态部门数据
					deptList := ([]*entity.SysDept)(nil)
					deptList, err = service.SysDept().GetList(ctx, &system.DeptSearchReq{Status: "1", ShowAll: true})
					liberr.ErrIsNil(ctx, err)
					var dList g.List
					for _, d := range deptList {
						m := g.Map{
							"id":    d.DeptId,
							"pid":   d.ParentId,
							"label": d.DeptName,
						}
						dList = append(dList, m)
					}
					l := libUtils.FindSonByParentId(dList, userInfo.DeptId, "pid", "id")
					for _, li := range l {
						deptIdArr.Add(gconv.Int64(li["id"]))
					}
				case 5: //仅本人数据权限
					whereJustMe = g.Map{createdUserField: userInfo.Id}
				}
			}
		}
	endLoop:
		if deptIdArr.Size() > 0 {
			where = g.Map{createdUserField + " IN(select id from sys_user where dept_id in(?))": deptIdArr.Slice()}
		} else if len(whereJustMe) > 0 {
			where = whereJustMe
		}
	})
	return
}

// GetAuthDeptWhere 获取部门数据权限判断条件-按部门id获取(dept_id),数据权限依赖部门，不会随创建人转移
func (s *sSysUser) GetAuthDeptWhere(ctx context.Context, userInfo *model.ContextUser, field ...string) g.Map {
	var (
		//当前请求api接口对应的菜单
		url    = gstr.TrimLeft(ghttp.RequestFromCtx(ctx).Request.URL.Path, "/")
		menuId uint
		err    error
		nm     g.Map
	)
	//获取菜单ID
	menuId, err = service.SysAuthRule().GetIdByName(ctx, url)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}
	nm, err = s.GetAuthDeptDataWhere(ctx, userInfo, menuId, field...)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}
	return nm
}

// GetAuthDeptDataWhere 获取数据权限判断条件-按部门
func (s *sSysUser) GetAuthDeptDataWhere(ctx context.Context, userInfo *model.ContextUser, menuId uint, field ...string) (data g.Map, err error) {
	whereJustMe := g.Map{} //本人数据权限
	deptField := "dept_id"
	createdUserField := "created_by"
	//表别名
	tableAlias := ""
	if len(field) > 0 && field[0] != "" {
		tableAlias = field[0]
	}
	if len(field) > 1 && field[1] != "" {
		deptField = field[1]
	}

	if len(field) > 2 && field[2] != "" {
		createdUserField = field[2]
	}

	if tableAlias != "" {
		deptField = tableAlias + "." + deptField
		createdUserField = tableAlias + "." + createdUserField
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//若存在用户id的字段，则生成判断数据权限的条件
		//1、获取当前用户所属角色Ids
		var (
			roleIds   []uint
			scope     []*model.ScopeAuthData
			deptIdArr = gset.New()
		)
		roleIds, err = s.GetAdminRoleIds(ctx, userInfo.Id)
		liberr.ErrIsNil(ctx, err)
		scope, err = service.SysRole().GetRoleMenuScope(ctx, roleIds, menuId)
		liberr.ErrIsNil(ctx, err)
		if scope == nil {
			//角色未设置数据权限，默认仅本人数据权限
			whereJustMe = g.Map{deptField: userInfo.DeptId, createdUserField: userInfo.Id}
		} else {
			//2获取角色对应数据权限
			for _, sv := range scope {
				switch sv.DataScope {
				case 1: //全部数据权限
					goto endLoop
				case 2: //自定数据权限
					deptIdArr.Add(gconv.Interfaces(sv.DeptIds)...)
				case 3: //本部门数据权限
					deptIdArr.Add(gconv.Int64(userInfo.DeptId))
				case 4: //本部门及以下数据权限
					deptIdArr.Add(gconv.Int64(userInfo.DeptId))
					//获取正常状态部门数据
					deptList := ([]*entity.SysDept)(nil)
					deptList, err = service.SysDept().GetList(ctx, &system.DeptSearchReq{Status: "1", ShowAll: true})
					liberr.ErrIsNil(ctx, err)
					var dList g.List
					for _, d := range deptList {
						m := g.Map{
							"id":    d.DeptId,
							"pid":   d.ParentId,
							"label": d.DeptName,
						}
						dList = append(dList, m)
					}
					l := libUtils.FindSonByParentId(dList, userInfo.DeptId, "pid", "id")
					for _, li := range l {
						deptIdArr.Add(gconv.Int64(li["id"]))
					}
				case 5: //仅本人数据权限
					whereJustMe = g.Map{deptField: userInfo.DeptId, createdUserField: userInfo.Id}
				}
			}
		}
	endLoop:
		if deptIdArr.Size() > 0 {
			data = g.Map{deptField: deptIdArr.Slice()}
		} else if len(whereJustMe) > 0 {
			data = whereJustMe
		}
	})
	return
}

// HasAccessByDataWhere 判断用户是否有数据权限
func (s *sSysUser) HasAccessByDataWhere(ctx context.Context, where g.Map, uid interface{}) bool {
	err := g.Try(ctx, func(ctx context.Context) {
		rec, err := dao.SysUser.Ctx(ctx).As("user").
			Fields("user.id").
			Where("user.id", uid).Where(where).One()
		liberr.ErrIsNil(ctx, err)
		if rec.IsEmpty() {
			liberr.ErrIsNil(ctx, gerror.New("没有数据"))
		}
	})
	return err == nil
}

// AccessRule 判断用户是否有某一菜单规则权限
func (s *sSysUser) AccessRule(ctx context.Context, userId uint64, rule string) bool {
	//获取无需验证权限的用户id
	tagSuperAdmin := s.IsSupperAdmin(ctx, userId)
	if tagSuperAdmin {
		return true
	}
	menuList, err := service.SysAuthRule().GetMenuList(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return false
	}
	var menu *model.SysAuthRuleInfoRes
	for _, m := range menuList {
		ms := gstr.SubStr(m.Name, 0, gstr.Pos(m.Name, "?"))
		if m.Name == rule || ms == rule {
			menu = m
			break
		}
	}
	// 不存在的规则直接false
	if menu == nil {
		return false
	}
	enforcer, err := commonService.CasbinEnforcer(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return false
	}
	hasAccess, err := enforcer.Enforce(fmt.Sprintf("%s%d", service.SysUser().GetCasBinUserPrefix(), userId), gconv.String(menu.Id), "All")
	if err != nil {
		g.Log().Error(ctx, err)
		return false
	}
	return hasAccess
}
