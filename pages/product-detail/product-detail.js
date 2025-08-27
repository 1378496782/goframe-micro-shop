Page({
  data: {
    currentImageIndex: 0,
    selectedSpecs: {},
    product: {
      id: 1,
      name: '高品质智能手机 8GB+256GB 全网通5G',
      price: '2999.00',
      originalPrice: '3999.00',
      discount: '7.5',
      sales: 12560,
      reviews: 3489,
      stock: 500,
      images: [
        'https://via.placeholder.com/400x400/19aecc/ffffff?text=主图1',
        'https://via.placeholder.com/400x400/19aecc/ffffff?text=主图2',
        'https://via.placeholder.com/400x400/19aecc/ffffff?text=主图3',
        'https://via.placeholder.com/400x400/19aecc/ffffff?text=主图4'
      ],
      specs: [
        {
          name: '颜色',
          values: ['黑色', '白色', '蓝色', '绿色']
        },
        {
          name: '内存',
          values: ['8GB+128GB', '8GB+256GB', '12GB+256GB', '12GB+512GB']
        }
      ]
    }
  },

  onLoad(options) {
    // 可以根据options中的参数加载不同的商品数据
    console.log('商品详情页面加载', options)
  },

  // 计算当前显示的图片
  currentImage() {
    return this.data.product.images[this.data.currentImageIndex]
  },

  // 选择图片
  selectImage(e) {
    const index = e.currentTarget.dataset.index
    this.setData({
      currentImageIndex: index
    })
  },

  // 选择规格
  selectSpec(e) {
    const { specName, value } = e.currentTarget.dataset
    const selectedSpecs = { ...this.data.selectedSpecs }
    selectedSpecs[specName] = value
    
    this.setData({
      selectedSpecs
    })
  },

  // 添加到收藏
  addToFavorites() {
    wx.showToast({
      title: '已添加到收藏',
      icon: 'success'
    })
  },

  // 前往购物车
  goToCart() {
    wx.switchTab({
      url: '/pages/cart/cart'
    })
  },

  // 加入购物车
  addToCart() {
    if (Object.keys(this.data.selectedSpecs).length === 0) {
      wx.showToast({
        title: '请选择规格',
        icon: 'none'
      })
      return
    }

    wx.showToast({
      title: '已加入购物车',
      icon: 'success'
    })
    console.log('加入购物车', this.data.selectedSpecs)
  },

  // 立即购买
  buyNow() {
    if (Object.keys(this.data.selectedSpecs).length === 0) {
      wx.showToast({
        title: '请选择规格',
        icon: 'none'
      })
      return
    }

    wx.navigateTo({
      url: `/pages/order/confirm?productId=${this.data.product.id}`
    })
    console.log('立即购买', this.data.selectedSpecs)
  }
})