#!/bin/bash

# 部署脚本
set -e

# 配置变量
BACKUP_DIR="./backups/$(date +%Y%m%d_%H%M%S)"
LOG_FILE="./logs/deploy_$(date +%Y%m%d_%H%M%S).log"
DEPLOY_START_TIME=$(date +%s)

# 创建必要的目录
mkdir -p backups logs

# 日志函数
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" | tee -a "$LOG_FILE"
}

log "=========================================="
log "商城微服务系统部署脚本（生产环境版本）"
log "特点：数据库备份 + 删除旧镜像 + 强制重新编译 + 监控告警"
log "=========================================="

# 数据库备份
log "开始数据库备份..."
mkdir -p "$BACKUP_DIR"

# 备份MySQL数据库
log "备份MySQL数据库..."
docker-compose -f docker-compose.prod.yml exec -T mysql mysqldump -uroot -proot --all-databases > "$BACKUP_DIR/mysql_all_databases.sql" 2>>"$LOG_FILE" || {
    log "警告：MySQL数据库备份失败，继续部署..."
}

# 备份重要数据卷（如果有的话）
log "备份数据卷..."
docker run --rm -v shop-goframe-micro-service-refacotor_mysql_data:/data -v "$BACKUP_DIR":/backup alpine tar czf /backup/mysql_data.tar.gz -C /data . 2>>"$LOG_FILE" || {
    log "警告：MySQL数据卷备份失败，继续部署..."
}

log "数据库备份完成，备份文件保存在：$BACKUP_DIR"

# 健康检查函数
check_service_health() {
    local service=$1
    local url=$2
    local max_retries=30
    local retry_count=0
    
    log "正在检查 $service 服务健康状态..."
    while [ $retry_count -lt $max_retries ]; do
        if curl -f -s "$url" > /dev/null 2>&1; then
            log "$service 服务健康检查通过"
            return 0
        fi
        retry_count=$((retry_count + 1))
        log "$service 服务检查失败，重试 $retry_count/$max_retries..."
        sleep 10
    done
    
    log "错误：$service 服务健康检查失败，部署可能有问题！"
    return 1
}

# 发送通知函数
send_notification() {
    local status=$1
    local message=$2
    local duration=$3
    
    # 这里可以集成钉钉、企业微信、邮件等通知
    log "发送部署通知：$status - $message"
    
    # 示例：发送到钉钉webhook（需要配置实际的webhook URL）
    # curl -H "Content-Type: application/json" \
    #      -d "{\"msgtype\":\"text\",\"text\":{\"content\":\"商城部署通知：$status\n$message\n耗时：${duration}s\"}}" \
    #      "YOUR_DINGTALK_WEBHOOK_URL" 2>>"$LOG_FILE" || true
}

# 部署开始通知
send_notification "START" "商城微服务部署开始" "0"

# 删除旧的镜像
log "删除旧的 Docker 镜像..."
log "停止并删除所有相关容器..."
docker-compose -f docker-compose.prod.yml down || true

log "删除前端管理 UI 镜像..."
docker rmi manage-ui:latest 2>/dev/null || true

log "删除后端微服务镜像..."
# 删除所有后端服务的镜像
docker images --format "table {{.Repository}}\t{{.Tag}}\t{{.ID}}" | grep -E "(shop-goframe|gateway|admin|user|goods|order|interaction|search|banner|worker)" | awk '{print $3}' | xargs -r docker rmi 2>/dev/null || true

log "清理未使用的镜像和缓存..."
docker system prune -f

log "镜像清理完成，开始重新构建..."

# 构建前端项目
log "构建前端管理 UI..."
cd ../shop-goframe-micro-service-manage-ui-gfast
log "使用 --no-cache 参数强制重新构建前端镜像..."
docker build --no-cache -t manage-ui:latest . 2>>"$LOG_FILE"

# 返回到 refacotor 目录
cd ../shop-goframe-micro-service-refacotor

