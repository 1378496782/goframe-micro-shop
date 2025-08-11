package main

import (
	_ "shop-goframe-micro-service-refacotor/app/gateway-admin/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
