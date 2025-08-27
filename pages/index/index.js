Page({
  data: {
    banners: [
      { id: 1, image: 'http://wangzhongyang.com/images/logo_removebg.png', url: '/pages/category/category' },
      { id: 2, image: 'http://wangzhongyang.com/images/logo_removebg.png', url: '/pages/category/category' },
      { id: 3, image: 'http://wangzhongyang.com/images/logo_removebg.png', url: '/pages/category/category' }
    ],
    categories: [
      { id: 1, name: '手机', icon: 'http://wangzhongyang.com/images/logo_removebg.png' },
      { id: 2, name: '电脑', icon: 'http://wangzhongyang.com/images/logo_removebg.png' },
      { id: 3, name: '配件', icon: 'http://wangzhongyang.com/images/logo_removebg.png' },
      { id: 4, name: '家电', icon: 'http://wangzhongyang.com/images/logo_removebg.png' },
      { id: 5, name: '服饰', icon: 'http://wangzhongyang.com/images/logo_removebg.png' }
    ],
    products: [
      {
        id: 1,
        name: '高品质智能手机 8GB+256GB',
        price: '2999.00',
        originalPrice: '3999.00',
        image: 'http://wangzhongyang.com/images/logo_removebg.png',
        sales: 12560
      },
      {
        id: 2,
        name: '轻薄笔记本电脑 i7处理器',
        price: '5999.00',
        originalPrice: '6999.00',
        image: 'http://wangzhongyang.com/images/logo_removebg.png',
        sales: 8560
      },
      {
        id: 3,
        name: '无线蓝牙耳机 降噪版',
        price: '399.00',
        originalPrice: '499.00',
        image: 'http://wangzhongyang.com/images/logo_removebg.png',
        sales: 23450
      },
      {
        id: 4,
        name: '智能手表 运动健康版',
        price: '899.00',
        originalPrice: '1099.00',
        image: 'http://wangzhongyang.com/images/logo_removebg.png',
        sales: 15680
      }
    ],
    searchValue: ''
  },

  onLoad() {
    console.log('首页加载')
  },

  onShow() {
    console.log('首页显示')
  },

  onSearchInput(e) {
    this.setData({
      searchValue: e.detail.value
    })
  },

  onSearch() {
    if (!this.data.searchValue.trim()) {
      wx.showToast({
        title: '请输入搜索关键词',
        icon: 'none'
      })
      return
    }

    wx.navigateTo({
      url: `/pages/search/search?keyword=${this.data.searchValue}`
    })
  },

  onBannerClick(e) {
    const url = e.currentTarget.dataset.url
    wx.navigateTo({
      url: url
    })
  },

  onCategoryClick(e) {
    const categoryId = e.currentTarget.dataset.id
    wx.navigateTo({
      url: `/pages/category/category?id=${categoryId}`
    })
  },

  onViewMore() {
    wx.navigateTo({
      url: '/pages/category/category'
    })
  },

  onProductClick(e) {
    const productId = e.currentTarget.dataset.id
    wx.navigateTo({
      url: `/pages/product-detail/product-detail?id=${productId}`
    })
  },

  onPullDownRefresh() {
    console.log('下拉刷新')
    setTimeout(() => {
      wx.stopPullDownRefresh()
    }, 1000)
  }
})