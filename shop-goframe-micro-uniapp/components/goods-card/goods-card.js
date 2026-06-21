Component({
  properties: {
    goods: {
      type: Object,
      value: {}
    },
    index: {
      type: Number,
      value: 0
    }
  },

  data: {
    imageLoaded: false,
    imageError: false
  },

  methods: {
    onTap() {
      this.triggerEvent('tap', { id: this.data.goods.id })
    },
    
    onImageError(e) {
      console.log('图片加载失败:', e)
      this.setData({
        imageError: true
      })
      // 设置默认图片
      this.triggerEvent('imageerror', { 
        index: this.properties.index,
        goods: this.properties.goods
      })
    },
    
    onImageLoad(e) {
      console.log('图片加载成功')
      this.setData({
        imageLoaded: true
      })
    }
  },
  
  lifetimes: {
    attached() {
      // 组件初始化时设置默认图片
      if (!this.properties.goods.mainImage) {
        this.setData({
          'goods.mainImage': 'https://via.placeholder.com/400x400/19aecc/ffffff?text=商品图片'
        })
      }
    }
  }
})