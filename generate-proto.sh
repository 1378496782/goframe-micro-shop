#!/bin/bash
set -e

echo "生成protobuf文件..."

# 检查protoc是否安装
if ! command -v protoc &> /dev/null; then
    echo "错误: protoc未安装，请先安装protoc"
    exit 1
fi

# 遍历所有proto文件并生成
find . -name "*.proto" | while read proto_file; do
    echo "生成: $proto_file"
    protoc --go_out=. --go-grpc_out=. "$proto_file"
done

echo "protobuf文件生成完成"