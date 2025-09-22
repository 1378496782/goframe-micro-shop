const api = require('../../utils/api').api;
const constants = require('../../config/constants');

Page({
  data: {
    selectedSpecs: {},
    product: null,
    loading: true
  },
 
  async onLoad(options) {
    if (!options.id) {
      wx.showToast({ title: '商品ID无效', icon: 'none' })
      wx.navigateBack()
      return
    }

    // 直接加载商品详情，不检查登录状态
    await this.loadProductDetail(options.id)
  },
  
  // 加载商品详情
  async loadProductDetail(productId) {
    console.log('开始加载商品详情，商品ID:', productId)
    if (!productId) {
      wx.showToast({
        title: '商品ID无效',
        icon: 'none'
      })
      wx.navigateBack()
      return
    }

    // 移除此处的登录检查，允许未登录用户查看商品详情
    
    this.setData({ loading: true })
    
    try {
      console.log('调用API获取商品详情...')
      const res = await api.getGoodsDetail(productId)
      
      if (!res) {
        throw new Error('获取商品详情失败，请检查网络连接')
      }
      console.log('API响应:', res)
      
      if (res.code === 0) {
        const product = res.data
        // 处理图片URL，使用IMAGE_BASE_URL拼接
        const images = []
        if (product.pic_url) {
          images.push(product.pic_url.startsWith('http') ? product.pic_url : constants.IMAGE_BASE_URL + product.pic_url)
        }
        
        // 如果没有图片，添加默认占位图
        if (images.length === 0) {
          images.push(constants.IMAGE_BASE_URL + 'default-product.png');
        }
        
        // 格式化商品数据
        const formattedProduct = {
          id: product.id,
          name: product.name,
          price: (product.price / 100).toFixed(2),
          originalPrice: ((product.price * 1.2) / 100).toFixed(2), // 模拟原价
          discount: '8.3', // 模拟折扣
          sales: product.sale || 0,
          reviews: Math.floor((product.sale || 0) * 0.3), // 模拟评论数
          stock: product.stock || 0,
          images: images,
          detailInfo: product.detail_info || '',
          tags: product.tags || '',
          brand: product.brand || ''
        }
        
        console.log('格式化后的商品数据:', formattedProduct)
        this.setData({
          product: formattedProduct,
          loading: false
        })
        // 设置页面标题为商品名称
        wx.setNavigationBarTitle({
          title: formattedProduct.name
        })
        console.log('页面数据更新完成')
      } else {
        throw new Error(res.message || '获取商品详情失败')
      }
    } catch (error) {
      console.error('加载商品详情失败:', error)
      wx.showToast({
        title: '加载失败',
        icon: 'none'
      })
      this.setData({ loading: false })
      wx.navigateBack()
    }
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
  async addToCart() {
    const token = wx.getStorageSync('token')
    if (!token) {
      wx.showModal({
        title: '提示',
        content: '请先登录',
        success: (res) => {
          if (res.confirm) {
            wx.navigateTo({
              url: '/pages/login/login'
            })
          }
        }
      })
      return
    }

    try {
      const res = await api.addToCart({
        goods_id: this.data.product.id,
        count: 1
      })

      if (res.code === 0) {
        wx.showToast({
          title: '已加入购物车',
          icon: 'success'
        })
      } else {
        throw new Error(res.message || '加入购物车失败')
      }
    } catch (error) {
      console.error('加入购物车失败:', error)
      wx.showToast({
        title: error.message || '加入购物车失败',
        icon: 'none'
      })
    }
  },

  // 立即购买
  buyNow() {
    wx.navigateTo({
      url: `/pages/order/confirm?productId=${this.data.product.id}`
    })
    console.log('立即购买', this.data.product.id)
  }
})