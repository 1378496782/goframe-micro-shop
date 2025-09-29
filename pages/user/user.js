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
    tempAvatar: '',
    uploadedAvatarUrl: '',
    hasShownLoginTip: false
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
    } else {
      console.log('未登录，显示登录引导')
      // 只在首次进入页面时显示登录引导提示
      if (!this.data.hasShownLoginTip) {
        this.data.hasShownLoginTip = true
        // 显示登录引导提示
        wx.showToast({
          title: '请点击头像登录',
          icon: 'none',
          duration: 2000
        })
      }
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
   * 微信自动登录方法
   * 调用微信登录接口获取code，然后获取用户信息
   */
  autoWxLogin() {
    wx.showLoading({ title: '准备登录...' })
    wx.login({
      success: (loginRes) => {
        if (!loginRes.code) {
          wx.hideLoading()
          return wx.showToast({ title: '获取code失败', icon: 'none' })
        }
        
        // 获取用户信息
        wx.getUserProfile({
          desc: '用于完善会员资料',
          success: (profileRes) => {
            if (!profileRes.iv || !profileRes.encryptedData) {
              wx.hideLoading()
              return wx.showToast({ title: '获取用户信息不完整', icon: 'none' })
            }
            
            const wxLoginData = {
              code: loginRes.code,
              iv: profileRes.iv,
              encryptedData: profileRes.encryptedData,
              phoneNumber: '',
              nickname: profileRes.userInfo.nickName || '微信用户',
              avatar: profileRes.userInfo.avatarUrl || 'https://via.placeholder.com/100x100/19aecc/ffffff?text=微信用户'
            }
            this.wxMiniLogin(wxLoginData)
          },
          fail: (err) => {
            wx.hideLoading()
            console.error('获取用户信息失败:', err)
            wx.showToast({ title: '获取用户信息失败', icon: 'none' })
          }
        })
      },
      fail: (err) => {
        wx.hideLoading()
        console.error('微信登录失败:', err)
        wx.showToast({ title: '微信登录失败', icon: 'none' })
      }
    })
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
      url: 'http://192.168.1.5:8399/upload/image',
      filePath,
      name: 'File',
      formData: { uploader_id: 0, uploader_type: 1, file_type: 1 },
      success: (res) => {
        wx.hideLoading()
        const data = JSON.parse(res.data)
        if (data.code === 0 && data.data?.url) {
          this.setData({ uploadedAvatarUrl: data.data.url })
          wx.showToast({ title: '上传成功', icon: 'success' })
        } else wx.showToast({ title: '上传失败', icon: 'none' })
      },
      fail: () => { wx.hideLoading(); wx.showToast({ title: '上传失败', icon: 'none' }) }
    })
  },

  onNicknameInput(e) { this.setData({ tempNickname: e.detail.value }) },

  getPhoneNumber(e) {
    if (e.detail.errMsg !== 'getPhoneNumber:ok') {
      wx.showToast({ title: '获取手机号失败', icon: 'none' })
      return
    }

    const avatar = this.data.uploadedAvatarUrl || this.data.tempAvatar || 'https://via.placeholder.com/100x100/19aecc/ffffff?text=用户'
    const nickname = this.data.tempNickname || '用户' + Math.floor(Math.random() * 10000)

    wx.login({
      success: (loginRes) => {
        if (!loginRes.code) return wx.showToast({ title: '登录失败', icon: 'none' })
        const wxLoginData = {
          code: loginRes.code,
          iv: e.detail.iv,
          encryptedData: e.detail.encryptedData,
          phoneNumber: '',
          nickname,
          avatar
        }
        this.wxMiniLogin(wxLoginData)
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
      .then(loginData => {
        wx.hideLoading()
        wx.setStorageSync('token', loginData.token)
        wx.setStorageSync('userInfo', loginData.user_info)
        
        // 保存openId，用于微信支付
        if (loginData.openId) {
          wx.setStorageSync('openId', loginData.openId)
          console.log('已保存openId:', loginData.openId)
        }
        
        app.globalData.isLoggedIn = true
        app.globalData.userInfo = loginData.user_info
        this.setData({ isLoggedIn: true, userInfo: loginData.user_info, showWxLogin: false })
        wx.showToast({ title: '登录成功', icon: 'success' })
      }).catch(() => { wx.hideLoading(); wx.showToast({ title: '登录失败', icon: 'none' }) })
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
   * 未登录用户点击时显示登录引导
   */
  onProfileClick() {
    if (!this.data.isLoggedIn) {
      wx.showToast({
        title: '请点击"微信一键登录"按钮',
        icon: 'none'
      })
    }
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
          
          // 更新页面数据
          this.setData({
            isLoggedIn: false,
            userInfo: {}
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
    
    // 更新页面数据
    this.setData({
      isLoggedIn: false,
      userInfo: {}
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
