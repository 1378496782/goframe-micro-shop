// 环境配置
const config = {
  // 当前环境
  env: 'development', // development: 开发环境, production: 生产环境
  
  // API配置
  api: {
    baseURL: 'https://api.example.com', // 生产环境API地址
    mockBaseURL: '', // Mock模式不需要baseURL
    timeout: 10000
  },
  
  // 功能开关
  features: {
    useMock: true, // 是否使用Mock数据
    debug: true,   // 是否开启调试模式
    logger: true   // 是否开启日志
  },
  
  // 版本信息
  version: '1.0.0',
  buildTime: '2024-01-01'
};

// 开发环境配置
if (config.env === 'development') {
  config.features.useMock = true;
  config.features.debug = true;
}

// 生产环境配置
if (config.env === 'production') {
  config.features.useMock = false;
  config.features.debug = false;
  config.features.logger = false;
}

module.exports = config;