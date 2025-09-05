/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/28 14:56
 */

package router

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/controller"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	isEnable := g.Cfg().MustGet(ctx, "mqueue.enable").Bool()
	if isEnable{
		group.Group("/mqueue/demo", func(group *ghttp.RouterGroup) {
			group.POST("/produce", controller.Demo.Produce)
			group.ALL("/subscribe", controller.Demo.Subscribe)
		})
	}
}
