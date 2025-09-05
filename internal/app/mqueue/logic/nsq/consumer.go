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
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/model"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/service"
)

type nsqConsumer struct {
	consumer  *nsq.Consumer
	nsqConfig *nsqConsumerConnConfig
	Handler   model.MQConsumerHandlerCallback
}

type nsqConsumerConnConfig struct {
	Addr                 string // NsqLookupd 服务器地址
	Port                 uint   // 服务器端口号
	LocalAddr            string // 本地IP地址
	AuthSecret           string // 认证秘钥
	LookupdAuthorization bool   // 是否开启认证
}

func (c *nsqConsumer) CloseMqConsumer() {
	c.consumer.Stop()
}

// HandleMessage 处理go_nsq转发的消息，返回nil则表式处理完成
func (c *nsqConsumer) HandleMessage(m *nsq.Message) error {
	// 空消息不用处理
	if len(m.Body) == 0 {
		return nil
	}
	var id model.MessageID
	copy(id[:], m.ID[:16])
	mqMsgRes := &model.MQMessage{
		ID:        id,
		Body:      m.Body,
		Timestamp: m.Timestamp,
	}
	return c.Handler(mqMsgRes)
}

// NewNsqConsumer 创建NSQConsumer实例
func NewNsqConsumer(topic, channel string, handler model.MQConsumerHandlerCallback) (service.IConsumer, error) {
	addr := g.Cfg().MustGet(context.TODO(), "mqueue.nsq.address").String()
	port := g.Cfg().MustGet(context.TODO(), "mqueue.nsq.consumer_port").Uint()
	if addr == "" || port < 1 {
		return nil, errors.New("nsq 配置读取错误")
	}

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)

	consumer.SetLoggerLevel(nsq.LogLevel(2))
	if err != nil {
		return nil, err
	}
	nsqConsumerConfig := &nsqConsumerConnConfig{
		Addr: addr,
		Port: port,
	}
	nsqC := &nsqConsumer{
		nsqConfig: nsqConsumerConfig,
		consumer:  consumer,
		Handler:   handler,
	}
	// 添加消息处理器
	consumer.AddHandler(nsqC)
	err = consumer.ConnectToNSQLookupd(fmt.Sprintf("%s:%d", nsqConsumerConfig.Addr, nsqConsumerConfig.Port))
	if err != nil {
		return nil, err
	}
	stats := consumer.Stats()
	if stats.Connections <= 0 {
		nsqC = nil
		err = errors.New("未能连接到nsq lookup 服务器，请检查错误日志")
	}
	return nsqC, err
}
