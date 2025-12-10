package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/consts"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/logic"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/model"
	"shop-goframe-micro-service-refacotor/app/flash-sale/utility"
	v1 "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

// TestFlashSaleSystem 测试秒杀系统整体功能
func TestFlashSaleSystem(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.Background()
		
		// 初始化测试环境
		initTestEnv(ctx)
		
		// 测试1：基本秒杀流程
		t.Run("BasicFlashSaleFlow", func(t *testing.T) {
			testBasicFlashSaleFlow(ctx, t)
		})
		
		// 测试2：限流机制测试
		t.Run("RateLimitTest", func(t *testing.T) {
			testRateLimit(ctx, t)
		})
		
		// 测试3：防刷机制测试
		t.Run("AntiBrushTest", func(t *testing.T) {
			testAntiBrush(ctx, t)
		})
		
		// 测试4：并发秒杀测试
		t.Run("ConcurrentFlashSale", func(t *testing.T) {
			testConcurrentFlashSale(ctx, t)
		})
		
		// 测试5：库存管理测试
		t.Run("StockManagementTest", func(t *testing.T) {
			testStockManagement(ctx, t)
		})
		
		// 测试6：消息队列测试
		t.Run("MessageQueueTest", func(t *testing.T) {
			testMessageQueue(ctx, t)
		})
	})
}

// initTestEnv 初始化测试环境
func initTestEnv(ctx context.Context) {
	// 初始化缓存
	cache := utility.GetFlashSaleCache()
	if cache == nil {
		g.Log().Error(ctx, "缓存未初始化")
		return
	}
	
	// 初始化库存管理器
	stockManager := utility.GetFlashSaleStockManager()
	if stockManager == nil {
		g.Log().Error(ctx, "库存管理器未初始化")
		return
	}
	
	// 初始化消息队列
	mq := utility.GetRabbitMQ()
	if mq == nil {
		g.Log().Error(ctx, "消息队列未初始化")
		return
	}
	
	g.Log().Info(ctx, "测试环境初始化完成")
}

// testBasicFlashSaleFlow 测试基本秒杀流程
func testBasicFlashSaleFlow(ctx context.Context, t *gtest.T) {
	// 准备测试数据
	userId := uint32(10001)
	goodsId := uint32(20001)
	initialStock := 100
	
	// 初始化商品库存
	stockManager := utility.GetFlashSaleStockManager()
	err := stockManager.InitStock(ctx, goodsId, initialStock)
	t.AssertNil(err)
	
	// 创建秒杀服务
	flashSale := logic.NewFlashSale()
	
	// 创建秒杀订单请求
	req := &v1.CreateFlashSaleOrderReq{
		UserId:  userId,
		GoodsId: goodsId,
		Count:   1,
	}
	
	// 执行秒杀
	resp, err := flashSale.CreateFlashSaleOrder(ctx, req)
	t.AssertNil(err)
	t.AssertEQ(resp.Success, true)
	t.AssertNE(resp.OrderNo, "")
	t.AssertEQ(resp.Status, consts.FlashSaleStatusSuccess)
	
	// 验证库存减少
	remainingStock, err := stockManager.GetStock(ctx, goodsId)
	t.AssertNil(err)
	t.AssertEQ(remainingStock, initialStock-1)
	
	// 查询秒杀结果
	resultReq := &v1.GetFlashSaleResultReq{
		ResultId: resp.ResultId,
		UserId:   userId,
	}
	resultResp, err := flashSale.GetFlashSaleResult(ctx, resultReq)
	t.AssertNil(err)
	t.AssertEQ(resultResp.Status, consts.FlashSaleStatusSuccess)
	t.AssertEQ(resultResp.OrderNo, resp.OrderNo)
	
	g.Log().Infof(ctx, "基本秒杀流程测试通过，订单号：%s", resp.OrderNo)
}

// testRateLimit 测试限流机制
func testRateLimit(ctx context.Context, t *gtest.T) {
	userId := uint32(10002)
	goodsId := uint32(20002)
	
	// 初始化库存
	stockManager := utility.GetFlashSaleStockManager()
	err := stockManager.InitStock(ctx, goodsId, 1000)
	t.AssertNil(err)
	
	flashSale := logic.NewFlashSale()
	
	// 测试用户限流 - 每秒最多5次请求
	successCount := 0
	for i := 0; i < 10; i++ {
		req := &v1.CreateFlashSaleOrderReq{
			UserId:  userId,
			GoodsId: goodsId,
			Count:   1,
		}
		
		resp, err := flashSale.CreateFlashSaleOrder(ctx, req)
		if err == nil && resp.Success {
			successCount++
		}
	}
	
	// 应该只有前5次请求成功
	t.AssertLTE(successCount, 5)
	
	// 等待1秒后再次请求，应该成功
	time.Sleep(1 * time.Second)
	req := &v1.CreateFlashSaleOrderReq{
		UserId:  userId,
		GoodsId: goodsId,
		Count:   1,
	}
	resp, err := flashSale.CreateFlashSaleOrder(ctx, req)
	t.AssertNil(err)
	t.AssertEQ(resp.Success, true)
	
	g.Log().Infof(ctx, "限流机制测试通过，成功请求数：%d", successCount)
}

