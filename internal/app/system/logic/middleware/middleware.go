/*
* @desc:中间件
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/23 15:05
 */

package middleware

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libResponse"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

type sMiddleware struct{}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()
	// 初始化登录用户信息
	data, err := service.GfToken().ParseToken(r)
	if err != nil {
		// 执行下一步请求逻辑
		r.Middleware.Next()
		return
	}
	if data != nil {
		context := new(model.Context)
		err = gconv.Struct(data.Data, &context.User)
		if err != nil {
			g.Log().Error(ctx, err)
			// 执行下一步请求逻辑
			r.Middleware.Next()
			return
		}
		service.Context().Init(r, context)
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// Auth 权限判断处理中间件
func (s *sMiddleware) Auth(r *ghttp.Request) {
	ctx := r.GetCtx()
	//获取登陆用户id
	adminId := service.Context().GetUserId(ctx)
	url := gstr.TrimLeft(r.Request.URL.Path, "/")
	/*if r.Method != "GET" && adminId != 1 && url != "api/v1/system/login" {
		libResponse.FailJson(true, r, "对不起！演示系统，不能删改数据！")
	}*/
	//获取无需验证权限的用户id
	tagSuperAdmin := service.SysUser().IsSupperAdmin(ctx, service.Context().GetUserId(ctx))
	if tagSuperAdmin {
		r.Middleware.Next()
		//不要再往后面执行
		return
	}
	//获取地址对应的菜单id
	menuList, err := service.SysAuthRule().GetMenuList(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		libResponse.FailJson(true, r, "请求数据失败")
	}
	var menu *model.SysAuthRuleInfoRes
	for _, m := range menuList {
		ms := gstr.SubStr(m.Name, 0, gstr.Pos(m.Name, "?"))
		if m.Name == url || ms == url {
			menu = m
			break
		}
	}
	//只验证存在数据库中的规则
	if menu != nil {
		//若是不登录能访问的接口则不判断权限
		excludePaths := g.Cfg().MustGet(ctx, "gfToken.excludePaths").Strings()
		for _, p := range excludePaths {
			if gstr.Equal(menu.Name, gstr.TrimLeft(p, "/")) {
				r.Middleware.Next()
				return
			}
		}
		//若存在不需要验证的条件则跳过
		if gstr.Equal(menu.Condition, "nocheck") {
			r.Middleware.Next()
			return
		}
		menuId := menu.Id
		//菜单没存数据库不验证权限
		if menuId != 0 {
			//判断权限操作
			err = s.checkAuth(ctx, adminId, menuId)
			if err != nil {
				libResponse.FailJson(true, r, err.Error())
			}
		}
	}
	r.Middleware.Next()
}

func (s *sMiddleware) checkAuth(ctx context.Context, adminId uint64, menuId uint) (err error) {
	var (
		roleIds    []uint
		roleList   []*entity.SysRole
		roleIdsMap = gmap.New()
		enforcer   *casbin.SyncedEnforcer
		b          bool
	)
	err = g.Try(ctx, func(ctx context.Context) {
		roleIds, err = service.SysUser().GetAdminRoleIds(ctx, adminId)
		liberr.ErrIsNil(ctx, err)
		for _, v := range roleIds {
			roleIdsMap.Set(v, v)
		}
		//获取对应角色
		roleList, err = service.SysRole().GetRoleList(ctx)
		liberr.ErrIsNil(ctx, err)
		for _, v := range roleList {
			if roleIdsMap.Contains(v.Id) {
				//判断是否在角色有效时间段内
				if v.EffectiveTime != "" {
					var effective *model.EffectiveTimeInfo
					err = gconv.Struct(v.EffectiveTime, &effective)
					liberr.ErrIsNil(ctx, err, "获取角色有效时间段失败")
					if effective != nil && effective.EffectiveType == consts.EffectiveTypeStartEnd {
						//按起止日期
						now := gtime.Now()
						if now.Before(effective.DateRange[0]) || now.After(effective.DateRange[1]) {
							roleIdsMap.Remove(v.Id)
						}
					} else if effective != nil && effective.EffectiveType == consts.EffectiveTypeDate {
						//按时间段
						now := gtime.Now()
						arr := garray.NewIntArrayFrom(effective.WeekDay)
						if arr.Contains(gconv.Int(now.Format("w"))) {
							sHis := effective.DayRange[0].Format("H:i:s")
							eHis := effective.DayRange[1].Format("H:i:s")
							start := gtime.NewFromStr(now.Format("Y-m-d " + sHis))
							end := gtime.NewFromStr(now.Format("Y-m-d " + eHis))
							if now.Before(start) || now.After(end) {
								roleIdsMap.Remove(v.Id)
							}
						} else {
							roleIdsMap.Remove(v.Id)
						}
					}
				}
			}
		}
		enforcer, err = commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, err)
		roleIdsMap.Iterator(func(k interface{}, v interface{}) bool {
			b, err = enforcer.Enforce(gconv.String(v), gconv.String(menuId), "All")
			liberr.ErrIsNil(ctx, err)
			return !b
		})
		if !b {
			liberr.ErrIsNil(ctx, errors.New("没有权限"))
		}
	})
	return
}
