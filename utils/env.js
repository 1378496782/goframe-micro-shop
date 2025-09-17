// 环境配置
const isDev = true // 开发环境

// 开发环境配置 - 使用HTTP并关闭域名校验
const DEV_CONFIG = {
  BASE_URL: 'http://192.168.1.5:8199',
  SKIP_DOMAIN_CHECK: true
} 
 
// 生产环境配置 - 使用HTTPS
const PROD_CONFIG = {
  BASE_URL: 'https://shop.dayu.club',
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
  
  // 商品相关
  PRODUCT_LIST: `${config.BASE_URL}/frontend/product/list`,
  PRODUCT_DETAIL: `${config.BASE_URL}/frontend/product/detail`,
  
  // 搜索相关   
  SEARCH_GOODS: `http://101.42.249.106:8499/search/goods`,
  
  // 订单相关
  ORDER_LIST: `${config.BASE_URL}/frontend/order/list`,
  ORDER_CREATE: `${config.BASE_URL}/frontend/order/create`,
  
  // 购物车相关
  CART_LIST: `${config.BASE_URL}/frontend/cart/list`,
  CART_ADD: `${config.BASE_URL}/frontend/cart/add`,
  CART_UPDATE: `${config.BASE_URL}/frontend/cart/update`,
  CART_DELETE: `${config.BASE_URL}/frontend/cart/delete`,
  
  // 文件上传
  UPLOAD_IMAGE: `http://192.168.1.5:8399/upload/image` // 图片上传，使用端口8399
}

module.exports = {
  isDev,
  config,
  API
}