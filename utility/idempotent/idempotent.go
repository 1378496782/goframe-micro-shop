package idempotent

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// IdempotentService 幂等性服务接口
type IdempotentService interface {
	// TryLock 尝试获取幂等锁
	TryLock(ctx context.Context, key string, expiration time.Duration) (bool, error)
	// ReleaseLock 释放幂等锁
	ReleaseLock(ctx context.Context, key string) error
	// CheckAndLock 检查并加锁，如果已存在则返回false
	CheckAndLock(ctx context.Context, key string, expiration time.Duration) (bool, error)
	// GenerateMessageKey 为消息生成唯一的幂等键
	GenerateMessageKey(prefix string, messageID string, businessID string) string
}

// redisIdempotentService 基于Redis的幂等性服务实现
type redisIdempotentService struct {
	redisClient interface{}
}

// NewRedisIdempotentService 创建基于Redis的幂等性服务
func NewRedisIdempotentService(redisClient interface{}) IdempotentService {
	return &redisIdempotentService{
		redisClient: redisClient,
	}
}

// TryLock 尝试获取幂等锁
func (s *redisIdempotentService) TryLock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	// 使用Redis的SETNX命令实现幂等性
	result, err := s.redisClient.(interface {
		SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	}).SetNX(ctx, key, time.Now().UnixNano(), expiration)

	if err != nil {
		g.Log().Errorf(ctx, "获取幂等锁失败: key=%s, error=%v", key, err)
		return false, fmt.Errorf("获取幂等锁失败: %v", err)
	}

	return result, nil
}

// ReleaseLock 释放幂等锁
func (s *redisIdempotentService) ReleaseLock(ctx context.Context, key string) error {
	// 直接删除键来释放锁
	_, err := s.redisClient.(interface {
		Del(ctx context.Context, keys ...string) (int64, error)
	}).Del(ctx, key)

	if err != nil {
		g.Log().Errorf(ctx, "释放幂等锁失败: key=%s, error=%v", key, err)
		return fmt.Errorf("释放幂等锁失败: %v", err)
	}

	return nil
}

// CheckAndLock 检查并加锁，如果已存在则返回false
func (s *redisIdempotentService) CheckAndLock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	// 直接使用TryLock，因为SETNX已经包含了检查和设置的逻辑
	return s.TryLock(ctx, key, expiration)
}

// GenerateMessageKey 为消息生成唯一的幂等键
func (s *redisIdempotentService) GenerateMessageKey(prefix string, messageID string, businessID string) string {
	// 生成格式: prefix:messageID:businessID
	return fmt.Sprintf("%s:%s:%s", prefix, messageID, businessID)
}

// DefaultIdempotentService 默认的幂等性服务实例
var DefaultIdempotentService IdempotentService

// InitIdempotentService 初始化默认的幂等性服务
func InitIdempotentService() error {
	ctx := context.Background()
	// 使用默认的Redis客户端
	redisClient := g.Redis()
	if redisClient == nil {
		return fmt.Errorf("获取Redis客户端失败")
	}

	DefaultIdempotentService = NewRedisIdempotentService(redisClient)
	g.Log().Info(ctx, "幂等性服务初始化成功")
	return nil
}

// GetDefaultIdempotentService 获取默认的幂等性服务
func GetDefaultIdempotentService() IdempotentService {
	if DefaultIdempotentService == nil {
		// 如果服务未初始化，尝试初始化
		if err := InitIdempotentService(); err != nil {
			// 如果初始化失败，返回nil
			g.Log().Error(context.Background(), "获取默认幂等性服务失败: %v", err)
			return nil
		}
	}
	return DefaultIdempotentService
}

// TryLock 全局函数：尝试获取幂等锁
func TryLock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	service := GetDefaultIdempotentService()
	if service == nil {
		return false, fmt.Errorf("幂等性服务未初始化")
	}
	return service.TryLock(ctx, key, expiration)
}

// ReleaseLock 全局函数：释放幂等锁
func ReleaseLock(ctx context.Context, key string) error {
	service := GetDefaultIdempotentService()
	if service == nil {
		return fmt.Errorf("幂等性服务未初始化")
	}
	return service.ReleaseLock(ctx, key)
}

// CheckAndLock 全局函数：检查并加锁
func CheckAndLock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	service := GetDefaultIdempotentService()
	if service == nil {
		return false, fmt.Errorf("幂等性服务未初始化")
	}
	return service.CheckAndLock(ctx, key, expiration)
}

// GenerateMessageKey 全局函数：生成消息幂等键
func GenerateMessageKey(prefix string, messageID string, businessID string) string {
	service := GetDefaultIdempotentService()
	if service == nil {
		// 如果服务未初始化，使用简单的字符串拼接
		return fmt.Sprintf("%s:%s:%s", prefix, messageID, businessID)
	}
	return service.GenerateMessageKey(prefix, messageID, businessID)
}