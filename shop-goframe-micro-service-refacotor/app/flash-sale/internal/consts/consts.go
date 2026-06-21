package consts

const (
	// 缓存相关
	FlashSaleGoodsCacheKey  = "flash_sale:goods:%d"  // 秒杀商品缓存
	FlashSaleStockCacheKey  = "flash_sale:stock:%d"  // 秒杀库存缓存
	FlashSaleResultCacheKey = "flash_sale:result:%s" // 秒杀结果缓存

	// 限流相关
	FlashSaleUserRateLimitKey = "flash_sale:rate_limit:%d"    // 用户限流
	FlashSaleIPRateLimitKey   = "flash_sale:ip_rate_limit:%s" // IP限流

	// 防刷相关
	FlashSaleUserBehaviorKey  = "flash_sale:behavior:%d:%s"    // 用户行为记录
	FlashSaleUserBlackListKey = "flash_sale:blacklist:user:%d" // 用户黑名单
	FlashSaleIPBlackListKey   = "flash_sale:blacklist:ip:%s"   // IP黑名单

	// 消息队列
	FlashSaleExchange        = "flash_sale.exchange"    // 秒杀交换机
	FlashSaleOrderQueue      = "flash_sale.order.queue" // 秒杀订单队列
	FlashSaleOrderRoutingKey = "flash_sale.order"       // 秒杀订单路由键

	// 状态定义
	FlashSaleStatusPending = 0 // 处理中
	FlashSaleStatusSuccess = 1 // 成功
	FlashSaleStatusFailed  = 2 // 失败

	// 业务限制
	MaxFlashSaleCountPerUser    = 1 // 每用户限购数量
	FlashSaleRateLimitPerSecond = 5 // 每秒限流次数

	// 限流配置
	FlashSaleIPRateLimitPerSecond     = 10  // IP每秒限流次数
	FlashSaleGlobalRateLimitPerSecond = 100 // 全局限流次数
	FlashSaleUserMinuteRateLimit      = 10  // 用户每分钟限流次数

	// 防刷配置
	MaxUserRequestsPerMinute   = 20          // 用户每分钟最大请求数
	MaxIPRequestsPerMinute     = 50          // IP每分钟最大请求数
	SuspiciousRequestThreshold = 100         // 可疑请求阈值
	BlackListExpireTime        = 1 * 60 * 60 // 黑名单过期时间（1小时）
)
