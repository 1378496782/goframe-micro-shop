const app = getApp()
const { checkLoginStatus } = require('../../utils/request')

Page({
  data: {
    isLoggedIn: false,
    userInfo: {},
    orderCounts: {
      pending: 2,
      shipping: 1,
      delivered: 0,
      completed: 5,
      afterSale: 0
    }
  },

  onLoad() {
    this.checkLoginStatus()
  },

  onShow() {
    this.checkLoginStatus()
  },

  // 检查登录状态
  async checkLoginStatus() {
    const { isLoggedIn, userInfo } = checkLoginStatus()
    app.globalData.isLoggedIn = isLoggedIn
    app.globalData.userInfo = userInfo
    
    this.setData({
      isLoggedIn,
      userInfo: userInfo || {}
    })
    
    // 如果已登录，获取最新的用户信息
    if (isLoggedIn) {
      await this.getUserInfo()
    }
  },

  // 登录
  onLogin() {
    wx.navigateTo({
      url: '/pages/login/login'
    })
  },

  // 退出登录
  onLogout() {
    wx.showModal({
      title: '提示',
      content: '确定要退出登录吗？',
      success: (res) => {
        if (res.confirm) {
          // 清除登录状态
          wx.removeStorageSync('token')
          app.globalData.isLoggedIn = false
          app.globalData.userInfo = null
          
          this.setData({
            isLoggedIn: false,
            userInfo: {}
          })
          
          wx.showToast({
            title: '已退出登录',
            icon: 'success'
          })
        }
      }
    })
  },

  // 查看订单
  viewOrders(e) {
    const status = e.currentTarget.dataset.status
    wx.navigateTo({
      url: `/pages/orders/orders?status=${status}`
    })
  },

  // 查看全部订单
  viewAllOrders() {
    wx.navigateTo({
      url: '/pages/orders/orders'
    })
  },

  // 导航到页面
  navigateTo(e) {
    const url = e.currentTarget.dataset.url
    if (this.data.isLoggedIn) {
      wx.navigateTo({ url })
    } else {
      wx.showModal({
        title: '提示',
        content: '请先登录',
        success: (res) => {
          if (res.confirm) {
            this.onLogin()
          }
        }
      })
    }
  },

  // 获取用户信息
  async getUserInfo() {
    if (!this.data.isLoggedIn) {
      return
    }
    
    try {
      const { userAPI } = require('../../utils/request')
      const res = await userAPI.getInfo()
      
      if (res.code === 0) {
        const userInfo = {
          nickname: res.data.username || '微信用户',
          avatar: res.data.avatar || 'https://via.placeholder.com/100x100/19aecc/ffffff?text=用户',
          level: res.data.level || '普通会员'
        }
        
        // 保存用户信息到全局和本地
        app.globalData.userInfo = userInfo
        const { saveUserInfo } = require('../../utils/request')
        saveUserInfo(userInfo)
        
        this.setData({ userInfo })
      } else {
        console.error('获取用户信息失败:', res.message)
      }
    } catch (error) {
      console.error('获取用户信息异常:', error)
    }
  },

  // 修改密码
  onChangePassword() {
    if (!this.data.isLoggedIn) {
      wx.showModal({
        title: '提示',
        content: '请先登录',
        success: (res) => {
          if (res.confirm) {
            this.onLogin()
          }
        }
      })
      return
    }
    
    wx.navigateTo({
      url: '/pages/change-password/change-password'
    })
  }
})