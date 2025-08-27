App({
  onLaunch() {
    console.log('小程序初始化')
    
    // 检查登录状态
    this.checkLoginStatus()
    
    // 获取系统信息
    this.getSystemInfo()
  },
  
  onShow() {
    console.log('小程序显示')
  },
  
  onHide() {
    console.log('小程序隐藏')
  },
  
  // 检查登录状态
  checkLoginStatus() {
    const token = wx.getStorageSync('token')
    if (token) {
      console.log('用户已登录')
      this.globalData.isLoggedIn = true
    } else {
      console.log('用户未登录')
      this.globalData.isLoggedIn = false
    }
  },
  
  // 获取系统信息
  getSystemInfo() {
    wx.getSystemInfo({
      success: (res) => {
        this.globalData.systemInfo = res
        console.log('系统信息:', res)
      }
    })
  },
  
  // 全局数据
  globalData: {
    isLoggedIn: false,
    userInfo: null,
    systemInfo: null,
    cartCount: 0
  }
})