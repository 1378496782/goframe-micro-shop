// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderOutboxMessageDao is the data access object for the table order_outbox_message.
type OrderOutboxMessageDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  OrderOutboxMessageColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// OrderOutboxMessageColumns defines and stores column names for the table order_outbox_message.
type OrderOutboxMessageColumns struct {
	Id          string // 主键id
	EventId     string // 事件唯一id，用于幂等去重
	EventType   string // 事件类型
	AggregateId string // 聚合根id，如订单id
	Exchange    string // 消息投递的 exchange
	RoutingKey  string // 消息投递的 routing key
	Payload     string // 消息体内容
	Status      string // 发送状态：0待发送 1发送中 2发送成功 3发送失败
	RetryCount  string // 已重试次数
	NextRetryAt string // 下次重试时间
	LastError   string // 最近一次失败原因
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	SentAt      string // 发送成功时间
}

// DEMO_WECHAT_OPEN_ID holds the columns for the table order_outbox_message.
var DEMO_WECHAT_OPEN_ID = OrderOutboxMessageColumns{
	Id:          "id",
	EventId:     "event_id",
	EventType:   "event_type",
	AggregateId: "aggregate_id",
	Exchange:    "exchange",
	RoutingKey:  "routing_key",
	Payload:     "payload",
	Status:      "status",
	RetryCount:  "retry_count",
	NextRetryAt: "next_retry_at",
	LastError:   "last_error",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	SentAt:      "sent_at",
}

// NewOrderOutboxMessageDao creates and returns a new DAO object for table data access.
func NewOrderOutboxMessageDao(handlers ...gdb.ModelHandler) *OrderOutboxMessageDao {
	return &OrderOutboxMessageDao{
		group:    "default",
		table:    "order_outbox_message",
		columns:  DEMO_WECHAT_OPEN_ID,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrderOutboxMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrderOutboxMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrderOutboxMessageDao) Columns() OrderOutboxMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrderOutboxMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrderOutboxMessageDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OrderOutboxMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
