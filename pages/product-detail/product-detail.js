const { api } = require('../../utils/api');
const constants = require('../../config/constants');

Page({
  data: {
    currentImageIndex: 0,
    selectedSpecs: {},
    product: null,
    loading: true
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

  async onLoad(options) {
    console.log('商品详情页面加载参数:', options)
    console.log('商品ID:', options.id)
    
    // 检查用户是否登录
    const token = wx.getStorageSync('token')
    console.log('用户token:', token)
    if (!token) {
      wx.showModal({
        title: '提示',
        content: '请先登录',
        success: (res) => {
          if (res.confirm) {
            wx.navigateTo({
              url: '/pages/login/login'
            })
          } else {
            wx.navigateBack()
          }
        }
      })
      return
    }
    
    // 加载商品详情数据
    await this.loadProductDetail(options.id)
  },
  
  // 加载商品详情
  async loadProductDetail(product极速版Id) {
    console.log('开始加载商品详情，商品ID:', productId)
    if (!productId) {
      wx.showToast({
        title: '商品ID无效',
        icon: 'none'
      })
      wx.navigateBack()
      return
    }
    
    this.setData({ loading: true })
    
    try {
      console.log('调用API获取商品详情...')
      const res = await api.getGoodsDetail(productId)
      console.log('API响应:', res)
      
      if (res.code === 0) {
        const product = res.data
        // 处理图片URL
        const images = []
        if (product.pic_url) {
          images.push(constants.IMAGE_BASE_URL + product.pic_url)
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