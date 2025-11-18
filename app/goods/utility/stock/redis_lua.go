package stock

import (
	"context"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// RedisLuaStockManager 基于Redis Lua脚本的库存管理器
type RedisLuaStockManager struct {
	redisClient interface{} // Redis客户端
	once        sync.Once   // 确保Lua脚本只初始化一次
}

// NewRedisLuaStockManager 创建基于Redis Lua脚本的库存管理器
func NewRedisLuaStockManager(redisClient interface{}) *RedisLuaStockManager {
	return &RedisLuaStockManager{
		redisClient: redisClient,
	}
}

// getStockKey 获取库存key
func (m *RedisLuaStockManager) getStockKey(goodsId uint32) string {
	return fmt.Sprintf("goods:stock:%d", goodsId)
}

// reduceStockScript 扣减库存的Lua脚本
const reduceStockScript = `
local stockKey = KEYS[1]
local count = tonumber(ARGV[1])

-- 获取当前库存
local currentStock = redis.call("GET", stockKey)
if currentStock == false then
    currentStock = 0
else
    currentStock = tonumber(currentStock)
end

-- 检查库存是否足够
if currentStock < count then
    return -1 -- 库存不足
end

-- 扣减库存
local newStock = currentStock - count
redis.call("SET", stockKey, newStock)

return newStock -- 返回扣减后的库存
`

// returnStockScript 返还库存的Lua脚本
const returnStockScript = `
local stockKey = KEYS[1]
local count = tonumber(ARGV[1])

-- 获取当前库存
local currentStock = redis.call("GET", stockKey)
if currentStock == false then
    currentStock = 0
else
    currentStock = tonumber(currentStock)
end

-- 返还库存
local newStock = currentStock + count
redis.call("SET", stockKey, newStock)

return newStock -- 返回返还后的库存
`

// ReduceStock 扣减库存（使用Redis Lua脚本）
func (m *RedisLuaStockManager) ReduceStock(ctx context.Context, goodsId uint32, count int) (bool, error) {
	// 参数校验
	if count <= 0 {
		return false, gerror.New("扣减数量必须大于0")
	}

	// 获取库存key
	stockKey := m.getStockKey(goodsId)

	// 执行Lua脚本使用Do方法
	result, err := m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "EVAL", reduceStockScript, 1, stockKey, count)

	if err != nil {
		return false, gerror.Wrapf(err, "扣减库存失败，商品ID:%d，扣减数量:%d", goodsId, count)
	}

	// 解析结果
	newStock := gconv.Int(result)
	if newStock == -1 {
		// 库存不足
		return false, gerror.Newf("库存不足，商品ID:%d，请求数量:%d", goodsId, count)
	}

	return true, nil
}

// ReturnStock 返还库存（使用Redis Lua脚本）
func (m *RedisLuaStockManager) ReturnStock(ctx context.Context, goodsId uint32, count int) (bool, error) {
	// 参数校验
	if count <= 0 {
		return false, gerror.New("返还数量必须大于0")
	}

	// 获取库存key
	stockKey := m.getStockKey(goodsId)

	// 执行Lua脚本使用Do方法
	_, err := m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "EVAL", returnStockScript, 1, stockKey, count)

	if err != nil {
		return false, gerror.Wrapf(err, "返还库存失败，商品ID:%d，返还数量:%d", goodsId, count)
	}

	// 结果解析，直接返回成功
	return true, nil
}

// GetStock 获取当前库存
func (m *RedisLuaStockManager) GetStock(ctx context.Context, goodsId uint32) (int, error) {
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
func (m *RedisLuaStockManager) InitStock(ctx context.Context, goodsId uint32, count int) (bool, error) {
	// 参数校验
	if count < 0 {
		return false, gerror.New("初始库存不能为负数")
	}

	// 设置初始库存
	stockKey := m.getStockKey(goodsId)
	_, err := m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "SET", stockKey, count)

	if err != nil {
		return false, gerror.Wrapf(err, "初始化库存失败，商品ID:%d", goodsId)
	}

	return true, nil
}
