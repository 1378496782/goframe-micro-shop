<template>  
  <div class="shop-userInfo-edit">
    <!-- 添加或修改用户对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.shop-userInfo-edit .el-dialog', '.shop-userInfo-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'用户'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="用户名" prop="name">
          <el-input v-model="formData.name" placeholder="请输入用户名" />
        </el-form-item>        
        <el-form-item label="头像" prop="avatar">
          <el-input v-model="formData.avatar" placeholder="请输入头像" />
        </el-form-item>          
        <el-form-item label="1男 2女" prop="sex">
          <el-select filterable clearable v-model="formData.sex" placeholder="请选择1男 2女" >
            <el-option label="请选择字典生成" value="" />
          </el-select>
        </el-form-item>        
        <el-form-item label="1正常 2拉黑冻结" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio>请选择字典生成</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="个性签名" prop="sign">
          <el-input v-model="formData.sign" placeholder="请输入个性签名" />
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
defineOptions({ name: "ApiV1ShopUserInfoEdit"})
const emit = defineEmits(['userInfoList'])
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
    status: undefined,    
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
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: UserInfoInfoData) => {
  resetForm();
  if(row) {
    getUserInfo(row.id!).then((res:any)=>{
      const data = res.data;      
      data.sex = parseInt(data.sex)      
      data.status = parseInt(data.status)      
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
      addUserInfo(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('userInfoList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateUserInfo(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('userInfoList')
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
    name: undefined,    
    avatar: undefined,    
    password: undefined,    
    userSalt: undefined,    
    sex: undefined,    
    status: '' ,    
    sign: undefined,    
    secretAnswer: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
  }  
};
</script>
<style scoped>  
  .kv-label{margin-bottom: 15px;font-size: 14px;}
  .mini-btn i.el-icon{margin: unset;}
  .kv-row{margin-bottom: 12px;}
</style>