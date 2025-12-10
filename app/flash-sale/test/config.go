package test

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/flash-sale/utility"
	"github.com/gogf/gf/v2/frame/g"
)

// TestConfig 测试配置
var TestConfig = struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	RabbitMQAddr  string
	RabbitMQUser  string
	RabbitMQPass  string
}{
	RedisAddr:     "localhost:6379",
	RedisPassword: "",
	RedisDB:       1, // 使用不同的DB避免影响正式数据
	RabbitMQAddr:  "localhost:5672",
	RabbitMQUser:  "guest",
	RabbitMQPass:  "guest",
}

// InitTestConfig 初始化测试配置
func InitTestConfig(ctx context.Context) {
	// 设置测试模式
	g.Cfg().Set(ctx, "redis.default.address", TestConfig.RedisAddr)
	g.Cfg().Set(ctx, "redis.default.password", TestConfig.RedisPassword)
	g.Cfg().Set(ctx, "redis.default.db", TestConfig.RedisDB)
	
	g.Cfg().Set(ctx, "rabbitmq.default.address", TestConfig.RabbitMQAddr)
	g.Cfg().Set(ctx, "rabbitmq.default.user", TestConfig.RabbitMQUser)
	g.Cfg().Set(ctx, "rabbitmq.default.pass", TestConfig.RabbitMQPass)
	
	g.Log().Info(ctx, "测试配置初始化完成")
}

// CleanupTestData 清理测试数据
func CleanupTestData(ctx context.Context) {
	// 获取缓存实例
	cache := utility.GetFlashSaleCache()
	if cache == nil {
		g.Log().Warning(ctx, "缓存实例为空，无法清理测试数据")
		return
	}
	
	// 清理限流相关数据
	pattern := "flash_sale:*"
	keys, err := cache.Keys(ctx, pattern)
	if err != nil {
		g.Log().Warning(ctx, "获取测试数据键失败:", err)
		return
	}
	
	for _, key := range keys {
		if err := cache.Delete(ctx, key); err != nil {
			g.Log().Warningf(ctx, "删除测试数据失败，键：%s，错误：%v", key, err)
		}
	}
	
	g.Log().Infof(ctx, "测试数据清理完成，清理了 %d 个键", len(keys))
}

// GenerateTestData 生成测试数据
func GenerateTestData(ctx context.Context) map[string]interface{} {
	return map[string]interface{}{
		"goods": []map[string]interface{}{
			{
				"id":          20001,
				"name":        "测试商品1",
				"description": "测试商品描述1",
				"price":       99900,
				"stock":       100,
				"status":      1,
			},
			{
				"id":          20002,
				"name":        "测试商品2",
				"description": "测试商品描述2",
				"price":       199900,
				"stock":       50,
				"status":      1,
			},
			{
				"id":          20003,
				"name":        "测试商品3",
				"description": "测试商品描述3",
				"price":       299900,
				"stock":       200,
				"status":      1,
			},
		},
		"users": []map[string]interface{}{
			{
				"id":       10001,
				"username": "testuser1",
				"email":    "test1@example.com",
				"status":   1,
			},
			{
				"id":       10002,
				"username": "testuser2",
				"email":    "test2@example.com",
				"status":   1,
			},
			{
				"id":       10003,
				"username": "testuser3",
				"email":    "test3@example.com",
				"status":   1,
			},
		},
	}
}