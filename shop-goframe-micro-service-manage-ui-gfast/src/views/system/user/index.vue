<template>
	<div class="system-user-container">
    <el-row :gutter="10" style="width: 100%;">
      <el-col :span="4">
        <el-card shadow="hover">
          <el-aside>
            <el-scrollbar>
              <el-input :prefix-icon="search" v-model="filterText" placeholder="请输入部门名称" clearable style="width: 80%;"/>
              <el-tree
                  ref="treeRef"
                  class="filter-tree"
                  :data="deptData"
                  :props="deptProps"
                  default-expand-all
                  :filter-node-method="deptFilterNode"
                  @node-click="handleNodeClick"
              />
            </el-scrollbar>
          </el-aside>
        </el-card>
      </el-col>
      <el-col :span="20">
        <el-card shadow="hover">
          <div class="system-user-search mb15">
            <el-form :model="param" ref="queryRef" :inline="true" label-width="68px">
              <el-form-item label="关键字" prop="keyWords">
                <el-input
                    v-model="param.keyWords"
                    placeholder="请输入用户账号或姓名"
                    clearable
                    style="width: 240px"
                    @keyup.enter.native="userList"
                />
              </el-form-item>
              <el-form-item label="手机号码" prop="mobile">
                <el-input
                    v-model="param.mobile"
                    placeholder="请输入手机号码"
                    clearable
                    style="width: 240px"
                    @keyup.enter.native="userList"
                />
              </el-form-item>
              <el-form-item label="状态" prop="status" style="width: 200px;">
                <el-select
                    v-model="param.status"
                    placeholder="用户状态"
                    clearable
                    style="width: 240px"
                >
                  <el-option label="启用"  :value="1"/>
                  <el-option label="禁用"  :value="0"/>
                </el-select>
              </el-form-item>
              <el-form-item label="创建时间" prop="dateRange">
                <el-date-picker
                    v-model="param.dateRange"
                    style="width: 240px"
                    value-format="YYYY-MM-DD"
                    type="daterange"
                    range-separator="-"
                    start-placeholder="开始日期"
                    end-placeholder="结束日期"
                ></el-date-picker>
              </el-form-item>
              <el-form-item>
                <el-button size="default" type="primary" class="ml10" @click="userList">
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
                <el-button size="default" type="success" class="ml10" @click="onOpenAddUser">
                  <el-icon>
                    <ele-FolderAdd />
                  </el-icon>
                  新增用户
                </el-button>
                <el-button size="default" type="danger" class="ml10" @click="onRowDel()">
                  <el-icon>
                    <ele-Delete />
                  </el-icon>
                  删除用户
                </el-button>
              </el-form-item>
            </el-form>
          </div>

          <UserList ref="userListRef" :dept-data="deptData" :gender-data="sys_user_sex" :param="param" @getUserList="userList"/>
        </el-card>
      </el-col>
    </el-row>
	</div>
</template>

<script setup lang="ts">
import {toRefs, reactive, onMounted, ref, watch, getCurrentInstance} from 'vue';
import {ElTree,FormInstance} from 'element-plus';
	import { Search } from '@element-plus/icons-vue'
import UserList from '/@/views/system/user/component/userList.vue';
import {getDeptTree} from '/@/api/system/user';

interface QueryParam {
  ids:number[];
  deptProps:{};
  deptData:any[];
  param: {
    deptId:string;
    mobile:string;
    status:string;
    keyWords:string;
    dateRange: string[];
  };
}
defineOptions({ name: "systemUser"})
const {proxy} = <any>getCurrentInstance();
const {sys_user_sex} = proxy.useDict('sys_user_sex')
const userListRef = ref();
const queryRef = ref();
const filterText = ref('');
const treeRef = ref<InstanceType<typeof ElTree>>();
const search = Search
const state = reactive<QueryParam>({
  ids:[],
  deptProps:{
    id:"deptId",
    children: "children",
    label: "deptName"
  },
  deptData:[
    {
      label: '集团总部',
      children: [
        {
          label: '曲靖分部',
          children: [
            {
              label: '总经办',
            },
            {
              label: '市场部',
            },
            {
              label: '研发部',
            },
          ],
        },
      ],
    },
  ],
  param: {
    deptId:'',
    mobile:'',
    status:'',
    keyWords:'',
    dateRange:[]
  },
});
const { deptData,deptProps,param}=toRefs(state)
// 初始化表格数据
const initTableData = () => {
  getDeptTree(true).then((res:any)=>{
    state.deptData = res.data.deps
  })
};
const userList = ()=>{
  userListRef.value.userList();
};
// 打开新增用户弹窗
const onOpenAddUser = () => {
  userListRef.value.onOpenAddUser();
};
// 删除用户
const onRowDel = () => {
  userListRef.value.onRowDel(null);
};
// 页面加载时
onMounted(() => {
  initTableData();
});
watch(filterText, (val) => {
  treeRef.value!.filter(val)
});
const deptFilterNode = (value: string, data:any) => {
  if (!value) return true;
  return data.deptName.includes(value)
};
// 节点单击事件
const handleNodeClick = (data:any) => {
  state.param.deptId = data.deptId;
  userList();
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
  userList()
};
</script>
