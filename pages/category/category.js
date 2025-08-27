Page({
  data: {
    currentCategory: 1,
    categories: [
      { id: 1, name: '手机数码' },
      { id: 2, name: '电脑办公' },
      { id: 3, name: '家用电器' },
      { id: 4, name: '服装鞋帽' },
      { id: 5, name: '美妆护肤' },
      { id: 6, name: '食品生鲜' },
      { id: 7, name: '家居家装' },
      { id: 8, name: '母婴玩具' },
      { id: 9, name: '运动户外' }
    ],
    categoryData: {
      1: {
        banner: 'https://via.placeholder.com/750x200/19aecc/ffffff?text=手机数码',
        subCategories: [
          { id: 101, name: '手机', image: 'https://via.placeholder.com/80x80/19aecc/ffffff?text=📱' },
          { id: 102, name: '平板', image: 'https://via.placeholder.com/80x80/19aecc/ffffff?text=📱' },
          { id: 103, name: '耳机', image: 'https://via.placeholder.com/80x80/19aecc/ffffff?text=🎧' },
          { id: 104, name: '配件', image: 'https://via.placeholder.com/80x80/19aecc/ffffff?text=🔌' }
        ],
        products: [
          {
            id: 1001,
            name: '智能手机 8+256GB',
            price: '2999.00',
            originalPrice: '3999.00',
            image: 'https://via.placeholder.com/200x200/19aecc/ffffff?text=手机',
            sales: 12560
          },
          {
            id: 1002,
            name: '无线蓝牙耳机',
            price: '399.00',
            originalPrice: '499.00',
            image: 'https://via.placeholder.com/200x200/19aecc/ffffff?text=耳机',
            sales: 23450
          }
        ]
      }
    }
  },

  onLoad(options) {
    console.log('分类页面加载', options)
  },

  // 切换分类
  switchCategory(e) {
    const categoryId = e.currentTarget.dataset.id
    this.setData({
      currentCategory: categoryId
    })
  },

  // 获取当前分类数据
  currentCategoryData() {
    return this.data.categoryData[this.data.currentCategory] || {
      banner: '',
      subCategories: [],
      products: []
    }
  },

  // 点击子分类
  onSubCategoryClick(e) {
    const subCategoryId = e.currentTarget.dataset.id
    wx.showToast({
      title: `点击子分类 ${subCategoryId}`,
      icon: 'none'
    })
  },

  // 点击商品
  onProductClick(e) {
    const productId = e.currentTarget.dataset.id
    wx.navigateTo({
      url: `/pages/product-detail/product-detail?id=${productId}`
    })
  }
})