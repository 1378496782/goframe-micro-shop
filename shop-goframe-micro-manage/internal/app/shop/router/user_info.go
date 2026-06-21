// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-09-09 15:51:53
// 生成路径: internal/app/shop/router/user_info.go
// 生成人：gfast
// desc:用户
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/shop/controller"
)

func (router *Router) BindUserInfoController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/userInfo", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.UserInfo,
		)
	})
}
