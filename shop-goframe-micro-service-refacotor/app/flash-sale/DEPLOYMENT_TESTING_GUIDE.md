# 秒杀系统部署与测试指南

## 快速开始

### 1. 环境检查

在开始部署前，请确保以下环境已准备就绪：

```bash
# 检查Go环境
go version  # 需要1.18+

# 检查Redis
redis-cli ping  # 应该返回PONG

# 检查RabbitMQ
rabbitmqctl status

# 检查MySQL
mysql -u root -p -e "SELECT VERSION();"
```

### 2. 一键部署脚本

我们提供了自动化部署脚本，可以快速部署整个系统：

```bash
# Windows系统
cd c:\code\shop-goframe-micro-service-refacotor\app\flash-sale
deploy\deploy.bat

# Linux/macOS系统
cd /code/shop-goframe-micro-service-refacotor/app/flash-sale
chmod +x deploy/deploy.sh
./deploy/deploy.sh
```

### 3. 手动部署步骤

#### 3.1 数据库初始化
```bash
# 创建数据库
mysql -u root -p < deploy/sql/init_database.sql

# 创建表结构
mysql -u root -p flash_sale < deploy/sql/create_tables.sql

# 初始化测试数据
mysql -u root -p flash_sale < deploy/sql/test_data.sql
```

#### 3.2 Redis配置
```bash
# 启动Redis服务
redis-server /etc/redis/redis.conf

# 验证连接
redis-cli -h localhost -p 6379 ping
```

#### 3.3 RabbitMQ配置
```bash
# 启动RabbitMQ
rabbitmq-server -detached

# 创建用户和权限
rabbitmqctl add_user flash_sale flash_sale_pass
rabbitmqctl set_user_tags flash_sale administrator
rabbitmqctl set_permissions -p / flash_sale ".*" ".*" ".*"
```

#### 3.4 编译服务
```bash
# 编译秒杀服务
cd app/flash-sale
go build -ldflags "-s -w" -o bin/flash-sale-service main.go

# 编译网关服务
cd app/gateway-h5
go build -ldflags "-s -w" -o bin/gateway-h5-service main.go
```

#### 3.5 启动服务
```bash
# 启动秒杀服务
cd app/flash-sale
./bin/flash-sale-service &

# 启动网关服务
cd app/gateway-h5
./bin/gateway-h5-service &
```

## 功能测试

### 1. 基础功能测试

#### 1.1 API接口测试
```bash
# 获取秒杀商品列表
curl -X GET "http://localhost:8000/api/flash-sale/v1/goods/list"

# 获取商品详情
curl -X GET "http://localhost:8000/api/flash-sale/v1/goods/detail?id=1"

# 创建秒杀订单（需要JWT Token）
curl -X POST "http://localhost:8000/api/flash-sale/v1/order/create" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "goods_id": 1,
    "count": 1
  }'
```

#### 1.2 使用测试脚本
```bash
# 运行基础功能测试
cd test
python test_api.py --host localhost --port 8000

# 运行压力测试
python stress_test.py --users 100 --duration 60
```

### 2. 性能测试

#### 2.1 使用Apache Bench
```bash
# 测试商品列表接口
ab -n 10000 -c 100 http://localhost:8000/api/flash-sale/v1/goods/list

# 测试订单创建接口（需要添加认证头）
ab -n 5000 -c 50 -T 'application/json' -H 'Authorization: Bearer TOKEN' \
  -p order_request.json http://localhost:8000/api/flash-sale/v1/order/create
```

#### 2.2 使用Go测试工具
```bash
# 运行性能基准测试
cd app/flash-sale
go test -bench=. -benchmem ./test/

# 生成性能报告
go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof ./test/
go tool pprof cpu.prof
```

### 3. 并发测试

#### 3.1 使用JMeter
```bash
# 启动JMeter
jmeter -n -t test/jmeter/flash_sale_test.jmx -l test_results.jtl

# 生成测试报告
jmeter -g test_results.jtl -o test_report/
```

#### 3.2 使用自定义脚本
```bash
# 运行并发测试
python concurrent_test.py \
  --users 1000 \
  --goods-id 1 \
  --duration 30 \
  --ramp-up 5
```

## 监控与验证

### 1. 系统监控

#### 1.1 服务状态检查
```bash
# 检查服务进程
ps aux | grep flash-sale-service
ps aux | grep gateway-h5-service

# 检查端口监听
netstat -tlnp | grep :8000
netstat -tlnp | grep :8001
```

#### 1.2 日志监控
```bash
# 实时查看日志
tail -f logs/flash-sale-service.log
tail -f logs/gateway-h5-service.log

# 错误日志分析
grep "ERROR" logs/*.log | tail -20
```

### 2. 性能监控

#### 2.1 Redis监控
```bash
# 连接数监控
redis-cli info clients

# 内存使用监控
redis-cli info memory

# 命令统计
redis-cli info commandstats
```

#### 2.2 RabbitMQ监控
```bash
# 队列状态
rabbitmqctl list_queues

# 连接状态
rabbitmqctl list_connections

# 通道状态
rabbitmqctl list_channels
```

#### 2.3 MySQL监控
```bash
# 连接数查看
mysql -u root -p -e "SHOW STATUS LIKE 'Threads_connected';"

# 慢查询查看
mysql -u root -p -e "SHOW PROCESSLIST;"

# 性能指标
mysql -u root -p -e "SHOW STATUS LIKE 'Questions';"
```

