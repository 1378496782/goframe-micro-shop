<template>
	<div class="system-user-list-container">
    <el-table :data="tableData.data" style="width: 100%" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column type="index" label="序号" width="60" />
      <el-table-column prop="userName" label="账户名称" show-overflow-tooltip></el-table-column>
      <el-table-column prop="userNickname" label="用户昵称" show-overflow-tooltip></el-table-column>
      <el-table-column prop="dept.deptName" label="部门" show-overflow-tooltip></el-table-column>
      <el-table-column label="角色" align="center" prop="roleInfo" :show-overflow-tooltip="true" >
        <template #default="scope">
          <span v-for="(item,index) of scope.row.roleInfo" :key="'role-'+index">   {{item.name+'   '}}   </span>
        </template>
      </el-table-column>
      <el-table-column prop="mobile" label="手机号" show-overflow-tooltip></el-table-column>
      <el-table-column prop="userStatus" label="用户状态" show-overflow-tooltip>
        <template #default="scope">
          <el-switch
              v-model="scope.row.userStatus"
              inline-prompt
              :active-value="1"
              :inactive-value="0"
              active-text="启"
              inactive-text="禁"
              @change="handleStatusChange(scope.row)">
          </el-switch>
        </template>
      </el-table-column>
      <el-table-column prop="createdAt" label="创建时间" show-overflow-tooltip></el-table-column>
      <el-table-column label="操作" width="200">
        <template #default="scope">
          <el-button size="small" text type="primary" @click="onOpenEditUser(scope.row)">修改</el-button>
          <el-button size="small" text type="primary" @click="onRowDel(scope.row)">删除</el-button>
          <el-button size="small" text type="primary" @click="handleResetPwd(scope.row)">重置</el-button>
        </template>
      </el-table-column>
    </el-table>
    <pagination
        v-show="tableData.total>0"
        :total="tableData.total"
        v-model:page="tableData.param.pageNum"
        v-model:limit="tableData.param.pageSize"
        @pagination="userList"
    />
		<EditUser ref="editUserRef" :dept-data="deptData" :gender-data="sys_user_sex" @getUserList="userList"/>
	</div>
</template>

<script setup lang="ts">
import {toRefs, reactive, onMounted, ref, getCurrentInstance} from 'vue';
import {ElMessageBox, ElMessage} from 'element-plus';
import EditUser from '/@/views/system/user/component/editUser.vue';
import {getUserList, resetUserPwd, changeUserStatus, deleteUser} from '/@/api/system/user';

interface TableDataState {
  ids:number[];
  deptProps:{};
	tableData: {
		data: any[];
		total: number;
		loading: boolean;
		param: TableParam;
	};
}
interface TableParam {
  pageNum: number;
  pageSize: number;
  deptId:string;
  roleId:number | undefined;
  mobile:string;
  status:string;
  keyWords:string;
  dateRange: string[];
}
defineOptions({ name: "systemUserList"})
const props = defineProps({
  deptData:{
    type:Array,
    default:()=>[]
  },
  param:{
    type:Object,
    default:()=>{}
  },
  genderData:{
    type:Array,
    default:()=>[]
  }
 })
const {proxy} = <any>getCurrentInstance();
const {sys_user_sex} = proxy.useDict('sys_user_sex')
const editUserRef = ref();
const state = reactive<TableDataState>({
  ids:[],
  deptProps:{
    id:"deptId",
    children: "children",
    label: "deptName"
  },
  tableData: {
    data: [],
    total: 0,
    loading: false,
    param: {
      pageNum: 1,
      pageSize: 10,
      roleId:undefined,
      deptId:'',
      mobile:'',
      status:'',
      keyWords:'',
      dateRange:[]
    },
  },
});
const { tableData} = toRefs(state);
// 初始化表格数据
const initTableData = () => {
  userList();
};
const userList = ()=>{
  const param = {...state.tableData.param, ...props.param};
  getUserList(param).then((res:any)=>{
    state.tableData.data = res.data.userList??[];
    state.tableData.total = res.data.total;
  });
};
// 打开新增用户弹窗
const onOpenAddUser = () => {
  editUserRef.value.openDialog();
};
// 打开修改用户弹窗
const onOpenEditUser = (row:any) => {
  editUserRef.value.openDialog(row);
};
// 删除用户
const onRowDel = (row:any) => {
  let msg = '你确定要删除所选用户？';
  let ids:number[] = [] ;
  if(row){
    msg = `此操作将永久删除用户：“${row.userName}”，是否继续?`
    ids = [row.id]
  }else{
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
        deleteUser(ids).then(()=>{
          ElMessage.success('删除成功');
          userList();
        })
      })
      .catch(() => {});
};
// 页面加载时
onMounted(() => {
  initTableData();
});
// 多选框选中数据
const handleSelectionChange = (selection:any[])=> {
  state.ids = selection.map(item => item.id)
};
/** 重置密码按钮操作 */
const handleResetPwd = (row:any)=> {
  ElMessageBox.prompt('请输入"' + row.userName + '"的新密码', "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消"
  }).then(({ value }) => {
    if(!value || value==''){
      ElMessage.success('密码不能为空');
      return
    }
    resetUserPwd(row.id, value).then(() => {
        ElMessage.success("修改成功，新密码是：" + value);
    });
  }).catch(() => {});
};
// 用户状态修改
const handleStatusChange = (row:any)=> {
  let text = row.userStatus === 1 ? "启用" : "停用";
  ElMessageBox.confirm('确认要"' + text + '"："' + row.userName + '"用户吗?', "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(function() {
    return changeUserStatus(row.id, row.userStatus);
  }).then(() => {
    ElMessage.success(text + "成功");
  }).catch(function() {
    row.userStatus =row.userStatus === 0 ?1 : 0;
  });
};
defineExpose({userList,onOpenAddUser,onRowDel})
</script>
