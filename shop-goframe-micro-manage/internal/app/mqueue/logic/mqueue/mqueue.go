/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/6/28 11:14
 */

package mqueue

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gpool"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/consts"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/logic/diskqueue"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/logic/nsq"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/logic/rocketmq"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/model"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/service"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var ctx = context.TODO()
var mQueueMutex = sync.Mutex{}

type mQueue struct {
	producerPool        *gpool.Pool
	consumerInstanceMap map[string]service.IConsumer
	logger              glog.ILogger
}

func init() {
	isEnable := g.Cfg().MustGet(ctx, "mqueue.enable").Bool()
	if !isEnable {
		return
	}
	service.RegisterMQueue(New())
	g.Log().Info(ctx, "mqueue service is running")
}

func New() service.IMQueue {
	mq := &mQueue{
		producerPool: gpool.New(60*time.Second, func() (interface{}, error) {
			//g.Log().Info(context.Background(), "创建新的producerPool 对象")
			return createProducer()
		}, func(i interface{}) {
			i.(service.IProducer).Close()
			i = nil
			//g.Log().Warning(context.Background(), "producerPool 超时释放资源 ")
		}),
		consumerInstanceMap: make(map[string]service.IConsumer),
	}
	go func() {
		// 监听程序退出命令，清理队列
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Kill, os.Interrupt)
		for {
			select {
			case s := <-c:
				if s == syscall.SIGINT || s == syscall.SIGTERM || s == syscall.SIGHUP || s == syscall.SIGQUIT || s == os.Interrupt || s == os.Kill {
					fmt.Println("mqueue exiting...")
					mq.Close()
					closeWaitTime := g.Cfg().MustGet(ctx, "mqueue.closeWaitTime").Int64()
					if closeWaitTime <= 0 {
						closeWaitTime = 5
					}
					time.Sleep(time.Duration(closeWaitTime) * time.Second)
					os.Exit(0)
					return
				}
			}
		}
	}()
	return mq
}

// CreateProducer 创建生产者
func createProducer() (service.IProducer, error) {
	mqDriver, err := g.Cfg().Get(context.TODO(), "mqueue.driver")
	if err != nil {
		return nil, err
	}
	var mqProducer service.IProducer
	switch mqDriver.String() {
	case "diskQueue":
		mqProducer = diskqueue.NewDiskQueueProducer()
	case "nsq":
		mqProducer, err = nsq.NewNsqProducer()
	case "rocketmq":
		mqProducer, err = rocketmq.NewProducer()
	default:
		return nil, errors.New("没有找到消息队列驱动，请检查配置文件中的driver类型")
	}
	return mqProducer, nil
}

// SendMsg 发送消息
func (m *mQueue) SendMsg(msg *model.MQSendMsg) error {
	// 从生产者对象池取一个
	pd, err := m.producerPool.Get()
	if err != nil {
		return err
	}
	if pd == nil {
		return errors.New("从生产者对象池中获取生产者失败")
	}
	// 使用完放回到对象池
	defer func(producerPool *gpool.Pool, value interface{}) {
		err2 := producerPool.Put(value)
		if err2 != nil {
			g.Log("mqueue").Error(ctx, "sendMsg defer producerPool put error:")
			g.Log("mqueue").Error(ctx, err2.Error())
		}
	}(m.producerPool, pd)

	switch msg.SendMethod {
	case consts.SendMsgPublishAsync:
		err = pd.(service.IProducer).PublishAsync(msg.Topic, msg.Body)
	case consts.SendMsgDelay:
		err = pd.(service.IProducer).DelayPublish(msg.Topic, msg.Body, msg.Delay)
	default:
		err = pd.(service.IProducer).Publish(msg.Topic, msg.Body)
	}
	return err
}

// Subscribe 从指定的topic 和 channel 订阅消息，并使用回调函数来处理消息
func (m *mQueue) Subscribe(topic, channel string, handler model.MQConsumerHandlerCallback) error {
	mQueueMutex.Lock()
	defer mQueueMutex.Unlock()
	key := fmt.Sprintf("%s_%s", topic, channel)
	// 一个频道尽量一个处理器以保证消息的原子性，多个应用消费同一主题需最好在handler里面自行处理错误
	if _, ok := m.consumerInstanceMap[key]; ok {
		return errors.New(fmt.Sprintf("已经订阅过该主题 %s 和频道 %s", topic, channel))
	}
	mqDriver, err := g.Cfg().Get(ctx, "mqueue.driver")
	if err != nil {
		return err
	}
	var mqConsumer service.IConsumer
	switch mqDriver.String() {
	case "diskQueue":
		mqConsumer, err = diskqueue.NewDiskConsumer(topic, channel, handler)
	case "nsq":
		mqConsumer, err = nsq.NewNsqConsumer(topic, channel, handler)
	case "rocketmq":
		mqConsumer, err = rocketmq.NewPushConsumer(topic, channel, handler)
	default:
		return errors.New("没有找到消息队列驱动，请检查配置文件中的driver类型")
	}
	if err != nil {
		return err
	}
	m.consumerInstanceMap[key] = mqConsumer
	return err
}

// Close 关闭消息队列
func (m *mQueue) Close() {
	// 清空生产者
	m.producerPool.Clear()
	// 关闭消费者
	for key, c := range m.consumerInstanceMap {
		c.CloseMqConsumer()
		delete(m.consumerInstanceMap, key)
	}
	mqDriver, err := g.Cfg().Get(ctx, "mqueue.driver")
	if err != nil {
		g.Log().Error(ctx, "获取 mqueue.driver 失败")
		return
	}
	switch mqDriver.String() {
	case "diskQueue":
		diskqueue.CloseDiskQueueService()
	}
}

// Unsubscribe 取消订阅主题
func (m *mQueue) Unsubscribe(topic, channel string) error {
	mQueueMutex.Lock()
	defer mQueueMutex.Unlock()
	key := fmt.Sprintf("%s_%s", topic, channel)
	if _, ok := m.consumerInstanceMap[key]; ok {
		m.consumerInstanceMap[key].CloseMqConsumer()
		delete(m.consumerInstanceMap, key)
		return nil
	} else {
		return errors.New("没有找到订阅信息")
	}
}
