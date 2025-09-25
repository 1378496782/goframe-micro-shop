<template>
  <!-- 用户详情抽屉 -->  
  <div class="shop-userInfo-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>用户详情</h4>
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
                </div>
              </template>
              {{ formData.id }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  用户名
                </div>
              </template>
              {{ formData.name }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  头像
                </div>
              </template>
              {{ formData.avatar }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  1男 2女
                </div>
              </template>
              {{ formData.sex }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  1正常 2拉黑冻结
                </div>
              </template>
              {{ formData.status }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  个性签名
                </div>
              </template>
              {{ formData.sign }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">                
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
    listUserInfo,
    getUserInfo,
    delUserInfo,
    addUserInfo,
    updateUserInfo,    
  } from "/@/api/shop/userInfo";  
  import {
    UserInfoTableColumns,
    UserInfoInfoData,
    UserInfoTableDataState,
    UserInfoEditState
  } from "/@/views/shop/userInfo/list/component/model"
  defineOptions({ name: "ApiV1ShopUserInfoDetail"})  
  const {proxy} = <any>getCurrentInstance()
  const formRef = ref<HTMLElement | null>(null);
  const menuRef = ref();  
  const state = reactive<UserInfoEditState>({
    loading:false,
    isShowDialog: false,
    formData: {      
      id: undefined,      
      name: undefined,      
      avatar: undefined,      
      password: undefined,      
      userSalt: undefined,      
      sex: undefined,      
      status: false ,      
      sign: undefined,      
      secretAnswer: undefined,      
      createdAt: undefined,      
      updatedAt: undefined,      
      deletedAt: undefined,      
    },
    // 表单校验
    rules: {      
      id : [
          { required: true, message: "不能为空", trigger: "blur" }
      ],      
      name : [
          { required: true, message: "用户名不能为空", trigger: "blur" }
      ],      
      status : [
          { required: true, message: "1正常 2拉黑冻结不能为空", trigger: "blur" }
      ],      
    }
  });
  const { isShowDialog,formData } = toRefs(state);
  // 打开弹窗
  const openDialog = (row?: UserInfoInfoData) => {
    resetForm();
    if(row) {
      getUserInfo(row.id!).then((res:any)=>{
        const data = res.data;        
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
      name: undefined,      
      avatar: undefined,      
      password: undefined,      
      userSalt: undefined,      
      sex: undefined,      
      status: false ,      
      sign: undefined,      
      secretAnswer: undefined,      
      createdAt: undefined,      
      updatedAt: undefined,      
      deletedAt: undefined,      
    }
  };  
</script>
<style scoped>  
  .shop-userInfo-detail :deep(.el-form-item--large .el-form-item__label){
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