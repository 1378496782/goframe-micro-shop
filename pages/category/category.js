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
        banner: 'http://wangzhongyang.com/images/logo_removebg.png',
        subCategories: [
          { id: 101, name: '手机', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 102, name: '平板', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 103, name: '耳机', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 104, name: '配件', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 105, name: '智能手表', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 106, name: '相机', image: 'http://wangzhongyang.com/images/logo_removebg.png' }
        ],
        products: [
          {
            id: 1001,
            name: '智能手机 8+256GB 全面屏',
            price: '2999.00',
            originalPrice: '3999.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 12560
          },
          {
            id: 1002,
            name: '无线蓝牙耳机 降噪版',
            price: '399.00',
            originalPrice: '499.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 23450
          },
          {
            id: 1003,
            name: '智能手表 运动健康监测',
            price: '899.00',
            originalPrice: '1099.00',
            image: 'http极速版://wangzhongyang.com/images/logo_removebg.png',
            sales: 15680
          },
          {
            id: 1004,
            name: '平板电脑 10.2英寸',
            price: '1999.00',
            originalPrice: '2499.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 8950
          }
        ]
      },
      2: {
        banner: 'http://wangzhongyang.com/images/logo_removebg.png',
        subCategories: [
          { id: 201, name: '笔记本', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 202, name: '台式机', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 203, name: '显示器', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 204, name: '外设', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 205, name: '办公设备', image: 'http://wangzhongyang.com/images/logo_removebg.png' }
        ],
        products: [
          {
            id: 2001,
            name: '轻薄笔记本电脑 i7处理器',
            price: '5999.00',
            originalPrice: '6999.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 8560
          },
          {
            id: 2002,
            name: '4K显示器 27英寸',
            price: '1999.00',
            originalPrice: '2499.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 4320
          }
        ]
      },
      3: {
        banner: 'http://wangzhongyang.com/images/logo_removebg.png',
        subCategories: [
          { id: 301, name: '电视', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 302, name: '空调', image: 'http://wangzhongyang.com/images/logo_removebg极速版.png' },
          { id: 303, name: '冰箱', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 304, name: '洗衣机', image极速版: 'http://wangzhongyang.com/images/logo_removebg.png' }
        ],
        products: [
          {
            id: 3001,
            name: '智能电视 55英寸 4K',
            price: '2999.00',
            originalPrice: '3999.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 5670
          },
          {
            id: 3002,
            name: '变频空调 1.5匹',
            price: '2599.00',
            originalPrice: '3299.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 3890
          }
        ]
      },
      4: {
        banner: 'http://wangzhongyang.com/images/logo_removebg.png',
        subCategories: [
          { id: 401, name: '男装', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 402, name: '女装', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 403, name: '鞋靴', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 404, name: '箱包', image: 'http://wangzhongyang.com/images/logo_removebg.png' }
        ],
        products: [
          {
            id: 4001,
            name: '男士休闲外套 春秋款',
            price: '299.00',
            originalPrice: '399.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 12340
          },
          {
            id: 4002,
            name: '女士连衣裙 夏季新款',
            price: '199.00',
            originalPrice: '299.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 15680
          }
        ]
      },
      5: {
        banner: 'http://wangzhongyang.com/images/logo_removebg.png',
        subCategories: [
          { id: 501, name: '护肤', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 502, name: '彩妆', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 503, name: '香水', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 504, name: '个护', image: 'http://wangzhongyang.com/images/logo_removebg.png' }
        ],
        products: [
          {
            id: 5001,
            name: '保湿面霜 50ml',
            price: '159.00',
            originalPrice: '199.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 23450
          },
          {
            id: 5002,
            name: '口红 丝绒质地',
            price: '99.00',
            originalPrice: '129.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 18760
          }
        ]
      },
      6: {
        banner: 'http://wangzhongyang.com/images/logo_removebg.png',
        subCategories: [
          { id: 601, name: '水果', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 602, name: '蔬菜', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 603, name: '肉类', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 604, name: '零食', image: 'http://wangzhongyang.com/images/logo_removebg.png' }
        ],
        products: [
          {
            id: 6001,
            name: '新鲜水果礼盒',
            price: '128.00',
            originalPrice: '158.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 45670
          },
          {
            id: 6002,
            name: '进口零食大礼包',
            price: '88.00',
            originalPrice: '108.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 28900
          }
        ]
      },
      7: {
        banner: 'http://wangzhongyang.com/images/logo_removebg.png',
        subCategories: [
          { id: 701, name: '家具', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 702, name: '家纺', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 703, name: '厨具', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 704, name: '装饰', image: 'http://wangzhongyang.com/images/logo_removebg.png' }
        ],
        products: [
          {
            id: 7001,
            name: '实木餐桌 1.4米',
            price: '1299.00',
            originalPrice: '1599.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 2340
          },
          {
            id: 7002,
            name: '床上四件套 纯棉',
            price: '299.00',
            originalPrice: '399.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 5670
          }
        ]
      },
      8: {
        banner: 'http://wangzhongyang.com/images/logo_removebg.png',
        subCategories: [
          { id: 801, name: '奶粉', image: 'http://wangzhongyang.com/images/logo_removebg.png极速版' },
          { id: 802, name: '尿裤', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 803, name: '玩具', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 804, name: '童装', image: 'http://wangzhongyang.com/images/logo_removebg.png' }
        ],
        products: [
          {
            id: 8001,
            name: '婴儿奶粉 1段',
            price: '298.00',
            originalPrice: '358.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 12340
          },
          {
            id: 8002,
            name: '益智玩具套装',
            price: '159.00',
            originalPrice: '199.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 8760
          }
        ]
      },
      9: {
        banner: 'http://wangzhongyang.com/images/logo_removebg.png',
        subCategories: [
          { id: 901, name: '运动鞋', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 902, name: '运动服', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 903, name: '器材', image: 'http://wangzhongyang.com/images/logo_removebg.png' },
          { id: 904, name: '户外', image: 'http://wangzhongyang.com/images/logo_removebg.png' }
        ],
        products: [
          {
            id: 9001,
            name: '跑步鞋 减震透气',
            price: '499.00',
            originalPrice: '599.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 9870
          },
          {
            id: 9002,
            name: '运动外套 防风',
            price: '299.00',
            originalPrice: '399.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 6540
          }
        ]
      }
    }
  },

  onLoad(options) {
    console.log('分类页面加载', options)
    // 模拟加载更多数据
    this.loadMoreProducts()
  },

  // 加载更多商品
  loadMoreProducts() {
    setTimeout(() => {
      const currentData = this.data.categoryData[this.data.currentCategory]
      if (currentData && currentData.products.length < 10) {
        const newProducts = [
          {
            id: 1005,
            name: '充电宝 20000mAh 快充',
            price: '129.00',
            originalPrice: '159.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 8760
          },
          {
            id: 1006,
            name: '手机壳 防摔保护',
            price: '49.00',
            originalPrice: '69.00',
            image: 'http://wangzhongyang.com/images/logo_removebg.png',
            sales: 23450
          }
        ]
        
        const categoryData = { ...this.data.categoryData }
        categoryData[this.data.currentCategory].products = [
          ...currentData.products,
          ...newProducts
        ]
        
        this.setData({ categoryData })
      }
    }, 1000)
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