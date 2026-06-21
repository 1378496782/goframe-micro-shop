# 秒杀系统测试文档

## 概述
本文档详细描述了秒杀系统的测试策略、测试用例、测试方法和测试结果验证方法。

## 测试环境
- **操作系统**: Windows/Linux/macOS
- **Go版本**: 1.18+
- **Redis**: 6.0+
- **RabbitMQ**: 3.8+
- **数据库**: MySQL 8.0+

## 测试分类

### 1. 单元测试 (Unit Tests)
测试单个组件的功能，确保每个模块正常工作。

#### 1.1 限流器测试 (TestRateLimiter)
- **测试目标**: 验证限流机制的正确性
- **测试内容**:
  - 用户级别限流（每秒最多5次请求）
  - IP级别限流（每秒最多10次请求）
  - 购买限制检查（每个用户每个商品限购1次）
  - 限流重置机制

#### 1.2 防刷检查器测试 (TestAntiBrushChecker)
- **测试目标**: 验证防刷机制的有效性
- **测试内容**:
  - 正常用户行为检查
  - 黑名单用户检测
  - 黑名单IP检测
  - 异常行为识别

### 2. 集成测试 (Integration Tests)
测试整个系统的工作流程，确保各组件协调工作。

#### 2.1 基本秒杀流程测试 (BasicFlashSaleFlow)
- **测试目标**: 验证完整的秒杀业务流程
- **测试步骤**:
  1. 初始化商品库存
  2. 创建秒杀订单
  3. 验证库存扣减
  4. 查询秒杀结果
  5. 验证订单信息

#### 2.2 限流机制集成测试 (RateLimitTest)
- **测试目标**: 验证限流在真实场景中的效果
- **测试内容**:
  - 短时间内大量请求的处理
  - 限流触发后的请求拒绝
  - 限流重置后的正常处理

#### 2.3 防刷机制集成测试 (AntiBrushTest)
- **测试目标**: 验证防刷机制的实际效果
- **测试内容**:
  - 黑名单用户的请求拦截
  - 黑名单IP的请求拦截
  - 异常行为的识别和处理

### 3. 并发测试 (Concurrent Tests)
测试系统在高并发场景下的表现。

#### 3.1 并发秒杀测试 (ConcurrentFlashSale)
- **测试目标**: 验证系统在并发场景下的正确性
- **测试参数**:
  - 并发用户数: 100
  - 初始库存: 50
  - 预期结果: 成功订单数 ≤ 初始库存
- **验证内容**:
  - 无超卖现象
  - 库存数据一致性
  - 订单生成正确性

### 4. 库存管理测试 (Stock Management Tests)
测试库存管理功能的正确性。

#### 4.1 库存初始化测试
- **测试目标**: 验证库存初始化的正确性
- **测试内容**:
  - 库存设置
  - 库存查询
  - 库存数据一致性

#### 4.2 库存扣减测试
- **测试目标**: 验证库存扣减的准确性
- **测试内容**:
  - 正常库存扣减
  - 库存不足时的处理
  - 并发库存扣减

### 5. 消息队列测试 (Message Queue Tests)
测试消息队列的功能和可靠性。

#### 5.1 消息发布测试
- **测试目标**: 验证消息发布的正确性
- **测试内容**:
  - 消息格式验证
  - 消息发布成功
  - 错误处理

#### 5.2 消息消费测试
- **测试目标**: 验证消息消费的可靠性
- **测试内容**:
  - 消息消费成功
  - 失败消息重试
  - 消息确认机制

### 6. 性能测试 (Performance Tests)
测试系统的性能指标。

#### 6.1 基准测试 (Benchmark)
- **测试目标**: 评估系统性能基线
- **测试指标**:
  - 请求处理速度 (QPS)
  - 内存使用情况
  - CPU使用率
  - 响应时间分布

## 测试执行方法

### 方法1: 使用测试脚本
```bash
# Windows
cd c:\code\shop-goframe-micro-service-refacotor\app\flash-sale\test
run_tests.bat

# Linux/macOS
cd /code/shop-goframe-micro-service-refacotor/app/flash-sale/test
chmod +x run_tests.sh
./run_tests.sh
```

