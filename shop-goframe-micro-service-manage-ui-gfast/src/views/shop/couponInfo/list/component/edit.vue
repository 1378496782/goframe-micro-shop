<template>  
  <div class="shop-couponInfo-edit">
    <!-- 添加或修改优惠券对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.shop-couponInfo-edit .el-dialog', '.shop-couponInfo-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'优惠券'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入名称" />
        </el-form-item>          
        <el-form-item label="类型" prop="type">
          <el-select filterable clearable v-model="formData.type" placeholder="请选择类型" >
            <el-option
              v-for="dict in typeOptions"
              :key="dict.value"
              :label="dict.label"              
                  :value="dict.value"              
            ></el-option>
          </el-select>
        </el-form-item>        
        <el-form-item label="优惠金额（元）" prop="amount">
          <el-input v-model="formData.amount" placeholder="请输入优惠金额（元）" />
        </el-form-item>        
        <el-form-item label="过期时间" prop="deadline">
          <el-date-picker clearable  style="width: 200px"
            v-model="formData.deadline"
            type="datetime"
            value-format="YYYY-MM-DD HH:mm:ss"
            placeholder="选择过期时间">
          </el-date-picker>
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
  listCouponInfo,
  getCouponInfo,
  delCouponInfo,
  addCouponInfo,
  updateCouponInfo,  
} from "/@/api/shop/couponInfo";
import {
  CouponInfoTableColumns,
  CouponInfoInfoData,
  CouponInfoTableDataState,
  CouponInfoEditState
} from "/@/views/shop/couponInfo/list/component/model"
defineOptions({ name: "ApiV1ShopCouponInfoEdit"})
const emit = defineEmits(['couponInfoList'])
  const props = defineProps({    
    typeOptions:{
      type:Array,
      default:()=>[]
    },    
  })
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<CouponInfoEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    goodsId: undefined,    
    name: undefined,    
    type: undefined,    
    amount: undefined,    
    deadline: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
  },
  // 表单校验
  rules: {    
    id : [
        { required: true, message: "ID不能为空", trigger: "blur" }
    ],    
    name : [
        { required: true, message: "名称不能为空", trigger: "blur" }
    ],    
    deadline : [
        { required: true, message: "过期时间不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: CouponInfoInfoData) => {
  resetForm();
  if(row) {
    getCouponInfo(row.id!).then((res:any)=>{
      const data = res.data;      
      data.type = ''+data.type      
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
      addCouponInfo(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('couponInfoList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateCouponInfo(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('couponInfoList')
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
    goodsId: undefined,    
    name: undefined,    
    type: undefined,    
    amount: undefined,    
    deadline: undefined,    
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