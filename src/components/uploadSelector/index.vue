<template>
  <div class="system-sysAttachment-container">
    <div class="up-selector" @click="openDialog">
      <el-icon class="uploader-icon"><ele-Plus /></el-icon>
    </div>
    <div class="img-list" v-if="dataFileList.length>0">
      <el-upload
          v-model:file-list="dataFileList"
          ref="upImageRef"
          :list-type="fileType=='image'?'picture-card':'text'"
          :on-preview="handlePictureCardPreview"
      ></el-upload>
      <el-dialog v-model="dialogVisible">
        <el-image :src="dialogImageUrl" fit="contain" />
      </el-dialog>
    </div>
    <el-dialog title="选择文件" v-model="isShowDialog" width="1000px" :close-on-click-modal="false" :destroy-on-close="true">
      <el-card shadow="hover">
        <div class="system-sysAttachment-search mb15">
          <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>
              <el-col :span="12" class="form-item">
                <el-form-item label="文件原始名" prop="name">
                  <el-input
                      v-model="tableData.param.name"
                      placeholder="请输入文件原始名"
                      clearable
                      @keyup.enter.native="sysAttachmentList"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="12" class="form-item">
                <el-form-item label="扩展类型" prop="mimeType">
                  <el-input
                      v-model="tableData.param.mimeType"
                      placeholder="请输入扩展类型"
                      clearable
                      @keyup.enter.native="sysAttachmentList"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="12" class="form-item">
                <el-form-item label="创建时间" prop="createdAt">
                  <el-date-picker
                      clearable  style="width: 200px"
                      v-model="tableData.param.createdAt"
                      format="YYYY-MM-DD HH:mm:ss"
                      value-format="YYYY-MM-DD HH:mm:ss"
                      type="datetimerange"
                      range-separator="至"
                      start-placeholder="开始时间"
                      end-placeholder="结束时间"
                  ></el-date-picker>
                </el-form-item>
              </el-col>
              <el-col :span="8" class="form-item">
                <el-form-item>
                  <el-button type="primary"  @click="sysAttachmentList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                  <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
          <el-row :gutter="10" class="mb8" style="margin-top: 8px;">
            <el-col :span="1.5">
              <el-button
                  type="primary"
                  @click="handleAdd('file')"
              ><el-icon><ele-Upload /></el-icon>上传文件</el-button>
            </el-col>
            <el-col :span="1.5">
              <el-button
                  color="#626aef"
                  :dark="true"
                  type="success"
                  @click="handleAdd('image')"
              ><el-icon><ele-PictureFilled /></el-icon>上传图片</el-button>
            </el-col>
            <el-col :span="1.5">
              <el-button
                  type="success"
                  @click="handleAdd('bigFile')"
              ><i class="iconfont icon-shangchuan"></i>上传大文件</el-button>
            </el-col>
            <el-col :span="1.5">
              <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleConfirm()"
                  v-auth="'api/v1/system/sysAttachment/delete'"
              ><el-icon><ele-Delete /></el-icon>确认返回</el-button>
            </el-col>
          </el-row>
        </div>
        <el-table ref="uploadTableRef" v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="文件原始名" align="center" prop="name" show-overflow-tooltip
                           min-width="260px"
          />
          <el-table-column label="上传类型" align="center" prop="kind" :formatter="kindFormat"
                           min-width="100px"
          />
          <el-table-column label="本地路径" align="center" prop="path" min-width="100px">
            <template #default="scope">
              <el-image
                  v-if="scope.row.kind=='image'"
                  style="width: 60px; height: 60px"
                  :src="proxy.getUpFileUrl(scope.row.path)"
                  :zoom-rate="1.2"
                  :max-scale="7"
                  :min-scale="0.2"
                  :preview-src-list="[proxy.getUpFileUrl(scope.row.path)]"
                  :initial-index="0"
                  preview-teleported
                  hide-on-click-modal
                  fit="fill"
              />
              <el-image
                  v-else
                  style="width: 60px; height: 60px">
                <template #error>
                  <div class="image-slot">
                    {{getExt(scope.row.name)}}
                  </div>
                </template>
              </el-image>
            </template>
          </el-table-column>
          <el-table-column label="文件大小" align="center" prop="size" :formatter="formatFileSize"
                           min-width="100px"
          />
          <el-table-column label="上传时间" align="center" prop="updatedAt"
                           min-width="150px"
          >
            <template #default="scope">
              <span>{{ proxy.parseTime(scope.row.updatedAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" class-name="small-padding" min-width="100px" fixed="right">
            <template #default="scope">
              <el-button
                  type="primary"
                  link
                  @click="handleSelect(scope.row)"
              ><el-icon><ele-Select /></el-icon>选择</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="sysAttachmentList"
        />
      </el-card>
      <ApiV1SystemSysAttachmentEdit
          ref="editRef"
          @sysAttachmentList="sysAttachmentList"
      ></ApiV1SystemSysAttachmentEdit>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import {toRefs, reactive, ref, computed,getCurrentInstance} from 'vue';
import {FormInstance, type UploadProps, UploadUserFile} from 'element-plus';
import {
    listSysAttachment,
} from "/@/api/system/sysAttachment";
import {
    SysAttachmentTableColumns,
    SysAttachmentInfoData,
    SysAttachmentTableDataState,
} from "/@/views/system/sysAttachment/list/component/model"
import ApiV1SystemSysAttachmentEdit from "/@/views/system/sysAttachment/list/component/edit.vue"
defineOptions({ name: "apiV1SystemSysAttachmentList"})
const props = defineProps({
  fileType:{
      type:String,
      default:function(){
        return 'image'
      }
    },
  modelValue:{
    type:Array,
    default:function(){
      return []
    }
  },
  limit:{
    type:Number,
    default:function(){
      return 10
    }
  }
})
const emit = defineEmits(['update:modelValue'])
const dataFileList = computed({
  get: () => {
    let value:Array<UploadUserFile> = props.modelValue as UploadUserFile[]|| []
    value.map((item: UploadUserFile)=>{
      if(item.url){
        item.url = proxy.getUpFileUrl(item.url)
      }
      return item
    })
    return value
  },
  set: val => {
    console.log('setVal',val)
    emit('update:modelValue', val)
  }
});
const uploadTableRef = ref()
const {proxy} = <any>getCurrentInstance()
const isShowDialog = ref(false)
const loading = ref(false)
const queryRef = ref()
const editRef = ref();
const dialogImageUrl = ref('')
const dialogVisible = ref(false)
// 是否显示所有搜索选项
const showAll =  ref(false)
// 非单个禁用
const single = ref(true)
// 非多个禁用
const multiple =ref(true)
// 字典选项数据
const {sys_upload_file_type} = proxy.useDict( 'sys_upload_file_type')
const state = reactive<SysAttachmentTableDataState>({
    ids:[],
    tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
            pageNum: 1,
            pageSize: 10,
            appId: undefined,
            drive: undefined,
            name: undefined,
            kind: props.fileType=='image'?'image':undefined,
            mimeType: undefined,
            status: undefined,
            createdAt: [],
            dateRange: []
        },
    },
});
const { tableData } = toRefs(state);

// 初始化表格数据
const initTableData = () => {
    sysAttachmentList()
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    state.tableData.param.status = true
    sysAttachmentList()
};
// 获取列表数据
const sysAttachmentList = ()=>{
  loading.value = true
  listSysAttachment(state.tableData.param).then((res:any)=>{
    state.tableData.data = res.data.list??[];
    state.tableData.total = res.data.total;
    loading.value = false
  })
};
// 上传类型字典翻译
const kindFormat = (row:SysAttachmentTableColumns) => {
    return proxy.selectDictLabel(sys_upload_file_type.value, row.kind);
}
// 多选框选中数据
const handleSelectionChange = (selection:Array<SysAttachmentInfoData>) => {
    state.ids = selection.map(item => item.id)
    single.value = selection.length!=1
    multiple.value = !selection.length
}
const handleAdd =  (upType:string)=>{
    editRef.value.openDialog(upType)
}
const handleSelect = (row: SysAttachmentTableColumns) => {
  let selected = true
  if(state.ids.includes(row.id)){
    selected = false
    state.ids.splice(state.ids.indexOf(row.id),1)
  }else{
    state.ids.push(row.id)
  }
  uploadTableRef.value.toggleRowSelection(row,selected)
};
const handleConfirm = () => {
  let ids:number[] = Array.from(new Set(state.ids));
  if(dataFileList.value.length+ids.length>props.limit){
    proxy.$message.error('最多只能选择'+props.limit+'个文件')
    return
  }
  let list :UploadUserFile[] = []
  state.tableData.data.map((item: SysAttachmentTableColumns)=>{
    if(ids.includes(item.id)){
      list.push({
        name: item.name,
        url:item.path
      })
    }
  })
  console.log('dataFileList.value=',dataFileList.value)
  dataFileList.value.push(...list)
  closeDialog()
}
const getExt = (fileName:string):string=>{
  // 查找最后一个点的位置
  const lastDotIndex = fileName.lastIndexOf('.');
  // 如果没有找到点，或者点在第一个字符（没有扩展名）
  if (lastDotIndex === -1 || lastDotIndex === 0) {
    return '';
  }
  // 提取并返回点后面部分的后缀（不带点）
  return fileName.slice(lastDotIndex + 1);
}
const formatFileSize = (row:SysAttachmentTableColumns):string =>{
  const bytes:number = row.size;
  if (bytes === 0) return '0 Bytes';
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
  const index = Math.floor(Math.log(bytes) / Math.log(1024)); // 计算索引
  const value = (bytes / Math.pow(1024, index)).toFixed(2); // 计算具体数值并保留两位小数
  return `${value} ${sizes[index]}`;
}
const openDialog = ()=>{
  reset()
  initTableData();
  isShowDialog.value = true
}
const closeDialog = ()=>{
  isShowDialog.value = false
}
const reset = ()=>{
  resetQuery(queryRef.value)
  state.ids = []
}
const handlePictureCardPreview: UploadProps['onPreview'] = (uploadFile) => {
  dialogImageUrl.value = uploadFile.url!
  dialogVisible.value = true
};
</script>
<style lang="scss" scoped>
    .system-sysAttachment-container{
      width: 100%;
    }
    .colBlock {
        display: block;
    }
    .colNone {
        display: none;
    }
    .ml-2{margin: 3px;}
    .form-item{
      margin-bottom: 8px;
    }
    .image-slot {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100%;
      height: 100%;
      background: var(--el-fill-color-light);
      color: var(--el-text-color-secondary);
      font-size: 20px;
    }

    .up-selector {
      width: 100px;
      height: 100px;
      border-radius: 6px;
      position: relative;
      overflow: hidden;
      transition: var(--el-transition-duration-fast);
      background-color: var(--el-fill-color-lighter);
      border: 1px dashed var(--el-border-color-darker);
      box-sizing: border-box;
      cursor: pointer;
      display: inline-flex;
      justify-content: center;
      vertical-align: top;

      .uploader-icon {
        font-size: 28px;
        color: #8c939d;
        width: 50px;
        height: 96px;
        text-align: center;
      }
    }
    .up-selector:hover {
      border-color: var(--el-color-primary);
    }
    .img-list{
      width: 100%;
      margin-top: 8px;
    }
    .img-list :deep(.el-upload.el-upload--picture-card){
      display: none;
    }
    .img-list:deep(.el-upload.el-upload--text){
      display: none;
    }
</style>
