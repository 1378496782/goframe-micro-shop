<template>
  <div class="excel-uploader">
    <!-- 添加或修改代码生成测试对话框 -->
    <el-dialog v-model="upload.open" width="500px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.excel-uploader .el-dialog', '.excel-uploader .el-dialog__header']">{{upload.title}}</div>
      </template>
      <el-upload
          ref="uploadRef"
          :limit="1"
          accept=".xlsx, .xls"
          :headers="upload.headers"
          :action="upload.url"
          :disabled="upload.isUploading"
          :on-progress="handleFileUploadProgress"
          :on-success="handleFileSuccess"
          :auto-upload="false"
          drag
          class="up-selector"
      >
        <el-icon class="el-icon--upload"><ele-UploadFilled /></el-icon>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        <template #tip>
          <div class="text-center el-upload__tip">
            <span>仅允许导入xls、xlsx格式文件。</span>
            <el-link type="primary" :underline="false" style="font-size:12px;vertical-align: baseline;" @click="importTemplate">下载模板</el-link>
          </div>
        </template>
      </el-upload>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitFileForm">确 定</el-button>
          <el-button @click="upload.open = false">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup name="loadExcel" lang="ts">
import {reactive, ref} from "vue";
import {getToken} from "/@/utils/gfast";
import {ElMessageBox, UploadFile} from "element-plus";
import {downLoadXml} from "/@/utils/zipdownload";
const baseURL:string|undefined|boolean = import.meta.env.VITE_API_URL
const uploadRef = ref<ElUploadInstance>();
const emit = defineEmits(['getList'])
const props = defineProps(['upUrl','tplUrl'])
const globalHeaders = ():object => {
  return {
    Authorization:'Bearer '+getToken()
  }
}
/*** 用户导入参数 */
const upload = reactive({
  // 是否显示弹出层（用户导入）
  open: false,
  // 弹出层标题（用户导入）
  title: "excel导入",
  // 是否禁用上传
  isUploading: false,
  // 设置上传的请求头部
  headers: globalHeaders(),
  // 上传的地址
  url: baseURL + props.upUrl
})
  /**文件上传中处理 */
  const handleFileUploadProgress = () => {
    upload.isUploading = true;
  }
  /** 文件上传成功处理 */
  const handleFileSuccess = (response: any, file: UploadFile) => {
    upload.open = false;
    upload.isUploading = false;
    uploadRef.value?.handleRemove(file);
    let message = '导入成功！'
    if(response.code!==0){
      message = response.message
    }
    ElMessageBox.alert("<div style='overflow: auto;overflow-x: hidden;max-height: 70vh;padding: 10px 20px 0;'>" + message + "</div>", "导入结果", { dangerouslyUseHTMLString: true });
    emit('getList')
  }
/** 下载模板操作 */
const importTemplate = () => {
  downLoadXml(props.tplUrl,{},'get')
}

/** 提交上传文件 */
const submitFileForm = () => {
  uploadRef.value?.submit();
}

const open = ()=>{
  upload.open = true
}
defineExpose({ open })
</script>

<style scoped lang="scss">
.excel-uploader :deep(.avatar-uploader .avatar) {
  width: 178px;
  height: 178px;
  display: block;
}
.excel-uploader :deep(.avatar-uploader .el-upload) {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}
.excel-uploader :deep(.avatar-uploader .el-upload:hover) {
  border-color: var(--el-color-primary);
}
.excel-uploader :deep(.el-icon.avatar-uploader-icon) {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
.kv-label{margin-bottom: 15px;font-size: 14px;}
.mini-btn i.el-icon{margin: unset;}
.kv-row{margin-bottom: 12px;}
:deep(.up-selector .el-upload.is-drag){
  margin: auto;
  width: 210px!important;
}

.text-center{
  text-align: center;
}
</style>
