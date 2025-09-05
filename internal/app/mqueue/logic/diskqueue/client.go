/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/7 14:54
 */

package diskqueue

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/consts"
	disk "github.com/tiger1103/gfast/v3/internal/app/mqueue/driver"
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/model"
	"io"
	"os"
	"sync"
	"time"
)

func logger(ctx context.Context) disk.AppLogFunc {
	return func(lvl disk.LogLevel, f string, args ...interface{}) {
		switch lvl {
		case disk.FATAL:
			g.Log("diskQueue").Fatal(ctx, fmt.Sprintf(f, args...))
		case disk.DEBUG:
			g.Log("diskQueue").Debug(ctx, fmt.Sprintf(f, args...))
		case disk.INFO:
			g.Log("diskQueue").Info(ctx, fmt.Sprintf(f, args...))
		case disk.WARN:
			g.Log("diskQueue").Warning(ctx, fmt.Sprintf(f, args...))
		case disk.ERROR:
			g.Log("diskQueue").Error(ctx, fmt.Sprintf(f, args...))
		}
	}
}

type client struct {
	name             string
	diskQueue        disk.IDiskQueue
	channelConsumers map[string][]*consumer
	exitChan         chan struct{}
	logger           disk.AppLogFunc
	sync.Mutex
}

// 磁盘消息队列实例
var clientMap = make(map[string]*client)
var clientMutex sync.Mutex

// registerDiskQueue 创建磁盘队列的实例
func registerDiskQueue(topic string) error {
	if _, ok := clientMap[topic]; ok {
		return errors.New(fmt.Sprintf("实例 %s 已经存在", topic))
	}
	var dc *model.DiskQueueConfig
	err := g.Cfg().MustGet(context.TODO(), "mqueue.diskQueue").Struct(&dc)
	if err != nil {
		return err
	}
	// 数据文件夹不存在则创建
	if stat, err2 := os.Stat(dc.DataPath); stat == nil || os.IsNotExist(err2) {
		err3 := os.MkdirAll(dc.DataPath, os.ModePerm)
		if err3 != nil {
			return err3
		}
	}

	clientMap[topic] = &client{
		name:             topic,
		diskQueue:        disk.NewDiskQueue(topic, dc.DataPath, dc.MaxBytesPerFile, dc.MinMsgSize, dc.MaxMsgSize, dc.SyncEvery, dc.SyncTimeout*time.Second, logger(context.TODO())),
		channelConsumers: make(map[string][]*consumer),
		exitChan:         make(chan struct{}),
		logger:           logger(context.TODO()),
	}
	// 开始消息循环
	go clientMap[topic].start()
	return nil
}

// 根据提供的topic 返回一个已注册的diskQueue实例，如果实例不存在，则创建一个返回
func getClient(topic string) (c *client, err error) {
	clientMutex.Lock()
	defer clientMutex.Unlock()
	if _, ok := clientMap[topic]; !ok {
		err = registerDiskQueue(topic)
		if err != nil {
			return nil, err
		}
	}
	return clientMap[topic], nil
}

// RegisterConsumer 注册消费者到对应频道
func (c *client) RegisterConsumer(channel string, consumer *consumer) {
	c.Lock()
	defer c.Unlock()
	c.channelConsumers[channel] = append(c.channelConsumers[channel], consumer)
}

// RemoveConsumer 移除消费者
func (c *client) RemoveConsumer(channel, id string) {
	c.Lock()
	c.Unlock()
	for i, ch := range c.channelConsumers[channel] {
		if ch.Id == id {
			if len(c.channelConsumers[channel]) <= 1 {
				delete(c.channelConsumers, channel)
				// 用空字符清理一下start中的ReadChan()管道
				c.diskQueue.Put([]byte(""))
			} else {
				c.channelConsumers[channel] = append(c.channelConsumers[channel][:i], c.channelConsumers[channel][i+1:]...)
			}
		}
	}
}

// 将消息MQMessage转换为[]byte
func (c *client) messageToByte(m *model.MQMessage, w io.Writer) (int64, error) {
	var buf [10]byte
	var total int64

	binary.BigEndian.PutUint64(buf[:8], uint64(m.Timestamp))
	binary.BigEndian.PutUint16(buf[8:10], uint16(m.Attempts))

	n, err := w.Write(buf[:])
	total += int64(n)
	if err != nil {
		return total, err
	}

	n, err = w.Write(m.ID[:])
	total += int64(n)
	if err != nil {
		return total, err
	}

	n, err = w.Write(m.Body)
	total += int64(n)
	if err != nil {
		return total, err
	}

	return total, nil
}

// Publish 生产者调用此方法进行消息发布
func (c *client) Publish(body []byte) error {
	var id model.MessageID
	copy(id[:], grand.B(16))
	m := c.newMessage(id, body)
	var b bytes.Buffer
	total, err := c.messageToByte(m, &b)
	if err != nil {
		return err
	}
	if total == 0 {
		return errors.New("发布信息失败，s.WriteTo 长度为 0")
	}
	return c.diskQueue.Put(b.Bytes())
}

// start 开始投递消息给消费者
func (c *client) start() {
	for {
		select {
		case m := <-c.diskQueue.ReadChan():
			if len(m) <= 0 {
				break
			}
			message, err := c.decodeMessage(m)
			if err != nil {
				c.logger(disk.ERROR, err.Error())
			}
			if len(c.channelConsumers) > 0 {
				// 消息广播到所有频道
				for _, channel := range c.channelConsumers {
					// 广播到当前频道下的所有消费者
					for _, ch := range channel {
						err = ch.Handler(message)
						if err != nil {
							c.logger(disk.ERROR, err.Error())
						}
					}
				}
			}
		case <-c.exitChan:
			g.Log().Debug(gctx.New(), "退出消息队列")
			return
		}
	}
}

// 解码消息
// message format:
//
//	[x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x][x]...
//	|       (int64)        ||    ||      (hex string encoded in ASCII)           || (binary)
//	|       8-byte         ||    ||                 16-byte                      || N-byte
//	------------------------------------------------------------------------------------------...
//	  nanosecond timestamp    ^^                   message ID                       message body
//	                       (uint16)
//	                        2-byte
//	                       attempts
func (c *client) decodeMessage(b []byte) (*model.MQMessage, error) {
	var msg model.MQMessage

	if len(b) < 10+consts.MsgIdLength {
		return nil, errors.New("not enough data to decode valid message")
	}

	msg.Timestamp = int64(binary.BigEndian.Uint64(b[:8]))
	msg.Attempts = binary.BigEndian.Uint16(b[8:10])
	copy(msg.ID[:], b[10:10+consts.MsgIdLength])
	msg.Body = b[10+consts.MsgIdLength:]

	return &msg, nil
}

// newMessage 创建消息
func (c *client) newMessage(id model.MessageID, body []byte) *model.MQMessage {
	return &model.MQMessage{
		ID:        id,
		Body:      body,
		Timestamp: time.Now().UnixNano(),
	}
}

func CloseDiskQueueService() {
	for _, v := range clientMap {
		v.diskQueue.Close()
		v.exitChan <- struct{}{}
	}
}

// Clear 清空所有磁盘文件，慎用
func Clear() {
	path := g.Cfg().MustGet(context.TODO(), "mqueue.diskQueue.dataPath").String()
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println(err)
	}
}
