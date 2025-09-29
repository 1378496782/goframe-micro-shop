const api = require('../../utils/api').api;
const constants = require('../../config/constants');

Page({
  data: {
    selectedSpecs: {},
    product: null,
    loading: true
  },
 
  async onLoad(options) {
    if (!options.id) {
      wx.showToast({ title: '商品ID无效', icon: 'none' })
      wx.navigateBack()
      return
    }

    // 直接加载商品详情，不检查登录状态
    await this.loadProductDetail(options.id)
  },
  
  // 加载商品详情
  async loadProductDetail(productId) {
    console.log('开始加载商品详情，商品ID:', productId)
    if (!productId) {
      wx.showToast({
        title: '商品ID无效',
        icon: 'none'
      })
      wx.navigateBack()
      return
    }

    // 移除此处的登录检查，允许未登录用户查看商品详情
    
    this.setData({ loading: true })
    
    try {
      console.log('调用API获取商品详情...')
      const res = await api.getGoodsDetail(productId)
      
      if (!res) {
        throw new Error('获取商品详情失败，请检查网络连接')
      }
      console.log('API响应:', res)
      
      if (res.code === 0) {
        const product = res.data
        // 处理图片URL，使用IMAGE_BASE_URL拼接
        const images = []
        
        // 添加主图
        if (product.pic_url) {
          images.push(product.pic_url.startsWith('http') ? product.pic_url : constants.IMAGE_BASE_URL + product.pic_url)
        }
        
        // 添加详情图
        if (product.images) {
          try {
            const detailImages = JSON.parse(product.images)
            detailImages.forEach(img => {
              if (img.url) {
                images.push(img.url.startsWith('http') ? img.url : constants.IMAGE_BASE_URL + img.url)
              }
            })
          } catch (e) {
            console.error('解析商品详情图片失败:', e)
          }
        }
        
        // 如果没有图片，添加默认占位图
        if (images.length === 0) {
          images.push(constants.IMAGE_BASE_URL + 'default-product.png')
        }
        
        // 格式化商品数据
        const formattedProduct = {
          id: product.id,
          name: product.name,
          price: (product.price / 100).toFixed(2),
          originalPrice: ((product.price * 1.2) / 100).toFixed(2), // 模拟原价
          discount: '8.3', // 模拟折扣
          sale: product.sale || 0,
          stock: product.stock || 0,
          images: images,
          tags: product.tags || '',
          brand: product.brand || ''
        }
        
        console.log('格式化后的商品数据:', formattedProduct)
        this.setData({
          product: formattedProduct,
          loading: false
        })
        // 设置页面标题为商品名称
        wx.setNavigationBarTitle({
          title: formattedProduct.name
        })
        console.log('页面数据更新完成')
      } else {
        throw new Error(res.message || '获取商品详情失败')
      }
    } catch (error) {
      console.error('加载商品详情失败:', error)
      wx.showToast({
        title: '加载失败',
        icon: 'none'
      })
      this.setData({ loading: false })
      wx.navigateBack()
    }
  },





  // 添加到收藏
  addToFavorites() {
    wx.showToast({
      title: '已添加到收藏',
      icon: 'success'
    })
  },

  // 前往购物车
  goToCart() {
    wx.switchTab({
      url: '/pages/cart/cart'
    })
  },

  // 加入购物车
  async addToCart() {
    const token = wx.getStorageSync('token')
    if (!token) {
      wx.showModal({
        title: '提示',
        content: '请先登录',
        success: (res) => {
          if (res.confirm) {
            wx.switchTab({
              url: '/pages/user/user'
            })
          }
        }
      })
      return
    }

    try {
      const res = await api.addToCart({
        goods_id: this.data.product.id,
        count: 1
      })

      if (res.code === 0) {
        wx.showToast({
          title: '已加入购物车',
          icon: 'success'
        })
      } else {
        throw new Error(res.message || '加入购物车失败')
      }
    } catch (error) {
      console.error('加入购物车失败:', error)
      wx.showToast({
        title: error.message || '加入购物车失败',
        icon: 'none'
      })
    }
  },

  // 立即购买
  async buyNow() {
    if (!wx.getStorageSync('token')) {
      wx.showModal({
        title: '提示',
        content: '请先登录',
        success: (res) => {
          if (res.confirm) {
            wx.switchTab({
              url: '/pages/user/user'
            })
          }
        }
      })
      return
    }

    // 获取商品信息
    const product = this.data.product;
    if (!product) {
      wx.showToast({
        title: '商品信息不完整',
        icon: 'none'
      })
      return;
    }

    // 显示加载提示
    wx.showLoading({
      title: '处理中...',
      mask: true
    });

    try {
      // 构建订单数据
      const orderData = {
        price: parseFloat(product.price) * 100, // 转换为分
        coupon_price: 0, // 不使用优惠券
        actual_price: parseFloat(product.price) * 100, // 实际价格，转换为分
        remark: '',
        order_goods_info: [{
          goods_id: product.id,
          count: 1,
          price: parseFloat(product.price) * 100, // 转换为分
          coupon_price: 0,
          actual_price: parseFloat(product.price) * 100 // 转换为分
        }]
      };

      console.log('提交订单参数:', orderData);
      
      // 调用提交订单接口
      const { api } = require('../../utils/api');
      const res = await api.submitOrder(orderData);
      
      if (res.code === 0 && res.data && res.data.id) {
        // 订单创建成功，获取订单编号
        const orderId = res.data.id;
        const orderNumber = res.data.number;
        
        console.log('订单创建成功，订单ID:', orderId, '订单编号:', orderNumber);
        
        // 调用微信支付接口
        await this.requestWxPayment(orderNumber);
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
      wx.hideLoading();
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
        amount: Math.round(parseFloat(this.data.product.price) * 100), // 转换为分
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
      const { config } = require('../../utils/env');
      
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