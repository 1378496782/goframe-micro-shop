package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// 退款状态常量
const (
	// 退款状态：待处理
	RefundStatusPending = "PENDING"
	// 退款状态：处理中
	RefundStatusProcessing = "PROCESSING"
	// 退款状态：退款成功
	RefundStatusSuccess = "SUCCESS"
	// 退款状态：退款失败
	RefundStatusFailed = "FAILED"
	// 退款状态：退款关闭
	RefundStatusClosed = "CLOSED"
	// 退款状态：退款异常
	RefundStatusError = "ERROR"
)

// 退款原因常量
const (
	// 退款原因：用户申请退款
	RefundReasonUserRequest = "用户申请退款"
	// 退款原因：商品质量问题
	RefundReasonQualityIssue = "商品质量问题"
	// 退款原因：商品缺货
	RefundReasonOutOfStock = "商品缺货"
	// 退款原因：多付退款
	RefundReasonOverpayment = "多付退款"
	// 退款原因：其他原因
	RefundReasonOther = "其他原因"
)

// 微信退款相关配置
const (
	// 微信退款超时时间（秒）
	WechatRefundTimeout = 300
	// 微信退款查询间隔（秒）
	WechatRefundQueryInterval = 5
	// 微信退款最大查询次数
	WechatRefundMaxQueryAttempts = 10
)

// 获取微信退款配置
func GetWechatRefundConfig() g.Map {
	ctx := gctx.New()
	refundNotifyUrl := ""
	mchID := ""
	apiV3Key := ""

	if val, err := g.Cfg().Get(ctx, "payment.wechat.refundNotifyUrl"); err == nil && val.String() != "" {
		refundNotifyUrl = val.String()
	}

	if val, err := g.Cfg().Get(ctx, "payment.wechat.mchID"); err == nil && val.String() != "" {
		mchID = val.String()
	}

	if val, err := g.Cfg().Get(ctx, "payment.wechat.apiV3Key"); err == nil && val.String() != "" {
		apiV3Key = val.String()
	}

	return g.Map{
		"refundNotifyUrl":  refundNotifyUrl,
		"mchID":            mchID,
		"apiV3Key":         apiV3Key,
		"timeout":          WechatRefundTimeout,
		"queryInterval":    WechatRefundQueryInterval,
		"maxQueryAttempts": WechatRefundMaxQueryAttempts,
	}
}

// 验证退款状态是否有效
func IsValidRefundStatus(status string) bool {
	validStatuses := map[string]bool{
		RefundStatusPending:    true,
		RefundStatusProcessing: true,
		RefundStatusSuccess:    true,
		RefundStatusFailed:     true,
		RefundStatusClosed:     true,
		RefundStatusError:      true,
	}
	return validStatuses[status]
}

// 获取退款状态显示文本
func GetRefundStatusText(status string) string {
	statusMap := map[string]string{
		RefundStatusPending:    "待处理",
		RefundStatusProcessing: "处理中",
		RefundStatusSuccess:    "退款成功",
		RefundStatusFailed:     "退款失败",
		RefundStatusClosed:     "退款关闭",
		RefundStatusError:      "退款异常",
	}
	if text, ok := statusMap[status]; ok {
		return text
	}
	return "未知状态"
}
