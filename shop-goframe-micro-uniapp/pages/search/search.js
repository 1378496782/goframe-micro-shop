// 使用项目封装的API方法
const { searchAPI } = require('../../utils/request')
const { CONSTANTS } = require('../../config/index')

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
    showFilter: false,
    
    // 搜索增强功能
    showSuggestions: false,
    searchHistory: [],
    hotKeywords: [
      { keyword: '手机', isHot: true },
      { keyword: '笔记本电脑', isHot: true },
      { keyword: '耳机', isHot: false },
      { keyword: '智能手表', isHot: false },
      { keyword: '平板电脑', isHot: true },
      { keyword: '相机', isHot: false }
    ],
    suggestions: []
  },

  onLoad(options) {
    // 加载搜索历史
    this.loadSearchHistory()
    
    if (options.keyword) {
      this.setData({ keyword: options.keyword })
      this.onSearch()
    }
  },

  // 搜索框聚焦事件
  onSearchFocus() {
    this.setData({ showSuggestions: true })
    this.loadSearchSuggestions()
  },

  // 搜索框失焦事件
  onSearchBlur() {
    setTimeout(() => {
      this.setData({ showSuggestions: false })
    }, 200)
  },

  // 加载搜索历史
  loadSearchHistory() {
    try {
      const history = wx.getStorageSync('searchHistory') || []
      this.setData({ searchHistory: history.slice(0, 10) }) // 最多显示10条
    } catch (error) {
      console.error('加载搜索历史失败:', error)
    }
  },

  // 保存搜索历史
  saveSearchHistory(keyword) {
    if (!keyword.trim()) return
    
    try {
      let history = wx.getStorageSync('searchHistory') || []
      // 移除重复项
      history = history.filter(item => item !== keyword)
      // 添加到开头
      history.unshift(keyword)
      // 限制最多保存20条
      history = history.slice(0, 20)
      wx.setStorageSync('searchHistory', history)
      this.setData({ searchHistory: history.slice(0, 10) })
    } catch (error) {
      console.error('保存搜索历史失败:', error)
    }
  },

  // 清空搜索历史
  clearSearchHistory() {
    wx.showModal({
      title: '提示',
      content: '确定要清空搜索历史吗？',
      success: (res) => {
        if (res.confirm) {
          try {
            wx.removeStorageSync('searchHistory')
            this.setData({ searchHistory: [] })
            wx.showToast({ title: '已清空搜索历史', icon: 'success' })
          } catch (error) {
            console.error('清空搜索历史失败:', error)
          }
        }
      }
    })
  },

  // 删除单条搜索历史
  deleteHistoryItem(e) {
    const index = e.currentTarget.dataset.index
    try {
      let history = wx.getStorageSync('searchHistory') || []
      history.splice(index, 1)
      wx.setStorageSync('searchHistory', history)
      this.setData({ searchHistory: history.slice(0, 10) })
    } catch (error) {
      console.error('删除搜索历史失败:', error)
    }
  },

  // 点击搜索历史
  onHistoryTap(e) {
    const keyword = e.currentTarget.dataset.keyword
    this.setData({ keyword })
    this.onSearch()
  },

  // 点击热门搜索
  onHotSearchTap(e) {
    const keyword = e.currentTarget.dataset.keyword
    this.setData({ keyword })
    this.onSearch()
  },

  // 点击搜索建议
  onSuggestionTap(e) {
    const suggestion = e.currentTarget.dataset.suggestion
    this.setData({ keyword: suggestion })
    this.onSearch()
  },

  // 加载搜索建议
  loadSearchSuggestions() {
    if (!this.data.keyword.trim()) {
      this.setData({ suggestions: [] })
      return
    }

    // 模拟搜索建议数据
    const mockSuggestions = [
      `${this.data.keyword} 手机`,
      `${this.data.keyword} 电脑`,
      `${this.data.keyword} 配件`,
      `${this.data.keyword} 新款`,
      `${this.data.keyword} 优惠`
    ].filter(item => item !== this.data.keyword).slice(0, 5)

    this.setData({ suggestions: mockSuggestions })
  },

  // 输入框事件
  onKeywordInput(e) {
    this.setData({ keyword: e.detail.value })
    // 实时加载搜索建议
    this.loadSearchSuggestions()
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

    // 保存搜索历史
    this.saveSearchHistory(this.data.keyword)

    this.setData({
      page: 1,
      goodsList: [],
      loading: true,
      hasMore: true,
      showSuggestions: false
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

      const res = await searchAPI.searchGoods(params)
      console.log('搜索响应:', res)
      
      // 检查返回数据（request函数已经返回data字段）
      if (!res || !res.list) {
        throw new Error('API返回数据格式不正确')
      }
      
      // 格式化商品数据
      const formattedList = res.list.map(item => {
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
        
        // 拼接完整的图片URL
        let fullImageUrl = imageUrl
        if (imageUrl && !imageUrl.startsWith('http')) {
          fullImageUrl = CONSTANTS.IMAGE_BASE_URL + imageUrl
        }
        
        return {
          id: item.id,
          name: item.name,
          price: item.price,
          priceFormatted: (item.price / 100).toFixed(2),
          mainImage: fullImageUrl || 'https://via.placeholder.com/200x200?text=商品图片',
          highlightName: item.highlight || item.name,
          sale: item.sale || 0,
          stock: item.stock || 0
        }
      })
      
      const newList = isRefresh ? formattedList : [...this.data.goodsList, ...formattedList]
      
      this.setData({
        goodsList: newList,
        total: res.total,
        hasMore: newList.length < res.total,
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