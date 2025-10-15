// API配置和请求封装
const { config, API } = require('../config/index');

// 使用统一环境配置
const isMock = false; // 禁用mock模式，使用真实API
const BASE_URL = config.BASE_URL; // 使用统一配置的API地址

// 错误状态模拟配置
const errorSimulation = {
  enabled: false, // 禁用错误模拟
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

// Mock数据
const mockData = {
  // 首页相关 - 轮播图
  '/frontend/rotation': {
    method: 'GET',
    response: {
      code: 0,
      message: 'OK',
      data: {
        list: [
          {
            "id": 1,
            "pic_url": "http://wangzhongyang.com/images/logo_removebg.png",
            "link": "https://example.com",
            "sort": 10,
            "created_at": {
              "seconds": 1658206381
            },
            "updated_at": {
              "seconds": 1658206764
            }
          }
        ],
        "page": 1,
        "size": 10,
        "total": 1
      }
    }
  },

  // 首页相关 - 热门推荐商品
  '/goods': {
    method: 'GET',
    response: {
      code: 0,
      message: 'OK',
      data: {
        list: [
          {
            "id": 1,
            "pic_url": "",
            "images": "{\"image\": \"http://dummyimage.com/400x400\"}",
            "name": "转群治",
            "price": 99,
            "level1_category_id": 3,
            "level2_category_id": 49,
            "level3_category_id": 14,
            "brand": "veniam",
            "stock": 69284016,
            "sale": 61,
            "tags": "velit ut proident",
            "detail_info": "officia tempor fugiat culpa",
            "created_at": {
              "seconds": 1756970352
            },
            "updated_at": {
              "seconds": 1756970352
            }
          },
          {
            "id": 2,
            "pic_url": "",
            "images": "{\"image\": \"http://wangzhongyang.com/images/logo_removebg.png\"}",
            "name": "转群治",
            "price": 99,
            "level1_category_id": 3,
            "level2_category_id": 49,
            "level3_category_id": 14,
            "brand": "veniam",
            "stock": 69284016,
            "sale": 61,
            "tags": "velit ut proident",
            "detail_info": "officia tempor fugiat culpa",
            "created_at": {
              "seconds": 1756970387
            },
            "updated_at": {
              "seconds": 1756970387
            }
          },
          {
            "id": 3,
            "pic_url": "",
            "images": "{\"image\": \"http://wangzhongyang.com/images/logo_removebg.png\"}",
            "name": "转群治",
            "price": 99,
            "level1_category_id": 3,
            "level2_category_id": 49,
            "level3_category_id": 14,
            "brand": "veniam",
            "stock": 69284016,
            "sale": 61,
            "tags": "velit ut proident",
            "detail_info": "officia tempor fugiat culpa",
            "created_at": {
              "seconds": 1756970388
            },
            "updated_at": {
              "seconds": 1756970388
            }
          },
          {
            "id": 4,
            "pic_url": "",
            "images": "{\"image\": \"http://wangzhongyang.com/images/logo_removebg.png\"}",
            "name": "转群治",
            "price": 99,
            "level1_category_id": 3,
            "level2_category_id": 49,
            "level3_category_id": 14,
            "brand": "veniam",
            "stock": 69284016,
            "sale": 61,
            "tags": "velit ut proident",
            "detail_info": "officia tempor fugiat culpa",
            "created_at": {
              "seconds": 1756970389
            },
            "updated_at": {
              "seconds": 1756970389
            }
          },
          {
            "id": 5,
            "pic_url": "",
            "images": "{\"image\": \"http://wangzhongyang.com/images/logo_removebg.png\"}",
            "name": "转群治",
            "price": 99,
            "level1_category_id": 3,
            "level2_category_id": 49,
            "level3_category_id": 14,
            "brand": "veniam",
            "stock": 69284016,
            "sale": 61,
            "tags": "velit ut proident",
            "detail_info": "officia tempor fugiat culpa",
            "created_at": {
              "seconds": 1756970390
            },
            "updated_at": {
              "seconds": 1756970390
            }
          },
          {
            "id": 6,
            "pic_url": "",
            "images": "{\"image\": \"http://wangzhongyang.com/images/logo_removebg.png\"}",
            "name": "转群治",
            "price": 99,
            "level1_category_id": 3,
            "level2_category_id": 49,
            "level3_category_id": 14,
            "brand": "veniam",
            "stock": 69284016,
            "sale": 61,
            "tags": "velit ut proident",
            "detail_info": "officia tempor fugiat culpa",
            "created_at": {
              "seconds": 1756970390
            },
            "updated_at": {
              "seconds": 1756970390
            }
          },
          {
            "id": 7,
            "pic_url": "",
            "images": "{\"image\": \"http://wangzhongyang.com/images/logo_removebg.png\"}",
            "name": "手机",
            "price": 99,
            "level1_category_id": 3,
            "level2_category_id": 49,
            "level3_category_id": 14,
            "brand": "veniam",
            "stock": 69284016,
            "sale": 61,
            "tags": "velit ut proident",
            "detail_info": "officia tempor fugiat culpa",
            "created_at": {
              "seconds": 1756970400
            },
            "updated_at": {
              "seconds": 1756970400
            }
          },
          {
            "id": 8,
            "pic_url": "",
            "images": "{\"image\": \"http://wangzhongyang.com/images/logo_removebg.png\"}",
            "name": "手机",
            "price": 99,
            "level1_category_id": 3,
            "level2_category_id": 49,
            "level3_category_id": 14,
            "brand": "veniam",
            "stock": 69284016,
            "sale": 61,
            "tags": "velit ut proident",
            "detail_info": "officia tempor fugiat culpa",
            "created_at": {
              "seconds": 1756970401
            },
            "updated_at": {
              "seconds": 1756970401
            }
          }
        ],
        "page": 1,
        "size": 10,
        "total": 8
      }
    }
  },

  '/goods/detail': {
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
  },

  // 修改密码
  '/frontend/user/update/password': {
    method: 'PUT',
    response: {
      code: 0,
      message: '密码修改成功',
      data: null
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
      code: 0,
      message: 'OK',
      data: {
        list: [
          {
            id: 1,
            number: 'ORD202412010001',
            user_id: 10,
            status: 2, // 1: 待付款, 2: 待发货, 3: 已发货, 4: 已完成, 5: 已取消
            price: 29900,
            actual_price: 26900,
            coupon_price: 3000,
            pay_type: 1,
            remark: '全栈开发课程',
            consignee_name: '张三',
            consignee_phone: '13800138000',
            consignee_address: '北京市朝阳区某某街道某某小区1号楼1单元101室',
            created_at: '2024-12-01 10:30:00',
            updated_at: '2024-12-01 10:30:00',
            pay_at: '2024-12-01 10:35:00'
          },
          {
            id: 2,
            number: 'ORD202412020002',
            user_id: 10,
            status: 1, // 待支付
            price: 9900,
            actual_price: 9900,
            coupon_price: 0,
            pay_type: 0,
            remark: 'xxsaf测试一下乱填',
            consignee_name: '',
            consignee_phone: '',
            consignee_address: '',
            created_at: '2024-12-02 15:45:00',
            updated_at: '2024-12-02 15:45:00',
            pay_at: null
          }
        ],
        total: 2,
        page: 1,
        size: 10
      }
    }
  },

  // 拼团砍价相关Mock数据
  '/frontend/group-buy': {
    method: 'GET',
    response: {
      code: 0,
      message: 'OK',
      data: {
        list: [
          {
            id: 101,
            name: 'GoFrame微服务实战课程',
            groupPrice: 9900, // 99元
            originalPrice: 19900, // 199元
            groupCount: 3,
            mainImage: 'http://wangzhongyang.com/images/goframe-course.jpg',
            participants: 128,
            endTime: '2025-09-15 23:59:59'
          },
          {
            id: 102,
            name: 'UniApp跨端开发实战',
            groupPrice: 7900, // 79元
            originalPrice: 14900, // 149元
            groupCount: 2,
            mainImage: 'http://wangzhongyang.com/images/uniapp-course.jpg',
            participants: 256,
            endTime: '2025-09-20 23:59:59'
          },
          {
            id: 103,
            name: '微信小程序高级开发',
            groupPrice: 12900, // 129元
            originalPrice: 25900, // 259元
            groupCount: 5,
            mainImage: 'http://wangzhongyang.com/images/wxapp-course.jpg',
            participants: 89,
            endTime: '2025-09-18 23:59:59'
          },
          {
            id: 104,
            name: '分布式系统架构设计',
            groupPrice: 19900, // 199元
            originalPrice: 39900, // 399元
            groupCount: 4,
            mainImage: 'http://wangzhongyang.com/images/distributed-course.jpg',
            participants: 67,
            endTime: '2025-09-25 23:59:59'
          },
          {
            id: 105,
            name: '云原生DevOps实战',
            groupPrice: 14900, // 149元
            originalPrice: 29900, // 299元
            groupCount: 3,
            mainImage: 'http://wangzhongyang.com/images/devops-course.jpg',
            participants: 182,
            endTime: '2025-09-22 23:59:59'
          }
        ],
        page: 1,
        size: 5,
        total: 15
      }
    }
  },

  // 砍价活动Mock数据
  '/frontend/bargain': {
    method: 'GET',
    response: {
      code: 0,
      message: 'OK',
      data: {
        list: [
          {
            id: 201,
            name: '全栈开发实战课程',
            bargainPrice: 4900, // 当前砍价后价格
            originalPrice: 9900, // 原价
            minPrice: 100, // 最低砍到1元
            mainImage: 'http://wangzhongyang.com/images/fullstack-course.jpg',
            bargainCount: 23, // 已砍次数
            totalBargainCount: 50, // 需要砍的总次数
            participants: 342
          }
        ]
      }
    }
  },

  // 订单详情Mock数据
  '/frontend/order/1': {
    method: 'GET',
    response: {
      code: 0,
      message: 'OK',
      data: {
        order_info: {
          id: 1,
          number: 'ORD202412010001',
          user_id: 10,
          status: 2,
          price: 29900, // 299元
          actual_price: 26900, // 269元（优惠后）
          coupon_price: 3000,
          pay_type: 1,
          remark: '全栈开发课程',
          consignee_name: '张三',
          consignee_phone: '13800138000',
          consignee_address: '北京市朝阳区某某街道某某小区1号楼1单元101室',
          created_at: '2024-12-01 10:30:00',
          updated_at: '2024-12-01 10:30:00',
          pay_at: '2024-12-01 10:35:00'
        },
        order_goods_infos: [
          {
            goods_id: 1,
            goods_options_id: 0,
            count: 1,
            remark: '',
            price: 9900, // 99元
            coupon_price: 0,
            actual_price: 9900,
            goods_name: '全栈开发实战课程',
            goods_pic_url: 'http://wangzhongyang.com/images/fullstack-course.jpg'
          },
          {
            goods_id: 2,
            goods_options_id: 0,
            count: 1,
            remark: '',
            price: 12900, // 129元
            coupon_price: 0,
            actual_price: 12900,
            goods_name: '微信小程序高级开发',
            goods_pic_url: 'http://wangzhongyang.com/images/wxapp-course.jpg'
          }
        ]
      }
    }
  },

  // 订单详情Mock数据 - 待支付状态
  '/frontend/order/2': {
    method: 'GET',
    response: {
      code: 0,
      message: 'OK',
      data: {
        order_info: {
          id: 2,
          number: 'ORD202412020002',
          user_id: 10,
          status: 1,
          price: 9900, // 99元
          actual_price: 9900,
          coupon_price: 0,
          pay_type: 0,
          remark: 'xxsaf测试一下乱填',
          consignee_name: '',
          consignee_phone: '',
          consignee_address: '',
          created_at: '2024-12-02 15:45:00',
          updated_at: '2024-12-02 15:45:00',
          pay_at: null
        },
        order_goods_infos: [
          {
            goods_id: 4,
            goods_options_id: 0,
            count: 1,
            remark: '',
            price: 0,
            coupon_price: 0,
            actual_price: 0,
            goods_name: '测试商品',
            goods_pic_url: 'http://wangzhongyang.com/images/test-product.jpg'
          }
        ]
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
      // 获取token
      const token = wx.getStorageSync('token') || '';
      
      // 设置请求头
      const header = {};
      if (token) {
        header['Authorization'] = `Bearer ${token}`;
      }
      
      // 真实网络请求
      wx.request({
        url: `${BASE_URL}${url}`,
        data: data,
        method: method,
        header: header,
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
  // 首页轮播图
  getBanners: () => request('/frontend/rotation', {}, 'GET'),

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
  submitOrder: (data) => request('/frontend/order', data, 'POST'), // 提交订单
  getOrders: (params) => request('/frontend/order', params, 'GET'),
  getOrderDetail: (id) => request(`/frontend/order/${id}`, {}, 'GET'),
  cancelOrder: (id) => request('/frontend/order/cancel', { id }, 'PUT'),

  // 拼团砍价相关
  getGroupBuyProducts: (params) => request('/frontend/group-buy', params, 'GET'),
  getBargainProducts: (params) => request('/frontend/bargain', params, 'GET'),
  createGroupOrder: (data) => request('/frontend/group-buy/order', data, 'POST'),
  createBargainOrder: (data) => request('/frontend/bargain/order', data, 'POST'),
  joinGroup: (data) => request('/frontend/group-buy/join', data, 'POST'),
  helpBargain: (data) => request('/frontend/bargain/help', data, 'POST'),
  
  // 用户优惠券
  getUserCoupons: (params) => request('/frontend/user_coupon', params, 'GET'),
  
  // 推荐商品
  getRecommendGoods: (params) => request('/frontend/recommend/goods', params, 'GET')
};

module.exports = {
  api,
  isMock,
  switchToRealAPI: () => { isMock = false; }
};