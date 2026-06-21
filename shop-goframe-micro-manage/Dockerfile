# build stage
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

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app/bin/main .

# final stage
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk --no-cache add ca-certificates tzdata

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/bin/main .

# 复制配置文件和资源
COPY resource ./resource
COPY manifest/config/config.prod.yaml ./config.yaml
COPY manifest/config/wechat.yaml ./wechat.yaml

# 暴露端口
EXPOSE 8808

# 设置默认启动命令
CMD ["./main"]