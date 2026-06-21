const app = getApp()
const { userAPI } = require('../../utils/request')

Page({
  data: {
    phone: '',
    password: '',
    loading: false
  },

  onPhoneInput(e) {
    this.setData({
      phone: e.detail.value
    });
  },

  onPasswordInput(e) {
    this.setData({
      password: e.detail.value
    });
  },

  onLogin() {
    const { phone, password } = this.data;

    if (!phone) {
      wx.showToast({
        title: '请输入手机号',
        icon: 'none'
      });
      return;
    }

    if (!/^1[3-9]\d{9}$/.test(phone)) {
      wx.showToast({
        title: '请输入正确的手机号',
        icon: 'none'
      });
      return;
    }

    if (!password) {
      wx.showToast({
        title: '请输入密码',
        icon: 'none'
      });
      return;
    }

    this.setData({ loading: true });

    // 使用封装的API进行登录
    userAPI.login({
      name: phone,
      password: password
    }).then((loginData) => {
      this.setData({ loading: false });
      
      // 更新全局登录状态和用户信息
      app.globalData.isLoggedIn = true
      const userInfo = loginData.userInfo || loginData.user_info || {}
      app.globalData.userInfo = userInfo
      
      console.log('登录成功，用户信息:', userInfo)
      console.log('全局用户信息:', app.globalData.userInfo)
      
      wx.showToast({
        title: '登录成功',
        icon: 'success'
      });
      
      // 登录成功后跳转到个人中心页面，确保用户信息显示更新
      setTimeout(() => {
        wx.switchTab({
          url: '/pages/user/user'
        });
      }, 1500);
    }).catch(err => {
      this.setData({ loading: false });
      // 错误信息已经在request中统一处理
      console.error('登录失败:', err)
    });
  },

  navigateToRegister() {
    wx.navigateTo({
      url: '/pages/register/register'
    });
  }
});