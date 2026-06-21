<template>  
  <div class="shop-userCouponInfo-edit">
    <!-- 添加或修改用户优惠券管理对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.shop-userCouponInfo-edit .el-dialog', '.shop-userCouponInfo-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'用户优惠券管理'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="用户id" prop="userId">
          <el-input v-model="formData.userId" placeholder="请输入用户id" />
        </el-form-item>        
        <el-form-item label="优惠券id" prop="couponId">
          <el-input v-model="formData.couponId" placeholder="请输入优惠券id" />
        </el-form-item>          
        <el-form-item label="状态" prop="status">
          <el-select filterable clearable v-model="formData.status" placeholder="请选择状态" >
            <el-option
              v-for="dict in statusOptions"
              :key="dict.value"
              :label="dict.label"              
                  :value="dict.value"              
            ></el-option>
          </el-select>
        </el-form-item>        
        <el-form-item label="优惠金额（元）" prop="amount">
          <el-input v-model="formData.amount" placeholder="请输入优惠金额（元）" />
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
  listUserCouponInfo,
  getUserCouponInfo,
  delUserCouponInfo,
  addUserCouponInfo,
  updateUserCouponInfo,  
} from "/@/api/shop/userCouponInfo";
import {
  UserCouponInfoTableColumns,
  UserCouponInfoInfoData,
  UserCouponInfoTableDataState,
  UserCouponInfoEditState
} from "/@/views/shop/userCouponInfo/list/component/model"
defineOptions({ name: "ApiV1ShopUserCouponInfoEdit"})
const emit = defineEmits(['userCouponInfoList'])
  const props = defineProps({    
    couponIdOptions:{
      type:Array,
      default:()=>[]
    },    
    statusOptions:{
      type:Array,
      default:()=>[]
    },    
  })
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<UserCouponInfoEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    userId: undefined,    
    couponId: undefined,    
    status: undefined,    
    amount: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
    linkedUserCouponInfoCouponInfo: {      
      id:undefined,    //      
      name:undefined,    // 优惠券名称      
    },    
  },
  // 表单校验
  rules: {    
    id : [
        { required: true, message: "不能为空", trigger: "blur" }
    ],    
    couponId : [
        { required: true, message: "优惠券id不能为空", trigger: "blur" }
    ],    
    status : [
        { required: true, message: "状态不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: UserCouponInfoInfoData) => {
  resetForm();
  if(row) {
    getUserCouponInfo(row.id!).then((res:any)=>{
      const data = res.data;      
      data.status = ''+data.status      
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
      addUserCouponInfo(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('userCouponInfoList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateUserCouponInfo(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('userCouponInfoList')
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
    userId: undefined,    
    couponId: undefined,    
    status: undefined,    
    amount: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
    linkedUserCouponInfoCouponInfo: {      
      id:undefined,    //      
      name:undefined,    // 优惠券名称      
    },    
  }  
};
</script>
<style scoped>  
  .kv-label{margin-bottom: 15px;font-size: 14px;}
  .mini-btn i.el-icon{margin: unset;}
  .kv-row{margin-bottom: 12px;}
</style>