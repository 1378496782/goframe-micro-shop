<template>
  <!-- 用户优惠券管理详情抽屉 -->  
  <div class="shop-userCouponInfo-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>用户优惠券管理详情</h4>
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
                  用户id
                </div>
              </template>
              {{ formData.userId }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">                  
                    <template #label>
                      <div class="cell-item">
                        优惠券id
                      </div>
                    </template>
                    {{ formData.linkedCouponId?formData.linkedCouponId.name:'' }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">              
                <template #label>
                  <div class="cell-item">
                    状态
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.status, statusOptions,'value','label') }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  优惠金额（元）
                </div>
              </template>
              {{ formData.amount }}            
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
  defineOptions({ name: "ApiV1ShopUserCouponInfoDetail"})  
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
      linkedCouponId:{id:undefined,name:undefined },      
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
  const { isShowDialog,formData } = toRefs(state);
  // 打开弹窗
  const openDialog = (row?: UserCouponInfoInfoData) => {
    resetForm();
    if(row) {
      getUserCouponInfo(row.id!).then((res:any)=>{
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
      userId: undefined,      
      couponId: undefined,      
      linkedCouponId:{id:undefined,name:undefined },      
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
  //关联coupon_info表选项
  const getCouponInfoItemsCouponId = () => {
    emit("getCouponInfoItemsCouponId")
  }
  const getCouponIdOp = computed(()=>{
    getCouponInfoItemsCouponId()
    return props.couponIdOptions
  })  
</script>
<style scoped>  
  .shop-userCouponInfo-detail :deep(.el-form-item--large .el-form-item__label){
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