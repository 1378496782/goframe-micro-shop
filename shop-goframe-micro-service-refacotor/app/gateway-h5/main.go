package main

import (
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/internal/cmd"
	"shop-goframe-micro-service-refacotor/utility/middleware"
	"shop-goframe-micro-service-refacotor/utility/metrics"
)

func main() {
	var ctx = gctx.New()
	conf, err := g.Cfg().Get(ctx, "etcd.address")
	if err != nil {
		panic(err)
	}

	var address = conf.String()
	grpcx.Resolver.Register(etcd.New(address))

	// 初始化 Prometheus 指标
	metrics.InitMetrics()

	// 创建 HTTP 服务
	s := g.Server()

	// 设置 CORS 头
	s.Use(middleware.MiddlewareCORS)
	// 注册 Prometheus 指标收集中间件
	s.Use(metrics.MetricsMiddleware)
	s.Use(metrics.ErrorMetricsMiddleware)
	// 注册 /metrics 端点
	metrics.RegisterHTTPHandler(s)
	cmd.Main.Run(ctx)
}
