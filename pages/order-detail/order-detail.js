const app = getApp();
const { request } = require('../../utils/request');
const { API } = require('../../config/index');

Page({
  data: {
    orderId: null,
    orderInfo: {},
    statusIcon: '📦',
    statusDesc: ''
  },

  onLoad: function (options) {
    console.log('订单详情页面onLoad参数:', options);
    const orderId = options.id || options.orderId;
    console.log('提取的订单ID:', orderId);
    
    if (!orderId) {
      console.error('订单ID为空');
      wx.showToast({
        title: '订单ID为空',
        icon: 'none'
      });
      return;
    }
    
    this.setData({
      orderId: orderId
    });
    this.loadOrderDetail();
  },

  onShow: function () {
    // 页面显示时重新加载数据
    if (this.data.orderId) {
      this.loadOrderDetail();
    }
  },

  /**
   * 加载订单详情
   */
  loadOrderDetail: function () {
    wx.showLoading({ title: '加载中...' });
    
    console.log('订单详情请求参数:', this.data.orderId);
    console.log('订单详情接口URL:', `${API.ORDER_DETAIL}/${this.data.orderId}`);
    
    request({
      url: `${API.ORDER_DETAIL}/${this.data.orderId}`,
      method: 'GET'
    }).then(res => {
      wx.hideLoading();
      console.log('订单详情接口返回:', res);
      console.log('接口返回数据结构:', JSON.stringify(res, null, 2));
      
      if (res.code === 0 && res.data) {
        console.log('开始格式化订单数据...');
        const orderInfo = this.formatOrderDetail(res.data);
        console.log('格式化后的订单数据:', orderInfo);
        console.log('订单商品数据:', orderInfo.goods);
        console.log('订单地址数据:', orderInfo.address);
        
        this.setData({
          orderInfo: orderInfo
        });
        this.updateStatusInfo(orderInfo.status);
      } else {
        wx.showToast({
          title: res.msg || '加载失败',
          icon: 'none'
        });
      }
    }).catch(err => {
      wx.hideLoading();
      console.error('加载订单详情失败:', err);
      wx.showToast({
        title: '加载失败，请重试',
        icon: 'none'
      });
    });
  },

  /**
   * 格式化订单详情数据
   */
  formatOrderDetail: function (data) {
    // 处理实际返回的数据结构
    const orderInfo = data.order_info || data;
    const orderGoods = data.order_goods_infos || data.goods || data.goods_info || [];
    
    console.log('格式化订单数据:', { orderInfo, orderGoods });

    // 格式化商品数据
    let goods = [];
    if (orderGoods && Array.isArray(orderGoods)) {
      goods = orderGoods.map(item => ({
        ...item,
        id: item.goods_id || item.id,
        name: item.goods_name || item.name,
        image: item.pic_url ? `${getApp().globalData.imageBaseUrl}${item.pic_url}` : '/images/default-product.png',
        price: item.goods_price || item.price || 0,
        quantity: item.count || item.quantity || 1
      }));
    }

    // 格式化地址数据 - 使用返回的收货人信息
    let address = null;
    if (orderInfo.consignee_name || orderInfo.consignee_address) {
      address = {
        name: orderInfo.consignee_name || '',
        phone: orderInfo.consignee_phone || '',
        detail: orderInfo.consignee_address || ''
      };
    }

    // 格式化时间
    const formatTime = (timeStr) => {
      if (!timeStr) return '';
      const date = new Date(timeStr);
      const year = date.getFullYear();
      const month = (date.getMonth() + 1).toString().padStart(2, '0');
      const day = date.getDate().toString().padStart(2, '0');
      const hour = date.getHours().toString().padStart(2, '0');
      const minute = date.getMinutes().toString().padStart(2, '0');
      return `${year}-${month}-${day} ${hour}:${minute}`;
    };

    return {
      ...orderInfo,
      goods: goods,
      address: address,
      createTime: formatTime(orderInfo.created_at || orderInfo.create_time),
      payTime: formatTime(orderInfo.pay_at || orderInfo.pay_time),
      shipTime: formatTime(orderInfo.ship_time),
      receiveTime: formatTime(orderInfo.receive_time),
      statusText: this.getStatusText(orderInfo.status),
      totalAmount: orderInfo.actual_price || orderInfo.price || 0,
      order_number: orderInfo.number || orderInfo.order_number
    };
  },

  /**
   * 获取订单状态文本
   */
  getStatusText: function (status) {
    const statusMap = {
      1: '待支付',
      2: '已支付待发货',
      3: '已发货',
      4: '已收货待评价',
      5: '已评价',
      6: '待确认',
      7: '已取消'
    };
    return statusMap[status] || '未知状态';
  },

  /**
   * 更新状态信息
   */
  updateStatusInfo: function (status) {
    const statusConfig = {
      1: { icon: '⏰', desc: '请在30分钟内完成支付' },
      2: { icon: '📦', desc: '商家正在准备发货' },
      3: { icon: '🚚', desc: '商品正在配送中' },
      4: { icon: '✅', desc: '已收货，快去评价吧' },
      5: { icon: '⭐', desc: '感谢您的评价' },
      6: { icon: '⏳', desc: '等待商家确认' },
      7: { icon: '❌', desc: '订单已取消' }
    };

    const config = statusConfig[status] || { icon: '📦', desc: '' };
    this.setData({
      statusIcon: config.icon,
      statusDesc: config.desc
    });
  },

  /**
   * 查看商品详情
   */
  viewGoodsDetail: function (e) {
    const goodsId = e.currentTarget.dataset.id;
    wx.navigateTo({
      url: `/pages/product-detail/product-detail?id=${goodsId}`
    });
  },

  /**
   * 去支付
   */
  goPay: function () {
    const orderId = this.data.orderId;
    const orderNumber = this.data.orderInfo.order_number || this.data.orderInfo.number;
    
    wx.navigateTo({
      url: `/pages/payment/payment?orderId=${orderId}&orderNumber=${orderNumber}`
    });
  },

  /**
   * 取消订单
   */
  cancelOrder: function () {
    const orderId = this.data.orderId;
    
    wx.showModal({
      title: '提示',
      content: '确定要取消该订单吗？',
      success: (res) => {
        if (res.confirm) {
          this.doCancelOrder(orderId);
        }
      }
    });
  },

  /**
   * 执行取消订单
   */
  doCancelOrder: function (orderId) {
    wx.showLoading({ title: '处理中...' });
    
    request({
      url: API.ORDER_CANCEL,
      method: 'POST',
      data: { id: orderId }
    }).then(res => {
      wx.hideLoading();
      
      if (res.code === 0) {
        wx.showToast({
          title: '订单已取消',
          icon: 'success'
        });
        
        // 返回上一页并刷新
        setTimeout(() => {
          wx.navigateBack();
        }, 1500);
      } else {
        wx.showToast({
          title: res.msg || '操作失败',
          icon: 'none'
        });
      }
    }).catch(err => {
      wx.hideLoading();
      console.error('取消订单失败:', err);
      wx.showToast({
        title: '操作失败，请重试',
        icon: 'none'
      });
    });
  },

  /**
   * 查看物流
   */
  viewLogistics: function () {
    const orderId = this.data.orderId;
    wx.navigateTo({
      url: `/pages/logistics/logistics?orderId=${orderId}`
    });
  },

  /**
   * 确认收货
   */
  confirmReceive: function () {
    const orderId = this.data.orderId;
    
    wx.showModal({
      title: '提示',
      content: '确认已收到商品吗？',
      success: (res) => {
        if (res.confirm) {
          this.doConfirmReceive(orderId);
        }
      }
    });
  },

  /**
   * 执行确认收货
   */
  doConfirmReceive: function (orderId) {
    wx.showLoading({ title: '处理中...' });
    
    request({
      url: API.ORDER_CONFIRM_RECEIVE,
      method: 'POST',
      data: { id: orderId }
    }).then(res => {
      wx.hideLoading();
      
      if (res.code === 0) {
        wx.showToast({
          title: '已确认收货',
          icon: 'success'
        });
        
        // 重新加载订单详情
        this.loadOrderDetail();
      } else {
        wx.showToast({
          title: res.msg || '操作失败',
          icon: 'none'
        });
      }
    }).catch(err => {
      wx.hideLoading();
      console.error('确认收货失败:', err);
      wx.showToast({
        title: '操作失败，请重试',
        icon: 'none'
      });
    });
  },

  /**
   * 申请退款
   */
  applyRefund: function () {
    const orderId = this.data.orderId;
    wx.showModal({
      title: '提示',
      content: '确定要申请退款吗？',
      success: (res) => {
        if (res.confirm) {
          wx.navigateTo({
            url: `/pages/refund/refund?orderId=${orderId}`
          });
        }
      }
    });
  },

  /**
   * 申请售后
   */
  applyAfterSale: function () {
    const orderId = this.data.orderId;
    wx.navigateTo({
      url: `/pages/after-sale/after-sale?orderId=${orderId}`
    });
  },

  /**
   * 去评价
   */
  evaluate: function () {
    const orderId = this.data.orderId;
    wx.navigateTo({
      url: `/pages/evaluate/evaluate?orderId=${orderId}`
    });
  },

  /**
   * 查看售后详情
   */
  viewAfterSaleDetail: function () {
    const orderId = this.data.orderId;
    wx.navigateTo({
      url: `/pages/after-sale-detail/after-sale-detail?orderId=${orderId}`
    });
  }
});