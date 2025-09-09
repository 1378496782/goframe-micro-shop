/*
* @desc:测试登录后才可以访问
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/11/3 16:05
 */

package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/wechat"
)

var Demo = new(demoController)

type demoController struct {
	BaseController
}

func (c *demoController) Demo(ctx context.Context,req *wechat.DemoReq)(res *wechat.DemoRes,err error){
	res = &wechat.DemoRes{Info: "hello word"}
	return
}
