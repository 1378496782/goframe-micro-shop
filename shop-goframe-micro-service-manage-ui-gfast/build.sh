#!/bin/bash

# 前端项目构建脚本
# 仅用于编译静态文件到 dist 目录

set -e

echo "=== 前端项目构建脚本 ==="
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
echo "构建成功！文件大小: $(du -sh dist 2>/dev/null | cut -f1 2>/dev/null || echo '无法获取大小')"
echo "文件列表:"
ls -la dist/

echo "=== 构建脚本完成 ==="
echo "静态文件已编译到 dist 目录"