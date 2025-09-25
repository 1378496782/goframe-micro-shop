<template>
  <!-- 优惠券详情抽屉 -->  
  <div class="shop-couponInfo-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>优惠券详情</h4>
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
                  名称
                </div>
              </template>
              {{ formData.name }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">              
                <template #label>
                  <div class="cell-item">
                    类型
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.type, typeOptions,'value','label') }}            
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
                过期时间
              </div>
            </template>
            {{ proxy.parseTime(formData.deadline, '{y}-{m}-{d} {h}:{i}:{s}') }}
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
  defineOptions({ name: "ApiV1ShopCouponInfoDetail"})  
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
  const { isShowDialog,formData } = toRefs(state);
  // 打开弹窗
  const openDialog = (row?: CouponInfoInfoData) => {
    resetForm();
    if(row) {
      getCouponInfo(row.id!).then((res:any)=>{
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
  .shop-couponInfo-detail :deep(.el-form-item--large .el-form-item__label){
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