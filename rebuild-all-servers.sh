#!/bin/bash

# 重新构建所有服务器脚本
set -e

# 配置变量
LOG_FILE="./logs/rebuild_$(date +%Y%m%d_%H%M%S).log"
COMPOSE_FILE="docker-compose.prod.yml"

# 创建日志目录
mkdir -p logs

# 日志函数
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" | tee -a "$LOG_FILE"
}

log "=========================================="
log "商城微服务系统 - 重新构建所有服务器脚本"
log "使用 Docker Compose V2 (docker compose)"
log "=========================================="

# 记录开始时间
START_TIME=$(date +%s)

# 步骤1：停止所有服务
log "[1/6] 停止所有服务..."
docker compose -f $COMPOSE_FILE down --volumes --remove-orphans --timeout 30
log "✅ 所有服务已停止"

# 步骤2：获取当前配置中的所有服务名称
log "[2/6] 获取服务列表..."
SERVICES=$(docker compose -f $COMPOSE_FILE config --services)
log "需要处理的服务: $SERVICES"

# 步骤3：删除相关镜像
log "[3/6] 清理相关镜像..."
for service in $SERVICES; do
    # 获取服务的镜像名（支持build和image两种模式）
    IMAGE_LINE=$(docker compose -f $COMPOSE_FILE config | grep -A 20 "^  $service:" | grep "image:" | head -1)
    if [ ! -z "$IMAGE_LINE" ]; then
        IMAGE_NAME=$(echo "$IMAGE_LINE" | awk '{print $2}')
        if [ ! -z "$IMAGE_NAME" ]; then
            log "删除镜像: $IMAGE_NAME"
            docker rmi -f "$IMAGE_NAME" 2>/dev/null || log "镜像 $IMAGE_NAME 不存在或已删除"
        fi
    fi
    
    # 检查是否有build配置
    BUILD_LINE=$(docker compose -f $COMPOSE_FILE config | grep -A 30 "^  $service:" | grep "build:" | head -1)
    if [ ! -z "$BUILD_LINE" ]; then
        log "服务 $service 使用本地构建"
    fi
done

# 清理悬空镜像和缓存
log "清理悬空镜像..."
docker image prune -f

# 步骤4：重新构建所有镜像
log "[4/6] 重新构建所有镜像..."
docker compose -f $COMPOSE_FILE build --no-cache --parallel
log "✅ 所有镜像构建完成"

# 步骤5：启动所有服务
log "[5/6] 启动所有服务..."
docker compose -f $COMPOSE_FILE up -d --force-recreate
log "✅ 所有服务已启动"

# 步骤6：等待服务启动并检查状态
log "[6/6] 等待服务启动..."
sleep 25

# 检查服务状态
log "检查服务状态..."
docker compose -f $COMPOSE_FILE ps | tee -a "$LOG_FILE"

# 检查健康状态
log "检查服务健康状态..."
HEALTHY_SERVICES=0
TOTAL_SERVICES=$(echo $SERVICES | wc -w)

for service in $SERVICES; do
    # 检查服务是否运行
    if docker compose -f $COMPOSE_FILE ps | grep -q "$service.*Up"; then
        log "✅ $service 正在运行"
        ((HEALTHY_SERVICES++))
    else
        log "❌ $service 未运行"
    fi
done

# 计算耗时
END_TIME=$(date +%s)
DURATION=$((END_TIME - START_TIME))

# 输出总结
log "=========================================="
log "重新构建完成！"
log "=========================================="
log "统计信息："
log "  总服务数: $TOTAL_SERVICES"
log "  运行服务数: $HEALTHY_SERVICES"
log "  总耗时: ${DURATION}秒"
log "  日志文件: $LOG_FILE"
log ""
log "服务访问地址："
log "  H5商城网关: http://localhost:8199"
log "  资源网关: http://localhost:8399"
log "  管理后台: http://localhost:8808"
log "  RabbitMQ管理: http://localhost:15672"
log "  Kibana: http://localhost:5601"
log "  Elasticsearch: http://localhost:9200"
log ""
log "常用命令："
log "  查看日志: docker compose -f $COMPOSE_FILE logs -f [服务名]"
log "  重启服务: docker compose -f $COMPOSE_FILE restart [服务名]"
log "  停止服务: docker compose -f $COMPOSE_FILE down"
log "  查看状态: docker compose -f $COMPOSE_FILE ps"
log "=========================================="

# 如果有过半服务未运行，返回错误码
if [ $HEALTHY_SERVICES -lt $((TOTAL_SERVICES / 2)) ]; then
    log "❌ 警告：超过半数服务未正常运行！"
    exit 1
fi

log "✅ 重新构建成功完成！"
exit 0