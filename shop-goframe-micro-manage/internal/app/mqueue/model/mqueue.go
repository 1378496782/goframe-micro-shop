/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/3 8:57
 */

package model

import (
	"github.com/tiger1103/gfast/v3/internal/app/mqueue/consts"
)

type MQConsumerHandlerCallback func(m *MQMessage) error

type MQSendMsg struct {
	// 主题
	Topic string
	// 消息体
	Body []byte
	// 发到频道
	Channel string
	// 消息发送类型
	SendMethod consts.SendMsgMethod
	// delayPublish 使用此参数，用于设置延迟消息等级，等级只能使用整数1-18 对应时间: 1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
	Delay consts.MsgDelayLevel
}

// MQMessage 消息返回结构体
type MQMessage struct {
	ID        MessageID
	Body      []byte
	Attempts  uint16
	Timestamp int64
}

type MessageID [consts.MsgIdLength]byte
