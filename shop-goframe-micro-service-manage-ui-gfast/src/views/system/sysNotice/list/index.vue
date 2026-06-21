<template>
  <div class="system-sysNotice-container">
    <el-card shadow="hover">
      <div class="system-sysNotice-search mb15">
        <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
          <el-row>
            <el-col :span="5" class="colBlock">
              <el-form-item label="标题" prop="title">
                <el-input
                    v-model="tableData.param.title"
                    placeholder="请输入标题"
                    clearable
                    @keyup.enter.native="sysNoticeList"
                />
              </el-form-item>
            </el-col>
            <el-col :span="5" class="colBlock">
              <el-form-item label="状态" prop="status">
                <el-select v-model="tableData.param.status" placeholder="请选择状态" clearable style="width: 160px">
                  <el-option label="正常" :value="1"/>
                  <el-option label="停用" :value="0"/>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="5" class="colBlock">
              <el-form-item label="类型" prop="type">
                <el-select v-model="tableData.param.type" placeholder="请选择类型" clearable style="width: 160px">
                  <el-option label="通知" value="1"/>
                  <el-option label="公告" value="2"/>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="5" class="colBlock">
              <el-form-item label="标签" prop="tag">
                <el-select v-model="tableData.param.tag" placeholder="请选择标签" clearable style="width: 160px">
                  <el-option
                      v-for="dict in notice_tag"
                      :key="dict.value"
                      :label="dict.label"
                      :value="dict.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="4" class="colBlock">
              <el-form-item>
                <el-button type="primary" @click="sysNoticeList">
                  <el-icon>
                    <ele-Search/>
                  </el-icon>
                  搜索
                </el-button>
                <el-button @click="resetQuery(queryRef)">
                  <el-icon>
                    <ele-Refresh/>
                  </el-icon>
                  重置
                </el-button>

              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
                type="primary"
                @click="handleAdd(1)"
                v-auth="'api/v1/system/sysNotice/noticeAdd'"
            >
              <el-icon>
                <ele-Plus/>
              </el-icon>
              发通知
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="primary"
                @click="handleAdd(2)"
                v-auth="'api/v1/system/sysNotice/messageAdd'"
            >
              <el-icon>
                <ele-Plus/>
              </el-icon>
              发私信
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="success"
                :disabled="single"
                @click="handleUpdate(null)"
                v-auth="'api/v1/system/sysNotice/edit'"
            >
              <el-icon>
                <ele-Edit/>
              </el-icon>
              修改
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="danger"
                :disabled="multiple"
                @click="handleDelete(null)"
                v-auth="'api/v1/system/sysNotice/delete'"
            >
              <el-icon>
                <ele-Delete/>
              </el-icon>
              删除
            </el-button>
          </el-col>
        </el-row>
      </div>
      <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center"/>
        <el-table-column label="ID" align="center" prop="id" min-width="50px"/>
        <el-table-column label="标题" align="left" prop="title" min-width="150px"/>
        <el-table-column label="类型" align="center" prop="type" min-width="150px">
          <template #default="scope">
            <el-tag type="warning" v-if="scope.row.type=='1'">通知</el-tag>
            <el-tag type="danger" v-if="scope.row.type=='2'">私信</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="标签" align="center" prop="tag" :formatter="tagFormat" min-width="80px">
          <template #default="scope">
            <el-tag type="info" v-if="scope.row.tag==0">无标签</el-tag>
            <el-tag type="success" v-if="scope.row.tag==1">提醒</el-tag>
            <el-tag type="success" v-if="scope.row.tag==2">一般</el-tag>
            <el-tag type="warning" v-if="scope.row.tag==3">次要</el-tag>
            <el-tag type="danger" v-if="scope.row.tag==4">重要</el-tag>
            <el-tag type="danger" v-if="scope.row.tag==5">紧急</el-tag>
          </template>
        </el-table-column>
        <!--        <el-table-column label="内容" align="center" prop="content" min-width="150px"/>-->
        <el-table-column label="备注" align="center" prop="remark" min-width="150px"/>
        <el-table-column label="阅读次数" align="center" prop="clickNumber" min-width="150px"/>
        <el-table-column label="排序" align="center" prop="sort" min-width="50px"/>
        <el-table-column label="状态" align="center" prop="status" min-width="80px">
          <template #default="scope">
            <el-tag type="success" v-if="scope.row.status==1">正常</el-tag>
            <el-tag type="danger" v-if="scope.row.status==0">停用</el-tag>
          </template>
        </el-table-column>
        <!--        <el-table-column label="发送人" align="center" prop="createdBy" min-width="150px"/>-->
        <el-table-column label="发送时间" align="center" prop="createdAt" min-width="100px">
          <template #default="scope">
            <span>{{ proxy.parseTime(scope.row.createdAt, '{y}-{m}-{d}') }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" class-name="small-padding" min-width="160px" fixed="right">
          <template #default="scope">
            <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/system/sysNotice/edit'"
            >
              <el-icon>
                <ele-EditPen/>
              </el-icon>
              修改
            </el-button>
            <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/system/sysNotice/delete'"
            >
              <el-icon>
                <ele-DeleteFilled/>
              </el-icon>
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <pagination
          v-show="tableData.total>0"
          :total="tableData.total"
          v-model:page="tableData.param.pageNum"
          v-model:limit="tableData.param.pageSize"
          @pagination="sysNoticeList"
      />
    </el-card>
    <NoticeMessageEdit
        ref="editRef"
        :tagOptions="notice_tag"
        @sysNoticeList="sysNoticeList"
    ></NoticeMessageEdit>

  </div>
</template>
<script setup lang="ts">

import {toRefs, reactive, onMounted, ref, defineComponent, computed, getCurrentInstance, toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
  listSysNotice,
  delSysNotice,
} from "/@/api/system/notice/sysNotice";
import {
  SysNoticeTableColumns,
  SysNoticeInfoData,
  SysNoticeTableDataState,
} from "/@/views/system/sysNotice/list/component/model"
import NoticeMessageEdit from "/@/views/system/sysNotice/list/component/NoticeMessageEdit.vue"
defineOptions({ name: "apiV1SystemSysNoticeList"})
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const queryRef = ref()
const editRef = ref();

// 是否显示所有搜索选项
const showAll = ref(false)
// 非单个禁用
const single = ref(true)
// 非多个禁用
const multiple = ref(true)
const word = computed(() => {
  if (showAll.value === false) {
    //对文字进行处理
    return "展开搜索";
  } else {
    return "收起搜索";
  }
})
// 字典选项数据
const {
  notice_tag,
} = proxy.useDict(
    'notice_tag',
)
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
      type: undefined,
      tag: undefined,
      status: undefined,
      createdAt: undefined,
      dateRange: []
    },
  },
});
const { tableData} = toRefs(state)
// 页面加载时
onMounted(() => {
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
  listSysNotice(state.tableData.param).then((res: any) => {
    let list = res.data.list ?? [];
    list.map((item: any) => {
      item.createdBy = item.createdUser?.userNickname
    })
    state.tableData.data = list;
    state.tableData.total = res.data.total;
    loading.value = false
  })
};
const toggleSearch = () => {
  showAll.value = !showAll.value;
}
// 标签字典翻译
const tagFormat = (row: SysNoticeTableColumns) => {
  return proxy.selectDictLabel(notice_tag.value, row.tag);
}
// 多选框选中数据
const handleSelectionChange = (selection: Array<SysNoticeInfoData>) => {
  state.ids = selection.map(item => item.id)
  single.value = selection.length != 1
  multiple.value = !selection.length
}

const handleAdd = (type: number) => {
  editRef.value.openDialog()
  editRef.value.setType(type)
}
const handleUpdate = (row: SysNoticeTableColumns|null) => {
  if (!row) {
    row = state.tableData.data.find((item: SysNoticeTableColumns) => {
      return item.id === state.ids[0]
    }) as SysNoticeTableColumns
  }
  editRef.value.openDialog(toRaw(row));
};
const handleDelete = (row: SysNoticeTableColumns|null) => {
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
</script>
<style lang="scss" scoped>
.colBlock {
  display: block;
}

.colNone {
  display: none;
}

.ml-2 {
  margin: 3px;
}
</style>
