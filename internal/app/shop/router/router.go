// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-09-08 11:37:29
// 生成路径: internal/app/shop/router/goods_info.go
// 生成人：王中阳
// desc:商品
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libRouter"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/shop", func(group *ghttp.RouterGroup) {
		//context拦截器
		group.Middleware(service.Middleware().Ctx)
		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBindBefore(ctx, router, group); err != nil {
			panic(err)
		}
		//登录验证拦截
		service.GfToken().Middleware(group)
		group.Middleware(service.Middleware().Auth)
		//后台操作日志记录
		group.Hook("/*", ghttp.HookAfterOutput, service.OperateLog().OperationLog)
		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
