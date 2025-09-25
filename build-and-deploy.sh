#!/bin/bash

# 前端项目构建和部署脚本
# 用于在微服务架构中构建和更新前端项目

set -e

echo "=== 前端项目构建和部署脚本 ==="
echo "当前目录: $(pwd)"

# 检查是否在正确的目录
if [ ! -f "package.json" ]; then
    echo "错误: 当前目录不是前端项目根目录"
    echo "请切换到 shop-goframe-micro-service-manage-ui-gfast 目录"
    exit 1
fi

echo "1. 安装依赖..."
npm install

echo "2. 构建项目..."
npm run build

echo "3. 检查构建结果..."
if [ ! -d "dist" ] || [ ! -f "dist/index.html" ]; then
    echo "错误: 构建失败，dist目录不存在或index.html不存在"
    exit 1
fi

echo "4. 显示构建结果..."
echo "构建成功！文件大小: $(du -sh dist | cut -f1)"
echo "文件列表:"
ls -la dist/

echo "5. 检查docker-compose状态..."
cd ../shop-goframe-micro-service-refacotor
if [ -f "docker-compose.prod.yml" ]; then
    echo "docker-compose文件存在，检查服务状态..."
    docker compose -f docker-compose.prod.yml ps manage-ui || echo "manage-ui服务未运行"
    
    echo ""
    echo "6. 更新提示:"
    echo "要更新前端服务，请执行:"
    echo "  docker compose -f docker-compose.prod.yml restart manage-ui"
    echo ""
    echo "或者直接重启所有服务:"
    echo "  docker compose -f docker-compose.prod.yml restart"
else
    echo "错误: 未找到docker-compose.prod.yml文件"
    exit 1
fi

echo "=== 构建和部署脚本完成 ==="