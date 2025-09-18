const { api } = require('../../utils/api');

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
    showCouponPopup: false,
    loadingCoupons: false
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
      actualPrice: actualPrice
    })
    
    // 加载用户优惠券
    this.loadUserCoupons();
  },

  // 加载用户优惠券
  async loadUserCoupons() {
    this.setData({ loadingCoupons: true });
    
    try {
      const res = await api.getUserCoupons({ page: 1, size: 10 });
      
      if (res.code === 0) {
        // 格式化优惠券数据
        const formattedCoupons = res.data.list
          .filter(coupon => coupon.status === 0) // 只显示未使用的优惠券
          .map(coupon => {
            let desc = '';
            let amountYuan = coupon.amount / 100; // 分转元
            
            // 根据coupon_id设置优惠券描述
            if (coupon.coupon_id === 1) {
              desc = `新人券 - 减${amountYuan.toFixed(2)}元`;
            } else if (coupon.coupon_id === 2) {
              desc = `满减券 - 减${amountYuan.toFixed(2)}元`;
            } else {
              desc = `优惠券 - 减${amountYuan.toFixed(2)}元`;
            }
            
            return {
              id: coupon.id,
              coupon_id: coupon.coupon_id,
              amount: coupon.amount,
              amountYuan: amountYuan,
              desc: desc
            };
          });
        
        this.setData({ coupons: formattedCoupons });
      }
    } catch (error) {
      console.error('加载优惠券失败:', error);
      wx.showToast({
        title: '加载优惠券失败',
        icon: 'none'
      });
    } finally {
      this.setData({ loadingCoupons: false });
    }
  },

  // 打开优惠券选择弹窗
  openCouponPopup() {
    if (this.data.coupons.length === 0) {
      wx.showToast({
        title: '暂无可用优惠券',
        icon: 'none'
      });
      return;
    }
    this.setData({ showCouponPopup: true });
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
    
    let couponPrice = coupon.amountYuan; // 直接使用元为单位
    let actualPrice = this.data.totalPrice - couponPrice;
    
    // 确保实际价格不会低于0
    if (actualPrice < 0) {
      actualPrice = 0;
      // 如果优惠券金额超过商品总价，调整实际抵扣金额为商品总价
      couponPrice = this.data.totalPrice;
    }
    
    this.setData({
      selectedCoupon: coupon, // 保持优惠券原始数据显示
      couponPrice: couponPrice, // 使用调整后的实际抵扣金额
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
  async submitOrder() {
    if (this.data.loading) return;
    
    this.setData({ loading: true });
    
    try {
      // 构建请求参数
      const orderData = {
        price: this.data.totalPrice,
        coupon_price: this.data.couponPrice,
        actual_price: this.data.actualPrice,
        remark: this.data.remark,
        order_goods_info: this.data.orderItems.map(item => ({
          goods_id: item.goods_id || item.id || 0,
          count: item.quantity || item.count || 1,
          price: item.price || item.goods_price || 0,
          coupon_price: this.data.couponPrice, // 使用订单级别的优惠金额
          actual_price: (item.price || item.goods_price || 0) * (item.quantity || item.count || 1) - this.data.couponPrice
        }))
      };
      
      console.log('提交订单参数:', orderData);
      
      // 调用提交订单接口
      const res = await api.submitOrder(orderData);
      
      if (res.code === 0 && res.data && res.data.id) {
        // 订单创建成功
        wx.showToast({
          title: '订单创建成功',
          icon: 'success',
          duration: 2000
        });
        
        // 2秒后返回首页
        setTimeout(() => {
          wx.switchTab({
            url: '/pages/index/index'
          });
        }, 2000);
      } else {
        // 订单创建失败
        throw new Error(res.msg || '订单创建失败');
      }
    } catch (error) {
      console.error('提交订单失败:', error);
      wx.showToast({
        title: error.message || '提交订单失败',
        icon: 'none'
      });
    } finally {
      this.setData({ loading: false });
    }
  }
})