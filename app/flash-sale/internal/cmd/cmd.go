package cmd

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/service"
	"shop-goframe-micro-service-refacotor/app/flash-sale/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func initTestData(ctx context.Context) {
	// 初始化商品库存
	cache := utility.GetFlashSaleCache()
	if cache == nil {
		return
	}
	// iPhone 15 Pro Max
	cache.Set(ctx, "flash_sale:stock:1001", 100, 0)
	// MacBook Pro M3
	cache.Set(ctx, "flash_sale:stock:1002", 50, 0)
	g.Log().Info(ctx, "测试库存数据已初始化")
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "Start flash sale service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化Redis
			if err := utility.InitFlashSaleRedis(ctx); err != nil {
				g.Log().Error(ctx, "Failed to init redis:", err)
				return err
			}

			// 初始化RabbitMQ（可选，失败时记录警告但不中断服务）
			if err := utility.InitFlashSaleRabbitMQ(ctx); err != nil {
				g.Log().Warning(ctx, "RabbitMQ初始化失败，消息队列功能将不可用:", err)
				// 不返回错误，让服务继续启动
			}

			// 初始化测试数据（MVP）
			initTestData(ctx)

			s := g.Server()

			// 注册中间件
			s.BindMiddlewareDefault(ghttp.MiddlewareCORS)

			// 注册控制器
			s.Group("/api/v1/flash-sale", func(group *ghttp.RouterGroup) {
				group.Bind(new(service.FlashSaleController))
			})

			// 启动服务
			s.Run()
			return nil
		},
	}
)
