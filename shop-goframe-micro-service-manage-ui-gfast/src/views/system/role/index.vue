<template>
	<div class="system-role-container">
		<el-card shadow="hover">
			<div class="system-user-search mb15">
        <el-form :inline="true">
          <el-form-item label="角色名称">
            <el-input v-model="tableData.param.roleName" placeholder="请输入角色名称" class="w-50 m-2" clearable/>
          </el-form-item>
          <el-form-item label="状态">
            <el-select placeholder="请选择状态" class="w-50 m-2" v-model="tableData.param.roleStatus" clearable style="width: 120px;">
              <el-option label="启用"  value="1" />
              <el-option label="禁用"  value="0" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button size="default" type="primary" class="ml10" @click="roleList">
              <el-icon>
                <ele-Search />
              </el-icon>
              查询
            </el-button>
            <el-button size="default" type="success" class="ml10" @click="onOpenAddRole">
              <el-icon>
                <ele-FolderAdd />
              </el-icon>
              新增角色
            </el-button>
          </el-form-item>
        </el-form>
			</div>
			<el-table :data="tableData.data" style="width: 100%"
        row-key="id"
        default-expand-all
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }">
				<el-table-column prop="name" label="角色名称" show-overflow-tooltip></el-table-column>
				<el-table-column prop="listOrder" label="排序" show-overflow-tooltip></el-table-column>
				<el-table-column prop="userCnt" label="用户数量" align="center">

					<template #default="scope">
						<el-link  type="primary" @click="onOpenUserList(scope.row)">{{scope.row.userCnt}}</el-link>
					</template>
				</el-table-column>
				<el-table-column prop="status" label="角色状态" show-overflow-tooltip>
					<template #default="scope">
						<el-tag type="success" v-if="scope.row.status===1">启用</el-tag>
						<el-tag type="info" v-else>禁用</el-tag>
					</template>
				</el-table-column>
				<el-table-column prop="remark" label="角色描述" show-overflow-tooltip></el-table-column>
        <el-table-column prop="createdAt" label="创建时间" show-overflow-tooltip></el-table-column>
				<el-table-column label="操作" width="220">
					<template #default="scope">
						<el-button size="small" text type="primary" @click="onOpenEditRole(scope.row)"><el-icon><ele-EditPen /></el-icon>修改</el-button>
            <el-dropdown class="auth-action" @command="handleCommand">
              <span class="el-dropdown-link">
                授权 <el-icon><ele-ArrowDown /></el-icon>
              </span>
              <template #dropdown>
                <el-dropdown-menu class="auth-action-menu">
                  <el-dropdown-item :command="'resource_'+scope.row.id">资源权限</el-dropdown-item>
                  <el-dropdown-item :command="'data_'+scope.row.id">数据权限</el-dropdown-item>
                  <el-dropdown-item :command="'user_'+scope.row.id">用户授权</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
						<el-button size="small" text type="primary" @click="onRowDel(scope.row)"><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-card>
		<EditRole ref="editRoleRef" @getRoleList="roleList" :roleData="tableData.data"/>
		<DataScope ref="dataScopeRef" @getRoleList="roleList"/>
		<el-dialog :title="selectRow.name+'-用户列表'" v-model="isShowDialog" width="70vw">
			<UserList v-if="isShowDialog" ref="userListRef" :dept-data="deptData" :gender-data="sys_user_sex" :param="userListParam" @getUserList="userList"/>
		</el-dialog>
    <select-user v-show="false" ref="selectUserRef" v-model="roleUsers"></select-user>
	</div>
</template>

<script setup lang="ts">
import {toRefs, reactive, onMounted, ref, defineComponent, toRaw, getCurrentInstance, watch} from 'vue';
import { ElMessageBox, ElMessage,ElLoading } from 'element-plus';
import EditRole from '/@/views/system/role/component/editRole.vue';
import DataScope from '/@/views/system/role/component/dataScope.vue';
import {deleteRole, getRoleList, setRoleUsers} from "/@/api/system/role";
import {getDeptTree, getUsersById} from '/@/api/system/user';
import UserList from '/@/views/system/user/component/userList.vue';
import selectUser from "/@/components/selectUser/index.vue";

