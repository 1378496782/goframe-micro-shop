// 百度地图组件
import { defineComponent, h } from 'vue';

export default defineComponent({
  name: 'BaiduMap',
  props: {
    src: {
      type: String,
      required: true
    }
  },
  setup(props) {
    return () => h('script', {
      src: props.src,
      type: 'text/javascript'
    });
  }
});