package stock

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
)

// 测试配置
const (
	TestGoodsID        = uint32(1001) // 测试商品ID
	TestInitialStock   = 100          // 初始库存
	TestConcurrency    = 200          // 并发请求数
	TestRequestCount   = 1            // 每个goroutine的请求次数
	TestStockDeduction = 1            // 每次扣减的库存数量
)

// 测试结果结构
type TestResult struct {
	TotalRequests int           // 总请求数
	SuccessCount  int           // 成功请求数
	FailureCount  int           // 失败请求数
	TotalDuration time.Duration // 总执行时间
}

// runConcurrencyTest 执行并发测试
func runConcurrencyTest(ctx context.Context, manager StockManager, concurrency int, reqCountPerGoroutine int) TestResult {
	var (
		wg            sync.WaitGroup
		result        TestResult
		mutex         sync.Mutex
		totalRequests = concurrency * reqCountPerGoroutine
		totalDuration time.Duration
		startTime     = time.Now()
	)

	// 设置总请求数
	result.TotalRequests = totalRequests

	// 启动并发测试
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 每个goroutine执行多次请求
			for j := 0; j < reqCountPerGoroutine; j++ {
				reqStart := time.Now()
				success, _ := manager.ReduceStock(ctx, TestGoodsID, TestStockDeduction)
				reqDuration := time.Since(reqStart)

				mutex.Lock()
				totalDuration += reqDuration
				if success {
					result.SuccessCount++
				} else {
					result.FailureCount++
				}
				mutex.Unlock()

				// 短暂休眠，避免过度请求
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	// 等待所有goroutine完成
	wg.Wait()
	result.TotalDuration = time.Since(startTime)

	return result
}

// TestStockManagerComparison 对比测试两种库存扣减方案
func TestStockManagerComparison(t *testing.T) {
	// 准备测试数据
	ctx := gctx.New()

	// 获取Redis客户端
	redisClient := g.Redis()
	if redisClient == nil {
		t.Fatal("获取Redis客户端失败")
	}

	// 确保测试前清理Redis中的测试数据
	defer func() {
		stockKey := fmt.Sprintf("goods:stock:%d", TestGoodsID)
		lockKey := fmt.Sprintf("goods:stock:lock:%d", TestGoodsID)
		redisClient.Del(ctx, stockKey)
		redisClient.Del(ctx, lockKey)
	}()

	gtest.C(t, func(t *gtest.T) {
		// 创建两种库存管理器实例
		distLockManager := NewDistributedLockStockManager(redisClient)
		luaManager := NewRedisLuaStockManager(redisClient)

		// 测试分布式锁方案
		fmt.Println("=== 分布式锁方案 测试开始 ===")
		func() {
			// 初始化库存
			_, err := distLockManager.InitStock(ctx, TestGoodsID, TestInitialStock)
			if err != nil {
				t.Fatalf("初始化库存失败: %v", err)
			}

			// 执行并发测试
			results := runConcurrencyTest(ctx, distLockManager, TestConcurrency, TestRequestCount)

			// 获取最终库存
			finalStock, err := distLockManager.GetStock(ctx, TestGoodsID)
			if err != nil {
				t.Fatalf("获取最终库存失败: %v", err)
			}

			// 计算理论上的最终库存（成功扣减的数量不能超过初始库存）
			successCount := results.SuccessCount
			if successCount > TestInitialStock {
				successCount = TestInitialStock
			}
			expectedFinalStock := TestInitialStock - successCount

			// 输出测试结果
			fmt.Printf("=== 分布式锁方案 测试结果 ===\n")
			fmt.Printf("初始库存: %d\n", TestInitialStock)
			fmt.Printf("并发请求数: %d\n", TestConcurrency)
			fmt.Printf("成功请求数: %d\n", results.SuccessCount)
			fmt.Printf("失败请求数: %d\n", results.FailureCount)
			fmt.Printf("平均响应时间: %.2f ms\n", float64(results.TotalDuration)/float64(results.TotalRequests)/float64(time.Millisecond))
			fmt.Printf("最终库存: %d\n", finalStock)
			fmt.Printf("期望库存: %d\n", expectedFinalStock)

			// 验证最终库存是否正确
			t.Assert(finalStock, expectedFinalStock)

			// 验证是否发生超卖
			hasOversold := results.SuccessCount > TestInitialStock
			if hasOversold {
				t.Fatal("发生超卖现象！")
			}
			fmt.Printf("是否发生超卖: %v\n\n", hasOversold)
		}()

		// 清理数据后测试Lua脚本方案
		stockKey := fmt.Sprintf("goods:stock:%d", TestGoodsID)
		lockKey := fmt.Sprintf("goods:stock:lock:%d", TestGoodsID)
		redisClient.Del(ctx, stockKey)
		redisClient.Del(ctx, lockKey)

		fmt.Println("=== Lua脚本方案 测试开始 ===")
		func() {
			// 初始化库存
			_, err := luaManager.InitStock(ctx, TestGoodsID, TestInitialStock)
			if err != nil {
				t.Fatalf("初始化库存失败: %v", err)
			}

			// 执行并发测试
			results := runConcurrencyTest(ctx, luaManager, TestConcurrency, TestRequestCount)

			// 获取最终库存
			finalStock, err := luaManager.GetStock(ctx, TestGoodsID)
			if err != nil {
				t.Fatalf("获取最终库存失败: %v", err)
			}

			// 计算理论上的最终库存（成功扣减的数量不能超过初始库存）
			successCount := results.SuccessCount
			if successCount > TestInitialStock {
				successCount = TestInitialStock
			}
			expectedFinalStock := TestInitialStock - successCount

			// 输出测试结果
			fmt.Printf("=== Lua脚本方案 测试结果 ===\n")
			fmt.Printf("初始库存: %d\n", TestInitialStock)
			fmt.Printf("并发请求数: %d\n", TestConcurrency)
			fmt.Printf("成功请求数: %d\n", results.SuccessCount)
			fmt.Printf("失败请求数: %d\n", results.FailureCount)
			fmt.Printf("平均响应时间: %.2f ms\n", float64(results.TotalDuration)/float64(results.TotalRequests)/float64(time.Millisecond))
			fmt.Printf("最终库存: %d\n", finalStock)
			fmt.Printf("期望库存: %d\n", expectedFinalStock)

			// 验证最终库存是否正确
			t.Assert(finalStock, expectedFinalStock)

			// 验证是否发生超卖
			hasOversold := results.SuccessCount > TestInitialStock
			if hasOversold {
				t.Fatal("发生超卖现象！")
			}
			fmt.Printf("是否发生超卖: %v\n\n", hasOversold)
		}()
	})
}

// TestStockManagerEdgeCases 测试库存管理器的边界情况
func TestStockManagerEdgeCases(t *testing.T) {
	// 获取Redis客户端
	redisClient := g.Redis()
	if redisClient == nil {
		t.Fatal("获取Redis客户端失败")
	}

	ctx := gctx.New()
	goodsID := uint32(1002) // 使用不同的商品ID

	// 确保测试前清理Redis中的测试数据
	defer func() {
		stockKey := fmt.Sprintf("goods:stock:%d", goodsID)
		lockKey := fmt.Sprintf("goods:stock:lock:%d", goodsID)
		redisClient.Del(ctx, stockKey)
		redisClient.Del(ctx, lockKey)
	}()

	gtest.C(t, func(t *gtest.T) {
		// 测试分布式锁方案的边界情况
		fmt.Println("=== 测试分布式锁方案边界情况 ===")
		distLockManager := NewDistributedLockStockManager(redisClient)
		testEdgeCases(t, distLockManager, ctx, goodsID)

		// 清理数据
		stockKey := fmt.Sprintf("goods:stock:%d", goodsID)
		lockKey := fmt.Sprintf("goods:stock:lock:%d", goodsID)
		redisClient.Del(ctx, stockKey)
		redisClient.Del(ctx, lockKey)

		// 测试Lua脚本方案的边界情况
		fmt.Println("=== 测试Lua脚本方案边界情况 ===")
		luaManager := NewRedisLuaStockManager(redisClient)
		testEdgeCases(t, luaManager, ctx, goodsID)
	})
}

// testEdgeCases 测试库存管理器的边界情况
func testEdgeCases(t *gtest.T, manager StockManager, ctx context.Context, goodsID uint32) {
	// 测试1：库存为0时扣减
	_, err := manager.InitStock(ctx, goodsID, 0)
	t.AssertNil(err)

	success, err := manager.ReduceStock(ctx, goodsID, 1)
	t.Assert(success, false)
	t.AssertNE(err, nil)

	// 测试2：负库存扣减
	success, err = manager.ReduceStock(ctx, goodsID, -1)
	t.Assert(success, false)

	// 测试3：库存充足时扣减和返还
	_, err = manager.InitStock(ctx, goodsID, 10)
	t.AssertNil(err)

	success, err = manager.ReduceStock(ctx, goodsID, 5)
	t.Assert(success, true)
	t.AssertNil(err)

	currentStock, err := manager.GetStock(ctx, goodsID)
	t.AssertNil(err)
	t.Assert(currentStock, 5)

	// 测试4：返还库存
	success, err = manager.ReturnStock(ctx, goodsID, 3)
	t.Assert(success, true)
	t.AssertNil(err)

	currentStock, err = manager.GetStock(ctx, goodsID)
	t.AssertNil(err)
	t.Assert(currentStock, 8)
}
