package goodsRedis

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

// 本地锁，用于防止缓存击穿
var localMutex sync.Map

// CacheStrategy 定义缓存策略接口
type CacheStrategy interface {
	GetWithLock(ctx context.Context, key string, expiration time.Duration, callback func() (interface{}, error)) (*gvar.Var, error)
	SetWithRandomExpiration(ctx context.Context, key string, value interface{}, baseExpiration time.Duration)
}

// 缓存策略实现
type cacheStrategyImpl struct{}

// NewCacheStrategy 创建缓存策略实例
func NewCacheStrategy() CacheStrategy {
	return &cacheStrategyImpl{}
}

// GetWithLock 带分布式锁的缓存获取，防止缓存击穿
func (s *cacheStrategyImpl) GetWithLock(ctx context.Context, key string, expiration time.Duration, callback func() (interface{}, error)) (*gvar.Var, error) {
	// 尝试从缓存获取
	result, err := goodsCache.Get(ctx, key)
	if err == nil && !result.IsEmpty() && result.String() != "null" && result.String() != EmptyValue {
		return result, nil
	}

	// 获取本地锁
	mutexKey := fmt.Sprintf("mutex:%s", key)
	var mutex *sync.Mutex
	actual, _ := localMutex.LoadOrStore(mutexKey, &sync.Mutex{})
	mutex = actual.(*sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()
	defer localMutex.Delete(mutexKey)

	// 双重检查，防止其他协程已经加载了缓存
	result, err = goodsCache.Get(ctx, key)
	if err == nil && !result.IsEmpty() && result.String() != "null" && result.String() != EmptyValue {
		return result, nil
	}

	// 调用回调函数加载数据
	data, err := callback()
	if err != nil {
		return nil, err
	}

	// 检查数据是否为空，防止缓存穿透
	if data == nil {
		// 设置空值缓存，短时间过期
		err = goodsCache.Set(ctx, key, EmptyValue, 1*time.Minute)
		return gvar.New(nil), err
	}

	// 序列化数据
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 设置缓存，带随机过期时间防止雪崩
	s.SetWithRandomExpiration(ctx, key, jsonData, expiration)

	return gvar.New(jsonData), nil
}

// SetWithRandomExpiration 设置带随机过期时间的缓存，防止缓存雪崩
func (s *cacheStrategyImpl) SetWithRandomExpiration(ctx context.Context, key string, value interface{}, baseExpiration time.Duration) {
	// 添加5%~15%的随机时间偏移
	randomFactor := 0.05 + rand.Float64()*0.10
	randomExpiration := time.Duration(float64(baseExpiration) * (1 + randomFactor))

	err := goodsCache.Set(ctx, key, value, randomExpiration)
	if err != nil {
		g.Log().Warningf(ctx, "设置带随机过期时间的缓存失败: %v, key=%s", err, key)
	}
}

// 初始化随机数种子
func init() {
	rand.Seed(time.Now().UnixNano())
}
