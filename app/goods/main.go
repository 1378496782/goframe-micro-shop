package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"os"
	"shop-goframe-micro-service-refacotor/app/goods/internal/cmd"
	"shop-goframe-micro-service-refacotor/app/goods/utility/consumer"
	"shop-goframe-micro-service-refacotor/app/goods/utility/goodsRedis"
)

func main() {
	var ctx = gctx.New()
	conf, err := g.Cfg().Get(ctx, "etcd.address")
	if err != nil {
		panic(err)
	}

	// 初始化Redis
	if err := goodsRedis.InitGoodsRedis(ctx); err != nil {
		g.Log().Fatal(ctx, "Redis初始化失败:", err)
		os.Exit(1)
	}
	// 启动消费者
	go consumer.StartUserRegisteredConsumer()

	var address = conf.String()
	grpcx.Resolver.Register(etcd.New(address))

	cmd.Main.Run(ctx)
}
