<template>
  <div class="system-sysNotice-container">
    <el-card shadow="hover">
      <el-tabs
          v-model="tableData.param.type"
          type="card"
          @tab-click="handleTabsClick"
      >
        <el-tab-pane label="通知" :name="1">
          <el-table :data="tableData.data" border style="width: 100%">
            <el-table-column>
              <template #default="scope">
                <el-row @click="handleRead(scope.row)">
                  <el-col :span="0.8">
                    <el-icon class="el_icon">
                      <ele-Bell class="icon"/>
                    </el-icon>
                  </el-col>
                  <el-col :span="23">
                    <span class="title" style="">{{ scope.row.title }}</span>
                    <span style="margin-left: 6px">
                  <el-tag type="success" effect="plain" v-if="scope.row.isRead==true">已读</el-tag>
                  <el-tag type="danger" effect="plain" v-else>未读</el-tag>
                  </span>
                    <span style="float: right">
                      <el-tag type="info" v-if="scope.row.tag==0">无标签</el-tag>
                        <el-tag type="success" v-if="scope.row.tag==1">提醒</el-tag>
                        <el-tag type="success" v-if="scope.row.tag==2">一般</el-tag>
                        <el-tag type="warning" v-if="scope.row.tag==3">次要</el-tag>
                        <el-tag type="danger" v-if="scope.row.tag==4">重要</el-tag>
                        <el-tag type="danger" v-if="scope.row.tag==5">紧急</el-tag>
                    </span>
                    <br>
                    <span>{{ scope.row.createdAt }}</span>
                    <br>
                    <span v-html="scope.row.content"></span>
                    <br>
                  </el-col>
                </el-row>
              </template>
            </el-table-column>

          </el-table>
        </el-tab-pane>
        <el-tab-pane label="私信" :name="2">
          <el-table :data="tableData.data" border style="width: 100%">
            <el-table-column>
              <template #default="scope">
                <el-row @click="handleRead(scope.row)">
                  <el-col :span="0.8">
                    <el-icon class="el_icon">
                      <ele-ChatDotRound class="icon"/>
                    </el-icon>
                  </el-col>
                  <el-col :span="23">
                    <span class="title" style="">{{ scope.row.title }}</span>
                    <span style="margin-left: 6px">

                  <el-tag type="success" effect="plain" v-if="scope.row.isRead==true">已读</el-tag>
                  <el-tag type="danger" effect="plain" v-else>未读</el-tag>
                  </span>
                    <span style="float: right">
                      <el-tag type="info" v-if="scope.row.tag==0">无标签</el-tag>
                        <el-tag type="success" v-if="scope.row.tag==1">提醒</el-tag>
                        <el-tag type="success" v-if="scope.row.tag==2">一般</el-tag>
                        <el-tag type="warning" v-if="scope.row.tag==3">次要</el-tag>
                        <el-tag type="danger" v-if="scope.row.tag==4">重要</el-tag>
                        <el-tag type="danger" v-if="scope.row.tag==5">紧急</el-tag>
                    </span>
                    <br>
                    <span>{{ scope.row.createdAt }}</span>
                    <span v-html="scope.row.content"></span>

                  </el-col>
                </el-row>
              </template>
            </el-table-column>

          </el-table>
        </el-tab-pane>
      </el-tabs>
      <pagination
          v-show="tableData.total>0"
          :total="tableData.total"
          v-model:page="tableData.param.pageNum"
          v-model:limit="tableData.param.pageSize"
          @pagination="sysNoticeList"
      />
    </el-card>

  </div>
</template>
<script setup lang="ts">
import {useRoute} from "vue-router";
import {toRefs, reactive, onMounted, ref, defineComponent, getCurrentInstance} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
  delSysNotice, listShowNotice,
} from "/@/api/system/notice/sysNotice";
import {
  SysNoticeTableColumns,
  SysNoticeTableDataState,
} from "/@/views/system/sysNotice/list/component/model"
import {readNotice} from "/@/api/system/notice/sysNoticeRead";
defineOptions({ name: "apiV1SystemSysNoticeShow"})
const route = useRoute();
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const state = reactive<SysNoticeTableDataState>({
  ids: [],
  tableData: {
    data: [],
    total: 0,
    loading: false,
    param: {
      pageNum: 1,
      pageSize: 10,
      id: undefined,
      title: undefined,
      type: 1,
      tag: undefined,
      status: undefined,
      createdAt: undefined,
      dateRange: []
    },
  },
});
const { tableData}= toRefs(state)
// 页面加载时
onMounted(() => {
  if (route.query.type){
    state.tableData.param.type = parseInt(route.query.type as string)
  }
  initTableData();
});
// 初始化表格数据
const initTableData = () => {
  sysNoticeList()
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
  sysNoticeList()
};
// 获取列表数据
const sysNoticeList = () => {
  loading.value = true
  listShowNotice(state.tableData.param).then((res: any) => {
    let list = res.data.list ?? [];
    list.map((item: any) => {
      item.createdBy = item.createdUser?.userNickname
    })
    state.tableData.data = list;
    state.tableData.total = res.data.total;
    loading.value = false
  })
};


const handleTabsClick = (e: any) => {
  //console.log(e.props.name)
  state.tableData.param.type = e.props.name
  sysNoticeList()
}


const handleDelete = (row: SysNoticeTableColumns) => {
  let msg = '你确定要删除所选数据？';
  let id: number[] = [];
  if (row) {
    msg = `此操作将永久删除数据，是否继续?`
    id = [row.id]
  } else {
    id = state.ids
  }
  if (id.length === 0) {
    ElMessage.error('请选择要删除的数据。');
    return
  }
  ElMessageBox.confirm(msg, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        delSysNotice(id).then(() => {
          ElMessage.success('删除成功');
          sysNoticeList();
        })
      })
      .catch(() => {
      });
}
// eslint-disable-next-line no-undef
const handleRead = (item: any) => {
  // console.log("handleRead", item)
  let query = {
    noticeId: item.id
  }
  readNotice(query).then(() => {
    sysNoticeList()
    ElMessage.success("已读");
  })
}
</script>
<style lang="scss" scoped>
.el_icon {
  width: 28px;
  height: 28px;
  border-radius: 10px;
  background-color: #2d8cf0;
  margin-right: 10px;
}

.title {
  font-size: 16px;
  font-weight: 500;
  color: #0a0a0a
}

.icon {
  width: 28px;
  height: 28px;
  color: #F8F8FF
}

.colBlock {
  display: block;
}

.colNone {
  display: none;
}

.ml-2 {
  margin: 3px;
}


:deep(.el-table__header-wrapper) {
  display: none;
}
</style>