### 方法2: 手动执行测试
```bash
# 进入项目目录
cd c:\code\shop-goframe-micro-service-refacotor\app\flash-sale

# 运行所有测试
go test -v ./test/...

# 运行特定测试
go test -v ./test/... -run TestRateLimiter

# 运行性能测试
go test -bench=. ./test/...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./test/...
go tool cover -html=coverage.out -o coverage.html
```

## 测试数据准备

### 测试商品数据
```go
map[string]interface{}{
    "id": 20001,
    "name": "测试商品1",
    "description": "测试商品描述1",
    "price": 99900,  // 999.00元
    "stock": 100,
    "status": 1,
}
```

### 测试用户数据
```go
map[string]interface{}{
    "id": 10001,
    "username": "testuser1",
    "email": "test1@example.com",
    "status": 1,
}
```

## 预期测试结果

### 功能测试结果
- ✅ 基本秒杀流程: 100% 成功率
- ✅ 限流机制: 正确限制超出阈值的请求
- ✅ 防刷机制: 有效拦截异常请求
- ✅ 并发处理: 无超卖，数据一致性保证
- ✅ 库存管理: 准确扣减，无负库存
- ✅ 消息队列: 可靠发布和消费

### 性能测试结果
- **QPS**: ≥ 1000 (单实例)
- **响应时间**: ≤ 100ms (P99)
- **内存使用**: ≤ 500MB
- **CPU使用**: ≤ 80%

## 问题排查指南

### 常见问题

#### 1. 测试失败
- **现象**: 测试用例执行失败
- **排查**:
  1. 检查Redis连接是否正常
  2. 检查RabbitMQ服务状态
  3. 查看测试日志获取详细错误信息
  4. 验证测试数据是否正确初始化

#### 2. 限流测试不通过
- **现象**: 限流阈值不符合预期
- **排查**:
  1. 检查限流配置参数
  2. 验证Redis缓存是否正常
  3. 检查限流算法实现
  4. 确认测试并发度是否合理

#### 3. 并发测试超卖
- **现象**: 成功订单数超过初始库存
- **排查**:
  1. 检查库存扣减逻辑
  2. 验证Redis Lua脚本
  3. 检查事务边界
  4. 确认并发控制机制

#### 4. 消息队列测试失败
- **现象**: 消息发布或消费失败
- **排查**:
  1. 检查RabbitMQ连接配置
  2. 验证交换机、队列配置
  3. 检查消息格式
  4. 查看RabbitMQ日志

### 调试技巧

1. **日志分析**: 启用详细日志记录，分析执行流程
2. **断点调试**: 使用IDE断点调试功能
3. **性能分析**: 使用Go pprof工具分析性能瓶颈
4. **压力测试**: 逐步增加并发度，观察系统表现

## 测试覆盖率目标

- **语句覆盖率**: ≥ 80%
- **分支覆盖率**: ≥ 70%
- **函数覆盖率**: ≥ 90%

## 持续集成

建议将测试集成到CI/CD流程中：

```yaml
# .github/workflows/test.yml
name: Flash Sale Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    services:
      redis:
        image: redis:6.0
        ports:
          - 6379:6379
      rabbitmq:
        image: rabbitmq:3.8
        ports:
          - 5672:5672
    
    steps:
    - uses: actions/checkout@v2
    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.18
    
    - name: Run tests
      run: |
        cd app/flash-sale
        go test -v ./test/...
    
    - name: Generate coverage
      run: |
        cd app/flash-sale
        go test -coverprofile=coverage.out ./test/...
        go tool cover -func=coverage.out
```

## 总结

本测试方案覆盖了秒杀系统的核心功能、性能、并发、可靠性等方面。通过系统化的测试，可以确保：

1. **功能正确性**: 所有业务逻辑按预期工作
2. **性能达标**: 满足高并发场景的性能要求
3. **数据一致性**: 保证数据在各种场景下的准确性
4. **系统稳定性**: 在异常情况下系统仍能正常运行
5. **安全防护**: 有效防止恶意攻击和异常行为

建议定期执行完整测试套件，特别是在系统升级或配置变更后，确保系统持续稳定运行。