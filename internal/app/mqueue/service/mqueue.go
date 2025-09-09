/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/3 8:56
 */

package service

import (
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/consts"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/model"
)

type IConsumer interface {
	// CloseMqConsumer 关闭消费者
	CloseMqConsumer()
}

// IProducer 消息队列生产者接口
type IProducer interface {
	// Publish 发布单条消息并等待结果
	Publish(topic string, body []byte) error
	// PublishAsync 发布单条异步消息，不用等待服务器返回结果
	PublishAsync(topic string, body []byte) error
	// DelayPublish 延时发布消息并等待结果，delay单位延迟等级1-18级对应: 1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
	DelayPublish(topic string, body []byte, delay consts.MsgDelayLevel) error
	// Close 关闭生产者
	Close()
}

type IMQueue interface {
	SendMsg(m *model.MQSendMsg) error
	Subscribe(topic, channel string, handler model.MQConsumerHandlerCallback) error
	Close()
	Unsubscribe(topic, channel string) error
}

var localMQueue IMQueue

func MQueue() IMQueue {
	if localMQueue == nil {
		panic("implement not found for interface IMQueue, forgot register?")
	}
	return localMQueue
}

func RegisterMQueue(i IMQueue) {
	localMQueue = i
}
