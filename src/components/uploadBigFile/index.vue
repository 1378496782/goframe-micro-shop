<template>
    <el-upload
        ref="upBigFileRef"
        v-model:file-list="dataFileList"
        class="upload-demo"
        :multiple="multiple"
        :drag="drag"
        :on-remove="handleRemove"
        :before-remove="beforeRemove"
        :limit="limit"
        :on-exceed="handleExceed"
        :on-change="handleChange"
        :on-preview="handlePreview"
        :http-request="handleRequest"
        :disabled="uploadStatus!==0"
    >
    <template v-if="uploadStatus===0">
      <el-icon class="el-icon--upload"><ele-UploadFilled /></el-icon>
      <div class="el-upload__text">
        拖拽文件至此 或<em>点击上传</em>
      </div>
    </template>
    <template v-if="uploadStatus===1">
      <el-icon class="el-icon--upload"><ele-Loading /></el-icon>
      <div class="el-upload__text">
        文件解析中...{{uploadStatusContent}}
      </div>
    </template>
    <template v-if="uploadStatus===2">
      <el-icon class="el-icon--upload"><ele-Loading /></el-icon>
      <div class="el-upload__text">
        文件上传中...（{{progress}}%）文件大小：{{fileFormatSize}}
      </div>
    </template>
    <template #tip>
      <div class="progress-show" v-if="uploadStatus===2">
        <el-progress :percentage="progress" />
      </div>
      <div class="el-upload__tip">
        大文件上传：支持分片上传，断点续传及秒传功能
      </div>
    </template>
  </el-upload>
</template>

<script setup  lang="ts">
import SparkMD5 from 'spark-md5';
import {computed, getCurrentInstance, reactive, ref} from "vue";
import type {UploadFile, UploadProps, UploadRawFile, UploadRequestOptions, UploadUserFile} from 'element-plus'
import { ElMessage ,ElMessageBox} from 'element-plus'
import {getToken} from "/@/utils/gfast";
import _ from 'lodash'
import {checkMultipart, uploadPart} from "/@/api/system/sysAttachment";
import axios, {AxiosProgressEvent} from "axios";
defineOptions({ name: "uploadBigFile"})
const props = defineProps({
  name:         { type : String, default : 'file' },//上传文件类型
  method:       { type : String, default : 'post' },//设置上传请求方法
  multiple:     { type : Boolean, default : true },//是否支持多选文件
  showFileList: { type : Boolean, default : true },//是否显示已上传文件列表
  drag:         { type : Boolean, default : true },//是否启用拖拽上传
  disabled:     { type : Boolean, default : false },//是否禁止
  listType:     { type : String, default : 'picture-card' },//
  limit:        { type : Number, default : 5 },//上传最大数量
  modelValue:{
    type:Array,
    default:function(){
      return []
    }
  }
})
const emit = defineEmits(['update:modelValue'])
const upBigFileRef = ref()
const uploadStatus = ref(0); // 上传状态 0等待上传 1解析中 2上传中 3已取消
const uploadStatusContent= ref('')
const progress = ref(0)
const fileFormatSize = ref('0byte')
let  uploadedFile:Array<any> = [] ;
const {proxy} = <any>getCurrentInstance();
const dataFileList = computed({
  get: () => {
    let value:Array<UploadUserFile> = props.modelValue as UploadUserFile[]|| [];
    value.map((item: UploadUserFile)=>{
      if(item.url){
        item.url = proxy.getUpFileUrl(item.url)
      }
      return item
    })
    uploadedFile = _.cloneDeep(value)
    return value
  },
  set: val => {
    emit('update:modelValue', val)
  }
});
const handleRequest = (options: UploadRequestOptions) => {}
const cancelToken = axios.CancelToken.source();
const handleChange: UploadProps['onChange'] = async (uploadFile: UploadFile) => {
  const {name,size,raw:file} = uploadFile
  fileFormatSize.value = formatFileSize(size as number)
  try{
    //计算文件md5
    uploadStatus.value = 1 //文件解析中
    const md5 = await calculateMD5(file as File)
    uploadStatus.value = 2
    let currentChunkIndex = 1
    //分片大小
    const chunkSize = 1024 * 1024 * 2; // 2M
    //计算分片数量
    const shards = Array.from({length: Math.ceil(size as number / chunkSize)}, (_, index) => index+1);
    //检查分片
    const checkPartRes:any = await checkMultipart({
      fileName: name,
      size: size,
      md5: md5,
      shardsCount: shards.length,
    })
    // 已存在(秒传)
    if (!checkPartRes.data.waitUploadIndex || checkPartRes.data.waitUploadIndex.length == 0) {
      upOver(checkPartRes.data.attachment,uploadFile)
      return;
    }
    // 断点续传，过滤掉已上传成功的分片文件
    const upShards = shards.filter((shard) => checkPartRes.data.waitUploadIndex.includes(shard));
    if (upShards.length == 0) {
      upOver(checkPartRes.data.attachment,uploadFile)
      return;
    }
    let uploadedSize = (shards.length-upShards.length-1)*chunkSize;
    for (const index of upShards) {
      if (uploadStatus.value == 3) {
        break;
      }
      let start = (index-1) * chunkSize;
      let end = index * chunkSize;
      uploadedSize+=chunkSize
      const progressFn = (result: AxiosProgressEvent)=>{
        // 实时展示上传进度
        progress.value = Math.round((uploadedSize+result.loaded)/size!*100);
      }
      const res = await uploadPart({
        fileName: name,
        size: size,
        md5: md5,
        shardsCount: shards.length,
        index: index,
        file: file!.slice(start, end),
        token:getToken(),
      },progressFn,cancelToken);
      currentChunkIndex = index;
      if (res.data.data.finish){
        upOver(res.data.data.attachment,uploadFile)
      }
    }
  }catch (e:any){
    if(e.code==="ERR_CANCELED"){
      ElMessage.error("已关闭上传，下次上传将继续从断点处上传")
    }else{
      console.log('error:',e)
    }
    clearUpStatus();
  }
};
const upOver = (attachment:any,uploadFile: UploadFile)=>{
  uploadedFile=uploadedFile.filter((item:UploadUserFile)=>{
    return item.raw?.uid != uploadFile.raw?.uid
  })
  uploadedFile.push({
    name:attachment.name,
    url:attachment.path,
    fullUrl:attachment.fullPath,
    fileType:attachment.type,
    size:attachment.size
  })
  setDataFileList();
  clearUpStatus();
}
const handleExceed = () => {
  ElMessage.error('最多可上传'+props.limit+'个文件,已超出最大限制数。');
}

