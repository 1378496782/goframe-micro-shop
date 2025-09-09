package libWebsocket

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
)

var (
	mctx          = gctx.GetInitCtx()             // 上下文
	clientManager = NewClientManager()            // 客户端管理
	routers       = make(map[string]EventHandler) // 消息路由
	msgGo         = grpool.New(20)                // 消息处理协程池
)

// Start 启动
func Start() {
	go clientManager.start()
	go clientManager.ping()
	g.Log().Debug(mctx, "start websocket..")
}

// Stop 关闭
func Stop() {
	clientManager.closeSignal <- struct{}{}
}

func Register(client *Client) {
	clientManager.Register <- client
}
