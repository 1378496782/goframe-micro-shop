package consts

// OrderStatus 订单状态枚举
type OrderStatus int

const (
	_                         OrderStatus = iota
	OrderStatusPendingPayment             // 1 待支付
	OrderStatusPaid                       // 2 已支付待发货
	OrderStatusShipped                    // 3 已发货
	OrderStatusReceived                   // 4 已收货待评价
	OrderStatusCompleted                  // 5 已评价
	OrderStatusPendingConfirm             // 6 待确认 (使用优惠券)
	OrderStatusCancelled                  // 7 已取消
)
