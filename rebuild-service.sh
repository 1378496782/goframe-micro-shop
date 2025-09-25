#!/bin/bash

# 微服务重新构建脚本 - 支持单个服务或所有服务
set -e

echo "=========================================="
echo "微服务重新构建工具"
echo "=========================================="
echo ""

# 服务列表
VALID_SERVICES=(
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

# 显示使用方法
show_usage() {
    echo "使用方法："
    echo "  ./rebuild-service.sh [选项] [服务名称]"
    echo ""
    echo "选项："
    echo "  -h, --help          显示帮助信息"
    echo "  -a, --all           重新构建所有服务"
    echo "  -s, --single        重新构建单个服务（默认）"
    echo ""
    echo "可用的服务："
    for service in "${VALID_SERVICES[@]}"; do
        echo "  - $service"
    done
    echo ""
    echo "示例："
    echo "  ./rebuild-service.sh -a                    # 重新构建所有服务"
    echo "  ./rebuild-service.sh user-service           # 重新构建用户服务"
    echo "  ./rebuild-service.sh -s goods-service       # 重新构建商品服务"
    echo ""
    echo "如果不指定选项，默认使用单个服务模式"
}

# 解析参数
MODE="single"
SERVICE_NAME=""

while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_usage
            exit 0
            ;;
        -a|--all)
            MODE="all"
            shift
            ;;
        -s|--single)
            MODE="single"
            shift
            ;;
        -*)
            echo "❌ 未知选项: $1"
            show_usage
            exit 1
            ;;
        *)
            if [ -n "$SERVICE_NAME" ]; then
                echo "❌ 只能指定一个服务名称"
                show_usage
                exit 1
            fi
            SERVICE_NAME=$1
            shift
            ;;
    esac
done

# 验证服务名称（仅在单个服务模式下）
if [ "$MODE" = "single" ]; then
    if [ -z "$SERVICE_NAME" ]; then
        echo "❌ 请指定要重新构建的服务名称"
        echo ""
        show_usage
        exit 1
    fi
    
    is_valid_service=false
    for valid_service in "${VALID_SERVICES[@]}"; do
        if [ "$SERVICE_NAME" = "$valid_service" ]; then
            is_valid_service=true
            break
        fi
    done
    
    if [ "$is_valid_service" = false ]; then
        echo "❌ 无效的服务名称: $SERVICE_NAME"
        echo "请使用上述列表中的有效服务名称"
        exit 1
    fi
fi

echo "🔍 检查当前Docker环境..."
echo ""

# 显示当前状态
echo "当前运行的容器："
if [ "$MODE" = "single" ]; then
    docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Image}}" | grep "$SERVICE_NAME" || echo "$SERVICE_NAME 服务当前未运行"
else
    docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Image}}"
fi

echo ""
echo "📋 检查本项目相关的镜像："
docker images --format "table {{.Repository}}\t{{.Tag}}\t{{.Size}}" | grep "shop-goframe-micro-service-refacotor" || echo "没有找到本项目相关的镜像"

echo ""
if [ "$MODE" = "single" ]; then
    echo "📝 将要执行的操作："
    echo "1. 停止服务: $SERVICE_NAME"
    echo "2. 删除相关的旧镜像（仅本项目）"
    echo "3. 重新构建服务: $SERVICE_NAME"
    echo "4. 启动服务并进行健康检查"
    echo ""
    echo "⚠️  注意：此操作只会影响指定的服务，不会影响其他容器"
else
    echo "📝 将要执行的操作："
    echo "1. 构建以下服务（不会影响其他容器）："
    for service in "${VALID_SERVICES[@]}"; do
        echo "   - $service"
    done
    echo ""
    echo "2. 清理操作："
    echo "   - 只删除本项目相关的旧镜像"
    echo "   - 不会使用 'docker system prune' 影响其他容器"
    echo ""
    echo "3. 启动服务并进行健康检查"
fi

echo ""
read -p "是否继续执行？(y/N): " -n 1 -r
echo ""
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "操作已取消。"
    exit 1
fi

# 创建日志目录
mkdir -p logs
if [ "$MODE" = "single" ]; then
    LOG_FILE="./logs/rebuild_${SERVICE_NAME}_$(date +%Y%m%d_%H%M%S).log"
