const app = getApp()

Page({
  data: {
    result: '',
    loading: false
  },

  onTestRequest() {
    this.setData({ loading: true, result: '' })
    
    // 直接使用wx.request测试
    wx.request({
      url: 'http://shop.dayu.club:8199/frontend/user/login',
      method: 'POST',
      data: {
        name: 'testuser',
        password: 'testpass'
      },
      success: (res) => {
        console.log('请求成功:', res)
        this.setData({ 
          loading: false,
          result: JSON.stringify(res, null, 2)
        })
      },
      fail: (err) => {
        console.log('请求失败:', err)
        this.setData({ 
          loading: false,
          result: `请求失败: ${JSON.stringify(err, null, 2)}`
        })
      },
      complete: () => {
        console.log('请求完成')
      }
    })
  },

  onTestPing() {
    this.setData({ loading: true, result: '' })
    
    // 测试简单的HTTP请求
    wx.request({
      url: 'http://httpbin.org/get',
      method: 'GET',
      success: (res) => {
        console.log('Ping成功:', res)
        this.setData({ 
          loading: false,
          result: JSON.stringify(res, null, 2)
        })
      },
      fail: (err) => {
        console.log('Ping失败:', err)
        this.setData({ 
          loading: false,
          result: `Ping失败: ${JSON.stringify(err, null, 2)}`
        })
      }
    })
  }
})