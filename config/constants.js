// 公共常量配置
const { config, isDev } = require('../utils/env');

const constants = {
  // 图片域名配置 
  IMAGE_BASE_URL: isDev 
    ? 'http://127.0.0.1:8808/' 
    : 'https://shopadmin.dayu.club/',
  
  // 分页配置
  PAGINATION: {
    DEFAULT_PAGE: 1,
    DEFAULT_SIZE: 10
  },
  
  // 商品相关配置
  GOODS: {
    HOT_RECOMMEND_SIZE: 10 // 热门推荐商品数量
  }
};

module.exports = constants;