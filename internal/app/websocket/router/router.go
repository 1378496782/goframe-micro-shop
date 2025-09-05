/*
* @desc:websocket路由
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/11/30 15:06
 */

package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/gorilla/websocket"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/tiger1103/gfast/v3/library/libWebsocket"
	"net/http"
)

var R = new(Router)

type Router struct{}

// WsPage ws入口
func WsPage(r *ghttp.Request) {
	upGrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		return
	}
	currentTime := uint64(gtime.Now().Unix())
	client := NewClient(r, conn, currentTime)
	go client.Read()
	go client.Write()
	// 用户连接事件
	libWebsocket.Register(client)
}

// NewClient 初始化
func NewClient(r *ghttp.Request, socket *websocket.Conn, firstTime uint64) (client *libWebsocket.Client) {
	client = &libWebsocket.Client{
		Addr:          socket.RemoteAddr().String(),
		ID:            guid.S(),
		Socket:        socket,
		Send:          make(chan *libWebsocket.WResponse, 100),
		SendClose:     false,
		CloseSignal:   make(chan struct{}, 1),
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
		User:          service.Context().GetLoginUser(r.Context()),
		IP:            libUtils.GetClientIp(r.Context()),
		UserAgent:     r.UserAgent(),
	}
	return
}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/websocket", func(group *ghttp.RouterGroup) {
		//登录验证拦截
		service.GfToken().Middleware(group)
		//context拦截器
		group.Middleware(service.Middleware().Ctx)

		group.GET("/", WsPage)
	})
	// 启动websocket监听
	libWebsocket.Start()

	// 注册消息路由
	libWebsocket.RegisterMsg(libWebsocket.EventHandlers{
		"ping": commonController.Ping.Ping, // 心跳
	})
}
