# 秒杀系统开发文档

## 项目概述

本秒杀系统是基于GoFrame微服务架构开发的高并发、高可用、高性能的电商秒杀解决方案。系统采用Redis缓存、消息队列、限流防刷等技术，确保在秒杀场景下的系统稳定性和数据一致性。

## 架构设计

### 系统架构图
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Gateway-H5    │    │   Gateway-Admin │    │   Gateway-API   │
│   (用户端)       │    │   (管理端)       │    │   (API接口)     │
└─────────┬───────┘    └─────────┬───────┘    └─────────┬───────┘
          │                      │                      │
          └──────────────────────┼──────────────────────┘
                                 │
                    ┌────────────┴────────────┐
                    │    Flash Sale Service   │
                    │     (秒杀服务)          │
                    └────────────┬────────────┘
                                 │
        ┌────────────────────────┼────────────────────────┐
        │                        │                        │
┌───────┴───────┐    ┌───────────┴──────────┐    ┌──────┴──────┐
│    Redis      │    │     RabbitMQ         │    │   MySQL     │
│   (缓存)       │    │     (消息队列)        │    │  (数据库)   │
└───────────────┘    └───────────────────────┘    └─────────────┘
```

### 核心组件

#### 1. 限流组件 (Rate Limiter)
- **功能**: 多级别限流保护
- **级别**: 用户级、IP级、全局级
- **算法**: 令牌桶 + 计数器
- **配置**:
  - 用户: 5次/秒，10次/分钟
  - IP: 10次/秒，50次/分钟
  - 全局: 100次/秒

#### 2. 防刷组件 (Anti-Brush)
- **功能**: 防止恶意刷单
- **机制**: 
  - 黑名单检测
  - 行为模式分析
  - 异常请求识别
- **处理**: 自动封禁可疑用户/IP

#### 3. 库存管理 (Stock Manager)
- **功能**: 原子性库存扣减
- **技术**: Redis Lua脚本
- **特性**:
  - 原子操作
  - 超卖保护
  - 库存回滚

#### 4. 消息队列 (Message Queue)
- **功能**: 异步订单处理
- **技术**: RabbitMQ
- **优势**:
  - 削峰填谷
  - 可靠传输
  - 失败重试

## 开发流程

### 1. 环境准备

#### 1.1 安装依赖
```bash
# 安装Go 1.18+
# 安装Redis 6.0+
# 安装RabbitMQ 3.8+
# 安装MySQL 8.0+
```

#### 1.2 项目初始化
```bash
# 克隆项目
git clone <repository-url>

# 进入项目目录
cd shop-goframe-micro-service-refacotor

# 安装Go依赖
go mod download

