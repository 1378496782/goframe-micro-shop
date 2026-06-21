/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/24 9:54
 */

package rocketmq

import (
	"context"
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	rmq_consumer "github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/model"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/service"
)

type rocketmqConsumer struct {
	topic        string
	pushConsumer rocketmq.PushConsumer
}

// NewPushConsumer 创建MyRocketmq消费者实例
func NewPushConsumer(topic, groupName string, handler model.MQConsumerHandlerCallback) (service.IConsumer, error) {
	endPoint := g.Cfg().MustGet(context.TODO(), "mqueue.rocketmq.nameServers").Strings()
	logLevel := g.Cfg().MustGet(context.TODO(), "mqueue.rocketmq.logLevel").String()
	if len(endPoint) < 1 {
		return nil, errors.New("配置文件不正确，获取mqueue.rocketmq.nameServers失败")
	}
	// 创建rocketmq实例
	rlog.SetLogLevel(logLevel)
	rmqC, err := rocketmq.NewPushConsumer(
		rmq_consumer.WithNameServer(endPoint),
		rmq_consumer.WithGroupName(groupName),
	)

	if err != nil {
		return nil, err
	}
	err = rmqC.Subscribe(topic, rmq_consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (rmq_consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			var id model.MessageID
			copy(id[:], msg.MsgId[:32])
			handler(&model.MQMessage{
				ID:        id,
				Body:      msg.Body,
				Attempts:  0,
				Timestamp: msg.BornTimestamp,
			})
		}
		// 无论如何都返回成功，消费失败自行在处理器里处理
		return rmq_consumer.ConsumeSuccess, nil
	})
	err = rmqC.Start()
	return &rocketmqConsumer{pushConsumer: rmqC}, nil
}

func (c *rocketmqConsumer) CloseMqConsumer() {
	var err error
	if c.pushConsumer != nil {
		err = c.pushConsumer.Shutdown()
	}
	if err != nil {
		g.Log("rocketmq").Error(context.TODO(), errors.New("rocketmq CloseMqConsumer 失败："))
		g.Log("rocketmq").Error(context.TODO(), err)
	}
}
