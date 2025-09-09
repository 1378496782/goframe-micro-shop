/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/12 10:22
 */

package diskqueue

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/model"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/service"
)

type consumer struct {
	Topic   string
	Channel string
	Id      string
	Handler model.MQConsumerHandlerCallback
}

// NewDiskConsumer 创建一个消费者
func NewDiskConsumer(topic, channel string, handler model.MQConsumerHandlerCallback) (service.IConsumer, error) {
	dqs, err := getClient(topic)
	if err != nil {
		return nil, err
	}
	id := "dqc_" + grand.S(16)
	c := &consumer{
		Topic:   topic,
		Channel: channel,
		Id:      id,
		Handler: handler,
	}
	dqs.RegisterConsumer(channel, c)
	return c, nil
}

func (s *consumer) CloseMqConsumer() {
	dqs, err := getClient(s.Topic)
	if err != nil {
		g.Log("diskQueue").Error(context.TODO(), "执行 CloseMqConsumer 失败："+err.Error())
		return
	}
	dqs.RemoveConsumer(s.Channel, s.Id)
}
