<template>
  <div>
    <div v-if="deptUser.length > 0">
      <el-tag  closable  :disable-transitions="false" class="u-m-r-10" v-for="(item) in deptUser"  :key="'sel_'+item.id" @close="handleClose(item.id)" >{{item.userNickname}}</el-tag>
    </div>
    <el-button
        style="padding-left: 10px;"
        type="primary"
        link
        @click="handleSelectUser" >请选择</el-button>
    <el-dialog title="选择用户" v-model="visible" width="80%" top="5vh" append-to-body :close-on-click-modal="false">
      <div class="system-user-container">
        <el-row :gutter="10" style="width: 100%;">
          <el-col :span="6">
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
          <el-col :span="12">
            <el-card shadow="hover">
              <div class="system-user-search mb15">
                <el-form :model="param" ref="queryRef" :inline="true" label-width="68px">
                  <el-form-item label="关键字" prop="keyWords">
                    <el-input
                        v-model="param.keyWords"
                        placeholder="请输入姓名"
                        clearable
                        style="width: 240px"
                        @keyup.enter.native="getUserList"
                    />
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
                    <el-button size="default" type="primary" class="ml10" @click="getUserList">
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
              </div>
              <UserList ref="userListRef" :dept-data="deptData" :gender-data="sys_user_sex" :param="param" :multiple="multiple"
                        @ok="handleSelectUserOk" @okBatch="handleSelectUserOkBatch"/>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover">
              <el-row :gutter="10">
                <el-col :span="10">
                  <p style="height: 32px;line-height: 32px;">已选择：{{selectedUsers.length}}</p>
                </el-col>
                <el-col :span="7">
                  <el-button type="success" plain @click="goBack">确认返回</el-button>
                </el-col>
                <el-col :span="7">
                  <el-button type="danger" plain @click="removeAll">全部移除</el-button>
                </el-col>
                <el-col :span="24">
                  <el-table :data="selectedUserInfo">
                    <el-table-column  label="操作"  >
                      <template #default="scope">
                        <el-button type="danger" plain @click="remove(scope.row.id)">移除</el-button>
                      </template>
                    </el-table-column>
                    <el-table-column prop="userNickname" label="姓名"  />
                  </el-table>
                </el-col>
              </el-row>
              <el-pagination layout="prev, pager, next" :total="selectedUsers.length" v-model:current-page="selectedUsersPage"/>
            </el-card>
          </el-col>
        </el-row>
      </div>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import {toRefs, reactive, ref, defineComponent, watch, getCurrentInstance, nextTick, computed, onMounted} from 'vue';
import {ElTree,FormInstance} from 'element-plus';
import { Search } from '@element-plus/icons-vue'
import UserList from './component/userList.vue';
import {getDeptTree, getUserByIds} from '/@/api/system/user/index';

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

export default defineComponent({
  name: 'systemUser',
  components: { UserList },
  props:{
    multiple:{
      type:Boolean,
      default:true
    },
    modelValue:{
      type:Array<number>,
      default:()=>[]
    }
  },
  emits:['update:modelValue'],
  setup(prop,{emit}) {
    const selectedUsersPage = ref(1)
    const visible = ref(false)
    const {proxy} = <any>getCurrentInstance();
    const {sys_user_sex} = proxy.useDict('sys_user_sex')
    const userListRef = ref();
    const queryRef = ref();
    const filterText = ref('');
    const treeRef = ref<InstanceType<typeof ElTree>>();
    const search = Search
    const  deptUser = ref<any>([]);
    const selectedUsers = computed({
      get: ()=>{
        return prop.modelValue??[]
      },
      set:(val:any[])=>{
        emit('update:modelValue',val)
      }
    })
    const selectedUserInfo = computed(()=>{
      let start = (selectedUsersPage.value-1)*10
      let end = start+10
      return deptUser.value.slice(start,end)
    });
    const initData = ()=>{
      if(prop.modelValue&&prop.modelValue.length>0){
        getUserByIds({ids:prop.modelValue}).then((res:any)=>{
          if(res.code === 0){
            deptUser.value = res.data.userList??[];
          }
        });
      }else{
        deptUser.value = []
      }
    };
    onMounted(()=>{
      initData()
    })
    watch(selectedUsers,(value, oldValue)=>{
      if(value!=oldValue){
        initData()
      }
    })
    const state = reactive<QueryParam>({
      ids:[],
      deptProps:{
        id:"deptId",
        children: "children",
        label: "deptName"
      },
      deptData:[],
      param: {
        deptId:'',
        mobile:'',
        status:'',
        keyWords:'',
        dateRange:[]
      },
    });
    const getUserList = ()=>{
        userListRef.value.setUserList();
    };
    // 初始化表格数据
    const initTableData = () => {
      getDeptTree().then((res:any)=>{
        state.deptData = res.data.deps
      })
      getUserList()
    };

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
      getUserList();
    };
    /** 重置按钮操作 */
    const resetQuery = (formEl: FormInstance | undefined) => {
      if (!formEl) return
      formEl.resetFields()
      getUserList()
    };
    const openDialog = ()=>{
      visible.value = true
      nextTick(()=>{
        initTableData();
      })
    }
    const handleSelectUserOk = (row:any)=>{
      if(!prop.multiple){
        selectedUsers.value = [row.id]
      }else{
        if(!selectedUsers.value.includes(row.id)){
          selectedUsers.value = [...selectedUsers.value!,row.id]
        }
      }

    }
    const handleSelectUserOkBatch = (row:any[])=>{
      if(prop.multiple){
        const ids:any[] = []
        row.forEach((item:any)=>{
          if(!selectedUsers.value.includes(item.id)){
            ids.push(item.id)
          }
        })
        if(ids.length>0){
          selectedUsers.value = [...selectedUsers.value!,...ids]
        }
      }
    }
    const goBack = ()=>{
      visible.value = false;
    }
    const removeAll = ()=>{
      selectedUsers.value = []
    }
    const remove = (id:number)=>{
      selectedUsers.value = selectedUsers.value.filter((item:any)=>{
        return item !== id
      })
    }
    const handleSelectUser = ()=>{
      openDialog()
    }
    const handleClose = (id:number)=>{
      selectedUsers.value =  selectedUsers.value.filter((item:any)=>{
        return item !== id
      })
    }
    return {
      selectedUsersPage,
      visible,
      queryRef,
      userListRef,
      deptFilterNode,
      filterText,
      treeRef,
      search,
      sys_user_sex,
      selectedUsers,
      deptUser,
      selectedUserInfo,
      openDialog,
      getUserList,
      handleSelectUserOk,
      handleSelectUserOkBatch,
      handleNodeClick,
      resetQuery,
      goBack,
      removeAll,
      remove,
      handleClose,
      handleSelectUser,
      ...toRefs(state),
    };
  },
});
</script>
<style scoped>
.u-m-r-10{
  margin-right: 8px;
}
</style>