# 编译项目
go build ./...
```

### 2. 数据库设计

#### 2.1 秒杀商品表
```sql
CREATE TABLE `flash_sale_goods` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `goods_id` bigint(20) NOT NULL COMMENT '商品ID',
  `title` varchar(255) NOT NULL COMMENT '秒杀标题',
  `price` bigint(20) NOT NULL COMMENT '秒杀价格（分）',
  `stock` int(11) NOT NULL COMMENT '秒杀库存',
  `start_time` datetime NOT NULL COMMENT '开始时间',
  `end_time` datetime NOT NULL COMMENT '结束时间',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-正常，2-下架',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_goods_id` (`goods_id`),
  KEY `idx_status_time` (`status`, `start_time`, `end_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='秒杀商品表';
```

#### 2.2 秒杀订单表
```sql
CREATE TABLE `flash_sale_orders` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_no` varchar(32) NOT NULL COMMENT '订单号',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `goods_id` bigint(20) NOT NULL COMMENT '商品ID',
  `goods_title` varchar(255) NOT NULL COMMENT '商品标题',
  `price` bigint(20) NOT NULL COMMENT '成交价格（分）',
  `quantity` int(11) NOT NULL COMMENT '购买数量',
  `total_amount` bigint(20) NOT NULL COMMENT '总金额（分）',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '订单状态：1-待支付，2-已支付，3-已取消',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_order_no` (`order_no`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_goods_id` (`goods_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='秒杀订单表';
```

### 3. 核心功能开发

#### 3.1 限流器实现
```go
// 文件: app/flash-sale/utility/rate_limiter.go

// RateLimiter 限流器
type RateLimiter struct {
    cache CacheInterface
    mu    sync.RWMutex
}

// CheckUserLimit 检查用户限流
func (r *RateLimiter) CheckUserLimit(ctx context.Context, userId uint32) error {
    // 用户级别限流：每秒最多5次请求
    userKey := fmt.Sprintf(consts.FlashSaleUserRateLimitKey, userId)
    if err := r.checkLimit(ctx, userKey, consts.FlashSaleRateLimitPerSecond, 1*time.Second); err != nil {
        return fmt.Errorf("用户请求过于频繁，请稍后再试")
    }
    
    // 用户级别限流：每分钟最多10次请求
    userMinuteKey := fmt.Sprintf("flash_sale:user:%d:minute", userId)
    if err := r.checkLimit(ctx, userMinuteKey, consts.FlashSaleUserMinuteRateLimit, 1*time.Minute); err != nil {
        return fmt.Errorf("用户请求过于频繁，请稍后再试")
    }
    
    return nil
}
```

#### 3.2 库存扣减实现
```go
// 文件: app/goods/utility/stock/flash_sale_stock.go

// FlashSaleStockManager 秒杀库存管理器
type FlashSaleStockManager struct {
    redis *gredis.Redis
}

// DeductStock 扣减库存（原子操作）
func (m *FlashSaleStockManager) DeductStock(ctx context.Context, goodsId uint32, count int) (bool, error) {
    // 使用Lua脚本确保原子性
    script := `
        local key = KEYS[1]
        local count = tonumber(ARGV[1])
        local current = tonumber(redis.call('get', key) or 0)
        
        if current >= count then
            redis.call('decrby', key, count)
            return 1
        else
            return 0
        end
    `
    
    key := fmt.Sprintf(consts.FlashSaleStockCacheKey, goodsId)
    result, err := m.redis.Eval(ctx, script, []string{key}, count)
    if err != nil {
        return false, err
    }
    
    return gconv.Int(result) == 1, nil
}
```

#### 3.3 消息队列实现
```go
// 文件: app/flash-sale/utility/rabbitmq.go

// PublishFlashSaleOrder 发布秒杀订单消息
func PublishFlashSaleOrder(ctx context.Context, msg *model.FlashSaleOrderMessage) error {
    ch, err := rabbitMQ.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()
    
    // 声明交换机
    err = ch.ExchangeDeclare(
        consts.FlashSaleExchange, // name
        "direct",                 // type
        true,                     // durable
        false,                    // auto-deleted
        false,                    // internal
        false,                    // no-wait
        nil,                      // arguments
    )
    if err != nil {
        return err
    }
    
    // 序列化消息
    body, err := json.Marshal(msg)
    if err != nil {
        return err
    }
    
    // 发布消息
    err = ch.Publish(
        consts.FlashSaleExchange,        // exchange
        consts.FlashSaleOrderRoutingKey, // routing key
        false,                            // mandatory
        false,                            // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
    
    return err
}
```

### 4. 业务流程

#### 4.1 秒杀流程图
```
用户请求 → 参数验证 → 限流检查 → 防刷检查 → 库存检查 → 库存扣减 → 生成订单 → 发送消息 → 返回结果
   ↓         ↓         ↓         ↓         ↓         ↓         ↓         ↓         ↓
失败返回  失败返回  失败返回  失败返回  失败返回  失败返回  成功返回  异步处理  成功响应
```

#### 4.2 详细流程说明

1. **参数验证**
   - 验证用户ID、商品ID、购买数量
   - 验证商品状态和活动状态
   - 验证用户资格和限购条件

2. **限流检查**
   - 检查用户级别限流（5次/秒，10次/分钟）
   - 检查IP级别限流（10次/秒，50次/分钟）
   - 检查全局级别限流（100次/秒）

3. **防刷检查**
   - 检查用户是否在黑名单
   - 检查IP是否在黑名单
   - 分析用户行为模式
   - 识别异常请求特征

4. **库存检查**
   - 查询商品库存信息
   - 验证库存是否充足
   - 检查用户购买限制

5. **库存扣减**
   - 使用Redis Lua脚本原子扣减
   - 确保不超卖
   - 记录扣减结果

6. **生成订单**
   - 生成唯一订单号
   - 创建订单记录
   - 记录秒杀结果到缓存

7. **发送消息**
   - 封装订单消息
   - 发布到RabbitMQ
   - 异步处理订单

8. **返回结果**
   - 返回秒杀结果
   - 提供结果查询ID
   - 异步处理状态

## 部署指南

### 1. 环境配置

#### 1.1 Redis配置
```bash
# redis.conf
maxmemory 2gb
maxmemory-policy allkeys-lru
timeout 300
tcp-keepalive 60
```

#### 1.2 RabbitMQ配置
```bash
# rabbitmq.conf
loopback_users.guest = false
listeners.tcp.default = 5672
management.tcp.port = 15672
disk_free_limit.absolute = 1GB
```

#### 1.3 MySQL配置
```sql
-- 创建数据库
CREATE DATABASE flash_sale CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建用户
CREATE USER 'flash_sale'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON flash_sale.* TO 'flash_sale'@'%';
FLUSH PRIVILEGES;
```

### 2. 服务部署

#### 2.1 编译服务
```bash
# 编译秒杀服务
cd app/flash-sale
go build -o flash-sale-service main.go

# 编译网关服务
cd app/gateway-h5
go build -o gateway-h5-service main.go
```

#### 2.2 启动服务
```bash
# 启动Redis
redis-server redis.conf

# 启动RabbitMQ
rabbitmq-server

# 启动秒杀服务
cd app/flash-sale
./flash-sale-service

# 启动网关服务
cd app/gateway-h5
./gateway-h5-service
```

#### 2.3 配置服务
```yaml
# config.yaml
server:
  address: ":8000"
  
database:
  default:
    link: "mysql:flash_sale:password@tcp(127.0.0.1:3306)/flash_sale"
    
redis:
  default:
    address: "127.0.0.1:6379"
    db: 0
    
rabbitmq:
  default:
    address: "127.0.0.1:5672"
    user: "guest"
    pass: "guest"
```

### 3. 监控配置

#### 3.1 日志配置
```go
// 配置日志级别
g.Log().SetLevel(glog.LEVEL_INFO)

// 配置日志文件
g.Log().SetPath("/var/log/flash-sale")

// 配置日志轮转
g.Log().SetRotateSize(100 * 1024 * 1024) // 100MB
```

#### 3.2 性能监控
```go
// 启用性能监控
import _ "net/http/pprof"

// 启动监控服务
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

## 性能优化

### 1. 缓存优化
- **Redis集群**: 使用Redis Cluster提高并发能力
- **缓存预热**: 秒杀开始前预热热点数据
- **缓存分级**: L1本地缓存 + L2 Redis缓存
- **缓存更新**: 异步更新缓存，避免缓存击穿

### 2. 数据库优化
- **读写分离**: 主从复制，读写分离
- **分库分表**: 按用户ID或商品ID分片
- **索引优化**: 合理创建索引，避免全表扫描
- **连接池**: 使用连接池减少连接开销

### 3. 代码优化
- **对象池**: 复用对象，减少GC压力
- **异步处理**: 使用goroutine处理耗时操作
- **批量操作**: 批量处理数据库和缓存操作
- **算法优化**: 使用高效算法和数据结构

## 安全策略

### 1. 接口安全
- **鉴权认证**: JWT Token验证
- **参数校验**: 严格的参数验证和过滤
- **接口限流**: 防止接口被恶意调用
- **HTTPS**: 使用HTTPS加密传输

### 2. 数据安全
- **数据加密**: 敏感数据加密存储
- **SQL注入**: 使用参数化查询
- **XSS防护**: 输出内容转义
- **CSRF防护**: 使用CSRF Token

### 3. 业务安全
- **防刷机制**: 多维度防刷策略
- **风控系统**: 实时风险识别
- **异常监控**: 异常行为检测
- **人工审核**: 关键操作人工介入

## 故障处理

### 1. 常见故障

#### 1.1 Redis故障
- **现象**: 缓存不可用
- **处理**: 
  - 启用本地缓存
  - 降级到数据库查询
  - 快速恢复Redis服务

#### 1.2 RabbitMQ故障
- **现象**: 消息队列不可用
- **处理**:
  - 降级为同步处理
  - 记录日志后续补偿
  - 快速恢复MQ服务

#### 1.3 数据库故障
- **现象**: 数据库连接失败
- **处理**:
  - 启用读写分离
  - 使用缓存数据
  - 快速恢复数据库

### 2. 应急预案

#### 2.1 限流降级
```go
// 自动降级逻辑
func autoDegrade() {
    if errorRate > 0.1 { // 错误率超过10%
        enableDegrade()
        reduceTraffic()
    }
}
```

#### 2.2 熔断机制
```go
// 熔断器实现
type CircuitBreaker struct {
    failureCount    int
    failureThreshold int
    timeout         time.Duration
    state           string // closed, open, half-open
}
```

## 运维监控

### 1. 监控指标
- **业务指标**: QPS、成功率、订单量
- **系统指标**: CPU、内存、磁盘、网络
- **应用指标**: 响应时间、错误率、线程数
- **业务指标**: 库存变化、用户行为

### 2. 告警配置
```yaml
# 告警规则
groups:
- name: flash-sale
  rules:
  - alert: HighErrorRate
    expr: error_rate > 0.05
    for: 5m
    annotations:
      summary: "错误率过高"
      
  - alert: HighLatency
    expr: response_time_p99 > 500
    for: 5m
    annotations:
      summary: "响应时间过长"
```

### 3. 日志分析
```bash
# 错误日志分析
grep "ERROR" /var/log/flash-sale/*.log | awk '{print $5}' | sort | uniq -c

# 性能日志分析
grep "slow_query" /var/log/flash-sale/*.log | awk '{print $8}' | sort -n
```

## 总结

本秒杀系统采用微服务架构，通过多层限流、防刷机制、异步处理等技术，实现了高并发场景下的稳定运行。系统具有以下特点：

1. **高并发**: 支持万级并发请求
2. **高可用**: 多级容错和降级机制
3. **高性能**: 毫秒级响应时间
4. **数据一致性**: 原子操作确保数据准确
5. **安全防护**: 多维度安全策略
6. **可扩展**: 微服务架构便于扩展

通过合理的架构设计和优化策略，系统能够有效应对秒杀场景的各种挑战，为用户提供流畅的购物体验。