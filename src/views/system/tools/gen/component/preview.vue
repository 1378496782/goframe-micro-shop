<template>
  <div class="gen-viewer-container">
  <el-dialog :fullscreen="fullscreen" v-model="open" width="80%" :close-on-click-modal="false" :destroy-on-close="true">
    <template #header>
      <div class="custom-dialog-title">
        <div class="title-content">
          {{title}}
        </div>
      </div>
    </template>
    <el-tabs v-model="activeName">
      <el-tab-pane
          v-for="(value, key) in data"
          :label="key"
          :name="key"
          :key="key"
      >
<!--        <pre>{{ value }}</pre>-->
        <div class="gen-viewer">
          <my-code-mirror :height="height" :model-value="value"></my-code-mirror>
        </div>
      </el-tab-pane>
    </el-tabs>
  </el-dialog>
  </div>
</template>

<script setup lang="ts">
import {defineComponent, reactive, toRefs, ref, onBeforeMount, onMounted} from "vue";
import {previewState} from "/@/views/system/tools/gen/component/model";
import {showPreview} from "/@/api/system/tools/gen";
import MyCodeMirror from "/@/components/myCodeMirror/index.vue";
import { ElLoading } from 'element-plus'
defineOptions({ name: "genCodePreview"})
const height = ref(400)
// 预览参数
const preview:previewState = reactive({
  flush: true,
  fullscreen: false,
  open: false,
  title: "代码预览",
  data: {},
  activeName: "api"
})
const { data, activeName, fullscreen, open, title}=toRefs(preview)
const getWindowInfo = () => {
  height.value = document.getElementsByTagName('body')[0].clientHeight-130;
};
const showView = (tableId:number)=>{
  const loading = ElLoading.service({
    lock: true,
    text: 'Loading',
    background: 'rgba(0, 0, 0, 0.7)',
  })
  showPreview(tableId).then((res:any)=>{
    preview.data = res.data.data??{}
    preview.open = true
    loading.close()
  }).catch(()=>{
    loading.close()
  })
}
onMounted(()=>{
  height.value = document.getElementsByTagName('body')[0].clientHeight-130;
})
onBeforeMount(()=>{
  window.addEventListener('resize', getWindowInfo);
})
defineExpose({showView})
</script>

<style lang="scss" scoped>
.gen-viewer-container :deep(.el-dialog .el-dialog__body){
  max-height: none!important;
}
</style>

