Component({
  properties: {
    goods: {
      type: Object,
      value: {}
    }
  },

  methods: {
    onTap() {
      this.triggerEvent('tap', { id: this.data.goods.id })
    }
  }
})