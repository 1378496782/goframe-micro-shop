 const app = getApp()
const { checkLoginStatus, request } = require('../../utils/request')
const { API } = require('../../utils/env')

Page({
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
    uploadedAvatarUrl: ''
  },

  onLoad() {
    this.checkLoginStatus()
  },

  onShow() {
    this.checkLoginStatus()
    if (this.data.isLoggedIn) this.getUserInfo()
  },

  checkLoginStatus() {
    const { isLoggedIn, userInfo } = checkLoginStatus()
    app.globalData.isLoggedIn = isLoggedIn
    app.globalData.userInfo = userInfo || {}
    this.setData({ isLoggedIn, userInfo: userInfo || {} })
  },

  closeWxLoginPopup() {
    this.setData({ showWxLogin: false })
  },

  stopPropagation() { return },

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

  wxMiniLogin(wxLoginData) {
    wx.showLoading({ title: '登录中...' })
    request({ url: API.USER_WX_LOGIN, method: 'POST', data: wxLoginData, needAuth: false })
      .then(loginData => {
        wx.hideLoading()
        wx.setStorageSync('token', loginData.token)
        wx.setStorageSync('userInfo', loginData.user_info)
        app.globalData.isLoggedIn = true
        app.globalData.userInfo = loginData.user_info
        this.setData({ isLoggedIn: true, userInfo: loginData.user_info, showWxLogin: false })
        wx.showToast({ title: '登录成功', icon: 'success' })
      }).catch(() => { wx.hideLoading(); wx.showToast({ title: '登录失败', icon: 'none' }) })
  },

  // 新增头像/昵称点击处理
  onProfileClick() {
    if (!this.data.isLoggedIn) {
      wx.navigateTo({
        url: '/pages/login/login?type=wechat'
      })
    }
  },

  // 确保登录按钮功能
  onLogin() {
    wx.navigateTo({
      url: '/pages/login/login'
    })
  }
})
