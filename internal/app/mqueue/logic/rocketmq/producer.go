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
	"github.com/apache/rocketmq-client-go/v2/primitive"
	rmq_producer "github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/consts"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/service"
)

type producer struct {
	rocketmqProducer rocketmq.Producer
}

// NewProducer 创建rocketmq生产者实例
func NewProducer() (service.IProducer, error) {
	endPoint := g.Cfg().MustGet(context.TODO(), "mqueue.rocketmq.nameServers")
	retry := g.Cfg().MustGet(context.TODO(), "mqueue.rocketmq.retry").Int()
	groupName := g.Cfg().MustGet(context.TODO(), "mqueue.channel").String()
	rmqP, err := rocketmq.NewProducer(rmq_producer.WithNameServer(endPoint.Strings()), rmq_producer.WithRetry(retry), rmq_producer.WithGroupName(groupName))
	if err != nil {
		return nil, err
	}
	if rmqP == nil {
		return nil, errors.New("创建生产者失败")
	}
	err = rmqP.Start()
	if err != nil {
		return nil, err
	}
	return &producer{rocketmqProducer: rmqP}, nil
}

func (p *producer) Publish(topic string, body []byte) error {
	msg := &primitive.Message{
		Topic: topic,
		Body:  body,
	}
	_, err := p.rocketmqProducer.SendSync(context.TODO(), msg)
	return err
}

func (p *producer) PublishAsync(topic string, body []byte) error {
	msg := &primitive.Message{
		Topic: topic,
		Body:  body,
	}
	err := p.rocketmqProducer.SendAsync(context.TODO(), func(ctx context.Context, result *primitive.SendResult, err error) {
		if err != nil {
			g.Log("rocketmq").Error(context.TODO(), errors.New("rocketmq PublishAsync 失败："))
			g.Log("rocketmq").Error(context.TODO(), err)
		}
	}, msg)
	return err
}

func (p *producer) DelayPublish(topic string, body []byte, delay consts.MsgDelayLevel) error {
	msg := &primitive.Message{
		Topic: topic,
		Body:  body,
	}
	if delay < 1 || delay > 18 {
		delay = 1
	}
	msg.WithDelayTimeLevel(int(delay))
	_, err := p.rocketmqProducer.SendSync(context.TODO(), msg)
	return err
}

func (p *producer) Close() {
	if p.rocketmqProducer != nil {
		err := p.rocketmqProducer.Shutdown()
		if err != nil {
			g.Log("rocketmq").Error(context.TODO(), err)
		}
	}
}
