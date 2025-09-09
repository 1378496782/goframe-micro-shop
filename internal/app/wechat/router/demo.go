/*
* @desc:xxxx功能描述
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/11/3 16:23
 */

package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/wechat/controller"
)

func (router *Router) BindDemoController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/demo", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Demo,
		)
	})
}

