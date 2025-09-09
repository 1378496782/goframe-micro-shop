/*
* @desc:测试登录小程序后才能访问
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/11/3 16:06
 */

package wechat

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/common"
)

type DemoReq struct {
	g.Meta `path:"/demo" tags:"微信接口/小程序测试" method:"get" summary:"测试"`
	common.Author
}

type DemoRes struct {
	common.EmptyRes
	Info string `json:"info"`
}
