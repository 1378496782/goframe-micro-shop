// 环境配置
const config = {
  // 当前环境
  env: 'development', // development: 开发环境, production: 生产环境, test: 测试环境
  
  // API配置
  api: {
    baseURL: 'https://api.example.com', // 生产环境API地址
    mockBaseURL: '', // Mock模式不需要baseURL
    timeout: 10000,
    retryCount: 3,    // 请求重试次数
    retryDelay: 1000  // 重试延迟(ms)
  },
  
  // 功能开关
  features: {
    useMock: true,     // 是否使用Mock数据
    debug: true,       // 是否开启调试模式
    logger: true,      // 是否开启日志
    errorSimulation: true, // 是否开启错误模拟
    cacheEnabled: true // 是否开启缓存
  },
  
  // Mock配置
  mock: {
    networkDelay: 500,   // 网络延迟(ms)
    errorRate: 0.1,      // 错误率(0-1)
    cacheDuration: 5 * 60 * 1000 // 缓存持续时间(ms)
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