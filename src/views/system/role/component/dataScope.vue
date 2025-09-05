<template>
  <div class="system-sysJob-edit">
    <!-- 分配角色数据权限对话框 -->
    <el-dialog title="数据权限设置" v-model="openDataScope" width="90%" append-to-body :close-on-click-modal="false">
      <el-alert title="注：此功能界面需要与代码查询条件配合使用，并非所有接口都需设置数据权限，多用于业务模块！" :closable="false" type="warning" style="margin-bottom: 12px;"/>
      <el-form ref="formRef" :model="form" label-width="80px">
        <el-form-item label="角色名称">
          <el-input v-model="form.roleName" :disabled="true" />
        </el-form-item>
      </el-form>
      <el-table ref="tableRef" :data="menuTableData" style="width: 100%" row-key="path" :tree-props="{ children: 'children', hasChildren: 'hasChildren' }" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column label="菜单名称" >
          <template #default="scope">
            <SvgIcon :name="scope.row.icon" />
            <span class="ml10">{{ scope.row.title }}</span>
          </template>
        </el-table-column>
        <el-table-column label="api接口" >
          <template #default="scope">
            <span>{{ scope.row.name }}</span>
          </template>
        </el-table-column>

        <el-table-column label="类型" show-overflow-tooltip width="80">
          <template #default="scope">
            <el-tag :type="scope.row.menuType===0?'danger':(scope.row.menuType===1?'success':'warning')"
                    size="small">{{scope.row.menuType===0?'目录':(scope.row.menuType===1?'菜单':'按钮') }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column label="操作"  width="600">
          <template #header>
              <el-radio-group v-model="optionsActionValue" @change="setOptionsActionValueHandle">
                <template v-for="item in dataScopeOptions" :key="item.value">
                  <el-radio v-if="item.value.toString()!='2'" :value="item.value" >{{item.label}}</el-radio>
                </template>
              </el-radio-group>
          </template>
          <template #default="scope">
            <div class="option-con">
              <el-radio-group v-model="menuAuthData[scope.row.id]">
                <template v-for="item in dataScopeOptions" :key="item.value">
                  <el-radio :value="item.value" @change="setOptionsItemHandle(item.value,scope.row.id)">
                    <el-badge v-if="item.value.toString()=='2' && menuAuthData[scope.row.id]=='2'" type="success" :value="getDeptLen(scope.row.id)" :max="99" class="item" :show-zero="false">{{item.label}}</el-badge>
                    <span v-else>{{item.label}}</span>
                  </el-radio>
                </template>
              </el-radio-group>
              <el-link type="primary" @click="openSelDeptHandler(scope.row.id)" :underline="false" v-show="menuAuthData[scope.row.id] && menuAuthData[scope.row.id]=='2'">选择部门</el-link>
            </div>
          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitDataScope">确 定</el-button>
          <el-button @click="cancelDataScope">取 消</el-button>
        </div>
      </template>
    </el-dialog>
    <!-- 选择部门对话框 -->
    <el-dialog title="选择部门" v-model="openSelDept" width="800px" append-to-body :close-on-click-modal="false">
      <el-row :gutter="35">
        <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
          <el-checkbox v-model="deptExpand" @change="handleCheckedTreeExpand($event)">展开/折叠</el-checkbox>
          <el-checkbox v-model="deptNodeAll" @change="handleCheckedTreeNodeAll($event)">全选/全不选</el-checkbox>
          <el-checkbox v-model="deptCheckStrictly" @change="handleCheckedTreeConnect($event)">父子联动</el-checkbox>
        </el-col>
        <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
          <el-tree
              class="tree-border"
              :data="deptOptions"
              show-checkbox
              default-expand-all
              ref="deptRef"
              node-key="deptId"
              :check-strictly="!deptCheckStrictly"
              :props="deptProps"
          ></el-tree>
        </el-col>
      </el-row>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitDept">确 定</el-button>
          <el-button type="success" @click="submitSonMenuDept">确定并应用到子菜单</el-button>
          <el-button @click="cancelDept">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import {getCurrentInstance, nextTick, reactive, ref} from "vue";
import {dataScope, getRole, roleDeptTreeSelect, roleMenuTreeSelect} from "/@/api/system/role";
import {ElMessage} from "element-plus/es";
defineOptions({ name: "dataScope"})
const emit = defineEmits(["getRoleList"])
const tableRef = ref()
const openSelDept = ref(false);
const {proxy} = getCurrentInstance() as any;
const openDataScope = ref(false)
const deptExpand =  ref(true)
const deptNodeAll = ref(false)
const deptOptions = ref<Array<any>>([])
const menuOptions = ref([])
const deptRef = ref();
const formRef = ref();
const deptCheckStrictly = ref(false)
const optionsActionValue = ref("0")
const setDeptId = ref(0)
const deptProps = ref({
  children: "children",
      label: "deptName"
})
// 数据范围选项
const dataScopeOptions = ref([
  {
    value: 1,
    label: "全部"
  },
  {
    value: 3,
    label: "本部门"
  },
  {
    value: 4,
    label: "本部门及以下"
  },
  {
    value: 5,
    label: "仅本人"
  },
  {
    value: 2,
    label: "自定义"
  }
])
const menuAuthData = ref([])
const deptAuthData = ref([])
const form = reactive({
  roleId:undefined,
  roleName:'',
  authData:[{
    menuId:'0',
    scope:'0',
    deptIds:[]
  }]
})
const menuTableData = ref([])
const menuTableList = ref([])
// 树权限（展开/折叠）
const handleCheckedTreeExpand = (value:any) => {
  let treeList = deptOptions.value;
  for (let i = 0; i < treeList.length; i++) {
    deptRef.value.store.nodesMap[treeList[i].deptId].expanded = value;
  }
}
// 树权限（全选/全不选）
const handleCheckedTreeNodeAll = (value:any) => {
  deptRef.value.setCheckedNodes(value ? deptOptions.value: []);
}
/** 树权限（父子联动） */
const handleCheckedTreeConnect = (value:any) => {
  deptCheckStrictly.value = value ? true : false;
}
/** 提交按钮（数据权限） */
const submitDataScope = () => {
  form.authData = []
  //获取选中的菜单
  let rows = tableRef.value.getSelectionRows()
  rows.map((item:any)=>{
    let index = item.id
    form.authData.push({
      menuId:index,
      scope:menuAuthData.value[index],
      deptIds:deptAuthData.value[index]??[]
    })
  })
  dataScope(form).then((response:any) => {
    if (response.code === 0) {
      ElMessage.success("设置成功");
      openDataScope.value = false;
      emit('getRoleList')
    } else {
      ElMessage.error("设置失败")
    }
  });

}
const cancelDataScope = () => {
  closeDialog();
};
const submitDept = ()=>{
  deptAuthData.value[setDeptId.value] = getDeptAllCheckedKeys() as never;
  cancelDept()
}
const submitSonMenuDept = ()=>{
  const checkedDeptId = getDeptAllCheckedKeys() as never;
  deptAuthData.value[setDeptId.value] = getDeptAllCheckedKeys() as never;
  //查询所有子级
  const children = proxy.findChildrenByPid(setDeptId.value,menuTableList.value)
  children.map((item:any)=>{
    deptAuthData.value[item.id] = getDeptAllCheckedKeys() as never;
  })
  cancelDept()
}
const cancelDept = () => {
  deptCheckStrictly.value= false
  openSelDept.value = false
};
// 打开弹窗
const openDialog = (row: any) => {
  openDataScope.value = true;
  roleDeptTreeSelect().then(response => {
    deptOptions.value = response.data.depts||[];
  });
  resetForm();
  if(row) {
    getRole(row.id).then((res:any)=>{
      if(res.data.role){
        form.roleName = res.data.role.name;
        form.roleId = res.data.role.id;
      }
    })
    getMenuTreeselect(row.id)
  }
};
defineExpose({ openDialog})
/** 根据角色ID查询部门树结构 */
const getRoleDeptTreeselect = (menuId:any) =>{
  nextTick(()=>{
    deptRef.value?.setCheckedKeys([], false)
    if(deptAuthData.value[menuId]){
      deptRef.value.setCheckedKeys(deptAuthData.value[menuId], true);
    }
  })
}
const getMenuTreeselect = (roleId:number) =>{
  roleMenuTreeSelect(roleId).then(res=>{
    menuTableData.value = proxy.handleTree(res.data.rules??[], "id","pid");
    menuTableList.value = proxy.flattenTree(menuTableData.value)
    const menuData:any = []
    const deptData:any = []
    const rows:never[] = []
    if(res.data.dataScope){
      res.data.dataScope.map((item:any)=>{
        menuData[item.menuId] = item.dataScope
        deptData[item.menuId] = item.deptIds
      })
      menuAuthData.value = menuData
      deptAuthData.value = deptData
      //设置菜单行选中
      menuTableList.value.map((item:any)=>{
        if(menuAuthData.value[item.id]){
          rows.push(item as never)
        }
      })
      nextTick(()=>{
        //行选择
        rows.map((item:any)=>{
          tableRef.value.toggleRowSelection(item,true)
        })
      })
    }
  })
}

// 关闭弹窗
const closeDialog = () => {
  openDataScope.value = false;
};
// 所有部门节点数据
const getDeptAllCheckedKeys = () => {
  // 目前被选中的部门节点
  let checkedKeys = deptRef.value.getCheckedKeys();
  return checkedKeys;
}
const resetForm = ()=>{
  form.roleId=undefined
  form.roleName=''
  form.authData = []
  deptCheckStrictly.value= false
  menuAuthData.value = []
  deptAuthData.value = []
};
const handleSelectionChange = (val: any[]) => {

}
const setOptionsActionValueHandle = (value:any)=>{
  //遍历选择所有菜单
  menuTableList.value.map((item:any)=>{
    menuAuthData.value[item.id] = value as never
  })
}
const setOptionsItemHandle = (value:any,id:any)=>{
  const children = proxy.findChildrenByPid(id,menuTableList.value)
  //设置行选中
  menuTableList.value.some((item:any)=>{
    if(item.id == id){
      children.unshift(item)
      return true
    }
    return false
  })
  children.map((item:any)=>{
    tableRef.value.toggleRowSelection(item,true)
  })
  //设置子级
  children.map((item:any)=>{
    menuAuthData.value[item.id] = value as never
  })
}
const openSelDeptHandler = (id:any)=>{
  openSelDept.value = true
  setDeptId.value = id
  getRoleDeptTreeselect(id)
}
const getDeptLen = (id:any)=>{
  if(deptAuthData.value[id]){
    return (<Array<any>>deptAuthData.value[id]).length
  }
  return 0
}

</script>

<style scoped lang="scss">
.tree-border {
  margin-top: 5px;
  border: 1px solid #e5e6e7!important;
  border-radius: 4px;
}
.system-edit-role-container {
  .menu-data-tree {
    border: var(--el-input-border, var(--el-border-base));
    border-radius: var(--el-input-border-radius, var(--el-border-radius-base));
    padding: 5px;
  }
}
.action{
  margin-top: 5px;
}
.option-con{
  display: flex;
  height: 40px;
  .el-link{
    margin-left: 30px;
  }
}
</style>
