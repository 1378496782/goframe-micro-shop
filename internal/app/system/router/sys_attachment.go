// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2024-10-23 16:10:12
// 生成路径: internal/app/system/router/sys_attachment.go
// 生成人：gfast
// desc:附件管理
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

func (router *Router) BindSysAttachmentController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/sysAttachment", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SysAttachment,
		)
	})
}
