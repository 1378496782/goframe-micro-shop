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

# 复制项目文件（包括utility工具包）
COPY . .

# 构建所有服务和网关
RUN go build -o bin/goods ./app/goods
RUN go build -o bin/interaction ./app/interaction
RUN go build -o bin/order ./app/order
RUN go build -o bin/user ./app/user
RUN go build -o bin/search ./app/search
RUN go build -o bin/gateway-h5 ./app/gateway-h5
RUN go build -o bin/gateway-resource ./app/gateway-resource
RUN go build -o bin/banner ./app/banner
RUN go build -o bin/admin ./app/admin
RUN go build -o bin/worker ./app/worker
RUN go build -o bin/gateway-admin ./app/gateway-admin

# 使用更小的运行时镜像
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk --no-cache add ca-certificates tzdata

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/bin/ .

# 复制所有配置文件
COPY app/*/manifest/docker/ ./config/

# 设置默认启动命令
CMD ["./admin"]