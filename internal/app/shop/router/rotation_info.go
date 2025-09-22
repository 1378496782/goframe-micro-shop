// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-09-22 16:50:04
// 生成路径: internal/app/shop/router/rotation_info.go
// 生成人：gfast
// desc:轮播图
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/shop/controller"
)

func (router *Router) BindRotationInfoController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/rotationInfo", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.RotationInfo,
		)
	})
}
