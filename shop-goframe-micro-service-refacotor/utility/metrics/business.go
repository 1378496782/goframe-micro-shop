package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// 业务指标定义
var (
	// 订单创建计数
	OrderCreateCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "business_order_create_total",
			Help: "Total number of order creation attempts",
		},
		[]string{"status"},
	)

	// 订单创建成功率
	OrderSuccessRatio = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "business_order_success_ratio",
			Help: "Ratio of successful orders",
		},
	)

	// 库存指标
	InventoryGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "business_inventory_count",
			Help: "Current inventory count",
		},
		[]string{"product_id"},
	)
)

// 业务指标记录函数

// RecordOrderCreate 记录订单创建
func RecordOrderCreate(ctx context.Context, success bool) {
	status := "failed"
	if success {
		status = "success"
	}
	OrderCreateCount.WithLabelValues(status).Inc()
}

// UpdateOrderSuccessRatio 更新订单成功率
func UpdateOrderSuccessRatio(ctx context.Context, ratio float64) {
	OrderSuccessRatio.Set(ratio)
}

// UpdateInventory 更新库存指标
func UpdateInventory(ctx context.Context, productID string, count int64) {
	InventoryGauge.WithLabelValues(productID).Set(float64(count))
}

// 通用业务指标接口，允许各服务自定义指标

// BusinessMetric 业务指标接口
// 可以用于创建更复杂的业务指标收集逻辑
// 这个接口为将来的扩展提供了基础

// 示例：使用方法
// 在服务中直接调用：
// metrics.RecordOrderCreate(ctx, true)
// metrics.UpdateInventory(ctx, "product123", 100)
