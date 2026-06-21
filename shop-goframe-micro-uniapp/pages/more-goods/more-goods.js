const { api } = require('../../utils/api');
const { CONSTANTS } = require('../../config/index');

Page({
  data: {
    products: [],
    loading: false,
    refreshing: false,
    page: 1,
    size: 10,
    hasMore: true
  },

  onLoad(options) {
    console.log('更多商品页面加载');
    this.loadGoodsData(true);
  },

  async loadGoodsData(isRefresh = false) {
    if (this.data.loading && !isRefresh) return;
    
    const currentPage = isRefresh ? 1 : this.data.page;
    
    this.setData({ 
      loading: true,
      refreshing: isRefresh
    });
    
    try {
      const res = await api.getGoodsList({ 
        page: currentPage, 
        size: this.data.size 
      });
      
      // 检查API返回的数据结构
      console.log('API返回数据:', res);
      
      // 根据API实际返回的数据结构获取商品列表
      // mock数据中返回的是 res.data.list，真实API可能返回 res.data 或 res.list
      const productList = (res.data && res.data.list) || res.list || res.data || res.products || [];
      const formattedProducts = productList.map(item => {
        // 处理图片URL，确保是完整URL
        let mainImage = item.pic_url || item.image || item.mainImage || 'https://via.placeholder.com/200x200?text=商品图片';
        if (mainImage && !mainImage.startsWith('http')) {
          mainImage = CONSTANTS.IMAGE_BASE_URL + mainImage;
        }
        
        return {
          ...item,
          priceFormatted: (item.price / 100).toFixed(2),
          mainImage: mainImage
        };
      });
      
      const newProducts = isRefresh ? formattedProducts : [...this.data.products, ...formattedProducts];
      // 检查是否有更多数据，如果当前返回的数据数量小于请求的size，说明没有更多数据了
      const hasMore = formattedProducts.length >= this.data.size;
      
      this.setData({
        products: newProducts,
        page: currentPage + 1,
        hasMore: hasMore,
        loading: false,
        refreshing: false
      });
    } catch (error) {
      console.error('加载商品数据失败:', error);
      wx.showToast({
        title: '加载失败',
        icon: 'none'
      });
      this.setData({ 
        loading: false,
        refreshing: false 
      });
    }
  },

  onProductClick(e) {
    const productId = e.currentTarget.dataset.id;
    wx.navigateTo({
      url: `/pages/product-detail/product-detail?id=${productId}`
    });
  },

  onPullDownRefresh() {
    console.log('下拉刷新');
    this.loadGoodsData(true).then(() => {
      wx.stopPullDownRefresh();
    });
  },

  onReachBottom() {
    console.log('上拉加载更多');
    if (this.data.hasMore && !this.data.loading) {
      this.loadGoodsData(false);
    }
  },

  onShow() {
    console.log('更多商品页面显示');
  },

  // 分享给朋友
  onShareAppMessage() {
    return {
      title: '发现更多优质商品，快来选购吧！',
      path: '/pages/more-goods/more-goods',
      imageUrl: this.data.products.length > 0 ? this.data.products[0].mainImage : ''
    };
  },

  // 分享到朋友圈
  onShareTimeline() {
    return {
      title: '海量商品任你选，快来发现更多好物',
      imageUrl: this.data.products.length > 0 ? this.data.products[0].mainImage : ''
    };
  }
});