package idempotent

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// mockRedisClient 模拟Redis客户端，用于单元测试
type mockRedisClient struct {
	// 存储键值对
	keyValues map[string]interface{}
	// 记录调用次数
	callCount map[string]int
}

// 实现SetNX方法，直接返回匿名接口类型
func (m *mockRedisClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) interface { Result() (bool, error) } {
	m.callCount["SetNX"]++
	return &mockSetNXResult{
		do: func() (bool, error) {
			// 模拟SetNX行为：如果键不存在则设置并返回true，否则返回false
			if _, exists := m.keyValues[key]; !exists {
				m.keyValues[key] = value
				return true, nil
			}
			return false, nil
		},
	}
}

// 实现Del方法，直接返回匿名接口类型
func (m *mockRedisClient) Del(ctx context.Context, keys ...string) interface { Result() (int64, error) } {
	m.callCount["Del"]++
	return &mockDelResult{
		do: func() (int64, error) {
			var count int64
			for _, key := range keys {
				if _, exists := m.keyValues[key]; exists {
					delete(m.keyValues, key)
					count++
				}
			}
			return count, nil
		},
	}
}

// 为了处理不同的返回类型，我们需要分别定义两个结构体

// mockSetNXResult 模拟SetNX操作的结果，实现SetNXResult接口
type mockSetNXResult struct {
	do func() (bool, error)
}

// Result 实现SetNXResult接口的Result方法
func (r *mockSetNXResult) Result() (bool, error) {
	if r.do != nil {
		return r.do()
	}
	return false, nil
}

// mockDelResult 模拟Del操作的结果，实现DelResult接口
type mockDelResult struct {
	do func() (int64, error)
}

// Result 实现DelResult接口的Result方法
func (r *mockDelResult) Result() (int64, error) {
	if r.do != nil {
		return r.do()
	}
	return 0, nil
}

// newMockRedisClient 创建新的模拟Redis客户端
func newMockRedisClient() *mockRedisClient {
	return &mockRedisClient{
		keyValues: make(map[string]interface{}),
		callCount: make(map[string]int),
	}
}

// TestTryLock 测试TryLock方法
func TestTryLock(t *testing.T) {
	ctx := context.Background()
	mockClient := newMockRedisClient()
	service := NewRedisIdempotentService(mockClient)

	// 测试场景1：首次尝试获取锁应该成功
	key := "test:lock:1"
	expiration := 10 * time.Second
	result1, err := service.TryLock(ctx, key, expiration)
	assert.NoError(t, err)
	assert.True(t, result1)
	assert.Equal(t, 1, mockClient.callCount["SetNX"])

	// 测试场景2：重复尝试获取同一个锁应该失败
	result2, err := service.TryLock(ctx, key, expiration)
	assert.NoError(t, err)
	assert.False(t, result2)
	assert.Equal(t, 2, mockClient.callCount["SetNX"])

	// 测试场景3：获取不同的锁应该成功
	key2 := "test:lock:2"
	result3, err := service.TryLock(ctx, key2, expiration)
	assert.NoError(t, err)
	assert.True(t, result3)
	assert.Equal(t, 3, mockClient.callCount["SetNX"])
}

// TestReleaseLock 测试ReleaseLock方法
func TestReleaseLock(t *testing.T) {
	ctx := context.Background()
	mockClient := newMockRedisClient()
	service := NewRedisIdempotentService(mockClient)

	// 测试场景1：先获取锁，然后释放锁
	key := "test:lock:3"
	expiration := 10 * time.Second

	// 先获取锁
	result1, err := service.TryLock(ctx, key, expiration)
	assert.NoError(t, err)
	assert.True(t, result1)

	// 释放锁
	err = service.ReleaseLock(ctx, key)
	assert.NoError(t, err)
	assert.Equal(t, 1, mockClient.callCount["Del"])

	// 再次获取同一把锁应该成功
	result2, err := service.TryLock(ctx, key, expiration)
	assert.NoError(t, err)
	assert.True(t, result2)

	// 测试场景2：释放不存在的锁
	err = service.ReleaseLock(ctx, "non:existing:key")
	assert.NoError(t, err)
	assert.Equal(t, 2, mockClient.callCount["Del"])
}

// TestCheckAndLock 测试CheckAndLock方法
func TestCheckAndLock(t *testing.T) {
	ctx := context.Background()
	mockClient := newMockRedisClient()
	service := NewRedisIdempotentService(mockClient)

	// 测试场景1：检查并获取新锁应该成功
	key := "test:lock:4"
	expiration := 10 * time.Second
	result1, err := service.CheckAndLock(ctx, key, expiration)
	assert.NoError(t, err)
	assert.True(t, result1)

	// 测试场景2：检查并获取已存在的锁应该失败
	result2, err := service.CheckAndLock(ctx, key, expiration)
	assert.NoError(t, err)
	assert.False(t, result2)
}

// TestGenerateMessageKey 测试GenerateMessageKey方法
func TestGenerateMessageKey(t *testing.T) {
	service := NewRedisIdempotentService(nil)

	// 测试场景：生成消息幂等键
	prefix := "message"
	messageID := "msg-001"
	businessID := "order-123"
	key := service.GenerateMessageKey(prefix, messageID, businessID)

	// 验证生成的键格式是否正确
	expectedKey := "message:msg-001:order-123"
	assert.Equal(t, expectedKey, key)
}

// TestIntegration 集成测试：模拟完整的幂等操作流程
func TestIntegration(t *testing.T) {
	ctx := context.Background()
	mockClient := newMockRedisClient()
	service := NewRedisIdempotentService(mockClient)

	// 模拟业务场景：处理重复请求
	orderID := "order-456"
	userID := "user-789"

	// 生成幂等键
	key := service.GenerateMessageKey("order", orderID, userID)

	// 模拟第一次请求
	acquired, err := service.TryLock(ctx, key, 30*time.Second)
	assert.NoError(t, err)
	assert.True(t, acquired)

	// 执行业务逻辑（模拟）
	// 假设这里执行业务逻辑...

	// 释放锁
	err = service.ReleaseLock(ctx, key)
	assert.NoError(t, err)

	// 模拟并发重复请求
	acquired, err = service.TryLock(ctx, key, 30*time.Second)
	assert.NoError(t, err)
	assert.True(t, acquired) // 因为之前已经释放了锁

	// 再次释放锁
	err = service.ReleaseLock(ctx, key)
	assert.NoError(t, err)
}
