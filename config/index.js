// 统一配置文件 - 项目所有配置都在这里管理

// 环境配置
const isDev = false // 开发环境开关

// 基础配置
const BASE_CONFIG = {
  // 开发环境配置
  development: {
    BASE_URL: 'http://127.0.0.1:8199',
    UPLOAD_URL: 'http://127.0.0.1:8399',
    SEARCH_URL: 'http://127.0.0.1:8499',
    IMAGE_BASE_URL: 'http://127.0.0.1:8808/',
    SKIP_DOMAIN_CHECK: true
  },
  
  // 生产环境配置
  production: {
    BASE_URL: 'https://business.dayu.club',
    UPLOAD_URL: 'http://101.42.249.106:8399',
    SEARCH_URL: 'http://101.42.249.106:8499',
    IMAGE_BASE_URL: 'http://101.42.249.106:8808/',
    SKIP_DOMAIN_CHECK: false
  }
}

// 当前环境配置
const currentEnv = isDev ? 'development' : 'production'
const config = BASE_CONFIG[currentEnv]

// API接口路径配置
const API = {
  // 用户相关
  USER_REGISTER: `${config.BASE_URL}/frontend/user/register`,
  USER_LOGIN: `${config.BASE_URL}/frontend/user/login`,
  USER_INFO: `${config.BASE_URL}/frontend/user/info`,
  USER_WX_LOGIN: `${config.BASE_URL}/frontend/user/wxMiniLogin`,
  USER_WX_REGISTER: `${config.BASE_URL}/frontend/user/wxMiniRegister`,
  USER_BIND_PHONE: `${config.BASE_URL}/frontend/user/bindPhone`,
  USER_FILL_PHONE: `${config.BASE_URL}/frontend/user/fillPhone`,
  
  // 商品相关
  PRODUCT_LIST: `${config.BASE_URL}/goods`,
  PRODUCT_DETAIL: `${config.BASE_URL}/goods/detail`,
  
  // 搜索相关   
  SEARCH_GOODS: `${config.SEARCH_URL}/search/goods/mysql`,
  
  // 订单相关
  ORDER_LIST: `${config.BASE_URL}/frontend/order/list`,
  ORDER_CREATE: `${config.BASE_URL}/frontend/order/create`,
  ORDER_DETAIL: `${config.BASE_URL}/frontend/order`,
  ORDER_CANCEL: `${config.BASE_URL}/frontend/order/cancel`,
  ORDER_PAY: `${config.BASE_URL}/frontend/order/pay`,
  ORDER_CONFIRM_RECEIVE: `${config.BASE_URL}/frontend/order/confirm`,
  ORDER_COUNT: `${config.BASE_URL}/frontend/order/count`,
  
  // 购物车相关
  CART_LIST: `${config.BASE_URL}/frontend/cart`,
  CART_ADD: `${config.BASE_URL}/frontend/cart/add`,
  CART_UPDATE: `${config.BASE_URL}/frontend/cart/update`,
  CART_DELETE: `${config.BASE_URL}/frontend/cart`,
  
  // 文件上传
  UPLOAD_IMAGE: `${config.UPLOAD_URL}/upload/image`,
  Key_URL: `${config.UPLOAD_URL}/image/url`,
  
  // 轮播图
  BANNERS: `${config.BASE_URL}/frontend/rotation`
}

// 公共常量配置
const CONSTANTS = {
  // 图片域名配置
  IMAGE_BASE_URL: config.IMAGE_BASE_URL,
  
  // 分页配置
  PAGINATION: {
    DEFAULT_PAGE: 1,
    DEFAULT_SIZE: 10
  },
  
  // 商品相关配置
  GOODS: {
    HOT_RECOMMEND_SIZE: 10 // 热门推荐商品数量
  }
}

module.exports = {
  isDev,
  config,
  API,
  CONSTANTS
}