<template>
  <div class="shop-userCouponInfo-container">
    <el-card shadow="hover">
        <div class="shop-userCouponInfo-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="" prop="id">
                    <el-input
                        v-model="tableData.param.id"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="userCouponInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="用户id" prop="userId">
                    <el-input
                        v-model="tableData.param.userId"
                        placeholder="请输入用户id"
                        clearable                        
                        @keyup.enter.native="userCouponInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="userCouponInfoList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="优惠券id" prop="couponId">
                    <el-input
                        v-model="tableData.param.couponId"
                        placeholder="请输入优惠券id"
                        clearable                        
                        @keyup.enter.native="userCouponInfoList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="状态" prop="status">
                    <el-select filterable v-model="tableData.param.status" placeholder="请选择状态" clearable style="width:200px;">
                        <el-option
                            v-for="dict in shop_coupon_status"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="优惠金额（元）" prop="amount">
                    <el-input
                        v-model="tableData.param.amount"
                        placeholder="请输入优惠金额（元）"
                        clearable                        
                        @keyup.enter.native="userCouponInfoList"
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
                    <el-button type="primary"  @click="userCouponInfoList"><el-icon><ele-Search /></el-icon>搜索</el-button>
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
                  v-auth="'api/v1/shop/userCouponInfo/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/shop/userCouponInfo/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/shop/userCouponInfo/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>             
             <el-col :span="1.5">
                <el-button
                        type="warning"
                        @click="handleExport()"
                        v-auth="'api/v1/shop/userCouponInfo/export'"
                ><el-icon><ele-Download /></el-icon>导出Excel</el-button>
             </el-col>            
                <el-col :span="1.5">
                    <el-button
                            type="success"
                            @click="handleImport()"
                            v-auth="'api/v1/shop/userCouponInfo/import'"
                    ><el-icon><ele-Upload /></el-icon>导入Excel</el-button>
                </el-col>            
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="" align="center" prop="id"
            min-width="150px"            
             />          
          <el-table-column label="用户id" align="center" prop="userId"
            min-width="150px"            
             />          
          <el-table-column label="优惠券id" align="center" prop="linkedCouponId.name"
            min-width="150px"            
             />          
          <el-table-column label="状态" align="center" prop="status" :formatter="statusFormat"
            min-width="150px"            
             />          
          <el-table-column label="优惠金额（元）" align="center" prop="amount"
            min-width="150px"            
             />          
          <el-table-column label="创建时间" align="center" prop="createdAt"
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
                v-auth="'api/v1/shop/userCouponInfo/get'"
              ><el-icon><ele-View /></el-icon>详情</el-button>              
              <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/shop/userCouponInfo/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/shop/userCouponInfo/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="userCouponInfoList"
        />
    </el-card>
    <ApiV1ShopUserCouponInfoEdit
       ref="editRef"       
       :couponIdOptions="couponIdOptions"       
       :statusOptions="shop_coupon_status"       
       @userCouponInfoList="userCouponInfoList"
    ></ApiV1ShopUserCouponInfoEdit>
    <ApiV1ShopUserCouponInfoDetail
      ref="detailRef"      
      :couponIdOptions="couponIdOptions"      
      :statusOptions="shop_coupon_status"      
      @userCouponInfoList="userCouponInfoList"
    ></ApiV1ShopUserCouponInfoDetail>    
    <loadExcel ref="loadExcelUserCouponInfoRef" @getList="userCouponInfoList"
               upUrl="api/v1/shop/userCouponInfo/import"
               tplUrl="/api/v1/shop/userCouponInfo/excelTemplate"></loadExcel>    
  </div>
</template>
<script setup lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listUserCouponInfo,
    getUserCouponInfo,
    delUserCouponInfo,
    addUserCouponInfo,
    updateUserCouponInfo,    
    linkedDataSearch    
} from "/@/api/shop/userCouponInfo";
import {
    UserCouponInfoTableColumns,
    UserCouponInfoInfoData,
    UserCouponInfoTableDataState,    
    LinkedUserCouponInfoCouponInfo,    
} from "/@/views/shop/userCouponInfo/list/component/model"
import ApiV1ShopUserCouponInfoEdit from "/@/views/shop/userCouponInfo/list/component/edit.vue"
import ApiV1ShopUserCouponInfoDetail from "/@/views/shop/userCouponInfo/list/component/detail.vue"
import {downLoadXml} from "/@/utils/zipdownload";
import loadExcel from "/@/components/loadExcel/index.vue"
defineOptions({ name: "apiV1ShopUserCouponInfoList"})
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const queryRef = ref()
const editRef = ref();
const detailRef = ref();
const loadExcelUserCouponInfoRef = ref();
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
    shop_coupon_status,    
} = proxy.useDict(    
    'shop_coupon_status',    
)
// couponIdOptions关联表数据
const couponIdOptions = ref<Array<ItemOptions>>([])
const state = reactive<UserCouponInfoTableDataState>({
    ids:[],
    tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
            pageNum: 1,
            pageSize: 10,            
            id: undefined,            
            userId: undefined,            
            couponId: undefined,            
            status: undefined,            
            amount: undefined,            
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
    userCouponInfoList()
};
const linkedData = ()=>{
    linkedDataSearch().then((res:any)=>{        
        //关联coupon_info表选项        
        couponIdOptions.value = proxy.setItems(res, 'id', 'name','linkedUserCouponInfoCouponInfo')        
    })
}
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    userCouponInfoList()
};
// 获取列表数据
const userCouponInfoList = ()=>{
  loading.value = true
  listUserCouponInfo(state.tableData.param).then((res:any)=>{
    let list = res.data.list??[];    
    state.tableData.data = list;
    state.tableData.total = res.data.total;
    loading.value = false
  })
};
const toggleSearch = () => {
    showAll.value = !showAll.value;
}
// 状态字典翻译
const statusFormat = (row:UserCouponInfoTableColumns) => {
    return proxy.selectDictLabel(shop_coupon_status.value, row.status);
}
// 多选框选中数据
const handleSelectionChange = (selection:Array<UserCouponInfoInfoData>) => {
    state.ids = selection.map(item => item.id)
    single.value = selection.length!=1
    multiple.value = !selection.length
}
const handleAdd =  ()=>{
    editRef.value.openDialog()
}
const handleUpdate = (row: UserCouponInfoTableColumns|null) => {
    if(!row){
        row = state.tableData.data.find((item:UserCouponInfoTableColumns)=>{
            return item.id ===state.ids[0]
        }) as UserCouponInfoTableColumns
    }
    editRef.value.openDialog(toRaw(row));
};
const handleDelete = (row: UserCouponInfoTableColumns|null) => {
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
            delUserCouponInfo(id).then(()=>{
                ElMessage.success('删除成功');
                userCouponInfoList();
            })
        })
        .catch(() => {});
}
const handleView = (row:UserCouponInfoTableColumns)=>{
    detailRef.value.openDialog(toRaw(row));
}
//导出excel
const handleExport = ()=>{
    downLoadXml('/api/v1/shop/userCouponInfo/export',state.tableData.param,'get')
}
const handleImport=()=>{
    loadExcelUserCouponInfoRef.value.open()
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