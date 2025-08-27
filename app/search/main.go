package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	_ "shop-goframe-micro-service-refacotor/app/search/internal/packed"
	"shop-goframe-micro-service-refacotor/app/search/utility/binlog"
	"shop-goframe-micro-service-refacotor/app/search/utility/elasticsearch"

	"shop-goframe-micro-service-refacotor/app/search/internal/cmd"
)

func main() {
	ctx := gctx.New()

	// 初始化ES
	if err := elasticsearch.Init(ctx); err != nil {
		panic(err)
	}
	// 启动 binlog 监听（在后台运行）
	go binlog.StartBinlogSyncer(ctx)

	cmd.Main.Run(gctx.GetInitCtx())
}
