#!/bin/bash

# 安全的微服务重新构建测试脚本（不影响其他容器）
set -e

echo "=========================================="
echo "安全测试：微服务重新构建脚本"
echo "=========================================="
echo ""
echo "🔍 检查当前Docker环境..."

# 检查当前运行的容器
echo "当前运行的容器："
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Image}}"

echo ""
echo "📋 检查本项目相关的镜像："
docker images --format "table {{.Repository}}\t{{.Tag}}\t{{.Size}}" | grep "shop-goframe-micro-service-refacotor" || echo "没有找到本项目相关的镜像"

echo ""
echo "📝 将要执行的操作："
echo "1. 构建以下服务（不会影响其他容器）："
echo "   - gateway-h5"
echo "   - gateway-resource"
echo "   - manage-service"
echo "   - manage-ui"
echo "   - user-service"
echo "   - goods-service"
echo "   - order-service"
echo "   - interaction-service"
echo "   - search-service"
echo "   - banner-service"
echo ""
echo "2. 清理操作："
echo "   - 只删除本项目相关的旧镜像"
echo "   - 不会使用 'docker system prune' 影响其他容器"
echo ""
echo "3. 启动服务并进行健康检查"
echo ""

read -p "是否继续执行？(y/N): " -n 1 -r
echo ""
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "操作已取消。"
    exit 1
fi

# 创建日志目录
mkdir -p logs
LOG_FILE="./logs/test_rebuild_safe_$(date +%Y%m%d_%H%M%S).log"

# 日志函数
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" | tee -a "$LOG_FILE"
}

log "=========================================="
log "开始安全构建测试"
log "=========================================="

# 服务列表
SERVICES=(
    "gateway-h5"
    "gateway-resource"
    "manage-service"
    "manage-ui"
    "user-service"
    "goods-service"
    "order-service"
    "interaction-service"
    "search-service"
    "banner-service"
)

log "步骤1: 检查当前服务状态"
# 尝试使用 docker compose（新版本）或 docker-compose（旧版本）
if command -v docker-compose &> /dev/null; then
    COMPOSE_CMD="docker-compose"
elif docker compose version &> /dev/null; then
    COMPOSE_CMD="docker compose"
else
    log "错误：未找到 docker-compose 或 docker compose 命令"
    exit 1
fi

$COMPOSE_CMD -f docker-compose.prod.yml ps | tee -a "$LOG_FILE"

log "步骤2: 停止本项目的服务（如果存在）"
for service in "${SERVICES[@]}"; do
    log "停止 $service 服务..."
    $COMPOSE_CMD -f docker-compose.prod.yml stop "$service" 2>/dev/null || log "$service 服务未运行"
done

log "步骤3: 安全清理 - 只删除本项目相关的镜像"
log "查找本项目相关的镜像..."
PROJECT_IMAGES=$(docker images --format "{{.Repository}}:{{.Tag}}" | grep "shop-goframe-micro-service-refacotor" || echo "")

if [ -n "$PROJECT_IMAGES" ]; then
    log "找到以下本项目镜像，将被删除："
    echo "$PROJECT_IMAGES" | tee -a "$LOG_FILE"
    echo "$PROJECT_IMAGES" | while read -r image; do
        if [ -n "$image" ]; then
            log "删除镜像: $image"
            docker rmi "$image" 2>/dev/null || log "无法删除镜像 $image（可能正在被使用）"
        fi
    done
else
    log "没有找到本项目相关的镜像需要清理"
fi

log "步骤4: 构建服务（逐个构建，便于错误处理）"
BUILD_SUCCESS=true

for service in "${SERVICES[@]}"; do
    log "构建 $service 服务..."
    if $COMPOSE_CMD -f docker-compose.prod.yml build --no-cache "$service" 2>>"$LOG_FILE"; then
        log "✅ $service 服务构建成功"
    else
        log "❌ $service 服务构建失败"
        BUILD_SUCCESS=false
        break
    fi
done

if [ "$BUILD_SUCCESS" = true ]; then
    log "步骤5: 启动所有服务"
    for service in "${SERVICES[@]}"; do
        log "启动 $service 服务..."
        if $COMPOSE_CMD -f docker-compose.prod.yml up -d --force-recreate "$service" 2>>"$LOG_FILE"; then
            log "✅ $service 服务启动成功"
        else
            log "❌ $service 服务启动失败"
            BUILD_SUCCESS=false
        fi
    done
fi

log "步骤6: 等待服务启动..."

sleep 10

log "步骤7: 最终状态检查"
$COMPOSE_CMD -f docker-compose.prod.yml ps | tee -a "$LOG_FILE"

log "步骤8: 检查服务端口"
log "检查端口监听状态："
# 检查 netstat 或 ss 命令
if command -v netstat &> /dev/null; then
    PORT_CHECK_CMD="netstat -tulnp"
elif command -v ss &> /dev/null; then
    PORT_CHECK_CMD="ss -tulnp"
else
    log "警告：未找到 netstat 或 ss 命令，跳过端口检查"
fi

if [ -n "$PORT_CHECK_CMD" ]; then
    $PORT_CHECK_CMD 2>/dev/null | grep -E ":(8199|8399|8808|80|443|31001|31002|31004|31005|8499|31006)" | tee -a "$LOG_FILE"
fi

# 总结
echo ""
echo "=========================================="
if [ "$BUILD_SUCCESS" = true ]; then
    log "✅ 安全构建测试完成！所有服务构建成功"
else
    log "⚠️  构建测试完成，但部分服务可能存在问题"
fi

echo ""
echo "📊 构建统计："
echo "  日志文件：$LOG_FILE"
echo ""
echo "🌐 服务访问地址："
echo "  H5商城网关：http://localhost:8199"
echo "  资源网关：http://localhost:8399"
echo "  管理后台：http://localhost:8808"
echo "  管理UI：http://localhost"
echo ""
echo "🔍 检查命令："
echo "  docker-compose -f docker-compose.prod.yml ps"
echo "  docker-compose -f docker-compose.prod.yml logs [服务名]"
echo ""
echo "✨ 安全特性："
echo "  - 没有使用 'docker system prune'"
echo "  - 只影响本项目相关容器"
echo "  - 保留了所有其他项目的容器"