const { api } = require('../../utils/api');
const { CONSTANTS } = require('../../config/index');

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
      
      console.log('轮播图响应:', bannersRes)
      console.log('商品列表响应:', productsRes)
      
      if (bannersRes.code === 0 && bannersRes.data?.list) {
        this.setData({
          banners: bannersRes.data.list?.map(item => ({
            id: item.id,
            image: item.pic_url ? (item.pic_url.startsWith('http') ? item.pic_url : CONSTANTS.IMAGE_BASE_URL + item.pic_url) : 'https://via.placeholder.com/300x150?text=轮播图',
            url: item.link
          }))
        })
      }
       

      if (productsRes.code === 0 && productsRes.data?.list) {
        // 格式化商品数据：价格转换和图片提取
        const formattedProducts = productsRes.data.list?.map(item => {
          // 处理图片URL，优先使用pic_url，如果没有则尝试从images字段解析
          let mainImage = '';
          if (item.pic_url) {
            mainImage = item.pic_url;
          } else if (item.images) {
            try {
              const imagesObj = JSON.parse(item.images);
              if (imagesObj.image) {
                mainImage = imagesObj.image;
              }
            } catch (e) {
              console.log('解析images字段失败:', e)
            }
          }
          
          // 确保图片URL是完整的
          if (mainImage && !mainImage.startsWith('http')) {
            mainImage = CONSTANTS.IMAGE_BASE_URL + mainImage;
          }
          
          return {
            ...item,
            priceFormatted: (item.price / 100).toFixed(2), // 价格从分转换为元
            mainImage: mainImage || 'https://via.placeholder.com/200x200?text=商品图片'
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
  },

  // 分享给朋友
  onShareAppMessage() {
    return {
      title: '发现优质商品，快来选购吧！',
      path: '/pages/index/index',
      imageUrl: this.data.banners.length > 0 ? this.data.banners[0].image : ''
    };
  },

  // 分享到朋友圈
  onShareTimeline() {
    return {
      title: '优质电商平台，海量商品任你选',
      imageUrl: this.data.banners.length > 0 ? this.data.banners[0].image : ''
    };
  }
})