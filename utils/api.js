// API配置和请求封装
const config = require('../config/env');

// 环境配置
const isMock = config.features.useMock;
const BASE_URL = isMock ? '' : config.api.baseURL;

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
      code: 200,
      message: 'success',
      data: {
        id: 1,
        username: 'testuser',
        avatar: 'http://wangzhongyang.com/images/logo_removebg.png',
        email: 'test@example.com',
        phone: '13800138000'
      }
    }
  }
};

// 统一的请求封装
function request(url, data = {}, method = 'GET') {
  return new Promise((resolve, reject) => {
    if (isMock && mockData[url] && mockData[url].method === method) {
      // 使用Mock数据
      setTimeout(() => {
        resolve(mockData[url].response);
      }, 500); // 模拟网络延迟
    } else {
      // 真实网络请求
      wx.request({
        url: `${BASE_URL}${url}`,
        data: data,
        method: method,
        success: (res) => {
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
  removeCollection: (id) => request('/frontend/collection', { id }, 'DELETE')
};

module.exports = {
  api,
  isMock,
  switchToRealAPI: () => { isMock = false; }
};