// testAntiBrush 测试防刷机制
func testAntiBrush(ctx context.Context, t *gtest.T) {
	userId := uint32(10003)
	goodsId := uint32(20003)
	
	// 初始化库存
	stockManager := utility.GetFlashSaleStockManager()
	err := stockManager.InitStock(ctx, goodsId, 100)
	t.AssertNil(err)
	
	// 获取缓存实例
	cache := utility.GetFlashSaleCache()
	
	// 模拟将用户加入黑名单
	blacklistKey := fmt.Sprintf(consts.FlashSaleUserBlackListKey, userId)
	err = cache.Set(ctx, blacklistKey, 1, 1*time.Hour)
	t.AssertNil(err)
	
	flashSale := logic.NewFlashSale()
	
	// 尝试秒杀，应该失败
	req := &v1.CreateFlashSaleOrderReq{
		UserId:  userId,
		GoodsId: goodsId,
		Count:   1,
	}
	
	resp, err := flashSale.CreateFlashSaleOrder(ctx, req)
	t.AssertNil(err)
	t.AssertEQ(resp.Success, false)
	t.AssertContains(resp.Message, "异常行为")
	
	// 从黑名单中移除用户
	err = cache.Delete(ctx, blacklistKey)
	t.AssertNil(err)
	
	// 再次尝试，应该成功
	resp, err = flashSale.CreateFlashSaleOrder(ctx, req)
	t.AssertNil(err)
	t.AssertEQ(resp.Success, true)
	
	g.Log().Infof(ctx, "防刷机制测试通过")
}

// testConcurrentFlashSale 测试并发秒杀
func testConcurrentFlashSale(ctx context.Context, t *gtest.T) {
	userId := uint32(10004)
	goodsId := uint32(20004)
	initialStock := 50
	concurrentUsers := 100
	
	// 初始化库存
	stockManager := utility.GetFlashSaleStockManager()
	err := stockManager.InitStock(ctx, goodsId, initialStock)
	t.AssertNil(err)
	
	flashSale := logic.NewFlashSale()
	
	// 使用WaitGroup协调并发测试
	var wg sync.WaitGroup
	successOrders := make([]string, 0)
	var mu sync.Mutex
	
	// 模拟并发秒杀
	for i := 0; i < concurrentUsers; i++ {
		wg.Add(1)
		go func(userIndex int) {
			defer wg.Done()
			
			req := &v1.CreateFlashSaleOrderReq{
				UserId:  userId + uint32(userIndex),
				GoodsId: goodsId,
				Count:   1,
			}
			
			resp, err := flashSale.CreateFlashSaleOrder(ctx, req)
			if err == nil && resp.Success {
				mu.Lock()
				successOrders = append(successOrders, resp.OrderNo)
				mu.Unlock()
			}
		}(i)
	}
	
	wg.Wait()
	
	// 验证成功订单数不超过初始库存
	t.AssertLTE(len(successOrders), initialStock)
	
	// 验证最终库存
	remainingStock, err := stockManager.GetStock(ctx, goodsId)
	t.AssertNil(err)
	t.AssertEQ(remainingStock, initialStock-len(successOrders))
	
	g.Log().Infof(ctx, "并发秒杀测试通过，成功订单数：%d，剩余库存：%d", len(successOrders), remainingStock)
}

// testStockManagement 测试库存管理
func testStockManagement(ctx context.Context, t *gtest.T) {
	goodsId := uint32(20005)
	initialStock := 100
	
	stockManager := utility.GetFlashSaleStockManager()
	
	// 测试初始化库存
	err := stockManager.InitStock(ctx, goodsId, initialStock)
	t.AssertNil(err)
	
	// 验证库存
	stock, err := stockManager.GetStock(ctx, goodsId)
	t.AssertNil(err)
	t.AssertEQ(stock, initialStock)
	
	// 测试扣减库存
	success, err := stockManager.DeductStock(ctx, goodsId, 10)
	t.AssertNil(err)
	t.AssertEQ(success, true)
	
	// 验证库存减少
	stock, err = stockManager.GetStock(ctx, goodsId)
	t.AssertNil(err)
	t.AssertEQ(stock, initialStock-10)
	
	// 测试超卖保护
	success, err = stockManager.DeductStock(ctx, goodsId, initialStock)
	t.AssertNil(err)
	t.AssertEQ(success, false) // 不应该成功，因为库存不足
	
	// 验证库存未变化
	stock, err = stockManager.GetStock(ctx, goodsId)
	t.AssertNil(err)
	t.AssertEQ(stock, initialStock-10)
	
	g.Log().Infof(ctx, "库存管理测试通过")
}

