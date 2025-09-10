const { api } = require('../../utils/api');

Page({
  data: {
    groupBuyList: [],
    loading: true,
    page: 1,
    size: 10,
    hasMore: true,
    loadingMore: false
  },

  onLoad() {
    this.loadGroupBuyList();
  },

  async loadGroupBuyList() {
    this.setData({ loading: true });
    
    try {
      // 模拟获取拼团列表数据
      const mockData = {
        list: [
          {
            id: 101,
            name: 'GoFrame微服务实战课程',
            groupPrice: 9900,
            originalPrice: 19900,
            groupCount: 3,
            currentParticipants: 2,
            mainImage: 'http://wangzhongyang.com/images/goframe-course.jpg',
            endTime: '2025-09-15 23:59:59',
            participants: 128
          },
          {
            id: 102,
            name: 'UniApp跨端开发实战',
            groupPrice: 7900,
            originalPrice: 14900,
            groupCount: 2,
            currentParticipants: 1,
            mainImage: 'http://wangzhongyang.com/images/uniapp-course.jpg',
            endTime: '2025-09-20 23:59:59',
            participants: 256
          },
          {
            id: 103,
            name: '微信小程序高级开发',
            groupPrice: 12900,
            originalPrice: 25900,
            groupCount: 5,
            currentParticipants: 3,
            mainImage: 'http://wangzhongyang.com/images/wxapp-course.jpg',
            endTime: '2025-09-18 23:59:59',
            participants: 89
          },
          {
            id: 104,
            name: '分布式系统架构设计',
            groupPrice: 19900,
            originalPrice: 39900,
            groupCount: 4,
            currentParticipants: 2,
            mainImage: 'http://wangzhongyang.com/images/distributed-course.jpg',
            endTime: '2025-09-25 23:59:59',
            participants: 67
          },
          {
            id: 105,
            name: '云原生DevOps实战',
            groupPrice: 14900,
            originalPrice: 29900,
            groupCount: 3,
            currentParticipants: 1,
            mainImage: 'http://wangzhongyang.com/images/devops-course.jpg',
            endTime: '2025-09-22 23:59:59',
            participants: 182
          }
        ],
        total: 15,
        page: 1,
        size: 10
      };

      const formattedList = mockData.list.map(item => ({
        ...item,
        groupPriceFormatted: (item.groupPrice / 100).toFixed(2),
        originalPriceFormatted: (item.originalPrice / 100).toFixed(2),
        progress: Math.round((item.currentParticipants / item.groupCount) * 100),
        timeLeft: this.calculateTimeLeft(item.endTime)
      }));

      this.setData({
        groupBuyList: formattedList,
        loading: false,
        hasMore: mockData.page < Math.ceil(mockData.total / mockData.size)
      });

    } catch (error) {
      console.error('加载拼团列表失败:', error);
      wx.showToast({
        title: '加载失败',
        icon: 'none'
      });
      this.setData({ loading: false });
    }
  },

  calculateTimeLeft(endTime) {
    const end = new Date(endTime);
    const now = new Date();
    const diff = end - now;
    
    if (diff <= 0) return '已结束';
    
    const days = Math.floor(diff / (1000 * 60 * 60 * 24));
    const hours = Math.floor((diff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
    
    if (days > 0) return `${days}天${hours}小时`;
    if (hours > 0) return `${hours}小时`;
    
    const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60));
    return `${minutes}分钟`;
  },

  onGroupBuyClick(e) {
    const { id } = e.currentTarget.dataset;
    wx.navigateTo({
      url: `/pages/group-buy/group-buy?id=${id}`
    });
  },

  onLoadMore() {
    if (!this.data.hasMore || this.data.loadingMore) return;
    
    this.setData({ loadingMore: true });
    
    // 模拟加载更多
    setTimeout(() => {
      const newPage = this.data.page + 1;
      if (newPage > 2) { // 模拟只有2页数据
        this.setData({ hasMore: false, loadingMore: false });
        return;
      }
      
      // 模拟新增数据
      const newItems = [
        {
          id: 106,
          name: '前端性能优化实战',
          groupPrice: 8900,
          originalPrice: 16900,
          groupCount: 2,
          currentParticipants: 1,
          mainImage: 'http://wangzhongyang.com/images/performance-course.jpg',
          endTime: '2025-09-28 23:59:59',
          participants: 94
        },
        {
          id: 107,
          name: '后端架构设计模式',
          groupPrice: 15900,
          originalPrice: 29900,
          groupCount: 4,
          currentParticipants: 2,
          mainImage: 'http://wangzhongyang.com/images/architecture-course.jpg',
          endTime: '2025-09-30 23:59:59',
          participants: 123
        }
      ].map(item => ({
        ...item,
        groupPriceFormatted: (item.groupPrice / 100).toFixed(2),
        originalPriceFormatted: (item.originalPrice / 100).toFixed(2),
        progress: Math.round((item.currentParticipants / item.groupCount) * 100),
        timeLeft: this.calculateTimeLeft(item.endTime)
      }));
      
      this.setData({
        groupBuyList: [...this.data.groupBuyList, ...newItems],
        page: newPage,
        loadingMore: false,
        hasMore: newPage < 2 // 模拟只有2页
      });
    }, 1000);
  },

  onPullDownRefresh() {
    this.setData({
      page: 1,
      hasMore: true
    });
    this.loadGroupBuyList().then(() => {
      wx.stopPullDownRefresh();
    });
  },

  onReachBottom() {
    this.onLoadMore();
  }
});