else
    LOG_FILE="./logs/rebuild_all_$(date +%Y%m%d_%H%M%S).log"
fi

# 日志函数
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" | tee -a "$LOG_FILE"
}

log "=========================================="
if [ "$MODE" = "single" ]; then
    log "开始重新构建单个服务: $SERVICE_NAME"
else
    log "开始重新构建所有服务"
fi
log "=========================================="

# 尝试使用 docker compose（新版本）或 docker-compose（旧版本）
if command -v docker-compose &> /dev/null; then
    COMPOSE_CMD="docker-compose"
elif docker compose version &> /dev/null; then
    COMPOSE_CMD="docker compose"
else
    log "错误：未找到 docker-compose 或 docker compose 命令"
    exit 1
fi

# 获取服务的镜像名称（从docker-compose文件中提取）
get_service_image() {
    local service=$1
    case $service in
        "manage-service")
            echo "shop-goframe-micro-manage"
            ;;
        "manage-ui")
            echo "shop-goframe-micro-service-manage-ui-gfast"
            ;;
        *)
            echo "shop-goframe-micro-service-refacotor_${service}"
            ;;
    esac
}

# 获取服务的健康检查URL
get_health_check_url() {
    local service=$1
    case $service in
        "gateway-h5")
            echo "http://localhost:8199/api.json"
            ;;
        "gateway-resource")
            echo "http://localhost:8399/"
            ;;
        "manage-service")
            echo "http://localhost:8808/"
            ;;
        "manage-ui")
            echo "http://localhost/health"
            ;;
        "user-service")
            echo "http://localhost:31001/"
            ;;
        "goods-service")
            echo "http://localhost:31004/"
            ;;
        "order-service")
            echo "http://localhost:31005/"
            ;;
        "interaction-service")
            echo "http://localhost:31002/"
            ;;
        "search-service")
            echo "http://localhost:8499/"
            ;;
        "banner-service")
            echo "http://localhost:31006/"
            ;;
    esac
}

# 停止服务
stop_service() {
    local service=$1
    log "停止 $service 服务..."
    if $COMPOSE_CMD -f docker-compose.prod.yml stop "$service" 2>>"$LOG_FILE"; then
        log "✅ $service 服务停止成功"
    else
        log "⚠️  $service 服务停止失败或不存在"
    fi
}

# 清理镜像
cleanup_images() {
    local service=$1
    local service_image=$(get_service_image "$service")
    
    log "查找 $service 相关的镜像..."
    PROJECT_IMAGES=$(docker images --format "{{.Repository}}:{{.Tag}}" | grep "$service_image" || echo "")
    
    if [ -n "$PROJECT_IMAGES" ]; then
        log "找到以下相关镜像，将被删除："
        echo "$PROJECT_IMAGES" | tee -a "$LOG_FILE"
        echo "$PROJECT_IMAGES" | while read -r image; do
            if [ -n "$image" ]; then
                log "删除镜像: $image"
                docker rmi "$image" 2>/dev/null || log "无法删除镜像 $image（可能正在被使用）"
            fi
        done
    else
        log "没有找到 $service 相关的镜像需要清理"
    fi
}

# 构建服务
build_service() {
    local service=$1
    log "构建 $service 服务..."
    if $COMPOSE_CMD -f docker-compose.prod.yml build --no-cache "$service" 2>>"$LOG_FILE"; then
        log "✅ $service 服务构建成功"
        return 0
    else
        log "❌ $service 服务构建失败"
        return 1
    fi
}

# 启动服务
start_service() {
    local service=$1
    log "启动 $service 服务..."
    if $COMPOSE_CMD -f docker-compose.prod.yml up -d --force-recreate "$service" 2>>"$LOG_FILE"; then
        log "✅ $service 服务启动成功"
        return 0
    else
        log "❌ $service 服务启动失败"
        return 1
    fi
}

