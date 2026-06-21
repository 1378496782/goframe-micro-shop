package stock

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// DistributedLockStockManager 基于分布式锁的库存管理器
type DistributedLockStockManager struct {
	redisClient   interface{}   // Redis客户端
	lockTimeout   time.Duration // 锁超时时间
	retryTimes    int           // 重试次数
	retryInterval time.Duration // 重试间隔
}

// NewDistributedLockStockManager 创建基于分布式锁的库存管理器
func NewDistributedLockStockManager(redisClient interface{}) *DistributedLockStockManager {
	return &DistributedLockStockManager{
		redisClient:   redisClient,
		lockTimeout:   3 * time.Second,
		retryTimes:    3,
		retryInterval: 100 * time.Millisecond,
	}
}

// getStockKey 获取库存key
func (m *DistributedLockStockManager) getStockKey(goodsId uint32) string {
	return fmt.Sprintf("goods:stock:%d", goodsId)
}

// getLockKey 获取锁key
func (m *DistributedLockStockManager) getLockKey(goodsId uint32) string {
	return fmt.Sprintf("lock:stock:%d", goodsId)
}

// generateLockValue 生成锁值
func (m *DistributedLockStockManager) generateLockValue() string {
	return fmt.Sprintf("%d", rand.Int63())
}

// acquireLock 获取分布式锁
func (m *DistributedLockStockManager) acquireLock(ctx context.Context, goodsId uint32) (string, bool, error) {
	lockKey := m.getLockKey(goodsId)
	lockValue := m.generateLockValue()

	// 使用传入的Redis客户端
	result, err := m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "SET", lockKey, lockValue, "NX", "EX", int(m.lockTimeout.Seconds()))

	if err != nil {
		return "", false, gerror.Wrap(err, "获取分布式锁失败")
	}

	// 检查结果
	success := result != nil

	return lockValue, success, nil
}

// releaseLock 释放分布式锁
func (m *DistributedLockStockManager) releaseLock(ctx context.Context, goodsId uint32, lockValue string) error {
	lockKey := m.getLockKey(goodsId)

	// 使用Lua脚本安全释放锁
	luaScript := `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	else
		return 0
	end
	`

	// 使用传入的Redis客户端
	_, err := m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "EVAL", luaScript, 1, lockKey, lockValue)

	if err != nil {
		return gerror.Wrap(err, "释放分布式锁失败")
	}

	return nil
}

// ReduceStock 扣减库存
func (m *DistributedLockStockManager) ReduceStock(ctx context.Context, goodsId uint32, count int) (bool, error) {
	// 参数校验
	if count <= 0 {
		return false, gerror.New("扣减数量必须大于0")
	}

	// 尝试获取锁
	lockValue, success, err := m.acquireLock(ctx, goodsId)
	if err != nil {
		return false, err
	}

	if !success {
		// 获取锁失败，尝试重试
		for i := 0; i < m.retryTimes; i++ {
			time.Sleep(m.retryInterval)
			lockValue, success, err = m.acquireLock(ctx, goodsId)
			if err != nil {
				return false, err
			}
			if success {
				break
			}
		}
	}

	if !success {
		return false, gerror.New("获取分布式锁失败，请稍后重试")
	}

	// 确保锁被释放
	defer func() {
		_ = m.releaseLock(ctx, goodsId, lockValue)
	}()

	// 获取当前库存
	stockKey := m.getStockKey(goodsId)
	currentStockStr, err := m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "GET", stockKey)

	if err != nil {
		return false, gerror.Wrapf(err, "获取商品库存失败，商品ID:%d", goodsId)
	}

	// 解析库存
	currentStock := 0
	if currentStockStr != nil {
		currentStock = gconv.Int(currentStockStr)
	}

	// 检查库存是否足够
	if currentStock < count {
		return false, gerror.Newf("库存不足，商品ID:%d，当前库存:%d，请求数量:%d", goodsId, currentStock, count)
	}

	// 扣减库存
	newStock := currentStock - count
	_, err = m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "SET", stockKey, newStock)

	if err != nil {
		return false, gerror.Wrapf(err, "更新库存失败，商品ID:%d", goodsId)
	}

	return true, nil
}

// ReturnStock 返还库存
func (m *DistributedLockStockManager) ReturnStock(ctx context.Context, goodsId uint32, count int) (bool, error) {
	// 参数校验
	if count <= 0 {
		return false, gerror.New("返还数量必须大于0")
	}

	// 获取锁
	lockValue, success, err := m.acquireLock(ctx, goodsId)
	if err != nil {
		return false, err
	}

	if !success {
		return false, gerror.New("获取分布式锁失败，请稍后重试")
	}

	// 确保锁被释放
	defer func() {
		_ = m.releaseLock(ctx, goodsId, lockValue)
	}()

	// 获取当前库存
	stockKey := m.getStockKey(goodsId)
	currentStockStr, err := m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "GET", stockKey)

	if err != nil {
		return false, gerror.Wrapf(err, "获取商品库存失败，商品ID:%d", goodsId)
	}

	// 解析库存
	currentStock := 0
	if currentStockStr != nil {
		currentStock = gconv.Int(currentStockStr)
	}

	// 返还库存
	newStock := currentStock + count
	_, err = m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "SET", stockKey, newStock)

	if err != nil {
		return false, gerror.Wrapf(err, "更新库存失败，商品ID:%d", goodsId)
	}

	return true, nil
}

// GetStock 获取当前库存
func (m *DistributedLockStockManager) GetStock(ctx context.Context, goodsId uint32) (int, error) {
	stockKey := m.getStockKey(goodsId)
	currentStockStr, err := m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "GET", stockKey)

	if err != nil {
		return 0, gerror.Wrapf(err, "获取商品库存失败，商品ID:%d", goodsId)
	}

	// 如果库存不存在，返回0
	if currentStockStr == nil {
		return 0, nil
	}

	// 解析并返回库存
	return gconv.Int(currentStockStr), nil
}

// InitStock 初始化库存
func (m *DistributedLockStockManager) InitStock(ctx context.Context, goodsId uint32, count int) (bool, error) {
	// 参数校验
	if count < 0 {
		return false, gerror.New("初始库存不能为负数")
	}

	// 获取锁
	lockValue, success, err := m.acquireLock(ctx, goodsId)
	if err != nil {
		return false, err
	}

	if !success {
		return false, gerror.New("获取分布式锁失败，请稍后重试")
	}

	// 确保锁被释放
	defer func() {
		_ = m.releaseLock(ctx, goodsId, lockValue)
	}()

	// 设置库存
	stockKey := m.getStockKey(goodsId)
	_, err = m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "SET", stockKey, count)

	if err != nil {
		return false, gerror.Wrapf(err, "初始化库存失败，商品ID:%d", goodsId)
	}

	return true, nil
}
