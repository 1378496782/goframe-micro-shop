const { api } = require('../../utils/api');

Page({
  data: {
    groupBuyInfo: null,
    loading: true,
    participants: [],
    currentUserJoined: false,
    groupProgress: 0,
    groupPriceFormatted: '0.00',
    originalPriceFormatted: '0.00'
  },

  onLoad(options) {
    const { id } = options;
    if (id) {
      this.loadGroupBuyDetail(id);
    }
  },

  async loadGroupBuyDetail(id) {
    this.setData({ loading: true });
    
    try {
      // 模拟获取拼团详情数据
      const mockData = {
        id: parseInt(id),
        name: 'GoFrame微服务实战课程',
        groupPrice: 9900,
        originalPrice: 19900,
        groupCount: 3,
        mainImage: 'http://wangzhongyang.com/images/goframe-course.jpg',
        description: '深入学习GoFrame框架，掌握微服务架构设计与实战',
        endTime: '2025-09-15 23:59:59',
        currentParticipants: 2,
        rules: '邀请2位好友参团，即可享受超值优惠价'
      };

      // 模拟参团用户数据
      const mockParticipants = [
        { id: 1, avatar: 'http://wangzhongyang.com/images/avatar1.jpg', nickname: '程序员小王', joinTime: '2025-09-10 10:30:00' },
        { id: 2, avatar: 'http://wangzhongyang.com/images/avatar2.jpg', nickname: '开发者小李', joinTime: '2025-09-10 11:15:00' }
      ];

      this.setData({
        groupBuyInfo: mockData,
        participants: mockParticipants,
        groupProgress: Math.round((mockData.currentParticipants / mockData.groupCount) * 100),
        groupPriceFormatted: (mockData.groupPrice / 100).toFixed(2),
        originalPriceFormatted: (mockData.originalPrice / 100).toFixed(2),
        loading: false
      });

    } catch (error) {
      console.error('加载拼团详情失败:', error);
      wx.showToast({
        title: '加载失败',
        icon: 'none'
      });
      this.setData({ loading: false });
    }
  },

  // 参团按钮点击
  onJoinGroup() {
    const { groupBuyInfo, currentUserJoined } = this.data;
    
    if (currentUserJoined) {
      wx.showToast({
        title: '您已参团',
        icon: 'none'
      });
      return;
    }

    // 检查是否满团
    if (groupBuyInfo.currentParticipants >= groupBuyInfo.groupCount) {
      wx.showToast({
        title: '该团已满',
        icon: 'none'
      });
      return;
    }

    // 跳转到支付页面或直接参团
    wx.showModal({
      title: '确认参团',
      content: `确定以¥${(groupBuyInfo.groupPrice / 100).toFixed(2)}的价格参团吗？`,
      success: (res) => {
        if (res.confirm) {
          this.handleJoinGroup();
        }
      }
    });
  },

  async handleJoinGroup() {
    try {
      // 模拟参团成功
      const newParticipant = {
        id: 3,
        avatar: 'http://wangzhongyang.com/images/avatar3.jpg',
        nickname: '当前用户',
        joinTime: new Date().toLocaleString()
      };

      this.setData({
        participants: [...this.data.participants, newParticipant],
        currentUserJoined: true,
        groupProgress: Math.round((3 / this.data.groupBuyInfo.groupCount) * 100),
        'groupBuyInfo.currentParticipants': 3
      });

      wx.showToast({
        title: '参团成功',
        icon: 'success'
      });

      // 如果满团，提示用户
      if (this.data.groupBuyInfo.currentParticipants >= this.data.groupBuyInfo.groupCount) {
        setTimeout(() => {
          wx.showModal({
            title: '拼团成功',
            content: '恭喜！拼团成功，课程已添加到您的账户',
            showCancel: false
          });
        }, 1500);
      }

    } catch (error) {
      console.error('参团失败:', error);
      wx.showToast({
        title: '参团失败',
        icon: 'none'
      });
    }
  },

  // 分享功能
  onShareAppMessage() {
    const { groupBuyInfo } = this.data;
    return {
      title: `快来和我一起拼团！${groupBuyInfo.name}`,
      path: `/pages/group-buy/group-buy?id=${groupBuyInfo.id}`,
      imageUrl: groupBuyInfo.mainImage
    };
  },

  // 分享到朋友圈
  onShareTimeline() {
    const { groupBuyInfo } = this.data;
    return {
      title: `超值拼团：${groupBuyInfo.name}`,
      imageUrl: groupBuyInfo.mainImage
    };
  }
});