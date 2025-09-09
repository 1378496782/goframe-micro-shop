/*
* @desc:扶苗工作路由
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/2/18 17:34
 */

package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/wechat/controller"
	"github.com/tiger1103/gfast/v3/internal/app/wechat/service"
	"github.com/tiger1103/gfast/v3/library/libRouter"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/wechat", func(group *ghttp.RouterGroup) {
		group.Bind(
			//登录
			controller.Login,
		)
		//context拦截器
		group.Middleware(service.Middleware().Ctx)
		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBindBefore(ctx, router, group); err != nil {
			panic(err)
		}
		//登录验证拦截
		service.GfToken().Middleware(group)

		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
