// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-09-05 11:52:54
// 生成路径: internal/app/shop/router/category_info.go
// 生成人：王中阳
// desc:商品分类
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/shop/controller"
)

func (router *Router) BindCategoryInfoController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/categoryInfo", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.CategoryInfo,
		)
	})
}
