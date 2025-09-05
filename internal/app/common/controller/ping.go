/*
* @desc:ping
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/11/30 15:59
 */

package controller

import (
	"github.com/tiger1103/gfast/v3/library/libWebsocket"
)

var Ping = new(pingController)

type pingController struct{}

func (c *pingController) Ping(client *libWebsocket.Client, req *libWebsocket.WRequest) {
	libWebsocket.SendSuccess(client, req.Event)
}
