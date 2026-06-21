// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderOutboxMessage is the golang structure for table order_outbox_message.
type OrderOutboxMessage struct {
	Id          int64       `json:"id"          orm:"id"            description:"主键id"`                       // 主键id
	EventId     string      `json:"eventId"     orm:"event_id"      description:"事件唯一id，用于幂等去重"`              // 事件唯一id，用于幂等去重
	EventType   string      `json:"eventType"   orm:"event_type"    description:"事件类型"`                       // 事件类型
	AggregateId string      `json:"aggregateId" orm:"aggregate_id"  description:"聚合根id，如订单id"`                // 聚合根id，如订单id
	Exchange    string      `json:"exchange"    orm:"exchange"      description:"消息投递的 exchange"`             // 消息投递的 exchange
	RoutingKey  string      `json:"routingKey"  orm:"routing_key"   description:"消息投递的 routing key"`          // 消息投递的 routing key
	Payload     string      `json:"payload"     orm:"payload"       description:"消息体内容"`                      // 消息体内容
	Status      int         `json:"status"      orm:"status"        description:"发送状态：0待发送 1发送中 2发送成功 3发送失败"` // 发送状态：0待发送 1发送中 2发送成功 3发送失败
	RetryCount  int         `json:"retryCount"  orm:"retry_count"   description:"已重试次数"`                      // 已重试次数
	NextRetryAt *gtime.Time `json:"nextRetryAt" orm:"next_retry_at" description:"下次重试时间"`                     // 下次重试时间
	LastError   string      `json:"lastError"   orm:"last_error"    description:"最近一次失败原因"`                   // 最近一次失败原因
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:"创建时间"`                       // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:"更新时间"`                       // 更新时间
	SentAt      *gtime.Time `json:"sentAt"      orm:"sent_at"       description:"发送成功时间"`                     // 发送成功时间
}
