package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	_ "shop-goframe-micro-service-refacotor/app/search/internal/packed"
	"shop-goframe-micro-service-refacotor/app/search/utility/elasticsearch"

	"shop-goframe-micro-service-refacotor/app/search/internal/cmd"
)

func main() {
	ctx := gctx.New()

	// 初始化ES
	if err := elasticsearch.Init(ctx); err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
