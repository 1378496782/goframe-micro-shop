package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	websocketRouter "github.com/tiger1103/gfast/v3/internal/app/websocket/router"
)

func (router *Router) BindWebsocketModuleController(ctx context.Context, group *ghttp.RouterGroup) {
	websocketRouter.R.BindController(ctx, group)
}
