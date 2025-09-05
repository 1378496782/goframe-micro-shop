<template>
  <div class="ue-content">
    <vue-ueditor-wrap v-model="html" :config="editorConfig" :editor-id="editorId"
                      :editorDependencies="['ueditor.config.js','ueditor.all.js']"></vue-ueditor-wrap>
  </div>
</template>


<script lang="ts">
import { getToken,baseURL } from '/@/utils/gfast';
import {defineComponent, reactive, computed} from "vue";
import {toolbar} from "/@/components/ueditor/model";

export default defineComponent({
  name:'GfUeditor',
  props: {
    editorId: {
      type: String,
      default: 'vueEditorId01'
    },
    modelValue: {
      type: String,
      default: ''
    },
    ueditorConfig: {
      type: Object,
      default(){
        return {
          initialFrameWidth: '100%',
          initialFrameHeight: 400,
          maximumWords: 5000,
          topOffset: 80,
          zIndex:2050
        }
      }
    },
    toolBars: {
      type: Array,
      default() {
        return toolbar
      }
    }
  },
  emits:['update:modelValue'],
  setup(props,{emit}){
    const config = Object.assign({
      elementPathEnabled: false,
      maximumWords: 10000,
      headers: {
        Authorization: 'Bearer ' + getToken()
      },
      toolbars: [props.toolBars],
      UEDITOR_HOME_URL: import.meta.env.BASE_URL + "js/ueditor/", // 访问 UEditor 静态资源的根路径，可参考常见问题1
      UEDITOR_CORS_URL: import.meta.env.BASE_URL + "js/ueditor/", // 访问 UEditor 静态资源的根路径，可参考常见问题1
      serverUrl: baseURL+'api/v1/system/uEditor/action?token='+encodeURIComponent(getToken()), // 服务端接口（这个地址是我为了方便各位体验文件上传功能搭建的临时接口，请勿在生产环境使用！！！）
    },props.ueditorConfig)
    const editorConfig = reactive(config)
    const html =computed({
      get:()=>{
        return props.modelValue
      },
      set:(newVal)=>{
        emit('update:modelValue',newVal)
      }
    })
    return {
      html,
      editorConfig
    }
  }
})
</script>

<style scoped>
.ue-content{line-height: normal!important;width: 100%;}
.ue-content .ue-container{

}
</style>
