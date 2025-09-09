package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	wechat "github.com/tiger1103/gfast/v3/internal/app/wechat/router"
)

func (router *Router) BindWechatModuleController(ctx context.Context, group *ghttp.RouterGroup) {
	wechat.R.BindController(ctx, group)
}
