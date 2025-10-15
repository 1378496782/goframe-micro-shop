// 直接请求搜索API
function requestSearchAPI(params) {
  return new Promise((resolve, reject) => {
    wx.request({
      url: 'http://101.42.249.106:8199/frontend/goods',
      data: params,
      method: 'GET',
      success: (res) => {
        if (res.statusCode === 200 && res.data) {
          resolve(res.data)
        } else {
          reject(new Error('请求失败'))
        }
      },
      fail: (err) => {
        reject(err)
      }
    })
  })
}

Page({
  data: {
    // 搜索参数
    keyword: '',
    brand: '',
    minPrice: '',
    maxPrice: '',
    sort: 'default',
    
    // 商品数据
    goodsList: [],
    loading: false,
    hasMore: true,
    page: 1,
    size: 10,
    total: 0,
    
    // 筛选面板
    showFilter: false
  },

  onLoad(options) {
    if (options.keyword) {
      this.setData({ keyword: options.keyword })
      this.onSearch()
    }
  },

  // 输入框事件
  onKeywordInput(e) {
    this.setData({ keyword: e.detail.value })
  },

  onBrandInput(e) {
    this.setData({ brand: e.detail.value })
  },

  onMinPriceInput(e) {
    this.setData({ minPrice: e.detail.value })
  },

  onMaxPriceInput(e) {
    this.setData({ maxPrice: e.detail.value })
  },

  onSortChange(e) {
    this.setData({ sort: e.detail.value })
    this.onSearch()
  },

  // 搜索操作
  onSearch() {
    if (!this.data.keyword.trim()) {
      wx.showToast({
        title: '请输入搜索关键词',
        icon: 'none'
      })
      return
    }

    this.setData({
      page: 1,
      goodsList: [],
      loading: true,
      hasMore: true
    })

    this.loadGoods(true)
  },

  // 加载更多
  onLoadMore() {
    if (this.data.loading || !this.data.hasMore) return
    
    this.setData({
      page: this.data.page + 1,
      loading: true
    })
    
    this.loadGoods(false)
  },

  // 加载商品数据
  async loadGoods(isRefresh) {
    try {
      const params = {
        keyword: this.data.keyword,
        brand: this.data.brand,
        sort: this.data.sort,
        page: this.data.page,
        size: this.data.size
      }

      // 添加价格筛选
      if (this.data.minPrice) {
        params.min_price = Math.floor(this.data.minPrice * 100) // 转换为分
      }
      if (this.data.maxPrice) {
        params.max_price = Math.floor(this.data.maxPrice * 100) // 转换为分
      }

      const res = await requestSearchAPI(params)
      console.log('搜索响应:', res)
      
      // 检查返回数据
      if (!res || res.code !== 0 || !res.data || !res.data.list) {
        throw new Error(res?.message || 'API返回数据格式不正确')
      }
      
      // 格式化商品数据
      const formattedList = res.data.list.map(item => {
        // 解析图片数据
        let imageUrl = item.pic_url
        if (!imageUrl && item.images) {
          try {
            const imagesObj = JSON.parse(item.images)
            imageUrl = imagesObj.image || ''
          } catch (e) {
            console.warn('解析图片数据失败:', e)
          }
        }
        
        return {
          id: item.id,
          name: item.name,
          price: item.price,
          priceFormatted: (item.price / 100).toFixed(2),
          mainImage: imageUrl || 'https://via.placeholder.com/200x200?text=商品图片',
          highlightName: item.highlight || item.name,
          sale: item.sale || 0,
          stock: item.stock || 0
        }
      })
      
      const newList = isRefresh ? formattedList : [...this.data.goodsList, ...formattedList]
      
      this.setData({
        goodsList: newList,
        total: res.data.total,
        hasMore: newList.length < res.data.total,
        loading: false
      })
      console.log('设置后的商品列表:', newList)
      console.log('商品数量:', newList.length)
 
    } catch (error) {
      console.error('搜索失败:', error)
      console.error('错误详情:', error.message)
      this.setData({ 
        loading: false,
        goodsList: [],
        hasMore: false
      })
      
      wx.showToast({
        title: error.message || '搜索失败，请重试',
        icon: 'none'
      })
    }
  },

  // 切换筛选面板
  toggleFilter() {
    this.setData({ showFilter: !this.data.showFilter })
  },

  // 应用筛选
  applyFilter() {
    this.setData({ showFilter: false })
    this.onSearch()
  },

  // 重置筛选
  resetFilter() {
    this.setData({
      brand: '',
      minPrice: '',
      maxPrice: '',
      sort: 'default'
    })
  },

  // 跳转到商品详情
  goToProductDetail(e) {
    const { id } = e.detail
    wx.navigateTo({
      url: `/pages/product-detail/product-detail?id=${id}`
    })
  },

  // 下拉刷新
  onPullDownRefresh() {
    this.onSearch()
    wx.stopPullDownRefresh()
  },

  // 上拉加载更多
  onReachBottom() {
    this.onLoadMore()
  }
})