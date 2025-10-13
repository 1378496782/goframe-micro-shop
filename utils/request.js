// 统一的请求封装，包含token管理和错误处理
const { API, isDev } = require('../config/index')

// 获取token
function getToken() {
  return wx.getStorageSync('token') || ''
}

// 获取用户信息
function getUserInfo() {
  return wx.getStorageSync('userInfo') || {}
}

// 保存用户信息
function saveUserInfo(userInfo) {
  wx.setStorageSync('userInfo', userInfo)
}

// 保存token
function saveToken(token) {
  wx.setStorageSync('token', token)
}

// 清除登录状态
function clearAuth() {
  wx.removeStorageSync('token')
  wx.removeStorageSync('userInfo')
}

// 统一的请求方法
function request(options) {
  return new Promise((resolve, reject) => {
    const { url, data = {}, method = 'GET', needAuth = true } = options
    
    console.log(`[API Request] ${method} ${url}`, data)

    // 添加token到header
    const header = {}
    if (needAuth) {
      const token = getToken()
      if (token) {
        header['Authorization'] = `Bearer ${token}`
      }
    }

    wx.request({
      url,
      data,
      method,
      header,
      success: (res) => {
        console.log(`[API Response] ${method} ${url}`, res)
        
        // 统一处理响应格式
        if (res.statusCode === 200) {
          const response = res.data
          
          // 处理标准响应格式 {code, message, data}
          if (response.code === 0) {
            resolve(response.data)
          } else {
            // 业务错误
            const errorMsg = response.message || '操作失败'
            
            // 特殊处理：用户信息接口返回用户不存在错误时不显示错误提示
            const isUserInfoApi = url.includes('/frontend/user/info')
            const isUserNotExistError = response.code === 52 || errorMsg.includes('用户不存在')
            
            if (!(isUserInfoApi && isUserNotExistError)) {
              // 只在非用户信息获取失败时显示错误提示
              if (isDev) {
                wx.showModal({
                  title: '请求错误',
                  content: errorMsg,
                  showCancel: false
                })
              } else {
                wx.showToast({
                  title: errorMsg,
                  icon: 'none'
                })
              }
            }
            reject(new Error(errorMsg))
          }
        } else {
          // HTTP错误
          const errorMsg = `HTTP错误: ${res.statusCode}`
          console.error(errorMsg, res)
          if (isDev) {
            wx.showModal({
              title: 'HTTP错误',
              content: errorMsg,
              showCancel: false
            })
          }
          reject(new Error(errorMsg))
        }
      },
      fail: (err) => {
        // 网络错误
        console.error('[API Error] 网络请求失败:', err)
        if (isDev) {
          wx.showModal({
            title: '网络错误',
            content: `请检查:
1. 网络连接
2. 域名配置
3. 服务器状态

错误详情: ${err.errMsg}`,
            showCancel: false
          })
        } else {
          wx.showToast({
            title: '网络错误，请重试',
            icon: 'none'
          })
        }
        reject(err)
      }
    })
  })
}

// 用户相关的API封装
const userAPI = {
  // 登录
  login(credentials) {
    return request({
      url: API.USER_LOGIN,
      data: credentials,
      method: 'POST',
      needAuth: false
    }).then(data => {
      // 保存token和用户信息
      saveToken(data.token)
      saveUserInfo(data.userInfo || data.user_info || {})
      return data
    })
  },

  // 注册
  register(userData) {
    return request({
      url: API.USER_REGISTER,
      data: userData,
      method: 'POST',
      needAuth: false
    })
  },

  // 获取用户信息
  getInfo() {
    return request({
      url: API.USER_INFO,
      method: 'GET'
    }).then(userInfo => {
      saveUserInfo(userInfo)
      return userInfo
    })
  },

  // 更新用户信息
  updateInfo(userData) {
    return request({
      url: API.USER_INFO,
      data: userData,
      method: 'PUT'
    }).then(userInfo => {
      saveUserInfo(userInfo)
      return userInfo
    })
  },

  // 填写手机号
  fillPhone(phoneData) {
    return request({
      url: API.USER_FILL_PHONE,
      data: phoneData,
      method: 'POST'
    })
  }
}

// 商品搜索相关的API封装
const searchAPI = {
  // 商品搜索
  searchGoods(params) {
    return request({
      url: API.SEARCH_GOODS,
      data: params,
      method: 'GET',
      needAuth: false
    })
  }
}

// 检查登录状态
function checkLoginStatus() {
  const token = getToken()
  const userInfo = getUserInfo()
  return {
    isLoggedIn: !!token,
    userInfo: userInfo
  }
}

module.exports = {
  request,
  userAPI,
  searchAPI,
  getToken,
  getUserInfo,
  saveUserInfo,
  saveToken,
  clearAuth,
  checkLoginStatus
}