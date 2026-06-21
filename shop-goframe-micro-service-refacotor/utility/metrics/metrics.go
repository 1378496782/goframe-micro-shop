package metrics

import (
	"context"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 基础指标定义
var (
	// 请求计数指标
	RequestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// 请求延迟指标
	RequestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
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

// RegisterHTTPHandler 注册Prometheus HTTP处理函数
func RegisterHTTPHandler(server *ghttp.Server) {
	// 在ghttp服务器上注册/metrics端点
	server.Group("/metrics", func(group *ghttp.RouterGroup) {
		group.GET("", func(r *ghttp.Request) {
			// 使用ghttp的Response对象直接写入Prometheus指标
			w := r.Response.Writer
			promhttp.Handler().ServeHTTP(w, r.Request)
		})
	})
}

// InitMetrics 初始化指标
func InitMetrics() {
	// 指标已经在定义时通过promauto自动注册，这里可以添加额外的初始化逻辑
}

// RecordRequest 记录HTTP请求
func RecordRequest(ctx context.Context, method, path string, status int, duration time.Duration) {
	RequestCount.WithLabelValues(method, path, http.StatusText(status)).Inc()
	RequestLatency.WithLabelValues(method, path, http.StatusText(status)).Observe(duration.Seconds())
}

// RecordError 记录错误
func RecordError(errorType, service string) {
	ErrorCount.WithLabelValues(errorType, service).Inc()
}