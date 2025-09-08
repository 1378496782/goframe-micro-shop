<template>  
  <div class="shop-goodsInfo-edit">
    <!-- 添加或修改商品对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.shop-goodsInfo-edit .el-dialog', '.shop-goodsInfo-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'商品'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入名称" />
        </el-form-item>        
        <el-form-item label="封面图" prop="picUrl">
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
        <el-form-item label="价格(分)" prop="price">
          <el-input v-model="formData.price" placeholder="请输入价格(分)" />
        </el-form-item>          
        <el-form-item label="一级分类" prop="level1CategoryId">
          <el-select filterable clearable v-model="formData.level1CategoryId" placeholder="请选择一级分类"  >
              <el-option              
                  v-for="item in level1CategoryIdOptions"              
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
        </el-form-item>          
        <el-form-item label="二级分类" prop="level2CategoryId">
          <el-select filterable clearable v-model="formData.level2CategoryId" placeholder="请选择二级分类"  >
              <el-option              
                  v-for="item in level2CategoryIdOptions"              
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
        </el-form-item>          
        <el-form-item label="三级分类" prop="level3CategoryId">
          <el-select filterable clearable v-model="formData.level3CategoryId" placeholder="请选择三级分类"  >
              <el-option              
                  v-for="item in level3CategoryIdOptions"              
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
        </el-form-item>        
        <el-form-item label="品牌" prop="brand">
          <el-input v-model="formData.brand" placeholder="请输入品牌" />
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
        <el-form-item label="商品详情">
          <gf-ueditor editorId="ueGoodsInfoDetailInfo" v-model="formData.detailInfo"></gf-ueditor>
        </el-form-item>        
        <el-form-item label="排序 倒叙" prop="sort">
          <el-input v-model="formData.sort" placeholder="请输入排序 倒叙" />
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
import GfUeditor from "/@/components/ueditor/index.vue"
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
    level1CategoryIdOptions:{
      type:Array,
      default:()=>[]
    },    
    level2CategoryIdOptions:{
      type:Array,
      default:()=>[]
    },    
    level3CategoryIdOptions:{
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
    images: [] ,    
    picUrl: undefined,    
    price: undefined,    
    level1CategoryId: undefined,    
    level2CategoryId: undefined,    
    level3CategoryId: undefined,    
    brand: undefined,    
    stock: undefined,    
    sale: undefined,    
    tags: undefined,    
    detailInfo: undefined,    
    createdAt: undefined,    
    sort: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
    linkedGoodsInfoCategoryInfo: {      
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
    price : [
        { required: true, message: "价格(分)不能为空", trigger: "blur" }
    ],    
    level1CategoryId : [
        { required: true, message: "一级分类不能为空", trigger: "blur" }
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
      data.level1CategoryId = parseInt(data.level1CategoryId)      
      data.level2CategoryId = parseInt(data.level2CategoryId)      
      data.level3CategoryId = parseInt(data.level3CategoryId)      
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
    images: [] ,    
    picUrl: undefined,    
    price: undefined,    
    level1CategoryId: undefined,    
    level2CategoryId: undefined,    
    level3CategoryId: undefined,    
    brand: undefined,    
    stock: undefined,    
    sale: undefined,    
    tags: undefined,    
    detailInfo: undefined,    
    createdAt: undefined,    
    sort: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
    linkedGoodsInfoCategoryInfo: {      
      id:undefined,    //      
      name:undefined,    //      
    },    
  }  
  imageUrlPicUrl.value = ''  
};
//单图上传封面图
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