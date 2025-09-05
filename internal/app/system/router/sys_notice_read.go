// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2023-11-09 17:37:25
// 生成路径: internal/app/system/router/sys_notice_read.go
// 生成人：gfast
// desc:已读记录
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

func (router *Router) BindSysNoticeReadController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/sysNoticeRead", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SysNoticeRead,
		)
	})
}
