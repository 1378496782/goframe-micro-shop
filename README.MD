# GoFrame 微服务电商项目

基于 [GoFrame v2](https://goframe.org/) 的微服务电商系统，采用 gRPC + HTTP 网关架构。

## 项目架构

| 服务 | 端口 | 说明 | 类型 |
|------|------|------|------|
| user | 31001 | 用户登录、注册、信息管理 | gRPC |
| interaction | 31002 | 点赞、评论、收藏 | gRPC |
| admin | 31003 | 管理员登录、注册 | gRPC |
| goods | 31004 | 商品增删改查、库存管理 | gRPC |
| order | 31005 | 订单、退款、微信支付 | gRPC |
| banner | 31006 | 轮播图、位置管理 | gRPC |
| gateway-h5 | 8199 | 用户端网关（H5/小程序） | HTTP |
| gateway-admin | 8299 | 管理端网关 | HTTP |
| gateway-resource | 8399 | 文件上传（七牛云） | HTTP |
| search | 8499 | 搜索服务（Elasticsearch） | HTTP |
| worker | 8599 | 异步任务处理（订单超时等） | HTTP |

---

## 快速启动（Docker Compose 一键部署）

### 前置条件

- [Docker](https://www.docker.com/) & Docker Compose

### 启动

```bash
docker-compose up -d
```

一条命令启动全部中间件 + 所有 Go 服务，包括：
- MySQL（自动执行 `init-db/` 下的 SQL 初始化脚本）
- Redis、RabbitMQ（含延迟消息插件）、etcd
- Elasticsearch + Kibana
- 全部 11 个 Go 微服务/网关

### 验证

| 服务 | 地址 |
|------|------|
| gateway-h5 Swagger | http://localhost:8199/swagger |
| gateway-admin Swagger | http://localhost:8299/swagger |
| search Swagger | http://localhost:8499/swagger |
| RabbitMQ 管理界面 | http://localhost:15672 (root/root) |
| Elasticsearch | http://localhost:9200 |
| Kibana | http://localhost:5601 |

---

## 本地开发模式（中间件用 Docker，代码本地运行）

适合需要断点调试、热重载的开发场景。

### 1. 仅启动中间件

```bash
docker-compose up -d mysql redis rabbitmq etcd elasticsearch
```

### 2. 准备配置文件

各服务的 `manifest/config/config.yaml.bak` 是本地开发配置模板，复制为 `config.yaml` 即可：

```powershell
# PowerShell 一键复制
Get-ChildItem -Path app -Recurse -Filter "config.yaml.bak" | ForEach-Object {
    $dest = Join-Path $_.DirectoryName "config.yaml"
    Copy-Item $_.FullName $dest
}
```

```bash
# Bash / Git Bash
find app -name "config.yaml.bak" | while read f; do cp "$f" "$(dirname $f)/config.yaml"; done
```

默认配置已适配本地环境（MySQL root/root@127.0.0.1:3306、Redis 127.0.0.1:6379、RabbitMQ root/root@127.0.0.1:5672、etcd 127.0.0.1:2379）。

### 3. 启动服务

先启动 gRPC 微服务，再启动网关（每个服务占一个终端窗口）：

```bash
# 启动你需要的微服务，例如：
cd app/user && go run main.go
cd app/goods && go run main.go
cd app/order && go run main.go

# 再启动对应的网关
cd app/gateway-h5 && go run main.go
```

> 最小启动集：etcd + 目标微服务 + 对应网关。例如只调试商品功能，启动 `goods` → `gateway-h5` 即可。

---

## 常见问题

**Q: docker-compose 启动后 Go 服务报错？**
A: MySQL 初始化需要时间，Go 服务可能先于 MySQL 就绪而启动失败。执行 `docker-compose restart` 重启服务即可。

**Q: 网关正常但调用微服务报错？**
A: 微服务需先于网关在 etcd 中完成注册，确保对应微服务已正常运行。

**Q: 本地开发模式连接不上中间件？**
A: 检查 `config.yaml` 中的地址是否为 `127.0.0.1`，而非 Docker 容器名（`mysql`、`etcd` 等是 `config.prod.yaml` 中的 Docker 网络地址）。