### 3. 业务指标验证

#### 3.1 成功率统计
```bash
# 分析成功率
grep "秒杀成功" logs/flash-sale-service.log | wc -l
grep "秒杀失败" logs/flash-sale-service.log | wc -l

# 计算成功率
total=$(grep "秒杀" logs/flash-sale-service.log | wc -l)
success=$(grep "秒杀成功" logs/flash-sale-service.log | wc -l)
echo "成功率: $(($success * 100 / $total))%"
```

#### 3.2 响应时间分析
```bash
# 提取响应时间
grep "响应时间" logs/flash-sale-service.log | awk '{print $6}' | sort -n

# 计算平均响应时间
grep "响应时间" logs/flash-sale-service.log | awk '{sum+=$6; count++} END {print "平均响应时间:", sum/count, "ms"}'
```

## 故障排查

### 1. 常见问题解决

#### 1.1 服务启动失败
```bash
# 检查端口占用
netstat -tlnp | grep :8000

# 检查配置文件
cat config/config.yaml | grep -v "^#"

# 检查依赖服务
redis-cli ping
rabbitmqctl status
```

#### 1.2 数据库连接失败
```bash
# 检查MySQL服务
systemctl status mysql

# 检查用户权限
mysql -u flash_sale -p -e "SHOW GRANTS;"

# 检查网络连接
telnet localhost 3306
```

#### 1.3 Redis连接失败
```bash
# 检查Redis服务
systemctl status redis

# 检查配置
redis-cli CONFIG GET bind
redis-cli CONFIG GET port

# 检查内存使用
redis-cli INFO memory
```

### 2. 性能问题诊断

#### 2.1 CPU使用率高
```bash
# 查看CPU使用
top -p $(pgrep flash-sale-service)

# 分析CPU热点
go tool pprof http://localhost:6060/debug/pprof/profile

# 查看goroutine
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

#### 2.2 内存使用高
```bash
# 查看内存使用
ps aux | grep flash-sale-service

# 分析内存分配
go tool pprof http://localhost:6060/debug/pprof/heap

# 查看内存统计
go tool pprof http://localhost:6060/debug/pprof/allocs
```

#### 2.3 响应时间慢
```bash
# 查看慢查询
grep "slow_query" logs/*.log

# 分析数据库慢查询
mysql -u root -p -e "SHOW VARIABLES LIKE 'slow_query_log%';"

# 检查Redis慢查询
redis-cli SLOWLOG GET 10
```

### 3. 业务问题诊断

#### 3.1 超卖问题
```bash
# 检查库存数据
redis-cli GET "flash_sale:stock:1"

# 检查订单数量
mysql -u flash_sale -p -e "SELECT COUNT(*) FROM flash_sale_orders WHERE goods_id = 1;"

# 分析日志
grep "库存不足\|超卖" logs/*.log
```

#### 3.2 限流失效
```bash
# 检查限流配置
grep "FlashSaleRateLimit" config/config.yaml

# 检查Redis限流数据
redis-cli KEYS "flash_sale:rate_limit:*"

# 分析限流日志
grep "限流" logs/*.log | tail -50
```

## 测试报告生成

### 1. 自动生成报告
```bash
# 运行完整测试套件
cd test
python generate_report.py

# 查看生成的报告
open test_report/index.html
```

### 2. 手动生成报告
```bash
# 收集测试数据
python collect_metrics.py > metrics.json

# 生成图表
python create_charts.py metrics.json

# 生成PDF报告
python create_pdf_report.py
```

## 验证清单

### ✅ 部署验证
- [ ] 所有服务正常启动
- [ ] 数据库连接正常
- [ ] Redis连接正常
- [ ] RabbitMQ连接正常
- [ ] 端口监听正常
- [ ] 配置文件正确

### ✅ 功能验证
- [ ] 商品列表接口正常
- [ ] 商品详情接口正常
- [ ] 订单创建接口正常
- [ ] 结果查询接口正常
- [ ] JWT认证正常
- [ ] 限流功能正常
- [ ] 防刷功能正常

### ✅ 性能验证
- [ ] 响应时间符合预期
- [ ] 并发处理能力达标
- [ ] 内存使用合理
- [ ] CPU使用正常
- [ ] 无内存泄漏
- [ ] 无goroutine泄漏

### ✅ 数据一致性验证
- [ ] 库存数据准确
- [ ] 订单数据完整
- [ ] 无超卖现象
- [ ] 消息队列数据一致
- [ ] 缓存数据同步

### ✅ 安全验证
- [ ] 接口鉴权有效
- [ ] SQL注入防护
- [ ] XSS攻击防护
- [ ] CSRF攻击防护
- [ ] 敏感数据加密

## 总结

通过本指南，您可以：

1. **快速部署**: 使用自动化脚本或手动步骤完成系统部署
2. **全面测试**: 执行功能、性能、并发等多维度测试
3. **有效监控**: 实时监控系统运行状态和性能指标
4. **快速排障**: 定位和解决常见的部署和运行问题
5. **生成报告**: 自动生成详细的测试报告和分析结果

建议在正式环境部署前，先在测试环境完整执行所有验证步骤，确保系统稳定可靠。