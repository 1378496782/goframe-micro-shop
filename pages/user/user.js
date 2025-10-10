/**
 * 用户中心页面
 * 功能：处理用户登录、个人信息展示、头像上传等功能
 */
const app = getApp()
const { checkLoginStatus, request } = require('../../utils/request')
const { API } = require('../../utils/env')

Page({
  /**
   * 页面数据对象
   * isLoggedIn: 登录状态标识
   * userInfo: 用户信息对象
   * orderCounts: 订单状态统计
   * tempAvatar: 临时头像路径
   * uploadedAvatarUrl: 上传后的头像URL
   */
  data: {
    isLoggedIn: false,
    userInfo: {},
    orderCounts: {
      pending: 0,
      shipping: 0,
      delivered: 0,
      completed: 0,
      afterSale: 0
    },
    hasShownLoginTip: false,
    showUserInfoForm: false,
    tempAvatar: '',
    uploadedAvatarUrl: '',
    tempNickname: '',
    wxLoginCode: '',
    wxIv: '',
    wxEncryptedData: '',
    showPhoneAuthPopup: false,
    phoneAuthData: {
      code: '',
      iv: '',
      encryptedData: ''
    }
  },

  /**
   * 监听orderCounts数据变化
   */
  onOrderCountsChange(newOrderCounts) {
    console.log('orderCounts数据发生变化:', newOrderCounts)
  },

  /**
   * 页面加载生命周期
   * 初始化时检查登录状态
   */
  onLoad() {
    this.checkLoginStatus()
  },

  /**
   * 页面显示生命周期
   * 每次页面显示时检查登录状态
   * 已登录则获取用户信息，未登录显示引导提示
   */
  onShow() {
    console.log('user页面 onShow 开始执行')
    this.checkLoginStatus()
    console.log('登录状态:', this.data.isLoggedIn)
    if (this.data.isLoggedIn) {
      console.log('已登录，开始获取用户信息和订单数量')
      this.getUserInfo()
      // 获取订单数量统计
      this.getOrderCounts()
    }
  },

  /**
   * 数据监听器
   */
  observers: {
    'orderCounts': function(orderCounts) {
      console.log('orderCounts数据变化观察者 - 新值:', orderCounts)
      this.onOrderCountsChange(orderCounts)
    }
  },

  /**
   * 微信登录方法 - 立即显示用户信息填写表单，异步获取授权参数
   */
  autoWxLogin() {
    // 立即显示用户信息填写表单，让用户可以开始操作
    this.setData({ showUserInfoForm: true })
    
    
    // 并行获取微信授权参数
    Promise.all([
      new Promise((resolve, reject) => {
        wx.getUserProfile({
          desc: '用于完善会员资料',
          success: (res) => resolve(res),
          fail: (err) => reject(err)
        })
      }),
      new Promise((resolve, reject) => {
        wx.login({
          success: (res) => resolve(res),
          fail: (err) => reject(err)
        })
      })
    ]).then(([userRes, loginRes]) => {
      wx.hideLoading()
      
      if (!loginRes.code) {
        return wx.showToast({ title: '获取code失败', icon: 'none' })
      }
      
      // 保存微信授权信息
      this.setData({
        wxLoginCode: loginRes.code,
        wxIv: userRes.iv,
        wxEncryptedData: userRes.encryptedData
      })
      
      console.log('微信授权参数获取完成')
    }).catch((err) => {
      wx.hideLoading()
      console.error('获取授权信息失败:', err)
      // 不显示错误提示，让用户可以继续操作
    })
  },
  
  /**
   * 触发手机号授权弹窗
   */
  // triggerPhoneAuth() {
  //   if (!this.data.isLoggedIn) {
  //     return wx.showToast({ title: '请先登录', icon: 'none' })
  //   }
    
  //   this.setData({ showPhoneAuthPopup: true })
  // },
  
  /**
   * 关闭手机号授权弹窗
   */
  closePhoneAuthPopup() {
    // 用户点击关闭按钮，立即退出弹窗页面，重新进入我的页面
    this.setData({
      showPhoneAuthPopup: false,
      wxLoginTempData: null
    })
    
    wx.showToast({ title: '已取消授权', icon: 'none' })
    
    // 立即重新启动到我的页面
    setTimeout(() => {
      wx.reLaunch({
        url: '/pages/user/user'
      })
<<<<<<< HEAD
    }, 200)
=======
    }, 100)
>>>>>>> 8925dc9075b97bd953a7803cde326b1e1c6bc549
  },
  
  /**
   * 获取手机号授权回调
   * @param {object} e - 事件对象，包含手机号加密数据
   */
  onGetPhoneNumber(e) {
    console.log('手机号授权回调:', e)
    
    // 判断是哪种场景的手机号授权
    const { wxLoginTempData } = this.data
<<<<<<< HEAD
    
    if (wxLoginTempData) {
      // 微信登录后需要手机号授权的场景
      this.handleWxLoginPhoneAuth(e)
    } else {
      // 用户信息填写场景的手机号授权
      this.handleUserInfoPhoneAuth(e)
    }
  },
  
  /**
   * 处理微信登录后的手机号授权
   * @param {object} e - 事件对象，包含手机号加密数据
   */
  handleWxLoginPhoneAuth(e) {
    console.log('处理微信登录后的手机号授权:', e)
    
    if (e.detail.errMsg === 'getPhoneNumber:ok') {
      // 用户同意授权，调用填写手机号接口
      this.fillPhoneNumber(e)
    } else if (e.detail.errMsg === 'getPhoneNumber:fail user deny') {
      // 用户拒绝授权，立即退出弹窗页面，重新进入我的页面
      this.setData({
        showPhoneAuthPopup: false,
        wxLoginTempData: null
      })
      
      wx.showToast({ title: '已取消授权', icon: 'none' })
      
      // 立即重新启动到我的页面
      setTimeout(() => {
        wx.reLaunch({
          url: '/pages/user/user'
        })
      }, 500)
    } else {
      // 其他错误
      console.error('手机号授权失败:', e.detail)
      wx.showToast({ title: '授权失败，请重试', icon: 'none' })
    }
  },
  
  /**
   * 处理用户信息填写场景的手机号授权
   * @param {object} e - 事件对象，包含手机号加密数据
   */
  handleUserInfoPhoneAuth(e) {
    console.log('处理用户信息填写场景的手机号授权:', e)
    
    // 先验证表单数据
    const { tempNickname, uploadedAvatarUrl, tempAvatar, wxLoginCode } = this.data
=======
>>>>>>> 8925dc9075b97bd953a7803cde326b1e1c6bc549
    
    if (wxLoginTempData) {
      // 微信登录后需要手机号授权的场景
      this.handleWxLoginPhoneAuth(e)
    } else {
      // 用户信息填写场景的手机号授权
      this.handleUserInfoPhoneAuth(e)
    }
  },
  
  /**
   * 处理微信登录后的手机号授权
   * @param {object} e - 事件对象，包含手机号加密数据
   */
  handleWxLoginPhoneAuth(e) {
    console.log('处理微信登录后的手机号授权:', e)
    
    if (e.detail.errMsg === 'getPhoneNumber:ok') {
      // 用户同意授权，调用填写手机号接口
      this.fillPhoneNumber(e)
    } else if (e.detail.errMsg === 'getPhoneNumber:fail user deny') {
      // 用户拒绝授权，立即退出弹窗页面，重新进入我的页面
      this.setData({
        showPhoneAuthPopup: false,
        wxLoginTempData: null
      })
      
      wx.showToast({ title: '已取消授权', icon: 'none' })
      
      // 立即重新启动到我的页面
      setTimeout(() => {
        wx.reLaunch({
          url: '/pages/user/user'
        })
      }, 500)
    } else {
      // 其他错误
      console.error('手机号授权失败:', e.detail)
      wx.showToast({ title: '授权失败，请重试', icon: 'none' })
    }
  },
  
  /**
   * 处理用户信息填写场景的手机号授权
   * @param {object} e - 事件对象，包含手机号加密数据
   */
  handleUserInfoPhoneAuth(e) {
    console.log('处理用户信息填写场景的手机号授权:', e)
    
    // 先验证表单数据
    const { tempNickname, uploadedAvatarUrl, tempAvatar, wxLoginCode } = this.data    
    if (wxLoginTempData) {
      // 微信登录后需要手机号授权的场景
      this.handleWxLoginPhoneAuth(e)
    } else {
      // 用户信息填写场景的手机号授权
      this.handleUserInfoPhoneAuth(e)
    }
  },


  
  /**
   * 处理登录成功逻辑
   * @param {object} loginData - 登录数据
   */
  handleLoginSuccess(loginData) {
    // 保存token到本地存储
    if (loginData.token) {
      wx.setStorageSync('token', loginData.token)
      console.log('已保存token:', loginData.token)
    }
    
    // 保存用户信息到本地存储
    if (loginData.user_info) {
      wx.setStorageSync('userInfo', loginData.user_info)
      console.log('已保存userInfo:', loginData.user_info)
    }
    
    // 保存openId，用于微信支付
    if (loginData.openId) {
      wx.setStorageSync('openId', loginData.openId)
      console.log('已保存openId:', loginData.openId)
    }
    
    // 保存token过期时间
    if (loginData.expire_in) {
      wx.setStorageSync('token_expire', Date.now() + loginData.expire_in * 1000)
      console.log('token过期时间:', new Date(Date.now() + loginData.expire_in * 1000))
    }
    
    // 更新全局状态
    app.globalData.isLoggedIn = true
    app.globalData.userInfo = loginData.user_info || {}
    app.globalData.token = loginData.token
    app.globalData.openId = loginData.openId
    
    // 更新页面数据
    this.setData({ 
      isLoggedIn: true, 
      userInfo: loginData.user_info || {},
      showUserInfoForm: false,
      tempAvatar: '',
      uploadedAvatarUrl: '',
      tempNickname: '',
      wxLoginCode: '',
      wxIv: '',
      wxEncryptedData: '',
      showPhoneAuthPopup: false,
      wxLoginTempData: null
    })
    
    wx.showToast({ title: '登录成功', icon: 'success' })
    
    // 登录成功后获取最新的用户信息和订单统计
    setTimeout(() => {
      this.getUserInfo()
      this.getOrderCounts()
    }, 500)
  },
  
  /**
   * 处理登录成功逻辑
   * @param {object} loginData - 登录数据
   */
  handleLoginSuccess(loginData) {
    // 保存token到本地存储
    if (loginData.token) {
      wx.setStorageSync('token', loginData.token)
      console.log('已保存token:', loginData.token)
    }
    
    // 保存用户信息到本地存储
    if (loginData.user_info) {
      wx.setStorageSync('userInfo', loginData.user_info)
      console.log('已保存userInfo:', loginData.user_info)
    }
    
    // 保存openId，用于微信支付
    if (loginData.openId) {
      wx.setStorageSync('openId', loginData.openId)
      console.log('已保存openId:', loginData.openId)
    }
    
    // 保存token过期时间
    if (loginData.expire_in) {
      wx.setStorageSync('token_expire', Date.now() + loginData.expire_in * 1000)
      console.log('token过期时间:', new Date(Date.now() + loginData.expire_in * 1000))
    }
    
    // 更新全局状态
    app.globalData.isLoggedIn = true
    app.globalData.userInfo = loginData.user_info || {}
    app.globalData.token = loginData.token
    app.globalData.openId = loginData.openId
    
    // 更新页面数据
    this.setData({ 
      isLoggedIn: true, 
      userInfo: loginData.user_info || {},
      showUserInfoForm: false,
      tempAvatar: '',
      uploadedAvatarUrl: '',
      tempNickname: '',
      wxLoginCode: '',
      wxIv: '',
      wxEncryptedData: '',
      showPhoneAuthPopup: false,
      wxLoginTempData: null
    })
    
    wx.showToast({ title: '登录成功', icon: 'success' })
    
    // 登录成功后获取最新的用户信息和订单统计
    setTimeout(() => {
      this.getUserInfo()
      this.getOrderCounts()
    }, 500)
  },
  
  /**
   * 使用手机号完成登录
   * @param {string} iv - 加密算法的初始向量
   * @param {string} encryptedData - 加密数据
   */
  completeLoginWithPhone(iv, encryptedData) {
    const { tempNickname, uploadedAvatarUrl, tempAvatar, wxLoginCode } = this.data
    
    // 如果有上传成功的头像URL，直接使用
    // 如果只有临时头像路径但没有上传，先上传头像再登录
    if (tempAvatar && !uploadedAvatarUrl) {
      wx.showLoading({ title: '正在上传头像...' })
      wx.uploadFile({
        url: API.UPLOAD_IMAGE,
        filePath: tempAvatar,
        name: 'File',
        formData: { 
          uploader_id: 0, 
          uploader_type: 1, // 1-H5用户
          file_type: 1 // 1-图片
        },
        success: (res) => {
          wx.hideLoading()
          const data = JSON.parse(res.data)
          if (data.code === 0 && data.data?.url) {
            // 上传成功后，使用返回的URL进行登录
            const wxLoginData = {
              code: wxLoginCode,
              iv: iv,
              encryptedData: encryptedData,
              phoneNumber: '', // 后端会从加密数据中解密手机号
              nickname: tempNickname,
              avatar: data.data.url
            }
            this.wxMiniLogin(wxLoginData)
          } else {
            wx.showToast({ 
              title: data.message || '头像上传失败', 
              icon: 'none' 
            })
          }
        },
        fail: (err) => {
          wx.hideLoading()
          console.error('头像上传失败:', err)
          wx.showToast({ title: '头像上传失败', icon: 'none' })
        }
      })
    } else {
      // 已有上传成功的头像URL，直接登录
      const avatar = uploadedAvatarUrl || 'https://via.placeholder.com/100x100/19aecc/ffffff?text=用户'
      const wxLoginData = {
        code: wxLoginCode,
        iv: iv,
        encryptedData: encryptedData,
        phoneNumber: '', // 后端会从加密数据中解密手机号
        nickname: tempNickname,
        avatar: avatar
      }
      this.wxMiniLogin(wxLoginData)
    }
  },
  
  /**
   * 填写手机号接口调用
   * @param {object} e - 事件对象，包含微信手机号授权返回的数据
   */
  fillPhoneNumber(e) {
    wx.showLoading({ title: '授权中...' })
    
    const { wxLoginTempData } = this.data
    if (!wxLoginTempData || !wxLoginTempData.token) {
      wx.hideLoading()
      return wx.showToast({ title: '登录信息已过期，请重新登录', icon: 'none' })
    }
    
    // 设置token用于接口调用
    wx.setStorageSync('token', wxLoginTempData.token)
    
    // 调试：查看微信返回的完整数据结构
    console.log('微信手机号授权返回的完整数据:', JSON.stringify(e.detail, null, 2))
    
    // 微信手机号授权返回的数据
    const { iv, encryptedData } = e.detail
    
    console.log('获取到的参数值:', {
      iv: iv,
      encryptedData: encryptedData
    })
    
    if (!iv || !encryptedData) {
      wx.hideLoading()
      console.error('授权数据不完整:', { iv, encryptedData })
      return wx.showToast({ title: '授权数据不完整', icon: 'none' })
    }
    
    // 重新获取有效的微信登录code
    wx.login({
      success: (loginRes) => {
        if (!loginRes.code) {
          wx.hideLoading()
          return wx.showToast({ title: '获取登录凭证失败', icon: 'none' })
        }
        
        console.log('重新获取的微信登录code:', loginRes.code)
        
        // 调用填写手机号接口，传递新获取的code以及iv、encryptedData
        request({
          url: API.USER_FILL_PHONE,
          method: 'POST',
          data: {
            code: loginRes.code,
            iv: iv,
            encryptedData: encryptedData
          }
        }).then(res => {
          wx.hideLoading()
          
          console.log('fillPhone接口完整响应:', JSON.stringify(res, null, 2))
          console.log('res.code:', res.code)
          console.log('res.data:', res.data)
          
          // 检查响应结构，支持多种可能的格式
          let success = false
          let message = ''
          
          if (res && typeof res === 'object') {
            // 格式1: {code: 0, message: "OK", data: {...}}
            if (res.code === 0) {
              success = true
              message = res.message || '授权成功'
            }
            // 格式2: {data: {code: 0, message: "OK", data: {...}}}
            else if (res.data && res.data.code === 0) {
              success = true
              message = res.data.message || '授权成功'
            }
            // 格式3: 直接返回成功数据
            else if (res.id) {
              success = true
              message = '授权成功'
            }
          }
          
          if (success) {
            console.log('手机号授权成功，返回数据:', res)
            
            // 填写手机号成功，完成登录流程
            // 使用wxLoginTempData中的用户信息完成登录
            this.handleLoginSuccess(wxLoginTempData)
            
            // 关闭授权弹窗
            this.setData({
              showPhoneAuthPopup: false,
              wxLoginTempData: null
            })
            
            wx.showToast({ title: message, icon: 'success' })
          } else {
            const errorMsg = res.message || res.data?.message || '授权失败'
            console.error('授权失败:', errorMsg)
            wx.showToast({ title: errorMsg, icon: 'none' })
          }
          
        }).catch(err => {
          wx.hideLoading()
          console.error('填写手机号失败:', err)
          wx.showToast({ title: '授权失败，请重试', icon: 'none' })
        })
      },
      fail: (err) => {
        wx.hideLoading()
        console.error('重新获取微信登录code失败:', err)
        wx.showToast({ title: '获取登录凭证失败', icon: 'none' })
      }
    })
  },
  
  /**
   * 从微信授权返回的数据中获取手机号
   * @param {object} detail - 微信getPhoneNumber返回的detail对象
   * @returns {string} 手机号
   */
  getPhoneNumberFromWxResponse(detail) {
    console.log('微信返回的完整数据结构:', JSON.stringify(detail, null, 2))
    
    // 微信手机号授权返回的数据格式可能有多种情况
    // 1. 直接包含手机号信息
    // 2. 只返回code，需要调用后端接口解密
    // 3. 返回加密数据，需要解密
    
    let phoneNumber = ''
    
    // 情况1: 直接包含手机号信息
    if (detail.phoneNumber) {
      phoneNumber = detail.phoneNumber
      console.log('从phoneNumber字段获取手机号:', phoneNumber)
    } else if (detail.purePhoneNumber) {
      phoneNumber = detail.purePhoneNumber
      console.log('从purePhoneNumber字段获取手机号:', phoneNumber)
    } else if (detail.data && detail.data.phoneNumber) {
      phoneNumber = detail.data.phoneNumber
      console.log('从data.phoneNumber字段获取手机号:', phoneNumber)
    } 
    // 情况2: 只有code，需要调用后端接口解密
    else if (detail.code) {
      console.log('微信返回的是code，需要调用后端接口解密:', detail.code)
      // 这里需要调用后端接口解密手机号
      // 暂时返回占位符，实际应该调用后端接口
      phoneNumber = '13812345678'
      console.log('使用临时手机号:', phoneNumber)
    }
    // 情况3: 返回加密数据，需要解密
    else if (detail.iv && detail.encryptedData) {
      console.log('微信返回的是加密数据，需要解密:', detail)
      // 这里需要调用后端接口解密手机号
      // 暂时返回占位符，实际应该调用后端接口
      phoneNumber = '13812345678'
      console.log('使用临时手机号:', phoneNumber)
    }
    else {
      console.error('无法从微信返回数据中获取手机号，完整数据结构:', JSON.stringify(detail, null, 2))
      return null
    }
    
    console.log('最终获取到的手机号:', phoneNumber)
    return phoneNumber
  },
  
  /**
   * 处理用户拒绝授权
   */
  handleAuthReject() {
    // 清空用户登录信息
    this.clearLoginStatus()
    
    // 关闭授权弹窗
    this.setData({
      showPhoneAuthPopup: false,
      wxLoginTempData: null
    })
    
    wx.showToast({ title: '已取消授权', icon: 'none' })
    app.globalData.userInfo = {}
    app.globalData.token = ''
    app.globalData.openId = ''
    
    // 更新页面数据
    this.setData({
      isLoggedIn: false,
      userInfo: {},
      showPhoneAuthPopup: false,
      wxLoginTempData: null
    })
    
    wx.showToast({ title: '已取消授权', icon: 'none' })
  },
  
  /**
   * 绑定手机号到用户账户
   */
  bindPhoneNumber() {
    const { phoneAuthData } = this.data
    
    if (!phoneAuthData.code || !phoneAuthData.iv || !phoneAuthData.encryptedData) {
      wx.showToast({ title: '授权数据不完整', icon: 'none' })
      return
    }
    
    wx.showLoading({ title: '绑定中...' })
    
    const bindData = {
      code: phoneAuthData.code,
      iv: phoneAuthData.iv,
      encryptedData: phoneAuthData.encryptedData
    }
    
    request({
      url: API.USER_BIND_PHONE,
      method: 'POST',
      data: bindData
    }).then(res => {
      wx.hideLoading()
      
      if (res.code === 0) {
        wx.showToast({ title: '手机号绑定成功', icon: 'success' })
        this.closePhoneAuthPopup()
        
        // 刷新用户信息
        this.getUserInfo()
        
        // 重置手机号授权数据
        this.setData({
          phoneAuthData: {
            code: '',
            iv: '',
            encryptedData: ''
          }
        })
      } else {
        wx.showToast({ title: res.message || '绑定失败', icon: 'none' })
      }
    }).catch(err => {
      wx.hideLoading()
      console.error('绑定手机号失败:', err)
      wx.showToast({ title: '绑定失败，请重试', icon: 'none' })
    })
  },
  
  /**
   * 关闭用户信息填写表单
   */
  closeUserInfoForm() {
    this.setData({ showUserInfoForm: false })
  },
  
  /**
   * 用户选择头像回调
   * 微信基础库 2.21.2 及以上版本支持
   */
  onChooseAvatar(e) {
    const { avatarUrl } = e.detail
    this.setData({
      tempAvatar: avatarUrl
    })
    // 上传头像
    this.uploadImage(avatarUrl)
  },
  
  /**
   * 昵称输入处理
   */
  onNicknameInput(e) { 
    this.setData({ tempNickname: e.detail.value }) 
  },
  
  /**
   * 提交用户信息验证（已废弃，验证逻辑移到 onGetPhoneNumber 中）
   */
  submitUserInfo() {
    // 这个方法现在不需要了，因为验证逻辑在 onGetPhoneNumber 中处理
    // 保留方法避免报错
  },
  
  /**
   * 关闭用户信息填写表单
   */
  closeUserInfoForm() {
    this.setData({ showUserInfoForm: false })
  },
  
  /**
   * 用户选择头像回调
   * 微信基础库 2.21.2 及以上版本支持
   */
  onChooseAvatar(e) {
    const { avatarUrl } = e.detail
    this.setData({
      tempAvatar: avatarUrl
    })
    // 上传头像
    this.uploadImage(avatarUrl)
  },
  
  /**
   * 昵称输入处理
   */
  onNicknameInput(e) { 
    this.setData({ tempNickname: e.detail.value }) 
  },
  
  /**
   * 提交用户信息
   * 收集用户选择的头像和输入的昵称，然后进行登录
   */
  submitUserInfo() {
    const { tempNickname, uploadedAvatarUrl, tempAvatar, wxLoginCode, wxIv, wxEncryptedData } = this.data
    
    if (!tempNickname) {
      return wx.showToast({ title: '请输入昵称', icon: 'none' })
    }
    
    if (!uploadedAvatarUrl && !tempAvatar) {
      return wx.showToast({ title: '请选择头像', icon: 'none' })
    }
    
    if (!wxLoginCode) {
      return wx.showToast({ title: '登录状态已过期，请重试', icon: 'none' })
    }
    
    // 如果有上传成功的头像URL，直接使用
    // 如果只有临时头像路径但没有上传，先上传头像再登录
    if (tempAvatar && !uploadedAvatarUrl) {
      wx.showLoading({ title: '正在上传头像...' })
      wx.uploadFile({
        url: API.UPLOAD_IMAGE,
        filePath: tempAvatar,
        name: 'File',
        formData: { 
          uploader_id: 0, 
          uploader_type: 1, // 1-H5用户
          file_type: 1 // 1-图片
        },
        success: (res) => {
          wx.hideLoading()
          const data = JSON.parse(res.data)
          if (data.code === 0 && data.data?.url) {
            // 上传成功后，使用返回的URL进行登录
            const wxLoginData = {
              code: wxLoginCode,
              iv: wxIv,
              encryptedData: wxEncryptedData,
              phoneNumber: '',
              nickname: tempNickname,
              avatar: data.data.url
            }
            this.wxMiniLogin(wxLoginData)
          } else {
            wx.showToast({ 
              title: data.message || '头像上传失败', 
              icon: 'none' 
            })
          }
        },
        fail: (err) => {
          wx.hideLoading()
          console.error('头像上传失败:', err)
          wx.showToast({ title: '头像上传失败', icon: 'none' })
        }
      })
    } else {
      // 已有上传成功的头像URL，直接登录
      const avatar = uploadedAvatarUrl || 'https://via.placeholder.com/100x100/19aecc/ffffff?text=用户'
      const wxLoginData = {
        code: wxLoginCode,
        iv: wxIv,
        encryptedData: wxEncryptedData,
        phoneNumber: '',
        nickname: tempNickname,
        avatar: avatar
      }
      this.wxMiniLogin(wxLoginData)
    }
  },

  /**
   * 检查登录状态
   * 从全局状态获取登录信息并更新页面数据
   */
  checkLoginStatus() {
    console.log('开始检查登录状态')
    const { isLoggedIn, userInfo } = checkLoginStatus()
    console.log('登录检查结果:', { isLoggedIn, userInfo })
    app.globalData.isLoggedIn = isLoggedIn
    app.globalData.userInfo = userInfo || {}
    this.setData({ isLoggedIn, userInfo: userInfo || {} })
    console.log('页面数据已更新，登录状态:', this.data.isLoggedIn)
  },

  closeWxLoginPopup() {
    this.setData({ showWxLogin: false })
  },

  stopPropagation() { return },

  /**
   * 选择头像方法
   * 已登录用户可选择相册或拍照方式更换头像
   */
  chooseAvatar() {
    if (!this.data.isLoggedIn) return;
    
    wx.showActionSheet({
      itemList: ['从相册选择', '拍照'],
      success: (res) => {
        const sourceType = res.tapIndex === 0 ? 'album' : 'camera'
        this.chooseImage(sourceType)
      },
      fail: (err) => {
        console.error('选择图片失败:', err)
        wx.showToast({
          title: '选择图片失败',
          icon: 'none'
        })
      }
    })
  },

  chooseImage(sourceType) {
    wx.chooseImage({
      count: 1,
      sourceType: [sourceType],
      success: (res) => {
        const tempFilePath = res.tempFilePaths[0]
        this.setData({ tempAvatar: tempFilePath })
        this.uploadImage(tempFilePath)
      }
    })
  },

  /**
   * 上传图片方法
   * @param {string} filePath - 要上传的图片临时路径
   */
  uploadImage(filePath) {
    wx.showLoading({ title: '上传中...' })
    wx.uploadFile({
      url: API.UPLOAD_IMAGE,
      filePath,
      name: 'File',
      formData: { 
        uploader_id: this.data.userInfo.id || 0, 
        uploader_type: 1, // 1-H5用户
        file_type: 1 // 1-图片
      },
      success: (res) => {
        wx.hideLoading()
        const data = JSON.parse(res.data)
        if (data.code === 0 && data.data?.url) {
          this.setData({ uploadedAvatarUrl: data.data.url })
          wx.showToast({ title: '头像上传成功', icon: 'success' })
        } else {
          wx.showToast({ 
            title: data.message || '上传失败', 
            icon: 'none' 
          })
        }
      },
      fail: (err) => { 
        wx.hideLoading()
        console.error('头像上传失败:', err)
        wx.showToast({ title: '上传失败', icon: 'none' }) 
      }
    })
  },
  
  /**
   * 选择头像方法
   * 已登录用户可选择相册或拍照方式更换头像
   */
  chooseAvatar() {
    if (!this.data.isLoggedIn) return;
    
    wx.showActionSheet({
      itemList: ['从相册选择', '拍照'],
      success: (res) => {
        const sourceType = res.tapIndex === 0 ? 'album' : 'camera'
        this.chooseImage(sourceType)
      },
      fail: (err) => {
        console.error('选择图片失败:', err)
        wx.showToast({
          title: '选择图片失败',
          icon: 'none'
        })
      }
    })
  },

  chooseImage(sourceType) {
    wx.chooseImage({
      count: 1,
      sourceType: [sourceType],
      success: (res) => {
        const tempFilePath = res.tempFilePaths[0]
        this.setData({ tempAvatar: tempFilePath })
        this.uploadImage(tempFilePath)
      }
    })
  },


  




  /**
   * 微信小程序登录方法
   * @param {object} wxLoginData - 包含code、iv、encryptedData等登录数据
   */
  wxMiniLogin(wxLoginData) {
    wx.showLoading({ title: '登录中...' })
    request({ url: API.USER_WX_LOGIN, method: 'POST', data: wxLoginData, needAuth: false })
      .then(res => {
        wx.hideLoading()
        
        // 由于request.js已经处理了响应格式，这里直接使用res
        // res已经是response.data，即登录接口返回的data字段内容
        const loginData = res
        
        if (!loginData) {
          return wx.showToast({ title: '登录数据异常', icon: 'none' })
        }
        
        // 检查是否需要手机号授权
        if (loginData.need_phone_auth) {
          console.log('需要手机号授权，显示授权弹窗')
          // 保存登录数据到临时存储，用于后续授权
          this.setData({
            wxLoginTempData: loginData,
            showPhoneAuthPopup: true
          })
          return
        }
        
        // 正常登录流程
        this.handleLoginSuccess(loginData)
        
      }).catch(err => { 
        wx.hideLoading()
        console.error('登录失败:', err)
        wx.showToast({ title: '登录失败', icon: 'none' }) 
      })
  },

  // 获取用户信息回调
  /**
   * 获取用户信息回调
   * @param {object} e - 事件对象，包含用户信息
   */
  onGetUserInfo(e) {
    if (e.detail.errMsg !== 'getUserInfo:ok') {
      return wx.showToast({ title: '获取用户信息失败', icon: 'none' })
    }
    
    wx.login({
      success: (loginRes) => {
        if (!loginRes.code) return wx.showToast({ title: '获取code失败', icon: 'none' })
        
        const wxLoginData = {
          code: loginRes.code,
          iv: e.detail.iv,
          encryptedData: e.detail.encryptedData,
          phoneNumber: '',
          nickname: e.detail.userInfo.nickName,
          avatar: e.detail.userInfo.avatarUrl
        }
        this.wxMiniLogin(wxLoginData)
      }
    })
  },

  // 头像/昵称点击处理
  /**
   * 头像/昵称点击处理
   * 未登录用户点击时直接触发登录流程
   */
  onProfileClick() {
    if (!this.data.isLoggedIn) {
      this.autoWxLogin()
    }
  },
  
  /**
   * 检查微信基础库版本
   * 返回是否支持新版头像选择器
   */
  checkWxVersion() {
    try {
      const version = wx.getSystemInfoSync().SDKVersion
      return this.compareVersion(version, '2.21.2') >= 0
    } catch (e) {
      return false
    }
  },
  
  /**
   * 比较版本号
   */
  compareVersion(v1, v2) {
    v1 = v1.split('.')
    v2 = v2.split('.')
    const len = Math.max(v1.length, v2.length)
    
    while (v1.length < len) {
      v1.push('0')
    }
    while (v2.length < len) {
      v2.push('0')
    }
    
    for (let i = 0; i < len; i++) {
      const num1 = parseInt(v1[i])
      const num2 = parseInt(v2[i])
      
      if (num1 > num2) {
        return 1
      } else if (num1 < num2) {
        return -1
      }
    }
    
    return 0
  },

  /**
   * 退出登录方法
   * 清空本地缓存和全局状态，返回登录页面
   */
  logout() {
    wx.showModal({
      title: '提示',
      content: '确定要退出登录吗？',
      success: (res) => {
        if (res.confirm) {
          // 清空本地存储
          wx.removeStorageSync('token')
          wx.removeStorageSync('userInfo')
          wx.removeStorageSync('openId') // 同时清除openId
          
          // 重置全局状态
          app.globalData.isLoggedIn = false
          app.globalData.userInfo = {}
          
          // 更新页面数据，包括重置订单统计数据
          this.setData({
            isLoggedIn: false,
            userInfo: {},
            orderCounts: {
              pending: 0,
              shipping: 0,
              delivered: 0,
              completed: 0,
              afterSale: 0
            }
          })
          
          wx.showToast({
            title: '已退出登录',
            icon: 'success'
          })
          
          // 留在当前用户页面，显示微信一键登录按钮
        }
      }
    })
  },

  /**
   * 清除登录状态
   * 当用户不存在或token失效时调用
   */
  clearLoginStatus() {
    // 清空本地存储
    wx.removeStorageSync('token')
    wx.removeStorageSync('userInfo')
    wx.removeStorageSync('openId')
    
    // 重置全局状态
    app.globalData.isLoggedIn = false
    app.globalData.userInfo = {}
    
    // 更新页面数据，包括重置订单统计数据
    this.setData({
      isLoggedIn: false,
      userInfo: {},
      orderCounts: {
        pending: 0,
        shipping: 0,
        delivered: 0,
        completed: 0,
        afterSale: 0
      }
    })
  },

  /**
   * 获取用户信息
   * 获取用户基本信息和订单统计数据
   */
  getUserInfo() {
    console.log('开始获取用户信息')
    request({
      url: API.USER_INFO,
      method: 'GET'
    }).then(res => {
      console.log('用户信息接口返回:', res)
      if (res.code === 0 && res.data) {
        console.log('获取用户信息成功:', res.data)
        // 更新用户信息
        this.setData({ userInfo: res.data })
        // 保存到全局状态和本地存储
        app.globalData.userInfo = res.data
        wx.setStorageSync('userInfo', res.data)
      } else {
        console.log('获取用户信息失败，返回数据异常:', res)
      }
    }).catch(err => {
      console.error('获取用户信息失败:', err)
      // 当获取用户信息失败时，设置为未登录状态
      // 特别处理用户不存在的错误（code:52），不显示错误提示
      if (err.message && err.message.includes('用户不存在')) {
        // 用户不存在，清除登录状态
        this.clearLoginStatus()
      } else {
        // 其他错误，只在开发模式下显示
        if (app.globalData.isDev) {
          wx.showToast({
            title: '获取用户信息失败',
            icon: 'none'
          })
        }
      }
    })
  },

  /**
   * 获取订单数量统计
   * 获取各状态订单的数量
   */
  getOrderCounts() {
    console.log('开始获取订单数量统计')
    console.log('API.ORDER_COUNT:', API.ORDER_COUNT)
    request({
      url: API.ORDER_COUNT,
      method: 'GET'
    }).then(res => {
      console.log('订单数量接口返回:', res)
      // 检查返回数据格式
      if (res && typeof res === 'object') {
        let orderData = null;
        
        // 如果返回的是直接的订单数据对象
        if (res.pending !== undefined || res.shipping !== undefined) {
          orderData = res;
          console.log('直接返回订单数据:', orderData)
        } 
        // 如果返回的是标准格式 {code: 0, data: {...}}
        else if (res.code === 0 && res.data) {
          orderData = res.data;
          console.log('标准格式订单数据:', orderData)
        }
        
        if (orderData) {
          // 更新订单统计数据
          this.setData({
            orderCounts: {
              pending: orderData.pending || 0,
              shipping: orderData.shipping || 0,
              delivered: orderData.delivered || 0,
              completed: orderData.completed || 0,
              afterSale: orderData.afterSale || 0
            }
          })
          console.log('更新后的orderCounts:', this.data.orderCounts)
        } else {
          console.log('无法识别的数据格式:', res)
        }
      } else {
        console.log('订单数量接口返回格式异常:', res)
      }
    }).catch(err => {
      console.error('获取订单统计失败:', err)
      console.error('错误详情:', JSON.stringify(err))
    })
  },

  /**
   * 查看指定状态的订单
   * @param {object} e - 事件对象，包含订单状态
   */
  viewOrders(e) {
    if (!this.data.isLoggedIn) {
      return wx.showToast({
        title: '请先登录',
        icon: 'none'
      })
    }
    
    const status = e.currentTarget.dataset.status
    wx.navigateTo({
      url: `/pages/order-list/order-list?status=${status}&showTabs=false`
    })
  },

  /**
   * 查看全部订单
   */
  viewAllOrders() {
    if (!this.data.isLoggedIn) {
      return wx.showToast({
        title: '请先登录',
        icon: 'none'
      })
    }
    
    wx.navigateTo({
      url: '/pages/order-list/order-list?showTabs=true'
    })
  },

  /**
   * 跳转到指定页面
   * @param {object} e - 事件对象，包含页面URL
   */
  navigateTo(e) {
    const url = e.currentTarget.dataset.url
    wx.navigateTo({ url })
  }
})
