/**
 * 我的订单列表页面
 * 功能：展示用户的订单列表，支持按状态筛选
 * 状态：1-待付款，2-待发货，3-待收货，4-已完成，5-售后
 */
const app = getApp()
const { request } = require('../../utils/request')
const { API } = require('../../utils/env')

Page({
  /**
   * 页面数据对象
   * activeTab: 当前激活的标签页（1-待付款，2-待发货，3-待收货，4-已完成，5-售后）
   * orderList: 订单列表数据
   * loading: 加载状态
   * page: 当前页码
   * limit: 每页数量
   * hasMore: 是否有更多数据
   */
  data: {
    activeTab: 1, // 默认显示待付款
    orderList: [],
    loading: false,
    page: 1,
    limit: 10,
    hasMore: true,
    showTabs: true, // 是否显示标签栏，默认显示
    tabs: [
      { id: 1, name: '待付款' },
      { id: 2, name: '待发货' },
      { id: 3, name: '待收货' },
      { id: 4, name: '已完成' }
    ],
    statusTitles: {
      1: '待付款订单',
      2: '待发货订单',
      3: '待收货订单',
      4: '已完成订单',
      5: '售后订单'
    }
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    // 处理页面参数
    const showTabs = options.showTabs === 'true' || options.showTabs === undefined;
    let status = 1; // 默认状态：待付款
    
    if (options.status) {
      status = parseInt(options.status);
    }
    
    // 设置页面数据
    this.setData({ 
      activeTab: status,
      showTabs: showTabs
    });
    
    // 设置页面标题
    if (!showTabs) {
      wx.setNavigationBarTitle({
        title: this.data.statusTitles[status] || '我的订单'
      });
    }
    
    // 加载订单数据
    this.loadOrderList(true);
  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {
    this.loadOrderList(true);
  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {
    if (this.data.hasMore && !this.data.loading) {
      this.loadOrderList();
    }
  },

  /**
   * 切换标签页
   */
  switchTab: function (e) {
    // 如果不显示标签栏，则不允许切换
    if (!this.data.showTabs) return;
    
    const tabId = e.currentTarget.dataset.id;
    if (tabId === this.data.activeTab) return;
    
    this.setData({
      activeTab: tabId,
      orderList: [],
      page: 1,
      hasMore: true
    });
    
    this.loadOrderList(true);
  },

  /**
   * 加载订单列表数据
   * @param {boolean} refresh 是否刷新（重置页码）
   */
  loadOrderList: function (refresh = false) {
    if (this.data.loading) return;
    
    const page = refresh ? 1 : this.data.page;
    
    this.setData({ loading: true });
    
    if (refresh) {
      wx.showLoading({ title: '加载中...' });
    }
    
    // 构建请求参数
    const params = {
      page: page,
      size: this.data.limit,
      status: this.data.activeTab,
      user_id: app.globalData.userInfo.id
    };
    
    // 调用订单列表接口
    request({
      url: 'https://business.dayu.club/frontend/order/list',
      method: 'GET',
      data: params
    }).then(res => {
      if (res && res.list) {
        const newList = res.list || [];
        const totalOrders = res.total || 0;
        
        // 格式化订单数据
        const formattedList = this.formatOrderList(newList);
        
        // 更新页面数据
        this.setData({
          orderList: refresh ? formattedList : [...this.data.orderList, ...formattedList],
          page: page + 1,
          hasMore: (page * this.data.limit) < totalOrders
        });
      } else {
        wx.showToast({
          title: res.msg || '加载失败',
          icon: 'none'
        });
      }
    }).catch(err => {
      console.error('加载订单列表失败:', err);
      wx.showToast({
        title: '加载失败，请重试',
        icon: 'none'
      });
    }).finally(() => {
      this.setData({ loading: false });
      wx.hideLoading();
      wx.stopPullDownRefresh();
    });
  },

  /**
   * 格式化订单列表数据
   * @param {Array} list 原始订单列表数据
   * @returns {Array} 格式化后的订单列表数据
   */
  formatOrderList: function (list) {
    return list.map(order => {
      // 处理商品信息
      const goods = (order.goods_info || []).map(item => ({
        id: item.goods_id,
        name: item.goods_name,
        count: item.count,
        price: item.price,
        pic_url: item.pic_url ? `https://shopadmin.dayu.club/upload_file/${item.pic_url}` : ''
      }));

      // 格式化订单状态文本
      const statusText = this.getStatusText(order.status);
      
      return {
        ...order,
        goods,
        totalAmount: order.actual_price,
        statusText,
        createTime: order.create_time ? this.formatTime(order.create_time) : ''
      };
    });
  },

  /**
   * 获取订单状态文本
   * @param {number} status 订单状态码
   * @returns {string} 状态文本
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
   * 格式化时间
   * @param {string} timeStr 时间字符串
   * @returns {string} 格式化后的时间字符串
   */
  formatTime: function (timeStr) {
    if (!timeStr) return '';
    
    const date = new Date(timeStr);
    const year = date.getFullYear();
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const day = date.getDate().toString().padStart(2, '0');
    const hour = date.getHours().toString().padStart(2, '0');
    const minute = date.getMinutes().toString().padStart(2, '0');
    
    return `${year}-${month}-${day} ${hour}:${minute}`;
  },

  /**
   * 去支付
   * @param {object} e 事件对象
   */
  goPay: function (e) {
    const orderId = e.currentTarget.dataset.id;
    const orderNumber = e.currentTarget.dataset.number;
    
    wx.navigateTo({
      url: `/pages/payment/payment?orderId=${orderId}&orderNumber=${orderNumber}`
    });
  },

  /**
   * 取消订单
   * @param {object} e 事件对象
   */
  cancelOrder: function (e) {
    const orderId = e.currentTarget.dataset.id;
    
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
   * 执行取消订单操作
   * @param {number} orderId 订单ID
   */
  doCancelOrder: function (orderId) {
    wx.showLoading({ title: '处理中...' });
    
    request({
      url: API.ORDER_CANCEL,
      method: 'POST',
      data: { id: orderId }
    }).then(res => {
      if (res.code === 0) {
        wx.showToast({
          title: '订单已取消',
          icon: 'success'
        });
        
        // 重新加载订单列表
        this.loadOrderList(true);
      } else {
        wx.showToast({
          title: res.msg || '操作失败',
          icon: 'none'
        });
      }
    }).catch(err => {
      console.error('取消订单失败:', err);
      wx.showToast({
        title: '操作失败，请重试',
        icon: 'none'
      });
    }).finally(() => {
      wx.hideLoading();
    });
  },

  /**
   * 确认收货
   * @param {object} e 事件对象
   */
  confirmReceive: function (e) {
    const orderId = e.currentTarget.dataset.id;
    
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
   * 执行确认收货操作
   * @param {number} orderId 订单ID
   */
  doConfirmReceive: function (orderId) {
    wx.showLoading({ title: '处理中...' });
    
    request({
      url: API.ORDER_CONFIRM_RECEIVE,
      method: 'POST',
      data: { id: orderId }
    }).then(res => {
      if (res.code === 0) {
        wx.showToast({
          title: '已确认收货',
          icon: 'success'
        });
        
        // 重新加载订单列表
        this.loadOrderList(true);
      } else {
        wx.showToast({
          title: res.msg || '操作失败',
          icon: 'none'
        });
      }
    }).catch(err => {
      console.error('确认收货失败:', err);
      wx.showToast({
        title: '操作失败，请重试',
        icon: 'none'
      });
    }).finally(() => {
      wx.hideLoading();
    });
  },

  /**
   * 查看订单详情
   * @param {object} e 事件对象
   */
  viewOrderDetail: function (e) {
    const orderId = e.currentTarget.dataset.id;
    
    wx.navigateTo({
      url: `/pages/order-detail/order-detail?id=${orderId}`
    });
  },

  /**
   * 申请售后
   * @param {object} e 事件对象
   */
  applyAfterSale: function (e) {
    const orderId = e.currentTarget.dataset.id;
    
    wx.navigateTo({
      url: `/pages/after-sale/after-sale?orderId=${orderId}`
    });
  }
})