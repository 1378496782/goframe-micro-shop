const { api } = require('../../utils/api');

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
      
      if (res.code === 0) {
        // 格式化商品数据
        const formattedProducts = res.data.list.map(item => {
          // 直接使用 pic_url 字段作为主图
          const mainImage = item.pic_url || 'https://via.placeholder.com/200x200?text=商品图片';
          
          return {
            ...item,
            priceFormatted: (item.price / 100).toFixed(2),
            mainImage: mainImage
          };
        });
        
        const newProducts = isRefresh ? formattedProducts : [...this.data.products, ...formattedProducts];
        const hasMore = newProducts.length < res.data.total;
        
        this.setData({
          products: newProducts,
          page: currentPage + 1,
          hasMore: hasMore,
          loading: false,
          refreshing: false
        });
      }
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
  }
});