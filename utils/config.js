// 服务器域名配置
const BASE_URL = 'http://shop.dayu.club:8199'

// API接口路径
const API = {
  // 用户相关
  USER_REGISTER: `${BASE_URL}/frontend/user/register`,
  USER_LOGIN: `${BASE_URL}/frontend/user/login`,
  USER_INFO: `${BASE_URL}/frontend/user/info`,
  
  // 商品相关
  PRODUCT_LIST: `${BASE_URL}/frontend/product/list`,
  PRODUCT_DETAIL: `${BASE_URL}/frontend/product/detail`,
  
  // 订单相关
  ORDER_LIST: `${BASE_URL}/frontend/order/list`,
  ORDER_CREATE: `${BASE_URL}/frontend/order/create`,
  
  // 购物车相关
  CART_LIST: `${BASE_URL}/frontend/cart/list`,
  CART_ADD: `${BASE_URL}/frontend/cart/add`,
  CART_UPDATE: `${BASE_URL}/frontend/cart/update`,
  CART_DELETE: `${BASE_URL}/frontend/cart/delete`
}

// 导出配置
module.exports = {
  BASE_URL,
  API
}