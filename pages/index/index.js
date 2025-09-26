const { api } = require('../../utils/api');
const constants = require('../../config/constants');

Page({
  data: {
    banners: [],
    categories: [],
    products: [],
    searchValue: '',
    loading: false
  },

  onLoad() {
    console.log('首页加载')
    this.loadHomeData()
  },

  async loadHomeData() {
    this.setData({ loading: true })
    
    try {
      // 并行加载轮播图和商品数据
      const [bannersRes, productsRes] = await Promise.all([
        api.getBanners(),
        api.getGoodsList({ page: 1, size: 10, is_hot: 1 })
      ])
      
      if (bannersRes.code === 0 && bannersRes.data?.list) {
        this.setData({
          banners: bannersRes.data.list?.map(item => ({
            id: item.id,
            image: constants.IMAGE_BASE_URL + item.pic_url,
            url: item.link
          }))
        })
      }
       

      if (productsRes.code === 0 && productsRes.data?.list) {
        // 格式化商品数据：价格转换和图片提取
        const formattedProducts = productsRes.data.list?.map(item => {
          // 处理图片URL，直接使用pic_url字段（已经是完整URL）
          let mainImage = '';
          if (item.pic_url) {
            mainImage = item.pic_url;
          }
          
          return {
            ...item,
            priceFormatted: (item.price / 100).toFixed(2), // 价格从分转换为元
            mainImage: mainImage ? (mainImage.startsWith('http') ? mainImage : constants.IMAGE_BASE_URL + mainImage) : 'https://via.placeholder.com/200x200?text=商品图片'
          }
        })
        
        this.setData({
          products: formattedProducts
        })
      }


    } catch (error) {
      console.error('加载首页数据失败:', error)
      wx.showToast({
        title: '数据加载失败',
        icon: 'none'
      })
    } finally {
      this.setData({ loading: false })
    }
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

  onCategoryClick(e) {
    const categoryId = e.currentTarget.dataset.id
    wx.navigateTo({
      url: `/pages/category/category?id=${categoryId}`
    })
  },

  onViewMore() {
    wx.navigateTo({
      url: '/pages/more-goods/more-goods'
    })
  },
  
  // 轮播图点击事件
  onBannerClick(e) {
    const { url } = e.currentTarget.dataset
    if (url) {
      wx.navigateTo({
        url: `/pages/webview/webview?url=${encodeURIComponent(url)}`
      })
    }
  },

  onProductClick(e) {
    const productId = e.currentTarget.dataset.id
    wx.navigateTo({
      url: `/pages/product-detail/product-detail?id=${productId}`
    })
  },


  onPullDownRefresh() {
    console.log('下拉刷新')
    this.loadHomeData().then(() => {
      wx.stopPullDownRefresh()
    })
  }
})