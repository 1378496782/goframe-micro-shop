# build stage
FROM node:18-alpine AS builder

# 设置工作目录
WORKDIR /app

# 设置 npm 镜像源为国内源，加速构建
RUN npm config set registry https://registry.npmmirror.com

# 复制包管理文件
COPY package*.json ./

# 安装依赖
RUN npm install

# 复制项目文件
COPY . .

# 构建生产版本
RUN npm run build

# production stage
FROM nginx:alpine

# 安装必要的工具
RUN apk add --no-cache tzdata curl

# 设置时区为上海
ENV TZ=Asia/Shanghai

# 复制自定义 nginx 配置
COPY nginx.conf /etc/nginx/nginx.conf

# 从构建阶段复制构建产物
COPY --from=builder /app/dist /usr/share/nginx/html

# 创建 nginx 缓存目录
RUN mkdir -p /var/cache/nginx/client_temp /var/cache/nginx/proxy_temp \
    /var/cache/nginx/fastcgi_temp /var/cache/nginx/uwsgi_temp /var/cache/nginx/scgi_temp

# 暴露端口
EXPOSE 80

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost/health || exit 1

# 启动 nginx
CMD ["nginx", "-g", "daemon off;"]