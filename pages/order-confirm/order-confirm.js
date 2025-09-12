Page({
  data: {
    orderItems: [],
    totalPrice: 0,
    couponPrice: 0,
    actualPrice: 0,
    remark: '',
    loading: false,
    coupons: [],
    selectedCoupon: null,
    showCouponPopup: false
  },

  onLoad(options) {
    let selectedItems = []
    
    // 从URL参数获取数据
    if (options.selectedItems) {
      try {
        selectedItems = JSON.parse(decodeURIComponent(options.selectedItems))
      } catch (error) {
        console.error('解析URL参数失败:', error)
      }
    }
    
    // 如果URL参数没有数据，使用全局数据
    if (selectedItems.length === 0) {
      const app = getApp()
      selectedItems = app.globalData.selectedCartItems || []
    }
    
    console.log('从URL参数获取的selectedItems:', selectedItems)
    
    // 获取全局应用实例
    const app = getApp()
    console.log('从全局数据获取的selectedCartItems:', app.globalData.selectedCartItems)
    
    if (selectedItems.length === 0) {
      wx.showToast({
        title: '请先选择商品',
        icon: 'none'
      })
      setTimeout(() => wx.navigateBack(), 1500)
      return
    }
    
    console.log('最终使用的商品数据:', selectedItems)
    console.log('第一个商品的字段:', Object.keys(selectedItems[0]))
    console.log('第一个商品的详细数据:', selectedItems[0])

    // 处理图片URL，添加基础URL - 兼容不同字段名
    const constants = require('../../config/constants')
    const processedItems = selectedItems.map(item => ({
      ...item,
      // 兼容不同字段名
      name: item.name || item.goods_name,
      spec: item.spec || item.goods_brand || '默认规格',
      price: item.price || item.goods_price || 0,
      quantity: item.quantity || item.count || 1,
      // 处理图片URL
      goods_pic_url: item.image || `${constants.IMAGE_BASE_URL}${item.goods_pic_url || ''}`
    }))

    // 计算总价 - 兼容不同字段名
    const totalPrice = processedItems.reduce((sum, item) => {
      const price = item.price || item.goods_price || 0
      const quantity = item.quantity || item.count || 1
      return sum + (price * quantity)
    }, 0)
    
    const actualPrice = totalPrice - this.data.couponPrice
    
    // 设置页面数据
    this.setData({
      orderItems: processedItems,
      totalPrice: totalPrice,
      actualPrice: actualPrice,
      coupons: [
        { id: 1, amount: 500, amountYuan: 5.00, desc: "满50减5" },
        { id: 2, amount: 1000, amountYuan: 10.00, desc: "满100减10" },
        { id: 3, amount: 2000, amountYuan: 20.00, desc: "满200减20" }
      ]
    })
  },

  // 打开优惠券选择弹窗
  openCouponPopup() {
    this.setData({ showCouponPopup: true })
  },

  // 关闭优惠券选择弹窗
  closeCouponPopup() {
    this.setData({ showCouponPopup: false })
  },

  // 阻止事件冒泡
  stopPropagation() {
    // 空方法，用于阻止事件冒泡
  },

  // 选择优惠券
  selectCoupon(e) {
    console.log('优惠券被点击，事件对象:', e)
    const coupon = e.currentTarget.dataset.coupon
    console.log('选中的优惠券数据:', coupon)
    
    if (!coupon) {
      console.error('优惠券数据为空')
      return
    }
    
    const couponPrice = coupon.amountYuan // 直接使用元为单位
    const actualPrice = this.data.totalPrice - couponPrice
    
    this.setData({
      selectedCoupon: coupon,
      couponPrice: couponPrice,
      actualPrice: actualPrice,
      showCouponPopup: false
    })
    
    console.log('选择优惠券成功:', coupon, '优惠金额:', couponPrice)
    console.log('更新后的价格 - 总价:', this.data.totalPrice, '优惠:', couponPrice, '实付:', actualPrice)
  },

  // 取消选择优惠券
  removeCoupon() {
    this.setData({
      selectedCoupon: null,
      couponPrice: 0,
      actualPrice: this.data.totalPrice
    })
  },

  // 处理备注输入
  onRemarkInput(e) {
    this.setData({
      remark: e.detail.value
    })
  },

  // 提交订单
  submitOrder() {
    wx.showToast({
      title: '订单提交成功',
      icon: 'success'
    })
  }
})