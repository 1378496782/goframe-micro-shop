<template>  
  <div class="shop-categoryInfo-edit">
    <!-- 添加或修改商品分类对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.shop-categoryInfo-edit .el-dialog', '.shop-categoryInfo-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'商品分类'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">          
        <el-form-item label="父级id" prop="parentId">
          <el-select filterable clearable v-model="formData.parentId" placeholder="请选择父级id"  >
              <el-option              
                  v-for="item in parentIdOptions"              
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
        </el-form-item>        
        <el-form-item label="名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入名称" />
        </el-form-item>        
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
        <el-form-item label="等级" prop="level">
          <el-input-number v-model="formData.level" placeholder="请输入等级" />
        </el-form-item>        
        <el-form-item label="排序" prop="sort">
          <el-input v-model="formData.sort" placeholder="请输入排序" />
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
  listCategoryInfo,
  getCategoryInfo,
  delCategoryInfo,
  addCategoryInfo,
  updateCategoryInfo,  
} from "/@/api/shop/categoryInfo";
import {getToken} from "/@/utils/gfast"
import {
  CategoryInfoTableColumns,
  CategoryInfoInfoData,
  CategoryInfoTableDataState,
  CategoryInfoEditState
} from "/@/views/shop/categoryInfo/list/component/model"
defineOptions({ name: "ApiV1ShopCategoryInfoEdit"})
const emit = defineEmits(['categoryInfoList'])
  const props = defineProps({    
    parentIdOptions:{
      type:Array,
      default:()=>[]
    },    
  })
const baseURL:string|undefined|boolean = import.meta.env.VITE_API_URL
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
//图片上传地址
const imageUrlPicUrl = ref('')
//上传加载
const upLoadingPicUrl = ref(false)
const state = reactive<CategoryInfoEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    parentId: undefined,    
    name: undefined,    
    picUrl: undefined,    
    level: undefined,    
    sort: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
    linkedCategoryInfoCategoryInfo: {      
      id:undefined,    //      
      name:undefined,    //      
    },    
  },
  // 表单校验
  rules: {    
    id : [
        { required: true, message: "ID不能为空", trigger: "blur" }
    ],    
    name : [
        { required: true, message: "名称不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: CategoryInfoInfoData) => {
  resetForm();
  if(row) {
    getCategoryInfo(row.id!).then((res:any)=>{
      const data = res.data;      
      data.parentId = parseInt(data.parentId)      
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
      addCategoryInfo(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('categoryInfoList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateCategoryInfo(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('categoryInfoList')
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
    parentId: undefined,    
    name: undefined,    
    picUrl: undefined,    
    level: undefined,    
    sort: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
    linkedCategoryInfoCategoryInfo: {      
      id:undefined,    //      
      name:undefined,    //      
    },    
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
  .shop-categoryInfo-edit :deep(.avatar-uploader .avatar) {
    width: 178px;
    height: 178px;
    display: block;
  }
  .shop-categoryInfo-edit :deep(.avatar-uploader .el-upload) {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
  }
  .shop-categoryInfo-edit :deep(.avatar-uploader .el-upload:hover) {
    border-color: var(--el-color-primary);
  }
  .shop-categoryInfo-edit :deep(.el-icon.avatar-uploader-icon) {
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