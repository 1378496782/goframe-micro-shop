package main

import (
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/cmd"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/mq"
	"shop-goframe-micro-service-refacotor/app/flash-sale/utility"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "shop-goframe-micro-service-refacotor/app/flash-sale/internal/packed"

	_ "shop-goframe-micro-service-refacotor/app/flash-sale/internal/logic"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()

	// 初始化RabbitMQ（改为非阻塞模式，失败仅记录警告）
	if err := utility.InitFlashSaleRabbitMQ(ctx); err != nil {
		g.Log().Warning(ctx, "初始化RabbitMQ失败，服务将以降级模式运行（无消息队列功能）:", err)
	} else {
		// 只有RabbitMQ初始化成功才启动消费者
		consumer := mq.NewFlashSaleOrderConsumer(ctx)
		if err := consumer.Start(); err != nil {
			g.Log().Error(ctx, "启动消息队列消费者失败:", err)
			// 消费者启动失败不影响主服务
		} else {
			g.Log().Info(ctx, "消息队列消费者启动成功")
		}
	}

	g.Log().Info(ctx, "秒杀服务启动中...")
	cmd.Main.Run(ctx)
}
