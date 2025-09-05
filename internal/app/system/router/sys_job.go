// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2023-01-12 17:43:50
// 生成路径: internal/app/system/router/sys_job.go
// 生成人：gfast
// desc:定时任务
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

func (router *Router) BindSysJobController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/sysJob", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SysJob,
		)
	})
}
