# 微信小程序商城项目

一个完整的微信小程序商城项目，包含商品展示、分类浏览、购物车、用户中心等核心功能。

## 🚀 功能特性

### 核心功能
- **首页** - 轮播图、分类导航、热门商品推荐
- **分类页** - 侧边栏分类、商品网格展示
- **商品详情** - 图片轮播、规格选择、加入购物车
- **购物车** - 商品管理、数量控制、价格计算
- **用户极速版中心** - 订单状态、个人资料、功能菜单

### 技术特性
- 纯微信小程序原生开发
- 符合微信小程序开发规范
- 响应式布局设计
- 简洁易用的UI界面

## 📁 项目结构

```
shop-miniprogram/
├── app.js              # 小程序入口文件
├── app.json            # 全局配置文件
├── app.wxss            # 全局样式文件
├── project.config.json # 项目配置文件
├── sitemap.json        # 搜索索引配置
└── pages/              # 页面目录
    ├── index/          # 首页
    │   ├── index.wxml
    │   ├── index.wxss
    │   ├── index.js
    │   └── index.json
    ├── category/       # 分类页
    ├── cart/          # 购物车
    ├── user/          # 用户中心
    └── product-detail/ # 商品详情
```

## 🛠️ 开发环境搭建

### 1. 安装微信开发者工具
- 访问[微信公众平台](https://mp.weixin.qq.com/)
- 下载并安装最新版微信开发者工具

### 2. 导入项目
1. 打开微信开发者工具
2. 点击"新建项目"
3. 选择项目目录（当前文件夹）
4. 填写AppID（可使用测试号）
5. 项目名称填写"shop-miniprogram"

### 3. 运行项目
- 点击"编译"按钮预览效果
- 使用"预览"功能在手机上测试
- 使用"真机调试"进行深度调试

## 📖 开发指南

### 页面开发
每个页面包含4个文件：
- `.wxml` - 页面结构（类似HTML）
- `.wx极速版ss` - 页面样式（类似CSS）
- `.js` - 页面逻辑
- `.json` - 页面配置

### 数据绑定
```javascript
// page.js
Page({
  data: {
    message: 'Hello World'
  }
})

// page.wxml
<view>{{message}}</view>
```

### 事件处理
```javascript
// 定义事件处理函数
onTap: function(e) {
  console.log('按钮被点击', e)
}

// wxml中绑定事件
<button bindtap="onTap">点击我</button>
```

### 页面跳转
```javascript
// 跳转到tab页
wx.switchTab({
  url: '/pages/index/index'
})

// 跳转到普通页面
wx.navigateTo({
  url: '/pages/product-detail/product-detail?id=123'
})
```

## 🎨 UI组件说明

### 常用组件
- `view` - 视图容器
- `text` - 文本
- `image` - 图片
- `scroll-view` - 可滚动视图
- `swiper` - 轮播图
- `button` - 按钮
- `input` - 输入框

### 样式规范
- 主色调：`#19aecc`
- 警示色：`#F53F3F`
- 成功色：`#07C160`
- 文字色：`#1D2129`（主要）、`#4E5969`（常规）、`#86909C`（次要）

## 🔧 常用API

### 数据存储
```javascript
// 同步存储
wx.setStorageSync('key', 'value')
const value = wx.getStorageSync('key')

// 异步存储
wx.setStorage({
  key: 'key',
  data: 'value'
})
```

### 网络请求
```javascript
wx.request({
  url: 'https://api.example.com/data',
  method: 'GET',
  success: (res) => {
    console.log(res.data)
  }
})
```

### 用户交互
```javascript
// 显示提示
wx.showToast({
  title: '操作成功',
  icon极速版: 'success'
})

// 显示加载中
wx.showLoading({
  title: '加载中...'
})

// 显示模态框
wx.showModal({
  title: '提示',
  content: '确定要删除吗？'
})
```

## 📱 页面功能详解

### 首页 (index)
- 搜索栏功能
- 轮播图展示
- 分类快捷入口
- 热门商品推荐

### 商品详情 (product-detail)
- 商品图片轮播
- 规格选择器
- 价格信息展示
- 加入购物车/立即购买

### 购物车 (cart)
- 商品列表管理
- 数量加减控制
- 全选/反选功能
- 价格实时计算

## 🚀 部署发布

### 测试阶段
1. 使用微信开发者工具预览功能
2. 生成体验版二维码进行测试
3. 收集用户反馈并修复问题

### 发布流程
1. 在微信公众平台提交审核
2. 等待审核通过（通常1-3个工作日）
3. 发布上线

## 🤝 贡献指南

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开Pull Request

## 📄 许可证

本项目仅用于学习交流，请遵守微信小程序平台相关规定。

## 📞 技术支持

如有问题请提交Issue或联系开发团队。

---

**Happy Coding!** 🎉