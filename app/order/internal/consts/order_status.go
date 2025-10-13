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
	OrderStatusRefund                     // 8 发起退款
)

// RefundStatus 对应字段：refund_info.status。审核状态
type RefundStatus int

const (
	_                    RefundStatus = iota
	RefundStatusPending               // 1 待处理（用户已申请，等待审核）
	RefundStatusApproved              // 2 同意退款（审核通过）
	RefundStatusRejected              // 3 拒绝退款（审核驳回）
)

// RefundOrderStatus 对应字段：refund_info.refund_status。退款状态
type RefundOrderStatus int

const (
	_                           RefundOrderStatus = iota // 0 未退款（初始状态）
	RefundOrderStatusNone                                // 0 未退款（初始状态）
	RefundOrderStatusProcessing                          // 1 退款中（已提交至支付平台）
	RefundOrderStatusSuccess                             // 2 退款成功
	RefundOrderStatusFailed                              // 3 退款失败
)
