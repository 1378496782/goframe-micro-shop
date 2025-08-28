package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"shop-goframe-micro-service-refacotor/app/gateway-admin/internal/controller/admin"
	"shop-goframe-micro-service-refacotor/app/gateway-admin/internal/controller/goods"
	"shop-goframe-micro-service-refacotor/app/gateway-admin/internal/controller/order"
	"shop-goframe-micro-service-refacotor/utility/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http gateway-admin server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Bind(
						admin.NewV1(),
					)
				})
				// 需要JWT验证的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTAuth)
					group.Bind(
						// 需要认证的接口
						goods.NewV1(),
						order.NewV1(),
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
