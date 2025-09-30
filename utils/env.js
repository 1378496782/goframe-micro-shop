// 环境配置
const isDev = false // 开发环境

// 开发环境配置 - 使用HTTP并关闭域名校验
const DEV_CONFIG = {
  BASE_URL: 'http://127.0.0.1:8199',
  UPLOAD_URL: 'http://127.0.0.1:8399',
  SEARCH_URL: 'http://127.0.0.1:8499',
  IMAGE_BASE_URL: 'http://127.0.0.1:8808/',
  SKIP_DOMAIN_CHECK: true
} 
 
// 生产环境配置 - 使用HTTP
const PROD_CONFIG = {
  BASE_URL: 'https://business.dayu.club',
  UPLOAD_URL: 'http://101.42.249.106:8399',
  // UPLOAD_URL: 'http://127.0.0.1:8399',
  SEARCH_URL: 'http://101.42.249.106:8499',
  IMAGE_BASE_URL: 'http://101.42.249.106:8808/',
  SKIP_DOMAIN_CHECK: false
}

const config = isDev ? DEV_CONFIG : PROD_CONFIG

// API接口路径
const API = {
  // 用户相关
  USER_REGISTER: `${config.BASE_URL}/frontend/user/register`,
  USER_LOGIN: `${config.BASE_URL}/frontend/user/login`,
  USER_INFO: `${config.BASE_URL}/frontend/user/info`,
  USER_WX_LOGIN: `${config.BASE_URL}/frontend/user/wxMiniAuth`, // 微信小程序登录
  USER_BIND_PHONE: `${config.BASE_URL}/frontend/user/bindPhone`, // 绑定手机号
  
  // 商品相关
  PRODUCT_LIST: `${config.BASE_URL}/goods`,
  PRODUCT_DETAIL: `${config.BASE_URL}/goods/detail`,
  
  // 搜索相关   
  SEARCH_GOODS: `${config.SEARCH_URL}/search/goods`,
  
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
  UPLOAD_IMAGE: `${config.UPLOAD_URL}/upload/image` 
}

module.exports = {
  isDev,
  config,
  API
}