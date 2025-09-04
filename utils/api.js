// API配置和请求封装
const config = require('../config/env');

// 环境配置
const isMock = config.features.useMock;
const BASE_URL = isMock ? '' : config.api.baseURL;

// 错误状态模拟配置
const errorSimulation = {
  enabled: config.features.debug, // 仅在调试模式开启
  rate: 0.1, // 10%的错误率
  errorCodes: [400, 401, 403, 404, 500],
  errorMessages: {
    400: '请求参数错误',
    401: '未授权访问',
    403: '访问被禁止',
    404: '资源不存在',
    500: '服务器内部错误'
  }
};

// 模拟错误响应
function simulateError() {
  if (!errorSimulation.enabled || Math.random() > errorSimulation.rate) {
    return null;
  }
  
  const errorCode = errorSimulation.errorCodes[
    Math.floor(Math.random() * errorSimulation.errorCodes.length)
  ];
  
  return {
    code: errorCode,
    message: errorSimulation.errorMessages[errorCode] || '未知错误',
    data: null
  };
}

// Mock数据
const mockData = {
  // 首页相关
  '/frontend/goods': {
    method: 'GET',
    response: {
      code: 200,
      message: 'success',
      data: {
        list: [
          {
            id: 1,
            name: '高品质智能手机 8GB+256GB',
            price: '2999.00',
            originalPrice: '3999.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 12560
          },
          {
            id: 2,
            name: '轻薄笔记本电脑 i7处理器',
            price: '5999.00',
            originalPrice: '6999.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 8560
          }
        ],
        total: 2,
        page: 1,
        size: 10
      }
    }
  },

  '/frontend/goods/detail': {
    method: 'GET',
    response: {
      code: 200,
      message: 'success',
      data: {
        id: 1,
        name: '高品质智能手机 8GB+256GB',
        price: '2999.00',
        originalPrice: '3999.00',
        images: [
          'http://wangzhongyang.com/images/logo_removebg.png',
          'http://wangzhongyang.com/images/logo_removebg.png'
        ],
        description: '高性能智能手机，8GB运行内存，256GB存储空间',
        specifications: [
          { name: '颜色', values: ['黑色', '白色', '蓝色'] },
          { name: '存储', values: ['128GB', '256GB', '512GB'] }
        ],
        stock: 100,
        sales: 12560
      }
    }
  },

  // 分类相关
  '/frontend/category/all': {
    method: 'GET',
    response: {
      code: 200,
      message: 'success',
      data: [
        { id: 1, name: '手机', icon: 'http://wangzhongyang.com/images/logo_removebg.png' },
        { id: 2, name: '电脑', icon: 'http://wangzhongyang.com/images/logo_removebg.png' },
        { id: 3, name: '配件', icon: 'http://wangzhongyang.com/images/logo_removebg.png' },
        { id: 4, name: '家电', icon: 'http://wangzhongyang.com/images/logo_removebg.png' },
        { id: 5, name: '服饰', icon: 'http://wangzhongyang.com/images/logo_removebg.png' }
      ]
    }
  },

  // 用户相关
  '/frontend/user/login': {
    method: 'POST',
    response: {
      code: 200,
      message: '登录成功',
      data: {
        token: 'mock_jwt_token_123456',
        userInfo: {
          id: 1,
          username: 'testuser',
          avatar: 'http://wangzhongyang.com/images/logo_removebg.png'
        }
      }
    }
  },

  '/frontend/user/info': {
    method: 'POST',
    response: {
      code极速版: 200,
      message: 'success',
      data: {
        id: 1,
        username: 'testuser',
        avatar: 'http://wangzhongyang.com/images/logo_removebg.png',
        email: 'test@example.com',
        phone: '13800138000'
      }
    }
  },

  // 购物车相关Mock数据
  '/frontend/cart': {
    method: 'GET',
    response: {
      code: 200,
      message: 'success',
      data: {
        items: [
          {
            id: 1,
            productId: 1,
            productName: '高品质智能手机 8GB+256GB',
            price: '2999.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            quantity: 2,
            selected: true
          },
          {
            id: 2,
            productId: 2,
            productName: '轻薄笔记本电脑 i7处理器',
            price: '5999.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            quantity: 1,
            selected: true
          }
        ],
        totalPrice: '11997.00',
        totalQuantity: 3
      }
    }
  },

  // 订单相关Mock数据
  '/frontend/order': {
    method: 'GET',
    response: {
      code: 200,
      message: 'success',
      data: {
        list: [
          {
            id: 'ORDER202401010001',
            status: 2, // 1: 待付款, 2: 待发货, 3: 已发货, 4: 已完成, 5: 已取消
            totalAmount: '11997.00',
            createTime: '2024-01-01 10:00:00',
            products: [
              {
                name: '高品质智能手机 8GB+256GB',
                image: 'http://wangzhongyang.com/images/logo_removebg.png',
                price: '2999.00',
                quantity: 2
              },
              {
                name: '轻薄笔记本电脑 i7处理器',
                image: 'http://wangzhongyang.com/images/logo_removebg.png',
                price: '5999.00',
                quantity: 1
              }
            ]
          }
        ],
        total: 1,
        page: 1,
        size: 10
      }
    }
  }
};

