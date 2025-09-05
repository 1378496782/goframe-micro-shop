<template>
  <!-- 附件管理详情抽屉 -->  
  <div class="system-sysAttachment-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>附件管理详情</h4>
      </template>
      <el-descriptions
              class="margin-top"
              :column="3"
              border
              style="margin: 8px;"
      >        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  文件ID
                </div>
              </template>
              {{ formData.id }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  应用ID
                </div>
              </template>
              {{ formData.appId }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">              
                <template #label>
                  <div class="cell-item">
                    上传驱动
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.drive, driveOptions,'value','label') }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  文件原始名
                </div>
              </template>
              {{ formData.name }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">              
                <template #label>
                  <div class="cell-item">
                    上传类型
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.kind, kindOptions,'value','label') }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  扩展类型
                </div>
              </template>
              {{ formData.mimeType }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  本地路径
                </div>
              </template>
              {{ formData.path }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  文件大小
                </div>
              </template>
              {{ formData.size }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  扩展名
                </div>
              </template>
              {{ formData.ext }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  md5校验码
                </div>
              </template>
              {{ formData.md5 }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  上传人ID
                </div>
              </template>
              {{ formData.createdBy }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                状态
              </div>
            </template>
            <el-switch  v-model="formData.status" class="ml-2" disabled />
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                创建时间
              </div>
            </template>
            {{ proxy.parseTime(formData.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}
          </el-descriptions-item>        
      </el-descriptions>
    </el-drawer>
  </div>
</template>
<script setup lang="ts">
  import { reactive, toRefs, defineComponent,ref,unref,getCurrentInstance,computed } from 'vue';
  import {ElMessageBox, ElMessage, FormInstance,UploadProps} from 'element-plus';  
  import {
    listSysAttachment,
    getSysAttachment,
    delSysAttachment,
    addSysAttachment,
    updateSysAttachment,    
  } from "/@/api/system/sysAttachment";  
  import {
    SysAttachmentTableColumns,
    SysAttachmentInfoData,
    SysAttachmentTableDataState,
    SysAttachmentEditState
  } from "/@/views/system/sysAttachment/list/component/model"
  defineOptions({ name: "ApiV1SystemSysAttachmentDetail"})  
  const props = defineProps({    
    driveOptions:{
      type:Array,
      default:()=>[]
    },    
    kindOptions:{
      type:Array,
      default:()=>[]
    },    
  })  
  const {proxy} = <any>getCurrentInstance()
  const formRef = ref<HTMLElement | null>(null);
  const menuRef = ref();  
  const state = reactive<SysAttachmentEditState>({
    loading:false,
    isShowDialog: false,
    formData: {      
      id: undefined,      
      appId: undefined,      
      drive: undefined,      
      name: undefined,      
      kind: undefined,      
      mimeType: undefined,      
      path: undefined,      
      size: undefined,      
      ext: undefined,      
      md5: undefined,      
      createdBy: undefined,      
      status: undefined,      
      createdAt: undefined,      
      updatedAt: undefined,      
    },
    // 表单校验
    rules: {      
      id : [
          { required: true, message: "文件ID不能为空", trigger: "blur" }
      ],      
      appId : [
          { required: true, message: "应用ID不能为空", trigger: "blur" }
      ],      
      name : [
          { required: true, message: "文件原始名不能为空", trigger: "blur" }
      ],      
      status : [
          { required: true, message: "状态不能为空", trigger: "blur" }
      ],      
    }
  });
  const { isShowDialog,formData } = toRefs(state);
  // 打开弹窗
  const openDialog = (row?: SysAttachmentInfoData) => {
    resetForm();
    if(row) {
      getSysAttachment(row.id!).then((res:any)=>{
        const data = res.data;        
        data.createdBy = data.createdUser?.userNickname        
        state.formData = data;
      })
    }
    state.isShowDialog = true;
  };
  // 关闭弹窗
  const closeDialog = () => {
    state.isShowDialog = false;
  };
  defineExpose({
    openDialog,
  });
  // 取消
  const onCancel = () => {
    closeDialog();
  };
  const resetForm = ()=>{
    state.formData = {      
      id: undefined,      
      appId: undefined,      
      drive: undefined,      
      name: undefined,      
      kind: undefined,      
      mimeType: undefined,      
      path: undefined,      
      size: undefined,      
      ext: undefined,      
      md5: undefined,      
      createdBy: undefined,      
      status: undefined,      
      createdAt: undefined,      
      updatedAt: undefined,      
    }
  };  
</script>
<style scoped>  
  .system-sysAttachment-detail :deep(.el-form-item--large .el-form-item__label){
    font-weight: bolder;
  }
  .pic-block{
    margin-right: 8px;
  }
  .file-block{
    width: 100%;
    border: 1px solid var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
    margin-bottom: 5px;
    padding: 3px 6px;
  }
  .ml-2{margin-right: 5px;}
</style>