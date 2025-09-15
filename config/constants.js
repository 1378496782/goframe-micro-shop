// 公共常量配置
const { config } = require('../utils/env');

const constants = {
  // 图片域名配置 - 使用统一配置
  IMAGE_BASE_URL: 'http://101.42.249.106:8808/', // 保持原有图片服务器配置
  
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