# 构建后端服务
log "构建后端微服务..."
log "使用 --no-cache 参数强制重新构建所有后端镜像..."
docker-compose -f docker-compose.prod.yml build --no-cache 2>>"$LOG_FILE"

# 启动所有服务
log "启动所有服务..."
log "使用 --force-recreate 参数强制重新创建所有容器..."
docker-compose -f docker-compose.prod.yml up -d --force-recreate 2>>"$LOG_FILE"

# 等待服务启动
log "等待服务启动..."
sleep 30

# 健康检查和监控
log "开始服务健康检查..."

# 检查关键服务
HEALTH_CHECK_PASSED=true

# 检查MySQL
if check_service_health "MySQL" "http://localhost:3306" 2>/dev/null; then
    log "✅ MySQL 服务正常"
else
    log "❌ MySQL 服务异常"
    HEALTH_CHECK_PASSED=false
fi

# 检查Redis
if docker-compose -f docker-compose.prod.yml exec -T redis redis-cli ping 2>/dev/null | grep -q "PONG"; then
    log "✅ Redis 服务正常"
else
    log "❌ Redis 服务异常"
    HEALTH_CHECK_PASSED=false
fi

# 检查RabbitMQ
if check_service_health "RabbitMQ" "http://localhost:15672" 2>/dev/null; then
    log "✅ RabbitMQ 服务正常"
else
    log "❌ RabbitMQ 服务异常"
    HEALTH_CHECK_PASSED=false
fi

# 检查主要网关服务
if check_service_health "Gateway-H5" "http://localhost:8199/health" 2>/dev/null; then
    log "✅ H5网关服务正常"
else
    log "❌ H5网关服务异常"
    HEALTH_CHECK_PASSED=false
fi

# 检查管理后台网关
if check_service_health "Gateway-Admin" "http://localhost:8299/health" 2>/dev/null; then
    log "✅ 管理网关服务正常"
else
    log "❌ 管理网关服务异常"
    HEALTH_CHECK_PASSED=false
fi

# 检查服务状态...
log "检查服务状态..."
docker-compose -f docker-compose.prod.yml ps | tee -a "$LOG_FILE"

# 部署完成和通知
DEPLOY_END_TIME=$(date '+%Y-%m-%d %H:%M:%S')
DEPLOY_DURATION=$(( $(date -d "$DEPLOY_END_TIME" +%s) - $(date -d "$DEPLOY_START_TIME" +%s) ))

if [ "$HEALTH_CHECK_PASSED" = true ]; then
    log "✅ 所有服务健康检查通过！"
    send_notification "部署成功" "商城微服务系统部署成功！\n部署时间：$DEPLOY_DURATION 秒\n备份目录：$BACKUP_DIR\n日志文件：$LOG_FILE"
else
    log "❌ 部分服务健康检查失败！"
    send_notification "部署警告" "商城微服务系统部署完成，但部分服务异常！\n部署时间：$DEPLOY_DURATION 秒\n请检查日志文件：$LOG_FILE"
fi

echo "=========================================="
echo "部署完成！（旧镜像已删除，所有镜像已重新编译）"
echo "=========================================="
echo ""
echo "部署统计："
echo "  开始时间：$DEPLOY_START_TIME"
echo "  结束时间：$DEPLOY_END_TIME"
echo "  总耗时：$DEPLOY_DURATION 秒"
echo "  备份目录：$BACKUP_DIR"
echo "  日志文件：$LOG_FILE"
echo ""
echo "服务访问地址："
echo "  管理后台：http://localhost:8299/admin"
echo "  H5商城：http://localhost:8199"
echo "  RabbitMQ管理界面：http://localhost:15672"
echo "  MySQL：localhost:3306"
echo ""
echo "常用命令："
echo "  查看日志：docker-compose logs -f [服务名]"
echo "  重启服务：docker-compose restart [服务名]"
echo "  停止服务：docker-compose down"
echo "  查看健康状态：docker-compose ps"
echo "=========================================="

# 记录部署完成
log "部署脚本执行完成 - 总耗时：${DEPLOY_DURATION}秒"