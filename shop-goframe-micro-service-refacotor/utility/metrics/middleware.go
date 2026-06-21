package metrics

import (
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
)

// MetricsMiddleware Prometheus指标收集中间件
func MetricsMiddleware(r *ghttp.Request) {
	// 记录请求开始时间
	startTime := time.Now()

	// 执行后续处理
	r.Middleware.Next()

	// 计算请求耗时
	duration := time.Since(startTime)

	// 获取请求路径，为了避免高基数问题，可以使用路由模式代替具体路径
	path := r.Router.Uri
	if path == "" {
		path = r.URL.Path
	}

	// 记录请求指标
	RecordRequest(
		r.Context(),
		r.Method,
		path,
		r.Response.Status,
		duration,
	)
}

// ErrorMetricsMiddleware 错误指标收集中间件
func ErrorMetricsMiddleware(r *ghttp.Request) {
	// 执行后续处理
	r.Middleware.Next()

	// 检查是否有错误
	if r.GetError() != nil {
		// 获取服务名称（可以从配置或上下文中获取）
		serviceName := "unknown_service"
		if service := r.Context().Value("service_name"); service != nil {
			if name, ok := service.(string); ok {
				serviceName = name
			}
		}

		// 记录错误指标
		errorType := "general_error"
		if r.Response.Status >= 500 {
			errorType = "server_error"
		} else if r.Response.Status >= 400 {
			errorType = "client_error"
		}

		RecordError(errorType, serviceName)
	}
}
