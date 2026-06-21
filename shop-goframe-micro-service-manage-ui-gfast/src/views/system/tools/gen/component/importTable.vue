<template>
  <!-- 导入表 -->
  <el-dialog title="导入表" v-model="visible" width="800px" top="5vh" append-to-body :close-on-click-modal="false">
    <el-form :model="tableData.param" ref="queryFormRef" :inline="true">
      <el-form-item label="数据库group" prop="group" style="width: 100%">
        <el-select v-model="tableData.param.group" placeholder="请选择数据库" @change="updateGroup">
          <el-option v-for="item in state.groups" :key="item" :label="item" :value="item" />
        </el-select>
      </el-form-item>
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
      <el-table @row-click="clickRow" ref="tableRef" :data="tableData.data" @selection-change="handleSelectionChange" height="260px">
        <el-table-column type="selection" width="55"></el-table-column>
        <el-table-column prop="tableName" label="表名称"></el-table-column>
        <el-table-column prop="tableComment" label="表描述"></el-table-column>
        <el-table-column prop="createTime" label="创建时间"></el-table-column>
        <el-table-column prop="updateTime" label="更新时间"></el-table-column>
      </el-table>
      <pagination
          v-show="tableData.total>0"
          :total="tableData.total"
          v-model:page="tableData.param.pageNum"
          v-model:limit="tableData.param.pageSize"
          @pagination="getList"
      />
    </el-row>
    <div footer class="dialog-footer">
      <el-button type="primary" @click="handleImportTable">确 定</el-button>
      <el-button @click="visible = false">取 消</el-button>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import {defineComponent, reactive, ref, toRefs} from "vue";
import {ImportTableDataState,TableData} from "/@/views/system/tools/gen/component/model"
import {getDBList, getImportTableList, importTable} from "/@/api/system/tools/gen";
import {ElMessage} from "element-plus/es";
defineOptions({ name: "importTable"})
const emit = defineEmits(['ok'])
const queryFormRef = ref()
const tableRef = ref()
const visible = ref(false)
// 选中数组值
let tables:string[] = []
const state = reactive<ImportTableDataState>({
  tableData:{
    data:[],
    total:0,
    loading:true,
    param:{
      group: '',
      pageNum: 1,
      pageSize: 10,
      tableName: '',
      tableComment: ''
    },
  },
  groups:[]
})
const {tableData} = toRefs(state)
const getList = ()=>{
  getDBList().then(res=>{
    state.groups = res.data.list??[]
    getImportTableList(state.tableData.param).then(res=>{
      state.tableData.data = res.data.list??[]
      state.tableData.total = res.data.total
    })
  })
}
const updateGroup = ()=> {
  getImportTableList(state.tableData.param).then(res=>{
    state.tableData.data = res.data.list??[]
    state.tableData.total = res.data.total
  })
}
const handleQuery = ()=>{
  state.tableData.param.pageNum = 1
  getList()
}
const resetQuery=()=>{
  queryFormRef.value.resetFields()
	state.tableData.param.group = state.groups[0]
  getList()
}
const clickRow=(row:TableData)=>{
  tableRef.value.toggleRowSelection(row);
}
const handleSelectionChange=(selection:Array<any>)=>{
  tables = selection.map(item => item.tableName);
}
const handleImportTable=()=>{
  if(tables.length==0){
    ElMessage.error("请选择要导入的表格");
    return
  }
  importTable(tables, state.tableData.param.group).then((res:any)=>{
    if (res.code === 0) {
      ElMessage.success('导入成功');
      visible.value = false;
      emit("ok");
    }
  })
}
const openDialog = ()=>{
  getList()
  visible.value = true
}
defineExpose({openDialog})
</script>

<style scoped>

</style>
