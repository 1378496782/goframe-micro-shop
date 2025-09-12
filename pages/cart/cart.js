Page({
  data: {
    cartItems: [],
    allSelected: false,
    page: 1,
    size: 10,
    total: 0,
    loading: false,
    hasMore: true
  },

  onLoad() {
    this.loadCartData()
  },

  onPullDownRefresh() {
    // 下拉刷新
    this.setData({
      page: 1,
      hasMore: true
    })
    this.loadCartData(true)
  },

  onReachBottom() {
    // 上拉加载更多
    if (this.data.hasMore && !this.data.loading) {
      this.setData({
        page: this.data.page + 1
      })
      this.loadCartData()
    }
  },

  onShow() {
    // 页面显示时如果购物车为空，则加载数据
    if (this.data.cartItems.length === 0) {
      this.setData({
        page: 1,
        hasMore: true
      })
      this.loadCartData()
    }
  },

  // 加载购物车数据
  loadCartData(isRefresh = false) {
    if (this.data.loading) return
    
    this.setData({ loading: true })
    
    const app = getApp()
    const { API, config } = app.globalData
    const { request } = require('../../utils/request')
    const constants = require('../../config/constants')
    
    console.log('正在请求购物车数据:', {
      url: API.CART_LIST,
      page: this.data.page,
      size: this.data.size
    })
    
    request({
      url: API.CART_LIST,
      method: 'GET',
      data: {
        page: this.data.page,
        size: this.data.size
      }
    }).then(cartData => {
      console.log('购物车数据:', cartData)
      
      // 合并相同商品的数量
      const itemMap = new Map()
      
      cartData.list.forEach(item => {
        const key = `${item.goods_id}_${item.goods_brand || 'default'}`
        if (itemMap.has(key)) {
          // 相同商品，合并数量
          const existingItem = itemMap.get(key)
          existingItem.count += item.count
          existingItem.id = item.id // 保留最新的ID
        } else {
          // 新商品，添加到map
          itemMap.set(key, {
            id: item.id,
            goods_id: item.goods_id,
            goods_name: item.goods_name,
            goods_brand: item.goods_brand,
            goods_price: item.goods_price,
            goods_pic_url: item.goods_pic_url,
            count: item.count,
            goods_stock: item.goods_stock,
            goods_sale: item.goods_sale
          })
        }
      })
      
      // 转换为前端需要的格式
      const newItems = Array.from(itemMap.values()).map(item => ({
        id: item.id,
        goods_id: item.goods_id,
        name: item.goods_name,
        spec: item.goods_brand || '默认规格',
        price: item.goods_price / 100, // 分转元
        originalPrice: (item.goods_price * 1.2) / 100, // 计算原价
        image: `${constants.IMAGE_BASE_URL}${item.goods_pic_url}`,
        quantity: item.count, // 合并后的数量
        selected: false,
        stock: item.goods_stock,
        sale: item.goods_sale
      }))
      
      console.log('转换后的商品列表:', newItems)
      
      if (isRefresh) {
        this.setData({
          cartItems: newItems,
          total: cartData.total,
          hasMore: this.data.page * this.data.size < cartData.total
        })
      } else {
        this.setData({
          cartItems: [...this.data.cartItems, ...newItems],
          total: cartData.total,
          hasMore: this.data.page * this.data.size < cartData.total
        })
      }
      
      this.calculateTotal()
      this.updateCartCount()
    }).catch(err => {
      console.log('购物车请求失败:', err)
      wx.showToast({
        title: err.message || '加载失败',
        icon: 'none'
      })
    }).finally(() => {
      this.setData({ loading: false })
      if (isRefresh) {
        wx.stopPullDownRefresh()
      }
    })
  },

  // 计算总价和选中数量
  calculateTotal() {
    const selectedItems = this.data.cartItems.filter(item => item.selected)
    const totalPrice = selectedItems.reduce((total, item) => total + (item.price * item.quantity), 0)
    const selectedCount = selectedItems.reduce((total, item) => total + item.quantity, 0)
    const allSelected = this.data.cartItems.length > 0 && this.data.cartItems.every(item => item.selected)

    this.setData({
      totalPrice: totalPrice.toFixed(2),
      selectedCount,
      allSelected
    })
  },

  // 切换单个商品选中状态
  toggleSelect(e) {
    const index = e.currentTarget.dataset.index
    const key = `cartItems[${index}].selected`
    this.setData({
      [key]: !this.data.cartItems[index].selected
    }, () => {
      this.calculateTotal()
    })
  },

  // 全选/取消全选
  toggleSelectAll() {
    const allSelected = !this.data.allSelected
    const cartItems = this.data.cartItems.map(item => ({
      ...item,
      selected: allSelected
    }))

    this.setData({
      cartItems,
      allSelected
    }, () => {
      this.calculateTotal()
    })
  },

  // 增加数量
  increaseQuantity(e) {
    const index = e.currentTarget.dataset.index
    const key = `cartItems[${index}].quantity`
    this.setData({
      [key]: this.data.cartItems[index].quantity + 1
    }, () => {
      this.calculateTotal()
      this.updateCartCount()
    })
  },

  // 减少数量
  decreaseQuantity(e) {
    const index = e.currentTarget.dataset.index
    if (this.data.cartItems[index].quantity > 1) {
      const key = `cartItems[${index}].quantity`
      this.setData({
        [key]: this.data.cartItems[index].quantity - 1
      }, () => {
        this.calculateTotal()
        this.updateCartCount()
      })
    }
  },

  // 删除商品
  deleteItem(e) {
    const index = e.currentTarget.dataset.index
    wx.showModal({
      title: '提示',
      content: '确定要删除这个商品吗？',
      success: (res) => {
        if (res.confirm) {
          const cartItems = this.data.cartItems.filter((_, i) => i !== index)
          this.setData({ cartItems }, () => {
            this.calculateTotal()
            this.updateCartCount()
          })
        }
      }
    })
  },

  // 更新购物车数量
  updateCartCount() {
    const cartCount = this.data.cartItems.reduce((total, item) => total + item.quantity, 0)
    const tabBar = this.getTabBar()
    if (tabBar) {
      tabBar.setData({ cartCount })
    }
  },

  // 去结算
  goToSettlement() {
    console.log('去结算按钮被点击')
    const selectedItems = this.data.cartItems.filter(item => item.selected)
    console.log('选中的商品:', selectedItems)
    
    if (selectedItems.length === 0) {
      console.log('没有选中商品，显示提示')
      wx.showToast({
        title: '请选择要结算的商品',
        icon: 'none'
      })
      return
    }

    // 保存选中商品到全局数据，供订单确认页面使用
    console.log('保存选中商品到全局数据:', selectedItems)
    const app = getApp()
    app.globalData.selectedCartItems = selectedItems
    console.log('全局数据已设置:', app.globalData.selectedCartItems)

    // 同时使用URL参数传递数据，确保数据不丢失
    const queryParams = `?selectedItems=${encodeURIComponent(JSON.stringify(selectedItems))}`
    
    console.log('跳转到订单确认页面，URL:', `/pages/order-confirm/order-confirm${queryParams}`)
    wx.navigateTo({
      url: `/pages/order-confirm/order-confirm${queryParams}`,
      success: (res) => {
        console.log('页面跳转成功:', res)
      },
      fail: (err) => {
        console.error('页面跳转失败:', err)
        console.error('错误详情:', err)
      }
    })
  },

  // 去首页
  goToIndex() {
    wx.switchTab({
      url: '/pages/index/index'
    })
  }
})