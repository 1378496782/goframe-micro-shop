/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/28 14:21
 */

package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/model"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/service"
	"github.com/tiger1103/gfast/v3/library/libResponse"
	"sync"
)

type demo struct {
}

var Demo = new(demo)

func (c *demo) Produce(r *ghttp.Request) {
	msg := &model.MQSendMsg{
		Topic: r.GetForm("topic").String(),
		Body:  []byte(r.GetForm("body").String()),
	}
	err := service.MQueue().SendMsg(msg)
	if err != nil {
		libResponse.FailJson(true, r, "error", err.Error())
	}
	libResponse.SusJson(true, r, "success")
}

func (c *demo) Subscribe(r *ghttp.Request) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var err error

	// 创建一个互斥锁
	var wsLock sync.Mutex
	go func() {
		topic := r.Get("topic").String()
		channel := r.Get("channel").String()
		ws, err := r.WebSocket()
		if err != nil {
			wg.Done()
			return
		}
		err = service.MQueue().Subscribe(topic, channel, func(m *model.MQMessage) error {
			fmt.Println(m)
			fmt.Println(string(m.Body))

			// 使用互斥锁保护写入操作
			wsLock.Lock()
			ws.WriteMessage(websocket.TextMessage, m.Body)
			wsLock.Unlock()

			return nil
		})
		if err != nil {
			wg.Done()
			return
		}
		wg.Done()
		for {
			_, _, err = ws.ReadMessage()
			if err != nil {
				g.Log().Warning(context.Background(), fmt.Sprintf("取消订阅 主题：%s，频道：%s", topic, channel))
				service.MQueue().Unsubscribe(topic, channel)
				return
			}
		}
	}()
	wg.Wait()
	if err != nil {
		libResponse.FailJson(true, r, "error", err.Error())
	}
	libResponse.SusJson(true, r, "success")
}
