# Mock数据使用指南

## 📋 概述

本文档介绍如何在微信小程序开发中使用Mock数据来模拟后端接口，实现前后端并行开发。

## 🚀 快速开始

### 1. 环境配置

在 `config/env.js` 中配置环境：

```javascript
const config = {
  env: 'development', // development | production
  features: {
    useMock: true,     // 是否使用Mock数据
    debug: true        // 是否开启调试
  }
};
```

### 2. 使用API方法

在页面中引入API工具：

```javascript
const { api } = require('../../utils/api');

// 获取商品列表
api.getGoodsList({ page: 1, size: 10 })
  .then(res => {
    console.log('商品数据:', res);
  })
  .catch(err => {
    console.error('请求失败:', err);
  });
```

## 📊 当前支持的Mock接口

### 首页相关
- `GET /frontend/goods` - 获取商品列表
- `GET /frontend/goods/detail` - 获取商品详情
- `GET /frontend/category/all` - 获取全部分类

### 用户相关
- `POST /frontend/user/login` - 用户登录
- `POST /frontend/user/info` - 获取用户信息
- `POST /frontend/user/register` - 用户注册
- `PUT /frontend/user/update/password` - 修改密码

### 收藏相关
- `GET /frontend/collection` - 获取收藏列表
- `POST /frontend/collection` - 添加收藏
- `DELETE /frontend/collection` - 删除收藏

## 🔧 添加新的Mock接口

### 1. 在 `utils/api.js` 中添加Mock数据

```javascript
const mockData = {
  '/frontend/your-api': {
    method: 'GET',
    response: {
      code: 200,
      message: 'success',
      data: {
        // 你的Mock数据
      }
    }
  }
};
```

### 2. 添加API方法

```javascript
const api = {
  yourApiMethod: (params) => request('/frontend/your-api', params, 'GET')
};
```

## 🔄 切换到真实接口

当后端接口准备好后：

1. 修改 `config/env.js`：
```javascript
features: {
  useMock: false  // 关闭Mock模式
}
```

2. 确保 `config.api.baseURL` 指向正确的后端地址

## 🎯 最佳实践

1. **保持数据结构一致**：Mock数据应该与真实接口返回的数据结构完全一致
2. **模拟网络延迟**：Mock请求有500ms延迟，模拟真实网络环境
3. **错误处理**：同时处理Mock和真实环境的错误情况
4. **类型安全**：确保请求参数和响应数据的类型正确

## 🐛 常见问题

### Q: Mock数据不生效？
A: 检查 `config/env.js` 中的 `useMock` 配置是否为 `true`

### Q: 如何添加新的接口？
A: 参考上面的"添加新的Mock接口"部分

### Q: 生产环境如何配置？
A: 设置 `env: 'production'` 和 `useMock: false`

## 📞 技术支持

如有问题，请查看具体的接口文档或联系开发团队。

---

**Happy Mocking!** 🎉