const beforeRemove: UploadProps['beforeRemove'] = (uploadFile) => {
  return ElMessageBox.confirm(
      `您确定要删除 ${uploadFile.name} ?`,
      '提示',
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      }
  ).then(
      () => true,
      () => false
  )
};
const handleRemove: UploadProps['onRemove'] = (file) => {//移除后
  uploadedFile.splice(uploadedFile.findIndex((item: any) => item.name === file.name),1)
  setDataFileList()
};
const setDataFileList = () => {
  dataFileList.value = uploadedFile
};
const handlePreview = (file:UploadUserFile)=>{
  window.open(file.url)
}

const stopUpBigFile = ()=>{
  uploadStatus.value = 3
  cancelToken.cancel()
}

const calculateMD5 = (file: File): Promise<string> => {
  return new Promise((resolve, reject) => {
    const chunkSize = 1024 * 1024; // 1 MB
    const chunks = Math.ceil(file.size / chunkSize);
    const spark = new SparkMD5.ArrayBuffer();
    let currentChunk = 0;

    const reader = new FileReader();

    reader.onload = (e: ProgressEvent<FileReader>) => {
      spark.append(e.target!.result as ArrayBuffer);
      currentChunk++;

      if (currentChunk < chunks) {
        uploadStatusContent.value = '校验MD5 '+ ((currentChunk/chunks)*100).toFixed(0)+'%'
        loadNextChunk();
      } else {
        // 所有块都读取完毕，生成最终的 MD5 值
        resolve(spark.end());
      }
    };

    reader.onerror = (error) => {
      reject(error);
    };

    function loadNextChunk() {
      const start = currentChunk * chunkSize;
      const end = Math.min(start + chunkSize, file.size);
      const blob = file.slice(start, end);
      reader.readAsArrayBuffer(blob);
    }
    loadNextChunk(); // 开始读取文件的第一个块
  });
}

const formatFileSize = (bytes:number):string => {
  if (bytes === 0) return '0 Bytes';
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
  const index = Math.floor(Math.log(bytes) / Math.log(1024)); // 计算索引
  const value = (bytes / Math.pow(1024, index)).toFixed(2); // 计算具体数值并保留两位小数
  return `${value} ${sizes[index]}`;
}

const clearUpStatus = ()=>{
  uploadStatus.value = 0
  uploadStatusContent.value = ''
  progress.value = 0
  fileFormatSize.value = '0byte'
}

defineExpose({stopUpBigFile})
</script>
<style>
.el-upload.is-drag {
  display: block;
  width: 200px;height: 200px;
}
.progress-show{
  margin-top: 8px;
}
</style>
