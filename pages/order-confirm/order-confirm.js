const { api } = require('../../utils/api');
const { config } = require('../../utils/env');

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
    console.log('当前用户token:', wx.getStorageSync('token'))
    let selectedItems = []
    
    // 场景1：从商品详情页直接购买(单商品)
    if (options.productId) {
      selectedItems = [{
        id: options.productId,
        quantity: options.quantity || 1,
        price: options.price || 0,
        name: options.name || '',
        image: options.image || ''
      }]
    } 
    // 场景2：从购物车多商品下单
    else if (options.selectedItems) {
      try {
        selectedItems = JSON.parse(decodeURIComponent(options.selectedItems))
      } catch (error) {
        console.error('解析URL参数失败:', error)
      }
    }
    
    // 场景3：使用全局数据(兼容旧逻辑)
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
    const { CONSTANTS } = require('../../config/index')
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
    
    // 如果点击的是已经选中的优惠券，则取消选中
    if (this.data.selectedCoupon && this.data.selectedCoupon.id === coupon.id) {
      console.log('取消选中优惠券')
      this.removeCoupon()
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
    const value = e.detail.value
    let error = ''
    
    // 实时校验
    if (value.trim().length === 0) {
      error = '备注信息不能为空'
    } else if (value.length < 2) {
      error = '备注信息至少需要2个字符'
    } else if (!/^[\u4e00-\u9fa5a-zA-Z0-9_-]+$/.test(value.trim())) {
      error = '只能包含中文、英文、数字、下划线和减号'
    }
    
    this.setData({
      remark: value,
      remarkError: error
    })
  },

  // 备注输入框失去焦点时的校验
  onRemarkBlur(e) {
    const value = e.detail.value.trim()
    let error = ''
    
    if (value.length === 0) {
      error = '备注信息不能为空，请填写微信号或手机号'
    } else if (value.length < 2) {
      error = '备注信息至少需要2个字符'
    } else if (!/^[\u4e00-\u9fa5a-zA-Z0-9_-]+$/.test(value)) {
      error = '只能包含中文、英文、数字、下划线和减号'
    }
    
    this.setData({
      remarkError: error
    })
  },

  // 提交订单
  async submitOrder() {
    if (this.data.loading) return;
    
    // 必填校验
    if (!this.data.remark || this.data.remark.trim().length === 0) {
      this.setData({
        remarkError: '备注信息不能为空，请填写微信号或手机号'
      });
      wx.showToast({
        title: '请填写备注信息',
        icon: 'none'
      });
      return;
    }
    
    // 格式校验
    const remark = this.data.remark.trim();
    if (remark.length < 2) {
      this.setData({
        remarkError: '备注信息至少需要2个字符'
      });
      wx.showToast({
        title: '备注信息太短',
        icon: 'none'
      });
      return;
    }
    
    if (!/^[\u4e00-\u9fa5a-zA-Z0-9_-]+$/.test(remark)) {
      this.setData({
        remarkError: '只能包含中文、英文、数字、下划线和减号'
      });
      wx.showToast({
        title: '备注信息格式不正确',
        icon: 'none'
      });
      return;
    }
    
    this.setData({ loading: true });
    
    try {
      // 构建请求参数
      const orderData = {
        price: this.data.totalPrice,
        coupon_price: this.data.couponPrice,
        actual_price: this.data.actualPrice,
        remark: this.data.remark,
        coupon_id: this.data.selectedCoupon ? this.data.selectedCoupon.coupon_id : null, // 添加优惠券ID
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
        // 订单创建成功，获取订单编号
        const orderId = res.data.id;
        const orderNumber = res.data.number;
        
        console.log('订单创建成功，订单ID:', orderId, '订单编号:', orderNumber);
        
        // 检查实际支付金额是否为0
        if (this.data.actualPrice === 0) {
          // 金额为0，直接跳转到订单列表页面
          wx.showToast({
            title: '订单创建成功',
            icon: 'success',
            duration: 2000
          });
          
          // 2秒后跳转到订单列表页面
          setTimeout(() => {
            wx.navigateTo({
              url: '/pages/order-list/order-list'
            });
          }, 2000);
        } else {
          // 金额不为0，调用微信支付接口
          await this.requestWxPayment(orderNumber);
        }
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
  },
  
  // 请求微信支付
  async requestWxPayment(orderNumber) {
    try {
      // 获取用户openId
      const openId = wx.getStorageSync('openId');
      if (!openId) {
        console.error('未找到openId，尝试重新获取');
        
        // 如果没有openId，提示用户重新登录
        wx.showModal({
          title: '登录信息不完整',
          content: '需要重新登录以完成支付',
          confirmText: '去登录',
          success: (res) => {
            if (res.confirm) {
              // 跳转到用户页面进行登录
              wx.switchTab({
                url: '/pages/user/user'
              });
            }
          }
        });
        throw new Error('未获取到用户openId，请重新登录');
      }
      
      // 构建支付请求参数
      const paymentData = {
        openId: openId,
        amount: Math.round(this.data.actualPrice * 100), // 转换为分
        total_price: Math.round(this.data.totalPrice * 100), // 总价（分）
        coupon_price: Math.round(this.data.couponPrice * 100), // 优惠价（分）
        actual_price: Math.round(this.data.actualPrice * 100), // 实际金额（分）
        number: orderNumber
      };
      
      console.log('发起支付请求参数:', paymentData);
      
      // 调用支付接口
      const paymentRes = await this.callPaymentAPI(paymentData);
      
      if (paymentRes.code === 0 && paymentRes.data) {
        // 调起微信支付
        await this.launchWxPayment(paymentRes.data);
        
        // 支付成功后跳转
        wx.showToast({
          title: '支付成功',
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
        throw new Error(paymentRes.msg || '获取支付参数失败');
      }
    } catch (error) {
      console.error('支付请求失败:', error);
      wx.showToast({
        title: error.message || '支付失败',
        icon: 'none'
      });
    }
  },
  
  // 调用支付API
  callPaymentAPI(paymentData) {
    return new Promise((resolve, reject) => {
      // 获取token
      const token = wx.getStorageSync('token') || '';
      
      // 设置请求头
      const header = {};
      if (token) {
        header['Authorization'] = `Bearer ${token}`;
      }
      
      wx.request({
        url: `${config.BASE_URL}/frontend/payment`,
        data: paymentData,
        method: 'POST',
        header: header,
        success: (res) => {
          resolve(res.data);
        },
        fail: (err) => {
          reject(err);
        }
      });
    });
  },
  
  // 调起微信支付
  launchWxPayment(payParams) {
    return new Promise((resolve, reject) => {
      wx.requestPayment({
        timeStamp: payParams.timeStamp,
        nonceStr: payParams.nonceStr,
        package: payParams.package,
        signType: payParams.signType,
        paySign: payParams.paySign,
        success: (res) => {
          console.log('支付成功:', res);
          resolve(res);
        },
        fail: (err) => {
          console.error('支付失败:', err);
          reject(new Error('支付失败或已取消'));
        }
      });
    });
  }
})