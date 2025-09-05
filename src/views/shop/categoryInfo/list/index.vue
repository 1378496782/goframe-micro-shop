<template>
  <div class="shop-categoryInfo-container">
    <el-card shadow="hover">
        <div class="shop-categoryInfo-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="ID" prop="id">
                    <el-input
                        v-model="tableData.param.id"
                        placeholder="请输入ID"
                        clearable                        
                        @keyup.enter.native="categoryInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="父级id" prop="parentId">
                    <el-select filterable v-model="tableData.param.parentId" placeholder="请选择父级id" clearable   style="width:200px;">
                      <el-option                      
                          v-for="item in parentIdOptions"                      
                          :key="item.key"
                          :label="item.value"
                          :value="item.key"
                      />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="categoryInfoList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="名称" prop="name">
                    <el-input
                        v-model="tableData.param.name"
                        placeholder="请输入名称"
                        clearable                        
                        @keyup.enter.native="categoryInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="图片" prop="picUrl">
                    <el-select filterable v-model="tableData.param.picUrl" placeholder="请选择图片" clearable style="width:200px;">
                        <el-option label="请选择字典生成" value="" />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="等级" prop="level">
                    <el-input
                        v-model="tableData.param.level"
                        placeholder="请输入等级"
                        clearable                        
                        @keyup.enter.native="categoryInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="排序" prop="sort">
                    <el-input
                        v-model="tableData.param.sort"
                        placeholder="请输入排序"
                        clearable                        
                        @keyup.enter.native="categoryInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="创建时间" prop="createdAt">
                    <el-date-picker
                        clearable  style="width: 200px"
                        v-model="tableData.param.createdAt"
                        format="YYYY-MM-DD HH:mm:ss"
                        value-format="YYYY-MM-DD HH:mm:ss"                    
                        type="datetime"
                        placeholder="选择创建时间"                    
                    ></el-date-picker>
                  </el-form-item>
                </el-col>            
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="categoryInfoList"><el-icon><ele-Search /></el-icon>搜索</el-button>
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
                  v-auth="'api/v1/shop/categoryInfo/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/shop/categoryInfo/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/shop/categoryInfo/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>             
             <el-col :span="1.5">
                <el-button
                        type="warning"
                        @click="handleExport()"
                        v-auth="'api/v1/shop/categoryInfo/export'"
                ><el-icon><ele-Download /></el-icon>导出Excel</el-button>
             </el-col>            
                <el-col :span="1.5">
                    <el-button
                            type="success"
                            @click="handleImport()"
                            v-auth="'api/v1/shop/categoryInfo/import'"
                    ><el-icon><ele-Upload /></el-icon>导入Excel</el-button>
                </el-col>            
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="ID" align="center" prop="id"
            min-width="150px"            
             />          
          <el-table-column label="父级id" align="center" prop="linkedParentId.name"
            min-width="150px"            
             />          
          <el-table-column label="名称" align="center" prop="name"
            min-width="150px"            
             />          
          <el-table-column align="center" label="图片"
            min-width="150px"            
            >
            <template #default="scope">
              <el-image
                style="width: 150px; height: 50px"
                v-if="!proxy.isEmpty(scope.row.picUrl)"
                :src="proxy.getUpFileUrl(scope.row.picUrl)"
                fit="contain"></el-image>
            </template>
          </el-table-column>          
          <el-table-column label="等级" align="center" prop="level"
            min-width="150px"            
             />          
          <el-table-column label="排序" align="center" prop="sort"
            min-width="150px"            
             />          
          <el-table-column label="创建时间" align="center" prop="createdAt"
            min-width="150px"            
            >
            <template #default="scope">
                <span>{{ proxy.parseTime(scope.row.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>        
          <el-table-column label="操作" align="center" class-name="small-padding" min-width="160px" fixed="right">
            <template #default="scope">            
              <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/shop/categoryInfo/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/shop/categoryInfo/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="categoryInfoList"
        />
    </el-card>
    <ApiV1ShopCategoryInfoEdit
       ref="editRef"       
       :parentIdOptions="parentIdOptions"       
       @categoryInfoList="categoryInfoList"
    ></ApiV1ShopCategoryInfoEdit>
    <ApiV1ShopCategoryInfoDetail
      ref="detailRef"      
      :parentIdOptions="parentIdOptions"      
      @categoryInfoList="categoryInfoList"
    ></ApiV1ShopCategoryInfoDetail>    
    <loadExcel ref="loadExcelCategoryInfoRef" @getList="categoryInfoList"
               upUrl="api/v1/shop/categoryInfo/import"
               tplUrl="/api/v1/shop/categoryInfo/excelTemplate"></loadExcel>    
  </div>
</template>
<script setup lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listCategoryInfo,
    getCategoryInfo,
    delCategoryInfo,
    addCategoryInfo,
    updateCategoryInfo,    
    linkedDataSearch    
} from "/@/api/shop/categoryInfo";
import {
    CategoryInfoTableColumns,
    CategoryInfoInfoData,
    CategoryInfoTableDataState,    
    LinkedCategoryInfoCategoryInfo,    
} from "/@/views/shop/categoryInfo/list/component/model"
import ApiV1ShopCategoryInfoEdit from "/@/views/shop/categoryInfo/list/component/edit.vue"
import ApiV1ShopCategoryInfoDetail from "/@/views/shop/categoryInfo/list/component/detail.vue"
import {downLoadXml} from "/@/utils/zipdownload";
import loadExcel from "/@/components/loadExcel/index.vue"
defineOptions({ name: "apiV1ShopCategoryInfoList"})
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const queryRef = ref()
const editRef = ref();
const detailRef = ref();
const loadExcelCategoryInfoRef = ref();
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
// parentIdOptions关联表数据
const parentIdOptions = ref<Array<ItemOptions>>([])
const state = reactive<CategoryInfoTableDataState>({
    ids:[],
    tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
            pageNum: 1,
            pageSize: 10,            
            id: undefined,            
            parentId: undefined,            
            name: undefined,            
            picUrl: undefined,            
            level: undefined,            
            sort: undefined,            
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
    categoryInfoList()
};
const linkedData = ()=>{
    linkedDataSearch().then((res:any)=>{        
        //关联category_info表选项        
        parentIdOptions.value = proxy.setItems(res, 'id', 'name','linkedCategoryInfoCategoryInfo')        
    })
}
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    categoryInfoList()
};
// 获取列表数据
const categoryInfoList = ()=>{
  loading.value = true
  listCategoryInfo(state.tableData.param).then((res:any)=>{
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
const handleSelectionChange = (selection:Array<CategoryInfoInfoData>) => {
    state.ids = selection.map(item => item.id)
    single.value = selection.length!=1
    multiple.value = !selection.length
}
const handleAdd =  ()=>{
    editRef.value.openDialog()
}
const handleUpdate = (row: CategoryInfoTableColumns|null) => {
    if(!row){
        row = state.tableData.data.find((item:CategoryInfoTableColumns)=>{
            return item.id ===state.ids[0]
        }) as CategoryInfoTableColumns
    }
    editRef.value.openDialog(toRaw(row));
};
const handleDelete = (row: CategoryInfoTableColumns|null) => {
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
            delCategoryInfo(id).then(()=>{
                ElMessage.success('删除成功');
                categoryInfoList();
            })
        })
        .catch(() => {});
}
const handleView = (row:CategoryInfoTableColumns)=>{
    detailRef.value.openDialog(toRaw(row));
}
//导出excel
const handleExport = ()=>{
    downLoadXml('/api/v1/shop/categoryInfo/export',state.tableData.param,'get')
}
const handleImport=()=>{
    loadExcelCategoryInfoRef.value.open()
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