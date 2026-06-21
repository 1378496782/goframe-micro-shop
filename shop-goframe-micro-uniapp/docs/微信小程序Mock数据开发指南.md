# 微信小程序Mock数据开发指南

## 📅 开发日期
2025年8月27日

by 王中阳

## 🎯 今日开发目标
1. 解决图片404错误问题
2. 构建完整的Mock数据架构
3. 实现前后端并行开发能力
4. 创建完善的开发文档

## 🔧 问题解决记录

### 1. 图片404错误修复
**问题**：网络图片链接失效，返回404错误
**解决方案**：
- 将所有网络图片替换为稳定的图片链接
- 统一使用：`http://wangzhongyang.com/images/logo_removebg.png`
- 修改文件：`pages/index/index.js`

### 2. Git提交冲突解决
**问题**：Git提交时出现子模块冲突和文件编码问题
**解决方案**：
- 排除有问题的子模块目录：`AI开发记录/source/`
- 逐个添加需要提交的文件
- 解决README.md合并冲突

## 🏗️ Mock数据架构设计

### 核心架构
```
src/
├── config/
│   └── env.js              # 环境配置
├── utils/
│   └── api.js              # API请求封装
├── types/
│   └── api.d.ts            # TypeScript类型定义
└── docs/
    └── MOCK_DATA_GUIDE.md  # 开发指南
```

### 1. 环境配置系统 (`config/env.js`)
```javascript
const config = {
  env: 'development', // 环境：development | test | production
  features: {
    useMock: true,     // Mock开关
    debug: true,       // 调试模式
    errorSimulation: true, // 错误模拟
    cacheEnabled: true // 缓存功能
  },
  mock: {
    networkDelay: 500,   // 网络延迟
    errorRate: 0.1,      // 错误率
    cacheDuration: 300000 // 缓存时长(5分钟)
  }
};
```

### 2. 统一的API封装 (`utils/api.js`)
**核心特性**：
- Mock数据与真实接口无缝切换
- 自动数据缓存（仅GET请求）
- 错误状态模拟
- 网络延迟模拟

**使用方法**：
```javascript
const { api } = require('../../utils/api');

// 获取商品列表
api.getGoodsList({ page: 1, size: 10 })
  .then(res => console.log(res))
  .catch(err => console.error(err));
```

### 3. TypeScript类型安全 (`types/api.d.ts`)
提供完整的类型定义：
- 基础响应类型 `BaseResponse<T>`
- 分页类型 `PaginatedResponse<T>`
- 业务数据类型：商品、用户、购物车、订单等

## 🚀 已实现的Mock接口

### 首页模块
- `GET /frontend/goods` - 商品列表
- `GET /frontend/goods/detail` - 商品详情
- `GET /frontend/category/all` - 全部分类

### 用户模块
- `POST /frontend/user/login` - 用户登录
- `POST /frontend/user/info` - 用户信息
- `POST /frontend/user/register` - 用户注册
- `PUT /frontend/user/update/password` - 修改密码

### 购物车模块
- `GET /frontend/cart` - 购物车列表
- `POST /frontend/cart` - 添加商品
- `PUT /frontend/cart` - 更新数量
- `DELETE /frontend/cart` - 删除商品

### 订单模块
- `POST /frontend/order` - 创建订单
- `GET /frontend/order` - 订单列表
- `GET /frontend/order/detail` - 订单详情
- `PUT /frontend/order/cancel` - 取消订单

## 💡 开发经验总结

### 1. 前后端并行开发策略
**优势**：
- 不阻塞前端开发进度
- 提前进行接口联调测试
- 确保数据结构一致性

**实施步骤**：
1. 根据接口文档创建Mock数据
2. 实现统一的请求封装
3. 开发阶段使用Mock数据
4. 后端接口完成后切换至真实环境

### 2. 错误处理最佳实践
```javascript
// 页面中的错误处理示例
async loadData() {
  try {
    const res = await api.getData();
    if (res.code === 200) {
      // 处理成功数据
    } else {
      // 处理业务错误
      wx.showToast({ title: res.message, icon: 'none' });
    }
  } catch (error) {
    // 处理网络错误
    console.error('请求失败:', error);
  }
}
```

### 3. 性能优化技巧
- **数据缓存**：对GET请求结果进行缓存
- **批量请求**：使用`Promise.all`并行请求
- **防抖处理**：对频繁操作进行防抖处理
- **图片优化**：使用CDN和合适的图片格式

## 🛠️ 快速上手指南

### 1. 环境设置
```bash
# 克隆项目
git clone <repository-url>

# 安装依赖（如果有）
npm install

# 启动微信开发者工具
```

### 2. 开发新功能步骤
1. **分析需求**：确定需要哪些接口
2. **查看文档**：阅读接口文档 (`api-docs.json`)
3. **添加Mock**：在 `utils/api.js` 中添加Mock数据
4. **定义类型**：在 `types/api.d.ts` 中添加类型定义
5. **页面开发**：在页面中使用API方法
6. **测试验证**：测试各种场景和错误状态

### 3. 常见问题解决
**Q: Mock数据不生效？**
A: 检查 `config/env.js` 中的 `useMock` 配置

**Q: 如何添加新接口？**
A: 参考现有接口的实现模式

**Q: 生产环境如何配置？**
A: 设置 `env: 'production'` 和 `useMock: false`

## 📈 后续开发建议

### 1. 接口扩展
- [ ] 支付相关接口
- [ ] 物流跟踪接口
- [ ] 消息通知接口
- [ ] 数据统计接口

### 2. 功能优化
- [ ] 请求重试机制
- [ ] 离线数据存储
- [ ] 接口版本管理
- [ ] 性能监控

### 3. 开发工具
- [ ] API文档自动生成
- [ ] Mock数据管理界面
- [ ] 接口测试工具

## 👥 团队协作规范

### 1. 代码提交
```bash
# 功能开发
git commit -m "feat: 添加购物车功能"

# Bug修复
git commit -m "fix: 修复图片加载问题"

# 文档更新
git commit -m "docs: 更新开发指南"
```

### 2. 代码审查
- 检查Mock数据与接口文档一致性
- 验证错误处理是否完整
- 确保类型定义准确
- 测试边界情况和异常场景

### 3. 版本发布
1. 开发环境：使用Mock数据，开启调试功能
2. 测试环境：连接测试服务器，关闭Mock
3. 生产环境：连接生产服务器，关闭调试功能

## 🎯 总结

今天的开发工作构建了一个完整的前后端并行开发体系，主要成果包括：

1. **问题解决**：修复了图片404错误和Git提交问题
2. **架构设计**：创建了可扩展的Mock数据架构
3. **接口实现**：完成了核心业务的Mock接口
4. **文档完善**：提供了详细的开发指南和最佳实践

这套架构可以让团队在前后端分离的情况下高效协作，确保项目进度和质量。

---

**让开发更高效，让协作更顺畅！** 🚀