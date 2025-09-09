/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/7 11:26
 */

package mqueue

import (
	"bytes"
	"fmt"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/consts"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/logic/diskqueue"
	_ "github.com/tiger1103/gfast/v3/internal/app/mqueue/logic/mqueue"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/model"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/service"
	"sync"
	"testing"
	"time"
)

var mBody = []byte("gfast-mqueue 测试消息队列内容")
var wg = sync.WaitGroup{}

const (
	TOPIC     = "producer_test6"
	SENDCOUNT = 10
	// 最好等待10秒来刷盘或更新rocketmq消费偏移
	TIMEOUT = 10
)

// channel 在rocketmq中可用,delay 1-18 对应时间: 1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
func producer(topic string, delay consts.MsgDelayLevel, timeout time.Duration, t *testing.T) {
	mq := service.MQueue()
	if mq == nil {
		t.Error("get mQueue failed")
		return
	}
	fmt.Println("start send msg")
	t1 := time.Now()
	pCount := 0
	for i := 0; i < SENDCOUNT; i++ {
		msg := &model.MQSendMsg{
			Topic: topic,
			Body:  mBody,
		}
		if delay > 0 {
			msg.Delay = delay
			msg.SendMethod = consts.SendMsgDelay
		}
		err := mq.SendMsg(msg)
		if err != nil {
			t.Error(err)
			return
		}
		pCount++
	}
	fmt.Println(fmt.Sprintf("发送数据 %d 条,耗时：%f", pCount, time.Since(t1).Seconds()))
	// 如果是diskqueue至少等待一次刷盘时，避免未刷就退出主线程
	time.Sleep(timeout * time.Second)
}

func consumer(topic, channel string, timeout time.Duration, t *testing.T) {
	fmt.Println(fmt.Sprintf("消费者 %s %s 启动", topic, channel))
	cChan := make(chan []byte)
	cCount := 0
	go func() {
		for {
			select {
			case <-cChan:
				cCount++
			}
		}
	}()
	mq := service.MQueue()
	if mq == nil {
		t.Error("get mQueue failed")
	}
	// 订阅消息
	err := mq.Subscribe(topic, channel, func(m *model.MQMessage) error {
		if !bytes.Equal(mBody, m.Body) {
			fmt.Println(fmt.Sprintf("消费者1第 %d 条数据错误", cCount))
		}
		cChan <- m.Body
		//fmt.Println(m.ID, m.Timestamp)
		return nil
	})
	if err != nil {
		t.Error("消息订阅失败：" + err.Error())
		return
	}

	// 至少等待一次刷盘或同步消费偏移，避免未刷就退出主线程
	time.Sleep(timeout * time.Second)
	fmt.Println(fmt.Sprintf("%s %s 消费数据 %d 条", topic, channel, cCount))
}

// 测试生产者
func TestProducer(t *testing.T) {
	producer(TOPIC, 0, 5, t)
}

// 测试生产者
func TestProducerDelay(t *testing.T) {
	fmt.Println("开始发送延迟消息")
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		producer("produce_delay_test1", consts.MsgDelay5s, 5, t)
		wg.Done()
	}()
	go func() {
		consumer("produce_delay_test1", "channel", 20, t)
		wg.Done()
	}()
	wg.Wait()
}

// 测试消费者
func TestConsumer(t *testing.T) {
	consumer(TOPIC, "channel", 120, t)
}

// 测试多个消费者
func TestMultiConsumer(t *testing.T) {
	wg.Add(3)
	go func() {
		consumer(TOPIC, "channel", TIMEOUT, t)
		wg.Done()
	}()
	go func() {
		consumer(TOPIC, "channel1", TIMEOUT, t)
		wg.Done()
	}()
	go func() {
		consumer(TOPIC, "channel2", TIMEOUT, t)
		wg.Done()
	}()
	wg.Wait()
}

// 同时测试生产者和消费者
func TestProducerAndConsumer(t *testing.T) {
	wg.Add(4)
	go func() {
		producer(TOPIC, 0, 5, t)
		wg.Done()
	}()
	go func() {
		consumer(TOPIC, "channel", TIMEOUT, t)
		wg.Done()
	}()
	go func() {
		consumer(TOPIC, "channel2", TIMEOUT, t)
		wg.Done()
	}()
	go func() {
		consumer(TOPIC, "channel3", TIMEOUT, t)
		wg.Done()
	}()
	wg.Wait()
}

// 测试删除diskQueue 的所有本地文件
func TestClearDiskQueueFiles(t *testing.T) {
	diskqueue.Clear()
}
