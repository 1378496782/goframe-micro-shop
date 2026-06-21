package utility

import (
	"context"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

// FlashSaleStockManager 秒杀库存管理器
type FlashSaleStockManager struct {
	cache *gcache.Cache
	mu    sync.RWMutex
}

var (
	stockManager     *FlashSaleStockManager
	stockManagerOnce sync.Once
)

// GetFlashSaleStockManager 获取秒杀库存管理器实例
func GetFlashSaleStockManager(cache *gcache.Cache) *FlashSaleStockManager {
	stockManagerOnce.Do(func() {
		stockManager = &FlashSaleStockManager{
			cache: cache,
		}
	})
	return stockManager
}

// CheckStock 检查库存
func (s *FlashSaleStockManager) CheckStock(ctx context.Context, goodsId uint32, count int) (bool, error) {
	key := fmt.Sprintf("flash_sale:stock:%d", goodsId)

	stock, err := s.cache.Get(ctx, key)
	if err != nil {
		return false, fmt.Errorf("获取库存失败: %v", err)
	}

	if stock == nil {
		return false, fmt.Errorf("商品库存信息不存在")
	}

	availableStock := gconv.Int(stock)
	return availableStock >= count, nil
}

// ReduceStock 减少库存
func (s *FlashSaleStockManager) ReduceStock(ctx context.Context, goodsId uint32, count int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	key := fmt.Sprintf("flash_sale:stock:%d", goodsId)

	stock, err := s.cache.Get(ctx, key)
	if err != nil {
		return fmt.Errorf("获取库存失败: %v", err)
	}

	if stock == nil {
		return fmt.Errorf("商品库存信息不存在")
	}

	availableStock := gconv.Int(stock)
	if availableStock < count {
		return fmt.Errorf("库存不足")
	}

	newStock := availableStock - count
	return s.cache.Set(ctx, key, newStock, 0) // 不设置过期时间
}

// GetStock 获取库存数量
func (s *FlashSaleStockManager) GetStock(ctx context.Context, goodsId uint32) (int, error) {
	key := fmt.Sprintf("flash_sale:stock:%d", goodsId)

	stock, err := s.cache.Get(ctx, key)
	if err != nil {
		return 0, fmt.Errorf("获取库存失败: %v", err)
	}

	if stock == nil {
		return 0, fmt.Errorf("商品库存信息不存在")
	}

	return gconv.Int(stock), nil
}
