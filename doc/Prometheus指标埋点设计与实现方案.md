# Prometheus指标埋点设计与实现方案

## 1. 设计目标

本方案旨在为Goframe微服务架构提供最小化的Prometheus指标埋点实现，包括基础HTTP请求指标和自定义业务指标，以便于服务监控、性能分析和业务状态追踪。

## 2. 设计原则

- **最小化实现**：只包含最核心的指标功能，避免过度设计
- **共享复用**：将指标功能作为公共组件，便于各服务复用
- **易扩展**：设计灵活的业务指标接口，支持后续扩展
- **性能优先**：确保指标采集对服务性能影响最小

## 3. 目录结构

```
/utility/metrics/
├── metrics.go        # 基础指标定义与初始化
├── middleware.go     # HTTP中间件，用于自动采集请求指标
└── business.go       # 业务指标接口，用于自定义业务指标采集
```

## 4. 功能实现

### 4.1 基础指标定义与初始化 (metrics.go)

定义了服务级别的基础指标，包括：

- **请求计数指标**：按请求方法、路径和状态码分类统计
- **请求延迟指标**：记录请求处理耗时的分布情况
- **错误计数指标**：记录服务错误情况

```go
// 核心指标定义\var (
	// 请求计数指标
	RequestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status_code"},
	)

	// 请求延迟指标
	RequestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// 错误计数指标
	ErrorCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_errors_total",
			Help: "Total number of service errors",
		},
		[]string{"error_type", "service"},
	)
)
```

### 4.2 HTTP中间件 (middleware.go)

提供两个核心中间件：

1. **MetricsMiddleware**：自动记录所有HTTP请求的方法、路径、状态码和处理时间
2. **ErrorMetricsMiddleware**：记录服务错误信息，包括错误类型和服务名称

```go
// MetricsMiddleware Prometheus指标收集中间件
func MetricsMiddleware(r *ghttp.Request) {
	// 记录请求开始时间
	startTime := time.Now()

	// 执行下一个中间件/处理函数
	r.Middleware.Next()

	// 计算请求处理时间
	duration := time.Since(startTime).Seconds()

	// 记录请求指标
	method := r.Method
	endpoint := r.RequestURI
	statusCode := r.Response.Status

	// 更新请求计数
	RequestCount.WithLabelValues(method, endpoint, statusCode).Inc()

	// 更新请求延迟分布
	RequestLatency.WithLabelValues(method, endpoint).Observe(duration)
}
```

### 4.3 业务指标接口 (business.go)

提供业务级别的指标接口，支持自定义业务指标的采集，包括：

- **订单创建计数**：统计订单创建的成功和失败次数
- **订单成功率**：记录订单创建的成功率
- **库存指标**：监控商品库存变化

```go
// RecordOrderCreate 记录订单创建事件
func RecordOrderCreate(ctx context.Context, success bool) {
	status := "failed"
	if success {
		status = "success"
	}
	OrderCreateCount.WithLabelValues(status).Inc()
}

// UpdateInventory 更新库存指标
func UpdateInventory(ctx context.Context, productID string, quantity int64) {
	ProductInventory.WithLabelValues(productID).Set(float64(quantity))
}
```

## 5. 集成方案

### 5.1 在服务中集成

在服务的main.go中集成Prometheus指标功能：

```go
// 初始化Prometheus指标
metrics.InitMetrics()

// 注册Prometheus中间件	s.Use(metrics.MetricsMiddleware)
	s.Use(metrics.ErrorMetricsMiddleware)

// 设置Prometheus指标端点
metrics.RegisterHTTPHandler(s)
```

### 5.2 在业务代码中使用

在业务逻辑中调用业务指标接口：

```go
// 订单创建成功时记录指标
metrics.RecordOrderCreate(ctx, true)

// 更新库存指标
metrics.UpdateInventory(ctx, productID, newStock)
```

## 6. 测试方法

### 6.1 编译验证

```bash
# 下载依赖
go mod tidy

# 编译服务
go build -o /dev/null ./app/order/internal/logic/order_info
go build -o /dev/null ./app/gateway-h5
```

### 6.2 运行时验证

1. 启动服务
2. 访问 `/metrics` 端点查看指标数据
3. 执行业务操作，观察指标变化

### 6.3 验证指标数据

示例Prometheus查询：

```
# 查询总请求数
sum(http_requests_total)

# 查询订单成功率
http_requests_total{status="success"} / http_requests_total

# 查询P95请求延迟
histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[5m])) by (le))
```

## 7. 扩展建议

1. 根据业务需求添加更多自定义指标
2. 为不同服务设置服务特定标签
3. 考虑添加告警规则，监控关键业务指标
4. 与Grafana集成，创建可视化监控面板

## 8. 注意事项

1. 指标标签值应避免高基数（如用户ID、订单号等）
2. 业务指标采集应在关键节点进行，避免过度采集
3. 生产环境部署时，应考虑指标采集对性能的影响
4. 定期清理过期指标，避免内存占用过高