package model

import "time"

// FlashSaleOrderMessage 秒杀订单消息
type FlashSaleOrderMessage struct {
	OrderId     string    `json:"order_id"`     // 订单ID
	UserId      uint32    `json:"user_id"`      // 用户ID
	GoodsId     uint32    `json:"goods_id"`     // 秒杀商品ID
	Count       uint32    `json:"count"`        // 购买数量
	Price       uint64    `json:"price"`        // 秒杀价格
	CreateTime  time.Time `json:"create_time"`  // 创建时间
	Status      uint32    `json:"status"`       // 状态：1-待处理，2-处理中，3-成功，4-失败
	RetryCount  uint32    `json:"retry_count"`  // 重试次数
	MaxRetry    uint32    `json:"max_retry"`    // 最大重试次数
}