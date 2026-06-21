<template>
  <!-- 商品分类详情抽屉 -->  
  <div class="shop-categoryInfo-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>商品分类详情</h4>
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
                  ID
                </div>
              </template>
              {{ formData.id }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">                  
                    <template #label>
                      <div class="cell-item">
                        父级id
                      </div>
                    </template>
                    {{ formData.linkedParentId?formData.linkedParentId.name:'' }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  名称
                </div>
              </template>
              {{ formData.name }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                图片
              </div>
            </template>
            <el-image
                    style="width: 150px; height: 150px"
                    v-if="!proxy.isEmpty(formData.picUrl)"
                    :src="proxy.getUpFileUrl(formData.picUrl)"
                    fit="contain"></el-image>
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  排序
                </div>
              </template>
              {{ formData.sort }}            
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
  defineOptions({ name: "ApiV1ShopCategoryInfoDetail"})  
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
      linkedParentId:{id:undefined,name:undefined },      
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
  const { isShowDialog,formData } = toRefs(state);
  // 打开弹窗
  const openDialog = (row?: CategoryInfoInfoData) => {
    resetForm();
    if(row) {
      getCategoryInfo(row.id!).then((res:any)=>{
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
  const resetForm = ()=>{
    state.formData = {      
      id: undefined,      
      parentId: undefined,      
      linkedParentId:{id:undefined,name:undefined },      
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
  };  
  //关联category_info表选项
  const getCategoryInfoItemsParentId = () => {
    emit("getCategoryInfoItemsParentId")
  }
  const getParentIdOp = computed(()=>{
    getCategoryInfoItemsParentId()
    return props.parentIdOptions
  })  
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
</script>
<style scoped>  
  .shop-categoryInfo-detail :deep(.avatar-uploader .avatar) {
    width: 178px;
    height: 178px;
    display: block;
  }
  .shop-categoryInfo-detail :deep(.avatar-uploader .el-upload) {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
  }
  .shop-categoryInfo-detail :deep(.avatar-uploader .el-upload:hover) {
    border-color: var(--el-color-primary);
  }
  .shop-categoryInfo-detail :deep(.el-icon.avatar-uploader-icon) {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
  }  
  .shop-categoryInfo-detail :deep(.el-form-item--large .el-form-item__label){
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