<template>
  <div class="system-sysAttachment-edit">
    <!-- 添加或修改附件管理对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.system-sysAttachment-edit .el-dialog', '.system-sysAttachment-edit .el-dialog__header']">上传</div>
      </template>
      <upload-img ref="upImageRef" v-if="upType==='image'" :action="baseURL+'api/v1/system/upload/singleImg'" v-model="images" :limit="10"></upload-img>
      <upload-file ref="upFileRef" v-else-if="upType==='file'" :action="baseURL+'api/v1/system/upload/singleFile'" v-model="files" :limit="10" :uploadSize="50"></upload-file>
      <upload-big-file ref="upBigFileRef" v-else-if="upType==='bigFile'" v-model="files" :limit="10"></upload-big-file>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="danger" @click="onCancel">关 闭</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import uploadImg from "/@/components/uploadImg/index.vue"
import uploadFile from "/@/components/uploadFile/index.vue"
import uploadBigFile from "/@/components/uploadBigFile/index.vue"
import {reactive, toRefs, ref, watch} from 'vue';
import {
  SysAttachmentEditState
} from "/@/views/system/sysAttachment/list/component/model"
defineOptions({ name: "ApiV1SystemSysAttachmentEdit"})
const baseURL:string|undefined|boolean = import.meta.env.VITE_API_URL
const emit = defineEmits(['sysAttachmentList'])
const upType = ref('')
const upFileRef = ref()
const upImageRef = ref()
const upBigFileRef = ref()
const state = reactive<SysAttachmentEditState>({
  isShowDialog: false,
  images: [],
  files:[]
});
const { isShowDialog,images,files } = toRefs(state);
// 打开弹窗
const openDialog = (upTp:string) => {
  reset()
  upType.value = upTp
  state.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
  state.isShowDialog = false;
  if(upType.value=='image'){
    upImageRef.value.stopUpImage()
  }else if(upType.value=='file'){
    upFileRef.value.stopUpFile()
  }else if(upType.value=='bigFile'){
    upBigFileRef.value.stopUpBigFile()
  }
};
defineExpose({
  openDialog,
});
// 取消
const onCancel = () => {
  closeDialog();
};
watch(()=>[state.images,state.files],()=>{
  emit('sysAttachmentList')
})
const reset = ()=>{
  state.images = []
  state.files = []
}
</script>
<style scoped>
  .kv-label{margin-bottom: 15px;font-size: 14px;}
  .mini-btn i.el-icon{margin: unset;}
  .kv-row{margin-bottom: 12px;}
</style>
