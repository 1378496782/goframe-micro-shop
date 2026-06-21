package stock

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

// FlashSaleStockManager 秒杀库存管理器接口
type FlashSaleStockManager interface {
	StockManager // 继承基础库存管理
	
	// ReduceFlashSaleStock 扣减秒杀库存（带用户限流）
	ReduceFlashSaleStock(ctx context.Context, goodsId uint32, userId uint32, count int) (bool, error)
	
	// CheckUserPurchaseLimit 检查用户购买限制
	CheckUserPurchaseLimit(ctx context.Context, goodsId uint32, userId uint32) (bool, error)
	
	// RecordUserPurchase 记录用户购买
	RecordUserPurchase(ctx context.Context, goodsId uint32, userId uint32) error
}

// flashSaleStockManager 秒杀库存管理器实现
type flashSaleStockManager struct {
	*RedisLuaStockManager
	cache *gcache.Cache
}

// NewFlashSaleStockManager 创建秒杀库存管理器
func NewFlashSaleStockManager(redisClient interface{}, cache *gcache.Cache) FlashSaleStockManager {
	return &flashSaleStockManager{
		RedisLuaStockManager: NewRedisLuaStockManager(redisClient),
		cache:                cache,
	}
}

// getFlashSaleStockKey 获取秒杀库存key
func (m *flashSaleStockManager) getFlashSaleStockKey(goodsId uint32) string {
	return fmt.Sprintf("flash_sale:stock:%d", goodsId)
}

// getUserPurchaseKey 获取用户购买记录key
func (m *flashSaleStockManager) getUserPurchaseKey(goodsId uint32, userId uint32) string {
	return fmt.Sprintf("flash_sale:user_purchase:%d:%d", goodsId, userId)
}

// ReduceFlashSaleStock 扣减秒杀库存（带用户限流）
func (m *flashSaleStockManager) ReduceFlashSaleStock(ctx context.Context, goodsId uint32, userId uint32, count int) (bool, error) {
	// 参数校验
	if count <= 0 {
		return false, gerror.New("购买数量必须大于0")
	}
	
	// 检查用户购买限制
	canPurchase, err := m.CheckUserPurchaseLimit(ctx, goodsId, userId)
	if err != nil {
		return false, gerror.Wrap(err, "检查用户购买限制失败")
	}
	if !canPurchase {
		return false, gerror.New("您已购买过该秒杀商品")
	}
	
	// 获取秒杀库存key
	stockKey := m.getFlashSaleStockKey(goodsId)
	
	// 执行Lua脚本扣减库存
	result, err := m.redisClient.(interface {
		Do(ctx context.Context, command string, args ...interface{}) (interface{}, error)
	}).Do(ctx, "EVAL", m.getReduceFlashSaleStockScript(), 1, stockKey, count)
	
	if err != nil {
		return false, gerror.Wrapf(err, "扣减秒杀库存失败，商品ID:%d，用户ID:%d，扣减数量:%d", goodsId, userId, count)
	}
	
	// 解析结果
	newStock := gconv.Int(result)
	if newStock == -1 {
		// 库存不足
		return false, gerror.New("秒杀商品库存不足")
	}
	
	// 记录用户购买（设置24小时过期）
	if err := m.RecordUserPurchase(ctx, goodsId, userId); err != nil {
		// 记录失败，需要回滚库存
		g.Log().Warning(ctx, "记录用户购买失败，准备回滚库存:", err)
		m.ReturnStock(ctx, goodsId, count)
		return false, gerror.Wrap(err, "记录用户购买失败")
	}
	
	g.Log().Infof(ctx, "秒杀库存扣减成功，商品ID:%d，用户ID:%d，扣减数量:%d，剩余库存:%d", 
		goodsId, userId, count, newStock)
	
	return true, nil
}

// CheckUserPurchaseLimit 检查用户购买限制
func (m *flashSaleStockManager) CheckUserPurchaseLimit(ctx context.Context, goodsId uint32, userId uint32) (bool, error) {
	purchaseKey := m.getUserPurchaseKey(goodsId, userId)
	
	// 从缓存检查用户是否已购买
	exists, err := m.cache.Contains(ctx, purchaseKey)
	if err != nil {
		return false, gerror.Wrap(err, "检查用户购买缓存失败")
	}
	
	return !exists, nil
}

// RecordUserPurchase 记录用户购买
func (m *flashSaleStockManager) RecordUserPurchase(ctx context.Context, goodsId uint32, userId uint32) error {
	purchaseKey := m.getUserPurchaseKey(goodsId, userId)
	
	// 记录用户购买，设置24小时过期
	err := m.cache.Set(ctx, purchaseKey, 1, 24*time.Hour)
	if err != nil {
		return gerror.Wrap(err, "记录用户购买失败")
	}
	
	return nil
}

// getReduceFlashSaleStockScript 扣减秒杀库存Lua脚本
func (m *flashSaleStockManager) getReduceFlashSaleStockScript() string {
	return `
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
}