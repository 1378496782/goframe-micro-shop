<template>  
  <div class="shop-rotationInfo-edit">
    <!-- 添加或修改轮播图对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.shop-rotationInfo-edit .el-dialog', '.shop-rotationInfo-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'轮播图'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="图片" prop="picUrl">
          <el-upload
            v-loading="upLoadingPicUrl"
            :action="baseURL+'api/v1/system/upload/singleImg'"
            :before-upload="beforeAvatarUploadPicUrl"
            :data="setUpData()"
            :on-success="handleAvatarSuccessPicUrl"
            :show-file-list="false"
            class="avatar-uploader"
            name="file"
          >
            <div v-if="!proxy.isEmpty(imageUrlPicUrl)">
              <el-link type="danger" style="position: absolute; right: 5px; top: 6px;font-size: 18px;" :underline="false" @click.stop="deleteImageUrlPicUrl" title="删除">
                <el-icon><ele-DeleteFilled /></el-icon>
              </el-link>
              <img :src="imageUrlPicUrl" class="avatar">
            </div>
            <el-icon v-else class="avatar-uploader-icon"><ele-Plus /></el-icon>
          </el-upload>
        </el-form-item>        
        <el-form-item label="跳转链接" prop="link">
          <el-input v-model="formData.link" placeholder="请输入跳转链接" />
        </el-form-item>        
        <el-form-item label="排序字段" prop="sort">
          <el-input v-model="formData.sort" placeholder="请输入排序字段" />
        </el-form-item>       
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="onSubmit" :disabled="loading">确 定</el-button>
          <el-button @click="onCancel">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import { reactive, toRefs, ref,unref,getCurrentInstance,computed } from 'vue';
import {ElMessageBox, ElMessage, FormInstance,UploadProps} from 'element-plus';
import {
  listRotationInfo,
  getRotationInfo,
  delRotationInfo,
  addRotationInfo,
  updateRotationInfo,  
} from "/@/api/shop/rotationInfo";
import {getToken} from "/@/utils/gfast"
import {
  RotationInfoTableColumns,
  RotationInfoInfoData,
  RotationInfoTableDataState,
  RotationInfoEditState
} from "/@/views/shop/rotationInfo/list/component/model"
defineOptions({ name: "ApiV1ShopRotationInfoEdit"})
const emit = defineEmits(['rotationInfoList'])
const baseURL:string|undefined|boolean = import.meta.env.VITE_API_URL
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
//图片上传地址
const imageUrlPicUrl = ref('')
//上传加载
const upLoadingPicUrl = ref(false)
const state = reactive<RotationInfoEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    picUrl: undefined,    
    link: undefined,    
    sort: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
  },
  // 表单校验
  rules: {    
    id : [
        { required: true, message: "ID不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: RotationInfoInfoData) => {
  resetForm();
  if(row) {
    getRotationInfo(row.id!).then((res:any)=>{
      const data = res.data;      
      //单图地址赋值
      imageUrlPicUrl.value = data.picUrl ? proxy.getUpFileUrl(data.picUrl) : ''      
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
// 提交
const onSubmit = () => {
  const formWrap = unref(formRef) as any;
  if (!formWrap) return;
  formWrap.validate((valid: boolean) => {
    if (valid) {
      state.loading = true;
      if(!state.formData.id || state.formData.id===0){
        //添加
      addRotationInfo(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('rotationInfoList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateRotationInfo(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('rotationInfoList')
        }).finally(()=>{
          state.loading = false;
        })
      }
    }
  });
};
const resetForm = ()=>{
  state.formData = {    
    id: undefined,    
    picUrl: undefined,    
    link: undefined,    
    sort: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
  }  
  imageUrlPicUrl.value = ''  
};
//单图上传图片
const handleAvatarSuccessPicUrl:UploadProps['onSuccess'] = (res, file) => {
  if (res.code === 0) {
    imageUrlPicUrl.value = URL.createObjectURL(file.raw!)
    state.formData.picUrl = res.data.path
  } else {
    ElMessage.error(res.msg)
  }
  upLoadingPicUrl.value = false
}
const beforeAvatarUploadPicUrl = () => {
  upLoadingPicUrl.value = true
  return true
}
const setUpData = () => {
  return { token: getToken() }
}
const deleteImageUrlPicUrl = ()=>{
  state.formData.picUrl = ''
  imageUrlPicUrl.value = ''
}
</script>
<style scoped>  
  .shop-rotationInfo-edit :deep(.avatar-uploader .avatar) {
    width: 178px;
    height: 178px;
    display: block;
  }
  .shop-rotationInfo-edit :deep(.avatar-uploader .el-upload) {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
  }
  .shop-rotationInfo-edit :deep(.avatar-uploader .el-upload:hover) {
    border-color: var(--el-color-primary);
  }
  .shop-rotationInfo-edit :deep(.el-icon.avatar-uploader-icon) {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
  }  
  .kv-label{margin-bottom: 15px;font-size: 14px;}
  .mini-btn i.el-icon{margin: unset;}
  .kv-row{margin-bottom: 12px;}
</style>