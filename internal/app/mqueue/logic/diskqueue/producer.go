/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/12 10:23
 */

package diskqueue

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/consts"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/service"
)

type producer struct {
	isClose bool
}

func NewDiskQueueProducer() service.IProducer {
	return &producer{}
}

func (p *producer) Publish(topic string, body []byte) error {
	dq, err := getClient(topic)
	if err != nil {
		return err
	}
	return dq.Publish(body)
}

func (p *producer) PublishAsync(topic string, body []byte) error {
	dq, err := getClient(topic)
	go func() {
		err = dq.Publish(body)
		if err != nil {
			g.Log("diskQueue").Error(context.TODO(), "diskQueue PublishAsync消息失败："+err.Error())
		}
	}()
	return err
}

func (p *producer) DelayPublish(topic string, body []byte, delay consts.MsgDelayLevel) error {
	g.Log("diskQueue").Warning(context.TODO(), "diskQueue 不支持延时消息，将使用publish发送")
	dq, err := getClient(topic)
	if err != nil {
		return err
	}
	return dq.Publish(body)
}

func (p *producer) Close() {
	p.isClose = true
}
