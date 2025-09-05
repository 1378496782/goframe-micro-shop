package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	shopRouter "github.com/tiger1103/gfast/v3/internal/app/shop/router"
)

func (router *Router) BindShopModuleController(ctx context.Context, group *ghttp.RouterGroup) {
	shopRouter.R.BindController(ctx, group)
}
