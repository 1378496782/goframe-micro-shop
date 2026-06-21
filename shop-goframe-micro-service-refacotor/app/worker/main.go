package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"os"
	"os/signal"
	_ "shop-goframe-micro-service-refacotor/app/worker/internal/packed"
	"shop-goframe-micro-service-refacotor/app/worker/utility/rabbitmq"
	"syscall"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()

	// 初始化配置
	g.Log().Info(ctx, "正在加载配置文件...")

	// 创建RabbitMQ客户端
	g.Log().Info(ctx, "正在初始化RabbitMQ客户端...")
	rabbitMQClient, err := rabbitmq.NewRabbitMQClient(ctx)
	if err != nil {
		g.Log().Fatalf(ctx, "初始化RabbitMQ客户端失败: %v", err)
	}
	defer rabbitMQClient.Close()

	// 测试连接
	g.Log().Info(ctx, "正在测试RabbitMQ连接...")
	err = rabbitMQClient.TestConnection(ctx)
	if err != nil {
		g.Log().Fatalf(ctx, "RabbitMQ连接测试失败: %v", err)
	}

	g.Log().Info(ctx, "Worker服务启动成功，等待处理消息...")

	// 等待中断信号优雅关闭
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞等待中断信号
	sig := <-sigChan
	g.Log().Infof(ctx, "收到信号: %v，开始关闭服务...", sig)

	// 执行清理操作
	g.Log().Info(ctx, "Worker服务已关闭")
}