// testMessageQueue 测试消息队列
func testMessageQueue(ctx context.Context, t *gtest.T) {
	// 创建测试消息
	testMsg := &model.FlashSaleOrderMessage{
		OrderId:    "TEST_ORDER_001",
		UserId:     10005,
		GoodsId:    20005,
		Count:      1,
		Price:      99900,
		CreateTime: time.Now(),
		Status:     1,
		RetryCount: 0,
		MaxRetry:   3,
	}
	
	// 发布消息
	err := mq.PublishFlashSaleOrder(ctx, testMsg)
	t.AssertNil(err)
	
	g.Log().Infof(ctx, "消息队列测试通过，消息ID：%s", testMsg.OrderId)
}

// TestRateLimiter 单独测试限流器
func TestRateLimiter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.Background()
		
		// 获取缓存实例
		cache := utility.GetFlashSaleCache()
		if cache == nil {
			t.Skip("缓存未初始化")
		}
		
		rateLimiter := utility.NewRateLimiter(cache)
		
		// 测试用户限流
		t.Run("UserRateLimit", func(t *gtest.T) {
			userId := uint32(20001)
			
			// 第一次检查应该通过
			err := rateLimiter.CheckUserLimit(ctx, userId)
			t.AssertNil(err)
			
			// 连续快速请求应该触发限流
			for i := 0; i < 7; i++ {
				rateLimiter.CheckUserLimit(ctx, userId)
			}
			
			// 第8次应该失败
			err = rateLimiter.CheckUserLimit(ctx, userId)
			t.AssertNE(err, nil)
		})
		
		// 测试IP限流
		t.Run("IPRateLimit", func(t *gtest.T) {
			ip := "192.168.1.100"
			
			// 第一次检查应该通过
			err := rateLimiter.CheckIPLimit(ctx, ip)
			t.AssertNil(err)
			
			// 连续快速请求应该触发限流
			for i := 0; i < 12; i++ {
				rateLimiter.CheckIPLimit(ctx, ip)
			}
			
			// 第13次应该失败
			err = rateLimiter.CheckIPLimit(ctx, ip)
			t.AssertNE(err, nil)
		})
		
		// 测试购买限制
		t.Run("PurchaseLimit", func(t *gtest.T) {
			userId := uint32(20002)
			goodsId := uint32(30001)
			
			// 第一次购买检查应该通过
			err := rateLimiter.CheckPurchaseLimit(ctx, userId, goodsId)
			t.AssertNil(err)
			
			// 记录购买
			err = rateLimiter.RecordPurchase(ctx, userId, goodsId)
			t.AssertNil(err)
			
			// 再次购买检查应该失败
			err = rateLimiter.CheckPurchaseLimit(ctx, userId, goodsId)
			t.AssertNE(err, nil)
		})
	})
}

// TestAntiBrushChecker 单独测试防刷检查器
func TestAntiBrushChecker(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.Background()
		
		// 获取缓存实例
		cache := utility.GetFlashSaleCache()
		if cache == nil {
			t.Skip("缓存未初始化")
		}
		
		antiBrush := utility.NewAntiBrushChecker(cache)
		
		// 测试正常用户行为
		t.Run("NormalUserBehavior", func(t *gtest.T) {
			userId := uint32(30001)
			ip := "192.168.1.50"
			
			err := antiBrush.CheckUserBehavior(ctx, userId, ip)
			t.AssertNil(err)
		})
		
		// 测试黑名单用户
		t.Run("BlacklistedUser", func(t *gtest.T) {
			userId := uint32(30002)
			ip := "192.168.1.51"
			
			// 将用户加入黑名单
			err := antiBrush.AddToBlacklist(ctx, "user", userId, 1*time.Hour)
			t.AssertNil(err)
			
			// 检查应该失败
			err = antiBrush.CheckUserBehavior(ctx, userId, ip)
			t.AssertNE(err, nil)
			t.AssertContains(err.Error(), "异常行为")
		})
		
		// 测试黑名单IP
		t.Run("BlacklistedIP", func(t *gtest.T) {
			userId := uint32(30003)
			ip := "192.168.1.52"
			
			// 将IP加入黑名单
			err := antiBrush.AddToBlacklist(ctx, "ip", ip, 1*time.Hour)
			t.AssertNil(err)
			
			// 检查应该失败
			err = antiBrush.CheckUserBehavior(ctx, userId, ip)
			t.AssertNE(err, nil)
			t.AssertContains(err.Error(), "网络环境")
		})
	})
}

// BenchmarkFlashSale 性能基准测试
func BenchmarkFlashSale(b *testing.B) {
	ctx := context.Background()
	
	// 初始化测试环境
	initTestEnv(ctx)
	
	userId := uint32(40001)
	goodsId := uint32(50001)
	
	// 初始化库存
	stockManager := utility.GetFlashSaleStockManager()
	stockManager.InitStock(ctx, goodsId, 10000)
	
	flashSale := logic.NewFlashSale()
	
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := &v1.CreateFlashSaleOrderReq{
				UserId:  userId,
				GoodsId: goodsId,
				Count:   1,
			}
			flashSale.CreateFlashSaleOrder(ctx, req)
		}
	})
}