// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-09-09 15:39:41
// 生成路径: internal/app/shop/router/user_coupon_info.go
// 生成人：gfast
// desc:用户优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/shop/controller"
)

func (router *Router) BindUserCouponInfoController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/userCouponInfo", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.UserCouponInfo,
		)
	})
}
