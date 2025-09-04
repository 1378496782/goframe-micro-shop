const { api } = require('../../utils/api')

Page({
  data: {
    oldPassword: '',
    newPassword: '',
    confirmPassword: '',
    loading: false
  },

  // 输入框事件
  onOldPasswordInput(e) {
    this.setData({ oldPassword: e.detail.value })
  },

  onNewPasswordInput(e) {
    this.setData({ newPassword: e.detail.value })
  },

  onConfirmPasswordInput(e) {
    this.setData({ confirmPassword: e.detail.value })
  },

  // 提交修改密码
  async onSubmit() {
    const { oldPassword, newPassword, confirmPassword } = this.data

    // 表单验证
    if (!oldPassword.trim()) {
      wx.showToast({
        title: '请输入原密码',
        icon: 'none'
      })
      return
    }

    if (!newPassword.trim()) {
      wx.showToast({
        title: '请输入新密码',
        icon: 'none'
      })
      return
    }

    if (newPassword.length < 6) {
      wx.showToast({
        title: '新密码至少6位',
        icon: 'none'
      })
      return
    }

    if (newPassword !== confirmPassword) {
      wx.showToast({
        title: '两次输入密码不一致',
        icon: 'none'
      })
      return
    }

    this.setData({ loading: true })

    try {
      const res = await api.updatePassword({
        old_password: oldPassword,
        new_password: newPassword
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