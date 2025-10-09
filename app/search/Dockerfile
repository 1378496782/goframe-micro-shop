# 使用官方 Go 镜像
FROM golang:1.24.5-alpine AS builder

# 设置工作目录
WORKDIR /app

# 设置环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目文件
COPY . .

# 构建 search 服务
RUN go build -o bin/search ./app/search

# 使用更小的运行时镜像
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk --no-cache add ca-certificates tzdata

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/bin/search ./

# 复制配置文件
COPY app/search/manifest/docker/ ./config/

# 设置启动命令
CMD ["./search"]