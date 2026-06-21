<template>
  <!-- 选择关联表 -->
  <el-dialog title="请选择关联表" v-model="visible" width="900px" top="5vh" append-to-body :close-on-click-modal="false">
    <el-form :model="tableData.param" ref="queryFormRef" :inline="true">
      <el-form-item label="表名称" prop="tableName">
        <el-input
            v-model="tableData.param.tableName"
            placeholder="请输入表名称"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="表描述" prop="tableComment">
        <el-input
            v-model="tableData.param.tableComment"
            placeholder="请输入表描述"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" size="small" @click="handleQuery"><el-icon><ele-Search /></el-icon>搜索</el-button>
        <el-button size="small" @click="resetQuery"><SvgIcon name="ele-Refresh" />重置</el-button>
      </el-form-item>
    </el-form>
    <el-row>
      <el-table @row-click="clickRow" ref="tableRef" :data="tableData.data" height="260px">
        <el-table-column prop="tableName" label="表名称"></el-table-column>
        <el-table-column prop="tableComment" label="表描述"></el-table-column>
        <el-table-column prop="packageName" label="关联表key">
          <template #default="scope">
            <el-select clearable filterable placeholder="请选择关联表key" v-model="scope.row.linkLabelId">
              <el-option v-for="column in scope.row.columns" :key="column.columnName" :label="column.columnName" :value="column.columnName">
                <span style="float: left">{{ column.htmlField }}</span>
                <span style="float: right; color: #8492a6; font-size: 13px">{{ column.columnComment }}</span>
              </el-option>
            </el-select>
          </template>
        </el-table-column>
        <el-table-column prop="updateTime" label="关联表value">
          <template #default="scope">
            <el-select clearable filterable placeholder="请选择关联表value" v-model="scope.row.linkLabelName">
              <el-option v-for="column in scope.row.columns" :key="column.columnName" :label="column.columnName" :value="column.columnName">
                <span style="float: left">{{ column.htmlField }}</span>
                <span style="float: right; color: #8492a6; font-size: 13px">{{ column.columnComment }}</span>
              </el-option>
            </el-select>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button size="small" type="primary" @click="handleImportTable(scope.row)"><el-icon><ele-Check /></el-icon>关联此表</el-button>
          </template>
        </el-table-column>
      </el-table>
      <pagination
          v-show="tableData.total>0"
          :total="tableData.total"
          v-model:page="tableData.param.pageNum"
          v-model:limit="tableData.param.pageSize"
          @pagination="getList"
      />
    </el-row>
  </el-dialog>
</template>

<script setup lang="ts">
import {defineComponent, reactive, ref, toRefs} from "vue";
import {ImportTableDataState, TableColumns, TableData} from "/@/views/system/tools/gen/component/model"
import {getRelationTable, importTable} from "/@/api/system/tools/gen";
import {ElMessage} from "element-plus/es";
defineOptions({ name: "relationTable"})
const emit = defineEmits(['ok'])
const queryFormRef = ref()
const tableRef = ref()
const visible = ref(false)
const columnId = ref<number>()
const state = reactive<ImportTableDataState>({
  tableData:{
    data:[],
    total:0,
    loading:true,
    param:{
      pageNum: 1,
      pageSize: 10,
      tableName: '',
      tableComment: ''
    },
  }
})
const { tableData} = toRefs(state)
const getList = ()=>{
  getRelationTable(state.tableData.param).then((res:any)=>{
    state.tableData.data = res.data.data??[]
    state.tableData.total = res.data.total

  })
}
const handleQuery = ()=>{
  state.tableData.param.pageNum = 1
  getList()
}
const resetQuery=()=>{
  queryFormRef.value.resetFields()
  getList()
}
const clickRow=(row:TableColumns)=>{
  tableRef.value.toggleRowSelection(row);
}

const handleImportTable=(row:TableColumns)=>{
  if(!row.linkLabelId){
    ElMessage.error("请选择关联表key")
    return
  }
  if(!row.linkLabelName){
    ElMessage.error("请选择关联表value")
    return
  }
  emit("ok",columnId.value,{
    linkTableName:row.tableName,
    linkLabelId:row.linkLabelId,
    linkLabelName:row.linkLabelName
  });
  visible.value = false;
}
const openDialog = (cid:number)=>{
  columnId.value = cid
  getList()
  visible.value = true
}
defineExpose({openDialog})
</script>

<style scoped lang="scss">
.check-key-label{
  width: 600px;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items:center;
  margin: 15px auto;
  .check-label{
    margin-top: 5px;
    text-align: right;
  }
}
</style>
