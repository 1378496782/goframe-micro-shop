const app = getApp()

Page({
  data: {
    phone: '',
    code: '',
    password: '',
    confirmPassword: '',
    agreed: false,
    loading: false,
    codeCountdown: 0
  },

  onPhoneInput(e) {
    this.setData({
      phone: e.detail.value
    });
  },

  onCodeInput(e) {
    this.setData({
      code: e.detail.value
    });
  },

  onPasswordInput(e) {
    this.setData({
      password: e.detail.value
    });
  },

  onConfirmPasswordInput(e) {
    this.setData({
      confirmPassword: e.detail.value
    });
  },

  onToggleAgreement() {
    this.setData({
      agreed: !this.data.agreed
    });
  },

  onSendCode() {
    const { phone, codeCountdown } = this.data;

    if (codeCountdown > 0) return;

    if (!phone) {
      wx.showToast({
        title: '请输入手机号',
        icon: 'none'
      });
      return;
    }

    if (!/^1[3-9]\d{9}$/.test(phone)) {
      wx.showToast({
        title: '请输入正确的手机号',
        icon: 'none'
      });
      return;
    }

    // 开始倒计时
    this.setData({ codeCountdown: 60 });
    this.startCountdown();

    // 模拟发送验证码
    wx.showToast({
      title: '验证码已发送',
      icon: 'success'
    });
  },

  startCountdown() {
    const timer = setInterval(() => {
      if (this.data.codeCountdown <= 1) {
        clearInterval(timer);
        this.setData({ codeCountdown: 0 });
        return;
      }
      this.setData({
        codeCountdown: this.data.codeCountdown - 1
      });
    }, 1000);
  },

  onRegister() {
    const { phone, code, password, confirmPassword, agreed } = this.data;

    if (!phone) {
      wx.showToast({
        title: '请输入手机号',
        icon: 'none'
      });
      return;
    }

    if (!/^1[3-9]\d{9}$/.test(phone)) {
      wx.showToast({
        title: '请输入正确的手机号',
        icon: 'none'
      });
      return;
    }

    if (!code) {
      wx.showToast({
        title: '请输入验证码',
        icon: 'none'
      });
      return;
    }

    if (code.length !== 6) {
      wx.showToast({
        title: '验证码格式不正确',
        icon: 'none'
      });
      return;
    }

    if (!password) {
      wx.showToast({
        title: '请输入密码',
        icon: 'none'
      });
      return;
    }

    if (password.length < 6 || password.length > 20) {
      wx.showToast({
        title: '密码长度应为6-20位',
        icon: 'none'
      });
      return;
    }

    if (password !== confirmPassword) {
      wx.showToast({
        title: '两次密码输入不一致',
        icon: 'none'
      });
      return;
    }

    if (!agreed) {
      wx.showToast({
        title: '请同意用户协议',
        icon: 'none'
      });
      return;
    }

    this.setData({ loading: true });

    // 调用真实注册API
    wx.request({
      url: app.globalData.API.USER_REGISTER,
      method: 'POST',
      data: {
        name: phone, // 使用手机号作为用户名
        password: password,
        avatar: 'http://dummyimage.com/100x100',
        sex: 16,
        sign: '新用户注册',
        secret_answer: code // 使用验证码作为密保答案
      },
      success: (res) => {
        this.setData({ loading: false });
        
        if (res.statusCode === 200 && res.data.code === 0) {
          wx.showToast({
            title: '注册成功',
            icon: 'success'
          });
          
          // 注册成功后跳转到登录页
          setTimeout(() => {
            wx.navigateBack();
          }, 1500);
        } else {
          wx.showToast({
            title: res.data.msg || '注册失败',
            icon: 'none'
          });
        }
      },
      fail: (err) => {
        this.setData({ loading: false });
        wx.showToast({
          title: '网络错误，请重试',
          icon: 'none'
        });
      }
    });
  },

  onViewAgreement() {
    wx.showModal({
      title: '用户协议',
      content: '这是用户协议内容...',
      showCancel: false
    });
  },

  navigateToLogin() {
    wx.navigateBack();
  }
});