# 健康检查
health_check() {
    local service=$1
    local health_url=$(get_health_check_url "$service")
    
    if command -v curl &> /dev/null; then
        if curl -f -s "$health_url" > /dev/null; then
            log "✅ 健康检查通过: $health_url"
        else
            log "⚠️  健康检查失败: $health_url"
        fi
    elif command -v wget &> /dev/null; then
        if wget --quiet --tries=1 --spider "$health_url" 2>>"$LOG_FILE"; then
            log "✅ 健康检查通过: $health_url"
        else
            log "⚠️  健康检查失败: $health_url"
        fi
    else
        log "⚠️  未找到 curl 或 wget 命令，跳过健康检查"
    fi
}

# 显示服务访问地址
show_service_url() {
    local service=$1
    case $service in
        "gateway-h5")
            echo "  H5商城网关：http://localhost:8199"
            ;;
        "gateway-resource")
            echo "  资源网关：http://localhost:8399"
            ;;
        "manage-service")
            echo "  管理后台：http://localhost:8808"
            ;;
        "manage-ui")
            echo "  管理UI：http://localhost"
            ;;
        "user-service")
            echo "  用户服务：http://localhost:31001"
            ;;
        "goods-service")
            echo "  商品服务：http://localhost:31004"
            ;;
        "order-service")
            echo "  订单服务：http://localhost:31005"
            ;;
        "interaction-service")
            echo "  交互服务：http://localhost:31002"
            ;;
        "search-service")
            echo "  搜索服务：http://localhost:8499"
            ;;
        "banner-service")
            echo "  横幅服务：http://localhost:31006"
            ;;
    esac
}

# 检查服务状态
log "步骤1: 检查服务当前状态"
if [ "$MODE" = "single" ]; then
    $COMPOSE_CMD -f docker-compose.prod.yml ps "$SERVICE_NAME" | tee -a "$LOG_FILE" || log "$SERVICE_NAME 服务未运行"
else
    $COMPOSE_CMD -f docker-compose.prod.yml ps | tee -a "$LOG_FILE"
fi

# 根据模式执行不同的操作
if [ "$MODE" = "single" ]; then
    # 单个服务模式
    stop_service "$SERVICE_NAME"
    cleanup_images "$SERVICE_NAME"
    
    if build_service "$SERVICE_NAME"; then
        if start_service "$SERVICE_NAME"; then
            log "步骤6: 等待服务启动..."
            sleep 15
            
            log "步骤7: 最终状态检查"
            echo ""
            echo "服务状态："
            $COMPOSE_CMD -f docker-compose.prod.yml ps "$SERVICE_NAME" | tee -a "$LOG_FILE"
            
            echo ""
            echo "健康检查："
            health_check "$SERVICE_NAME"
            
            # 总结
            echo ""
            echo "=========================================="
            echo "✅ 单个服务重新构建完成！"
            echo "=========================================="
            echo ""
            echo "📊 构建统计："
            echo "  服务名称: $SERVICE_NAME"
            echo "  日志文件: $LOG_FILE"
            echo ""
            echo "🔍 检查命令："
            echo "  docker-compose -f docker-compose.prod.yml ps $SERVICE_NAME"
            echo "  docker-compose -f docker-compose.prod.yml logs $SERVICE_NAME"
            echo ""
            echo "🌐 服务访问地址："
            show_service_url "$SERVICE_NAME"
        else
            echo ""
            echo "=========================================="
            echo "❌ 启动失败！请检查日志文件: $LOG_FILE"
            echo "=========================================="
            exit 1
        fi
    else
        echo ""
        echo "=========================================="
        echo "❌ 构建失败！请检查日志文件: $LOG_FILE"
        echo "=========================================="
        exit 1
    fi
else
    # 所有服务模式
    BUILD_SUCCESS=true
    
    log "步骤2: 停止本项目的服务（如果存在）"
    for service in "${VALID_SERVICES[@]}"; do
        stop_service "$service"
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
    for service in "${VALID_SERVICES[@]}"; do
        if ! build_service "$service"; then
            BUILD_SUCCESS=false
            break
        fi
    done
    
    if [ "$BUILD_SUCCESS" = true ]; then
        log "步骤5: 启动所有服务"
        for service in "${VALID_SERVICES[@]}"; do
            if ! start_service "$service"; then
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
        log "✅ 所有服务重新构建完成！"
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
fi

echo ""
echo "✨ 安全特性："
echo "  - 没有使用 'docker system prune'"
echo "  - 只影响本项目相关容器"
echo "  - 保留了所有其他项目的容器"