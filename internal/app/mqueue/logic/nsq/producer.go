/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/3 10:44
 */

package nsq

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/nsqio/go-nsq"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/consts"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/service"
	"time"
)

type nsqProducer struct {
	Producer *nsq.Producer
}

// NewNsqProducer 创建操作NSQProducer的实例
func NewNsqProducer() (service.IProducer, error) {
	config := nsq.NewConfig()
	address := g.Cfg().MustGet(context.TODO(), "mqueue.nsq.address").String()
	port := g.Cfg().MustGet(context.TODO(), "mqueue.nsq.producer_port").Uint()
	if address == "" || port <= 0 {
		return nil, errors.New("配置文件有错误")
	}
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%d", address, port), config)
	if err != nil {
		return nil, err
	}
	return &nsqProducer{Producer: producer}, nil
}

func (p *nsqProducer) Publish(topic string, body []byte) error {
	return p.Producer.Publish(topic, body)
}

func (p *nsqProducer) PublishAsync(topic string, body []byte) error {
	doneChan := make(chan *nsq.ProducerTransaction)
	go func() {
		timeout := time.NewTicker(5 * time.Second)
		select {
		case res := <-doneChan:
			if res.Error != nil {
				g.Log("nsq").Error(context.TODO(), "publishAsync error doneChan 返回错误："+res.Error.Error())
			}
		case <-timeout.C:
			g.Log("nsq").Error(context.TODO(), "PublishAsync error doneChan 返回超时")
		}

	}()
	return p.Producer.PublishAsync(topic, body, doneChan)
}

func (p *nsqProducer) DelayPublish(topic string, body []byte, delay consts.MsgDelayLevel) error {
	return p.Producer.DeferredPublish(topic, delay.Value(), body)
}

func (p *nsqProducer) Close() {
	p.Producer.Stop()
}
