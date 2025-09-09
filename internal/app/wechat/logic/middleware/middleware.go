/*
* @desc:中间件
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/23 15:05
 */

package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/wechat/service"
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

type sMiddleware struct{}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()
	// 初始化登录用户信息
	data, err := service.GfToken().ParseToken(r)
	if err != nil {
		// 执行下一步请求逻辑
		r.Middleware.Next()
		return
	}
	if data != nil {
		context := new(model.Context)
		err = gconv.Struct(data.Data, &context.User)
		if err != nil {
			g.Log().Error(ctx, err)
			// 执行下一步请求逻辑
			r.Middleware.Next()
			return
		}
		service.Context().Init(r, context)
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}
