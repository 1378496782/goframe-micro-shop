const { api } = require('../../utils/api')

Page({
  data: {
    password: '',
    confirmPassword: '',
    secretAnswer: '',
    loading: false
  },

  // 输入框事件
  onPasswordInput(e) {
    this.setData({ password: e.detail.value })
  },

  onConfirmPasswordInput(e) {
    this.setData({ confirmPassword: e.detail.value })
  },

  onSecretAnswerInput(e) {
    this.setData({ secretAnswer: e.detail.value })
  },

  // 提交修改密码
  async onSubmit() {
    const { password, confirmPassword, secretAnswer } = this.data

    // 表单验证
    if (!password.trim()) {
      wx.showToast({
        title: '请输入新密码',
        icon: 'none'
      })
      return
    }

    if (password.length < 6) {
      wx.showToast({
        title: '新密码至少6位',
        icon: 'none'
      })
      return
    }

    if (password !== confirmPassword) {
      wx.showToast({
        title: '两次输入密码不一致',
        icon: 'none'
      })
      return
    }

    if (!secretAnswer.trim()) {
      wx.showToast({
        title: '请输入密保答案',
        icon: 'none'
      })
      return
    }

    this.setData({ loading: true })

    try {
      const res = await api.updatePassword({
        password: password,
        secret_answer: secretAnswer
      })

      if (res.code === 0) {
        wx.showToast({
          title: '密码修改成功',
          icon: 'success',
          duration: 2000,
          success: () => {
            setTimeout(() => {
              wx.navigateBack()
            }, 2000)
          }
        })
      } else {
        wx.showToast({
          title: res.message || '修改失败',
          icon: 'none'
        })
      }
    } catch (error) {
      console.error('修改密码失败:', error)
      wx.showToast({
        title: '网络错误，请重试',
        icon: 'none'
      })
    } finally {
      this.setData({ loading: false })
    }
  },

  // 返回
  onBack() {
    wx.navigateBack()
  }
})