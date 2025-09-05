<template>
	<div class="app-container">
		<el-card shadow="hover">
			<div class="system-user-search mb15">
        <el-form :inline="true" ref="queryRef" :model="tableData.param">
          <el-form-item label="表名称" prop="tableName">
            <el-input v-model="tableData.param.tableName" placeholder="请输入表名称" class="w-50 m-2" clearable/>
          </el-form-item>
          <el-form-item label="表描述" prop="tableComment">
            <el-input v-model="tableData.param.tableComment" placeholder="请输入岗位编码" class="w-50 m-2" clearable/>
          </el-form-item>
          <el-form-item label="创建时间" prop="dateRange">
            <el-date-picker
                v-model="tableData.param.dateRange"
                style="width: 240px"
                value-format="YYYY-MM-DD"
                type="daterange"
                range-separator="-"
                start-placeholder="开始日期"
                end-placeholder="结束日期"
            ></el-date-picker>
          </el-form-item>
          <el-form-item>
            <el-button size="default" type="primary" class="ml10" @click="tableList">
              <el-icon>
                <ele-Search />
              </el-icon>
              查询
            </el-button>
            <el-button size="default" @click="resetQuery(queryRef)">
              <el-icon>
                <ele-Refresh />
              </el-icon>
              重置
            </el-button>
          </el-form-item>
        </el-form>
        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
                type="primary"
                size="small"
                @click="handleGenTable(null)"
                :disabled="multiple"
            ><el-icon><ele-Download /></el-icon>生成</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="info"
                size="small"
                @click="openImportTable"
            ><el-icon><ele-UploadFilled /></el-icon>导入</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="success"
                size="small"
                :disabled="single"
                @click="handleEditTable(null)"
            ><el-icon><ele-EditPen /></el-icon>修改</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="danger"
                size="small"
                :disabled="multiple"
                @click="onRowDel(null)"
            ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
          </el-col>
        </el-row>
			</div>
			<el-table :data="tableData.data" style="width: 100%" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center" />
				<el-table-column type="index" label="序号" width="60" />
				<el-table-column prop="tableName" label="表名称" show-overflow-tooltip></el-table-column>
				<el-table-column prop="tableComment" label="表描述" show-overflow-tooltip></el-table-column>
				<el-table-column prop="className" label="实体" show-overflow-tooltip></el-table-column>
        <el-table-column prop="createTime" label="创建时间" show-overflow-tooltip></el-table-column>
        <el-table-column prop="updateTime" label="更新时间" show-overflow-tooltip></el-table-column>
				<el-table-column label="操作" width="320">
					<template #default="scope">
						<el-button size="small" text type="primary" @click="handlePreview(scope.row)"><el-icon><ele-View /></el-icon>预览</el-button>
						<el-button size="small" text type="primary" @click="handleEditTable(scope.row)"><el-icon><ele-EditPen /></el-icon>编辑</el-button>
						<el-button size="small" text type="primary" @click="handleSyncTable(scope.row)"><el-icon><ele-Refresh /></el-icon>同步</el-button>
						<el-button size="small" text type="primary" @click="handleGenTable(scope.row)"><el-icon><ele-Download /></el-icon>生成代码</el-button>
						<el-button size="small" text type="primary" @click="onRowDel(scope.row)"><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
					</template>
				</el-table-column>
			</el-table>
      <pagination
          v-show="tableData.total>0"
          :total="tableData.total"
          v-model:page="tableData.param.pageNum"
          v-model:limit="tableData.param.pageSize"
          @pagination="tableList"
      />
		</el-card>
    <import-table ref="importRef" @ok="tableList" />
    <gen-code-preview ref="genCodePreviewRef"></gen-code-preview>
	</div>
</template>

