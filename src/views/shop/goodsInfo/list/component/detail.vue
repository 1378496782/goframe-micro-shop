<template>
  <!-- 商品详情抽屉 -->  
  <div class="shop-goodsInfo-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>商品详情</h4>
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
                封面图
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
                  价格(分)
                </div>
              </template>
              {{ formData.price }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">                  
                    <template #label>
                      <div class="cell-item">
                        一级分类
                      </div>
                    </template>
                    {{ formData.linkedLevel1CategoryId?formData.linkedLevel1CategoryId.name:'' }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">                  
                    <template #label>
                      <div class="cell-item">
                        二级分类
                      </div>
                    </template>
                    {{ formData.linkedLevel2CategoryId?formData.linkedLevel2CategoryId.name:'' }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">                  
                    <template #label>
                      <div class="cell-item">
                        三级分类
                      </div>
                    </template>
                    {{ formData.linkedLevel3CategoryId?formData.linkedLevel3CategoryId.name:'' }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  品牌
                </div>
              </template>
              {{ formData.brand }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  库存
                </div>
              </template>
              {{ formData.stock }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  销量
                </div>
              </template>
              {{ formData.sale }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  标签
                </div>
              </template>
              {{ formData.tags }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                商品详情
              </div>
            </template>
            <div v-html="formData.detailInfo"></div>
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">                
              </div>
            </template>
            {{ proxy.parseTime(formData.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  排序 倒叙
                </div>
              </template>
              {{ formData.sort }}            
          </el-descriptions-item>        
      </el-descriptions>
    </el-drawer>
  </div>
</template>
<script setup lang="ts">
  import { reactive, toRefs, defineComponent,ref,unref,getCurrentInstance,computed } from 'vue';
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
  defineOptions({ name: "ApiV1ShopGoodsInfoDetail"})  
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
      linkedLevel1CategoryId:{id:undefined,name:undefined },      
      level2CategoryId: undefined,      
      linkedLevel2CategoryId:{id:undefined,name:undefined },      
      level3CategoryId: undefined,      
      linkedLevel3CategoryId:{id:undefined,name:undefined },      
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
  const { isShowDialog,formData } = toRefs(state);
  // 打开弹窗
  const openDialog = (row?: GoodsInfoInfoData) => {
    resetForm();
    if(row) {
      getGoodsInfo(row.id!).then((res:any)=>{
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
      name: undefined,      
      images: [] ,      
      picUrl: undefined,      
      price: undefined,      
      level1CategoryId: undefined,      
      linkedLevel1CategoryId:{id:undefined,name:undefined },      
      level2CategoryId: undefined,      
      linkedLevel2CategoryId:{id:undefined,name:undefined },      
      level3CategoryId: undefined,      
      linkedLevel3CategoryId:{id:undefined,name:undefined },      
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
  };  
  const setUpImgListImages = (data:any)=>{
    state.formData.images = data
  }  
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
  //关联category_info表选项
  const getCategoryInfoItemsLevel1CategoryId = () => {
    emit("getCategoryInfoItemsLevel1CategoryId")
  }
  const getLevel1CategoryIdOp = computed(()=>{
    getCategoryInfoItemsLevel1CategoryId()
    return props.level1CategoryIdOptions
  })  
  //关联category_info表选项
  const getCategoryInfoItemsLevel2CategoryId = () => {
    emit("getCategoryInfoItemsLevel2CategoryId")
  }
  const getLevel2CategoryIdOp = computed(()=>{
    getCategoryInfoItemsLevel2CategoryId()
    return props.level2CategoryIdOptions
  })  
  //关联category_info表选项
  const getCategoryInfoItemsLevel3CategoryId = () => {
    emit("getCategoryInfoItemsLevel3CategoryId")
  }
  const getLevel3CategoryIdOp = computed(()=>{
    getCategoryInfoItemsLevel3CategoryId()
    return props.level3CategoryIdOptions
  })  
  //富文本编辑器商品详情
  const setDetailInfoEditContent = (data:string) => {
    state.formData.detailInfo = data
  }  
</script>
<style scoped>  
  .shop-goodsInfo-detail :deep(.avatar-uploader .avatar) {
    width: 178px;
    height: 178px;
    display: block;
  }
  .shop-goodsInfo-detail :deep(.avatar-uploader .el-upload) {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
  }
  .shop-goodsInfo-detail :deep(.avatar-uploader .el-upload:hover) {
    border-color: var(--el-color-primary);
  }
  .shop-goodsInfo-detail :deep(.el-icon.avatar-uploader-icon) {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
  }  
  .shop-goodsInfo-detail :deep(.el-form-item--large .el-form-item__label){
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