// 定义接口来定义对象的类型
interface TableData {
  id:number;
  pid:number;
	status: number;
	listOrder: number;
	name: string;
  remark: string;
  dataScope:number;
  createdAt: string;
  userCnt: number;
}
interface TableDataState {
	isShowDialog:boolean;
	selectRow:any;
	deptData:any[];
	userListParam: {
		roleId:number | undefined;
	};
	tableData: {
		data: Array<TableData>;
		total: number;
		loading: boolean;
		param: {
      roleName:string;
      roleStatus:string;
			pageNum: number;
			pageSize: number;
		};
	};
}
defineOptions({ name: "apiV1SystemRoleList"})
const selectUserRef = ref()
const {proxy} = getCurrentInstance() as any;
const {sys_user_sex} = proxy.useDict('sys_user_sex')
const addRoleRef = ref();
const userListRef = ref();
const editRoleRef = ref();
const dataScopeRef =ref();
const roleListData = ref<Array<TableData>>()
const roleUsers = ref([])
const setRole = ref(0)
const state = reactive<TableDataState>({
  isShowDialog: false,
  deptData: [],
  userListParam: {
    roleId: undefined,
  },
  selectRow:{},
  tableData: {
    data: [],
    total: 0,
    loading: false,
    param: {
      roleName: '',
      roleStatus: '',
      pageNum: 1,
      pageSize: 10,
    },
  },
});
const { tableData,selectRow,deptData,userListParam,isShowDialog} = toRefs(state);
// 初始化表格数据
const initTableData = () => {
  roleList()
};
const roleList = ()=>{
  const data: Array<TableData> = [];
  getRoleList(state.tableData.param).then(res=>{
    const list = res.data.list??[]
    list.map((item:TableData)=>{
      data.push({
        id:item.id,
        pid:item.pid,
        status: item.status,
        listOrder: item.listOrder,
        name: item.name,
        remark: item.remark,
        dataScope:item.dataScope,
        userCnt:item.userCnt,
        createdAt: item.createdAt,
      });
    })
    roleListData.value = data
    state.tableData.data = proxy.handleTree(data??[], "id","pid","children",true);
  })
};
// 打开角色用户列表
const onOpenUserList = (row: TableData) => {
  state.selectRow = row
  state.userListParam.roleId = row.id
  if (state.deptData.length == 0){
    getDeptTree().then((res:any)=>{
      state.deptData = res.data.deps
      state.isShowDialog = true
    })
  }else{
    state.isShowDialog = true
  }
};
// 打开新增角色弹窗
const onOpenAddRole = () => {
  editRoleRef.value.openDialog();
};
// 打开修改角色弹窗
const onOpenEditRole = (row: Object|undefined) => {
  editRoleRef.value.openDialog(toRaw(row));
};
//数据权限设置弹窗
const handleDataScope=(row:any)=>{
  dataScopeRef.value.openDialog(toRaw(row))
}
//用户授权
const handleUserScope = ((row:any)=>{
  const ld = ElLoading.service()
  setRole.value = row.id
  getUsersById(row.id).then((res:any)=>{
    ld.close()
    roleUsers.value=res.data.userList?res.data.userList.map((user:any)=>user.id):[]
    selectUserRef.value.openDialog()
  })
})
// 删除角色
const onRowDel = (row: any) => {
  ElMessageBox.confirm(`此操作将永久删除角色：“${row.name}”，是否继续?`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      deleteRole(row.id).then(()=>{
        ElMessage.success('删除成功');
        proxy.$refs['editRoleRef'].resetMenuSession();
        roleList();
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
const userList = ()=>{
  userListRef.value.userList();
};
// 页面加载时
onMounted(() => {
  initTableData();
});
const handleCommand = (command: string )=>{
  let commandArr = command.split('_')
  let row = roleListData.value?.filter((item:TableData)=>{
    return item.id==parseInt(commandArr[1])
  })?.[0]
  switch (commandArr[0]){
    case 'resource':
      onOpenEditRole(row)
      break
    case 'data':
      handleDataScope(row)
      break
    case 'user':
      handleUserScope(row)
      break
  }
}
watch(roleUsers,(newVal) => {
  setRoleUser(newVal)
})
const setRoleUser = (ids:any)=>{
  //用户授权提交
  setRoleUsers({roleId:setRole.value,userIds:ids}).then((res:any)=>{
    roleList()
  })
}
</script>
<style scoped lang="scss">
.auth-action{
  margin:6px 8px 0 8px;
  .el-dropdown-link {
    cursor: pointer;
    color: var(--el-color-primary);
    font-size: 12px;
    display: flex;
    align-items: center;
  }
  .auth-action-menu{
    font-size: 12px;
  }
}
</style>
