<template>  
  <div class="shop-goodsInfo-edit">
    <!-- 添加或修改商品对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.shop-goodsInfo-edit .el-dialog', '.shop-goodsInfo-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'商品'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="名字" prop="name">
          <el-input v-model="formData.name" placeholder="请输入名字" />
        </el-form-item>        
        <el-form-item label="主图" prop="picUrl">
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
        <el-form-item label="详情配图" prop="images" >
          <upload-img :action="baseURL+'api/v1/system/upload/singleImg'" v-model="formData.images" :limit="10"></upload-img>
        </el-form-item>        
        <el-form-item label="价格(分)" prop="price">
          <el-input v-model="formData.price" placeholder="请输入价格(分)" />
        </el-form-item>        
        <el-form-item label="库存" prop="stock">
          <el-input v-model="formData.stock" placeholder="请输入库存" />
        </el-form-item>        
        <el-form-item label="销量" prop="sale">
          <el-input v-model="formData.sale" placeholder="请输入销量" />
        </el-form-item>        
        <el-form-item label="标签" prop="tags">
          <el-input v-model="formData.tags" placeholder="请输入标签" />
        </el-form-item>        
        <el-form-item label="排序 倒叙" prop="sort">
          <el-input v-model="formData.sort" placeholder="请输入排序 倒叙" />
        </el-form-item>          
        <el-form-item label="允许砍价" prop="enableBargain">
          <el-select filterable clearable v-model="formData.enableBargain" placeholder="请选择允许砍价" >
            <el-option
              v-for="dict in enableBargainOptions"
              :key="dict.value"
              :label="dict.label"              
                  :value="dict.value"              
            ></el-option>
          </el-select>
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
  listGoodsInfo,
  getGoodsInfo,
  delGoodsInfo,
  addGoodsInfo,
  updateGoodsInfo,  
} from "/@/api/shop/goodsInfo";
import {getToken} from "/@/utils/gfast"
import uploadImg from "/@/components/uploadImg/index.vue"
import {
  GoodsInfoTableColumns,
  GoodsInfoInfoData,
  GoodsInfoTableDataState,
  GoodsInfoEditState
} from "/@/views/shop/goodsInfo/list/component/model"
defineOptions({ name: "ApiV1ShopGoodsInfoEdit"})
const emit = defineEmits(['goodsInfoList'])
  const props = defineProps({    
    enableBargainOptions:{
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
const state = reactive<GoodsInfoEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    name: undefined,    
    picUrl: undefined,    
    images: [] ,    
    price: undefined,    
    level1CategoryId: undefined,    
    level2CategoryId: undefined,    
    level3CategoryId: undefined,    
    brand: undefined,    
    stock: undefined,    
    sale: undefined,    
    tags: undefined,    
    sort: undefined,    
    detailInfo: undefined,    
    enableBargain: undefined,    
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
        { required: true, message: "名字不能为空", trigger: "blur" }
    ],    
    price : [
        { required: true, message: "价格(分)不能为空", trigger: "blur" }
    ],    
    level1CategoryId : [
        { required: true, message: "1级分类id不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: GoodsInfoInfoData) => {
  resetForm();
  if(row) {
    getGoodsInfo(row.id!).then((res:any)=>{
      const data = res.data;      
      //单图地址赋值
      imageUrlPicUrl.value = data.picUrl ? proxy.getUpFileUrl(data.picUrl) : ''      
      data.images =data.images?JSON.parse(data.images) : []      
      data.enableBargain = ''+data.enableBargain      
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
      addGoodsInfo(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('goodsInfoList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateGoodsInfo(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('goodsInfoList')
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
    picUrl: undefined,    
    images: [] ,    
    price: undefined,    
    level1CategoryId: undefined,    
    level2CategoryId: undefined,    
    level3CategoryId: undefined,    
    brand: undefined,    
    stock: undefined,    
    sale: undefined,    
    tags: undefined,    
    sort: undefined,    
    detailInfo: undefined,    
    enableBargain: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
  }  
  imageUrlPicUrl.value = ''  
};
//单图上传主图
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
  .shop-goodsInfo-edit :deep(.avatar-uploader .avatar) {
    width: 178px;
    height: 178px;
    display: block;
  }
  .shop-goodsInfo-edit :deep(.avatar-uploader .el-upload) {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
  }
  .shop-goodsInfo-edit :deep(.avatar-uploader .el-upload:hover) {
    border-color: var(--el-color-primary);
  }
  .shop-goodsInfo-edit :deep(.el-icon.avatar-uploader-icon) {
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