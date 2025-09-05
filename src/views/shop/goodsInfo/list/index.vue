<template>
  <div class="shop-goodsInfo-container">
    <el-card shadow="hover">
        <div class="shop-goodsInfo-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="ID" prop="id">
                    <el-input
                        v-model="tableData.param.id"
                        placeholder="请输入ID"
                        clearable                        
                        @keyup.enter.native="goodsInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="名称" prop="name">
                    <el-input
                        v-model="tableData.param.name"
                        placeholder="请输入名称"
                        clearable                        
                        @keyup.enter.native="goodsInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="goodsInfoList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="支持单图,多图" prop="images">
                    <el-select filterable v-model="tableData.param.images" placeholder="请选择支持单图,多图" clearable style="width:200px;">
                        <el-option label="请选择字典生成" value="" />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="价格(分)" prop="price">
                    <el-input
                        v-model="tableData.param.price"
                        placeholder="请输入价格(分)"
                        clearable                        
                        @keyup.enter.native="goodsInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="一级分类" prop="level1CategoryId">
                    <el-select filterable v-model="tableData.param.level1CategoryId" placeholder="请选择一级分类" clearable   style="width:200px;">
                      <el-option                      
                          v-for="item in level1CategoryIdOptions"                      
                          :key="item.key"
                          :label="item.value"
                          :value="item.key"
                      />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="二级分类" prop="level2CategoryId">
                    <el-select filterable v-model="tableData.param.level2CategoryId" placeholder="请选择二级分类" clearable   style="width:200px;">
                      <el-option                      
                          v-for="item in level2CategoryIdOptions"                      
                          :key="item.key"
                          :label="item.value"
                          :value="item.key"
                      />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="三级分类" prop="level3CategoryId">
                    <el-select filterable v-model="tableData.param.level3CategoryId" placeholder="请选择三级分类" clearable   style="width:200px;">
                      <el-option                      
                          v-for="item in level3CategoryIdOptions"                      
                          :key="item.key"
                          :label="item.value"
                          :value="item.key"
                      />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="品牌" prop="brand">
                    <el-input
                        v-model="tableData.param.brand"
                        placeholder="请输入品牌"
                        clearable                        
                        @keyup.enter.native="goodsInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="库存" prop="stock">
                    <el-input
                        v-model="tableData.param.stock"
                        placeholder="请输入库存"
                        clearable                        
                        @keyup.enter.native="goodsInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="销量" prop="sale">
                    <el-input
                        v-model="tableData.param.sale"
                        placeholder="请输入销量"
                        clearable                        
                        @keyup.enter.native="goodsInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="标签" prop="tags">
                    <el-input
                        v-model="tableData.param.tags"
                        placeholder="请输入标签"
                        clearable                        
                        @keyup.enter.native="goodsInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="商品详情" prop="detailInfo">
                    <el-select filterable v-model="tableData.param.detailInfo" placeholder="请选择商品详情" clearable style="width:200px;">
                        <el-option label="请选择字典生成" value="" />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="createdAt">
                    <el-date-picker
                        clearable  style="width: 200px"
                        v-model="tableData.param.createdAt"
                        format="YYYY-MM-DD HH:mm:ss"
                        value-format="YYYY-MM-DD HH:mm:ss"                    
                        type="datetime"
                        placeholder="选择"                    
                    ></el-date-picker>
                  </el-form-item>
                </el-col>            
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="goodsInfoList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                        {{ word }}
                        <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                        <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>            
              </el-row>
            </el-form>
            <el-row :gutter="10" class="mb8">
              <el-col :span="1.5">
                <el-button
                  type="primary"
                  @click="handleAdd"
                  v-auth="'api/v1/shop/goodsInfo/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/shop/goodsInfo/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/shop/goodsInfo/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>             
             <el-col :span="1.5">
                <el-button
                        type="warning"
                        @click="handleExport()"
                        v-auth="'api/v1/shop/goodsInfo/export'"
                ><el-icon><ele-Download /></el-icon>导出Excel</el-button>
             </el-col>            
                <el-col :span="1.5">
                    <el-button
                            type="success"
                            @click="handleImport()"
                            v-auth="'api/v1/shop/goodsInfo/import'"
                    ><el-icon><ele-Upload /></el-icon>导入Excel</el-button>
                </el-col>            
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="ID" align="center" prop="id"
            min-width="150px"            
             />          
          <el-table-column label="名称" align="center" prop="name"
            min-width="150px"            
             />          
          <el-table-column label="支持单图,多图" align="center" prop="images"
            min-width="150px"            
             />          
          <el-table-column label="价格(分)" align="center" prop="price"
            min-width="150px"            
             />          
          <el-table-column label="一级分类" align="center" prop="linkedLevel1CategoryId.name"
            min-width="150px"            
             />          
          <el-table-column label="二级分类" align="center" prop="linkedLevel2CategoryId.name"
            min-width="150px"            
             />          
          <el-table-column label="三级分类" align="center" prop="linkedLevel3CategoryId.name"
            min-width="150px"            
             />          
          <el-table-column label="品牌" align="center" prop="brand"
            min-width="150px"            
             />          
          <el-table-column label="库存" align="center" prop="stock"
            min-width="150px"            
             />          
          <el-table-column label="销量" align="center" prop="sale"
            min-width="150px"            
             />          
          <el-table-column label="标签" align="center" prop="tags"
            min-width="150px"            
             />          
          <el-table-column label="商品详情" align="center" prop="detailInfo"
            min-width="150px"            
             />          
          <el-table-column label="" align="center" prop="createdAt"
            min-width="150px"            
            >
            <template #default="scope">
                <span>{{ proxy.parseTime(scope.row.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>        
          <el-table-column label="操作" align="center" class-name="small-padding" min-width="200px" fixed="right">
            <template #default="scope">            
              <el-button
                type="primary"
                link
                @click="handleView(scope.row)"
                v-auth="'api/v1/shop/goodsInfo/get'"
              ><el-icon><ele-View /></el-icon>详情</el-button>              
              <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/shop/goodsInfo/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/shop/goodsInfo/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="goodsInfoList"
        />
    </el-card>
    <ApiV1ShopGoodsInfoEdit
       ref="editRef"       
       :level1CategoryIdOptions="level1CategoryIdOptions"       
       :level2CategoryIdOptions="level2CategoryIdOptions"       
       :level3CategoryIdOptions="level3CategoryIdOptions"       
       @goodsInfoList="goodsInfoList"
    ></ApiV1ShopGoodsInfoEdit>
    <ApiV1ShopGoodsInfoDetail
      ref="detailRef"      
      :level1CategoryIdOptions="level1CategoryIdOptions"      
      :level2CategoryIdOptions="level2CategoryIdOptions"      
      :level3CategoryIdOptions="level3CategoryIdOptions"      
      @goodsInfoList="goodsInfoList"
    ></ApiV1ShopGoodsInfoDetail>    
    <loadExcel ref="loadExcelGoodsInfoRef" @getList="goodsInfoList"
               upUrl="api/v1/shop/goodsInfo/import"
               tplUrl="/api/v1/shop/goodsInfo/excelTemplate"></loadExcel>    
  </div>
</template>
<script setup lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listGoodsInfo,
    getGoodsInfo,
    delGoodsInfo,
    addGoodsInfo,
    updateGoodsInfo,    
    linkedDataSearch    
} from "/@/api/shop/goodsInfo";
import {
    GoodsInfoTableColumns,
    GoodsInfoInfoData,
    GoodsInfoTableDataState,    
    LinkedGoodsInfoCategoryInfo,    
} from "/@/views/shop/goodsInfo/list/component/model"
import ApiV1ShopGoodsInfoEdit from "/@/views/shop/goodsInfo/list/component/edit.vue"
import ApiV1ShopGoodsInfoDetail from "/@/views/shop/goodsInfo/list/component/detail.vue"
import {downLoadXml} from "/@/utils/zipdownload";
import loadExcel from "/@/components/loadExcel/index.vue"
defineOptions({ name: "apiV1ShopGoodsInfoList"})
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const queryRef = ref()
const editRef = ref();
const detailRef = ref();
const loadExcelGoodsInfoRef = ref();
// 是否显示所有搜索选项
const showAll =  ref(false)
// 非单个禁用
const single = ref(true)
// 非多个禁用
const multiple =ref(true)
const word = computed(()=>{
    if(showAll.value === false) {
        //对文字进行处理
        return "展开搜索";
    } else {
        return "收起搜索";
    }
})
// 字典选项数据
const {    
} = proxy.useDict(    
)
// level1CategoryIdOptions关联表数据
const level1CategoryIdOptions = ref<Array<ItemOptions>>([])
// level2CategoryIdOptions关联表数据
const level2CategoryIdOptions = ref<Array<ItemOptions>>([])
// level3CategoryIdOptions关联表数据
const level3CategoryIdOptions = ref<Array<ItemOptions>>([])
const state = reactive<GoodsInfoTableDataState>({
    ids:[],
    tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
            pageNum: 1,
            pageSize: 10,            
            id: undefined,            
            name: undefined,            
            images: undefined,            
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
            dateRange: []
        },
    },
});
const { tableData } = toRefs(state);
// 页面加载时
onMounted(() => {
    initTableData();
});
// 初始化表格数据
const initTableData = () => {    
    linkedData()    
    goodsInfoList()
};
const linkedData = ()=>{
    linkedDataSearch().then((res:any)=>{        
        //关联category_info表选项        
        level1CategoryIdOptions.value = proxy.setItems(res, 'id', 'name','linkedGoodsInfoCategoryInfo')        
        //关联category_info表选项        
        level2CategoryIdOptions.value = proxy.setItems(res, 'id', 'name','linkedGoodsInfoCategoryInfo')        
        //关联category_info表选项        
        level3CategoryIdOptions.value = proxy.setItems(res, 'id', 'name','linkedGoodsInfoCategoryInfo')        
    })
}
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    goodsInfoList()
};
// 获取列表数据
const goodsInfoList = ()=>{
  loading.value = true
  listGoodsInfo(state.tableData.param).then((res:any)=>{
    let list = res.data.list??[];    
    state.tableData.data = list;
    state.tableData.total = res.data.total;
    loading.value = false
  })
};
const toggleSearch = () => {
    showAll.value = !showAll.value;
}
// 多选框选中数据
const handleSelectionChange = (selection:Array<GoodsInfoInfoData>) => {
    state.ids = selection.map(item => item.id)
    single.value = selection.length!=1
    multiple.value = !selection.length
}
const handleAdd =  ()=>{
    editRef.value.openDialog()
}
const handleUpdate = (row: GoodsInfoTableColumns|null) => {
    if(!row){
        row = state.tableData.data.find((item:GoodsInfoTableColumns)=>{
            return item.id ===state.ids[0]
        }) as GoodsInfoTableColumns
    }
    editRef.value.openDialog(toRaw(row));
};
const handleDelete = (row: GoodsInfoTableColumns|null) => {
    let msg = '你确定要删除所选数据？';
    let id:number[] = [] ;
    if(row){
    msg = `此操作将永久删除数据，是否继续?`
    id = [row.id]
    }else{
    id = state.ids
    }
    if(id.length===0){
        ElMessage.error('请选择要删除的数据。');
        return
    }
    ElMessageBox.confirm(msg, '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
    })
        .then(() => {
            delGoodsInfo(id).then(()=>{
                ElMessage.success('删除成功');
                goodsInfoList();
            })
        })
        .catch(() => {});
}
const handleView = (row:GoodsInfoTableColumns)=>{
    detailRef.value.openDialog(toRaw(row));
}
//导出excel
const handleExport = ()=>{
    downLoadXml('/api/v1/shop/goodsInfo/export',state.tableData.param,'get')
}
const handleImport=()=>{
    loadExcelGoodsInfoRef.value.open()
}
</script>
<style lang="scss" scoped>
    .colBlock {
        display: block;
    }
    .colNone {
        display: none;
    }
    .ml-2{margin: 3px;}
</style>