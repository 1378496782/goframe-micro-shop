Page({
  data: {
    cartItems: [
      {
        id: 1,
        name: '高品质智能手机 8GB+256GB',
        spec: '黑色, 8GB+256GB',
        price: '2999.00',
        originalPrice: '3999.00',
        image: 'https://via.placeholder.com/100x100/19aecc/ffffff?text=手机',
        quantity: 1,
        selected: true
      },
      {
        id: 2,
        name: '无线蓝牙耳机 降噪版',
        spec: '白色',
        price: '399.00',
        originalPrice: '499.00',
        image: 'https://via.placeholder.com/100x100/19aecc/ffffff?text=耳机',
        quantity: 2,
        selected: false
      }
    ],
    allSelected: false
  },

  onLoad() {
    this.calculateTotal()
  },

  onShow() {
    // 更新购物车数量
    this.getTabBar().setData({
      cartCount: this.data.cartItems.reduce((total, item) => total + item.quantity, 0)
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
    this.getTabBar().setData({ cartCount })
  },

  // 去结算
  goToSettlement() {
    const selectedItems = this.data.cartItems.filter(item => item.selected)
    if (selectedItems.length === 0) {
      wx.showToast({
        title: '请选择要结算的商品',
        icon: 'none'
      })
      return
    }

    wx.navigateTo({
      url: '/pages/order/confirm'
    })
  },

  // 去首页
  goToIndex() {
    wx.switchTab({
      url: '/pages/index/index'
    })
  }
})