// 数据缓存
const cache = new Map();
const CACHE_DURATION = 5 * 60 * 1000; // 5分钟缓存

// 统一的请求封装
function request(url, data = {}, method = 'GET') {
  return new Promise((resolve, reject) => {
    // 生成缓存键
    const cacheKey = `${method}:${url}:${JSON.stringify(data)}`;
    
    // 检查缓存（仅对GET请求）
    if (method === 'GET') {
      const cached = cache.get(cacheKey);
      if (cached && Date.now() - cached.timestamp < CACHE_DURATION) {
        console.log('[API] 使用缓存数据:', cacheKey);
        resolve(cached.data);
        return;
      }
    }
    
    if (isMock && mockData[url] && mockData[url].method === method) {
      // 使用Mock数据
      setTimeout(() => {
        // 模拟错误
        const errorResponse = simulateError();
        if (errorResponse) {
          reject(errorResponse);
          return;
        }
        
        const response = mockData[url].response;
        // 缓存GET请求结果
        if (method === 'GET') {
          cache.set(cacheKey, {
            data: response,
            timestamp: Date.now()
          });
        }
        resolve(response);
      }, 500); // 模拟网络延迟
    } else {
      // 真实网络请求
      wx.request({
        url: `${BASE_URL}${url}`,
        data: data,
        method: method,
        success: (res) => {
          // 缓存GET请求结果
          if (method === 'GET') {
            cache.set(cacheKey, {
              data: res.data,
              timestamp: Date.now()
            });
          }
          resolve(res.data);
        },
        fail: (err) => {
          reject(err);
        }
      });
    }
  });
}

// API方法封装
const api = {
  // 商品相关
  getGoodsList: (params) => request('/frontend/goods', params, 'GET'),
  getGoodsDetail: (id) => request('/frontend/goods/detail', { id }, 'GET'),

  // 分类相关
  getCategories: () => request('/frontend/category/all', {}, 'GET'),

  // 用户相关
  login: (data) => request('/frontend/user/login', data, 'POST'),
  getUserInfo: () => request('/frontend/user/info', {}, 'POST'),
  register: (data) => request('/frontend/user/register', data, 'POST'),
  updatePassword: (data) => request('/frontend/user/update/password', data, 'PUT'),

  // 收藏相关
  getCollections: () => request('/frontend/collection', {}, 'GET'),
  addCollection: (data) => request('/frontend/collection', data, 'POST'),
  removeCollection: (id) => request('/frontend/collection', { id }, 'DELETE'),

  // 购物车相关
  getCart: () => request('/frontend/cart', {}, 'GET'),
  addToCart: (data) => request('/frontend/cart', data, 'POST'),
  updateCartItem: (data) => request('/frontend/cart', data, 'PUT'),
  removeCartItem: (id) => request('/frontend/cart', { id }, 'DELETE'),

  // 订单相关
  createOrder: (data) => request('/frontend/order', data, 'POST'),
  getOrders: (params) => request('/frontend/order', params, 'GET'),
  getOrderDetail: (id) => request('/frontend/order/detail', { id }, 'GET'),
  cancelOrder: (id) => request('/frontend/order/cancel', { id }, 'PUT')
};

module.exports = {
  api,
  isMock,
  switchToRealAPI: () => { isMock = false; }
};