Page({
  data: {
    message: '测试页面加载成功!',
    items: [
      { name: '测试商品1', price: 9.99, quantity: 1 },
      { name: '测试商品2', price: 19.99, quantity: 2 }
    ]
  },

  onLoad() {
    console.log('测试页面加载')
    wx.showToast({
      title: '测试页面加载成功',
      icon: 'none'
    })
  }
})