# Protobuf 开发指南

## 问题背景
团队开发中由于不同开发者使用的protoc版本和插件版本不一致，导致生成的pb文件频繁冲突。

## 解决方案
1. **忽略生成的pb文件** - 已添加到.gitignore
2. **统一构建环境** - 使用版本控制文件
3. **自动化生成** - 使用构建脚本

## 开发流程

### 1. 安装依赖
```bash
# 查看版本要求
cat .protoc-version

# 安装指定版本的protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.9

# 安装指定版本的protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
```

### 2. 生成pb文件
```bash
# 使用构建脚本（推荐）
./generate-proto.sh

# 或手动生成
find . -name "*.proto" | while read proto_file; do
    protoc --go_out=. --go-grpc_out=. "$proto_file"
done
```

### 3. 开发流程
1. 修改proto文件
2. 运行生成脚本
3. 编译测试
4. 提交代码（不提交pb文件）

## 注意事项

### 必须遵守的规则
1. **不要提交pb文件** - 这些文件会被.gitignore忽略
2. **使用统一版本** - 严格按照.protoc-version中的版本
3. **构建前生成** - 每次修改proto文件后必须重新生成

### 版本要求
- protoc: 6.32.0
- protoc-gen-go: v1.36.9
- protoc-gen-go-grpc: v1.2.0

## 故障排除

### 常见问题
1. **生成失败**: 检查protoc和插件版本
2. **导入错误**: 确保所有依赖的proto文件都存在
3. **编译错误**: 重新生成所有pb文件

### 紧急情况
如果遇到无法解决的pb文件冲突：
```bash
# 删除所有pb文件重新生成
git rm -r --cached **/*.pb.go **/*_grpc.pb.go
./generate-proto.sh
```

## 联系方式
如有问题请联系项目负责人或查看项目文档。