<script setup lang="ts">
import {toRefs, reactive, onMounted, ref, defineComponent} from 'vue';
import {ElMessageBox, ElMessage, FormInstance,ElLoading } from 'element-plus';
import {getTableList, deleteTables, batchGenCode, syncTable} from "/@/api/system/tools/gen";
import {TableData,TableDataState} from "/@/views/system/tools/gen/component/model"
import importTable from "/@/views/system/tools/gen/component/importTable.vue";
import { useRouter } from 'vue-router';
import genCodePreview from '/@/views/system/tools/gen/component/preview.vue'
import {refreshBackEndControlRoutes} from "/@/router/backEnd";
defineOptions({ name: "apiV1SystemToolsGenTableList"})
const router =  useRouter()
const addPostRef = ref()
const editPostRef = ref()
const queryRef = ref()
const importRef = ref()
const genCodePreviewRef = ref()
const state = reactive<TableDataState>({
  ids:[],
  // 非单个禁用
  single: true,
  // 非多个禁用
  multiple: true,
  tableNames:[],
  tableData: {
    data: [],
    total: 0,
    loading: false,
    param: {
      tableName:'',
      tableComment:'',
      pageNum: 1,
      pageSize: 10,
      dateRange: []
    },
  },
});
const {single,multiple, tableData}=toRefs(state)
// 初始化表格数据
const initTableData = () => {
  tableList()
};
const tableList = ()=>{
  getTableList(state.tableData.param).then(res=>{
    state.tableData.data = res.data.list??[];
    state.tableData.total = res.data.total;
  })
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
  tableList()
};
// 删除导入的表
const onRowDel = (row: TableData|null) => {
  let msg = '你确定要删除所选数据？';
  let ids:number[] = [] ;
  if(row){
    msg = `此操作将永久删除：“${row.tableName}”，是否继续?`
    ids = [row.tableId]
  }else{
    msg = `此操作将永久删除：“${state.tableNames.join(',')}”，是否继续?`
    ids = state.ids
  }
  if(ids.length===0){
    ElMessage.error('请选择要删除的数据。');
    return
  }
  ElMessageBox.confirm(msg, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      deleteTables(ids).then(()=>{
        ElMessage.success('删除成功');
        tableList();
      })
    })
    .catch(() => {});
};
// 分页改变
const onHandleSizeChange = (val: number) => {
  state.tableData.param.pageSize = val;
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
  state.tableData.param.pageNum = val;
};
// 页面加载时
onMounted(() => {
  initTableData();
});
// 多选框选中数据
const handleSelectionChange = (selection:Array<TableData>)=> {
  state.ids = selection.map(item => item.tableId)
  state.multiple = !selection.length
  state.single = selection.length != 1
  state.tableNames = selection.map(item=>item.tableName)
};
const handleGenTable=(row: TableData|null)=>{
  let msg = '你确定要生成？';
  let ids:number[] = [] ;
  if(row){
    msg = `此操作将生成业务表：“${row.tableName}”相关代码，是否继续?`
    ids = [row.tableId]
  }else{
    msg = `此操作将生成业务表：“${state.tableNames.join(',')}”相关代码，是否继续?`
    ids = state.ids
  }
  if(ids.length===0){
    ElMessage.error('请选择要生成的业务表数据。');
    return
  }
  ElMessageBox.confirm(msg, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
				const loading = ElLoading.service({text:'生成中...'})
        batchGenCode(ids).then(()=>{
          ElMessage.success('生成成功');
          resetMenuSession()
          tableList();
        }).finally(() => {
					loading.close()
				})
      })
      .catch(() => {});
}
// 重置菜单session
const resetMenuSession = () => {
  refreshBackEndControlRoutes();
};
//打开导入页面
const openImportTable=()=>{
  importRef.value.openDialog()
}
const handlePreview = (row:TableData)=>{
  genCodePreviewRef.value.showView(row.tableId);
}
const handleEditTable=(row:TableData|null)=>{
  const tableId = row?.tableId || state.ids[0];
  router.push({ path: "/system/tools/gen/edit", query: { tableId: tableId } });
}
const handleSyncTable=((row:TableData)=>{
    const tableId = row?.tableId || state.ids[0];
    syncTable(tableId).then((res:any)=>{
        if (res.code === 0) {
            ElMessage.success('同步成功');
        }
    })
})
</script>
