<template>
	<div class="system-sysJob-container">
		<el-card shadow="hover">
			<div class="system-sysJob-search mb15">
				<el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
					<el-row>
						<el-col :span="8" class="colBlock">
							<el-form-item label="任务名称" prop="jobName">
								<el-input v-model="tableData.param.jobName" placeholder="请输入任务名称" clearable @keyup.enter.native="sysJobList" />
							</el-form-item>
						</el-col>
						<el-col :span="8" class="colBlock">
							<el-form-item label="任务组名" prop="jobGroup">
								<el-select v-model="tableData.param.jobGroup" placeholder="请选择任务组名" clearable style="width: 160px;">
									<el-option v-for="dict in sys_job_group" :key="dict.value" :label="dict.label" :value="dict.value" />
								</el-select>
							</el-form-item>
						</el-col>
						<el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
							<el-form-item>
								<el-button type="primary" size="small" @click="sysJobList"
									><el-icon><ele-Search /></el-icon>搜索</el-button
								>
								<el-button size="small" @click="resetQuery(queryRef)"
									><el-icon><ele-Refresh /></el-icon>重置</el-button
								>
								<el-button type="primary" link size="small" @click="toggleSearch">
									{{ word }}
									<el-icon v-show="showAll"><ele-ArrowUp /></el-icon>
									<el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
								</el-button>
							</el-form-item>
						</el-col>
						<el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
							<el-form-item label="状态" prop="status">
								<el-select v-model="tableData.param.status" placeholder="请选择状态" clearable style="width: 160px;">
									<el-option v-for="dict in sys_job_status" :key="dict.value" :label="dict.label" :value="dict.value" />
								</el-select>
							</el-form-item>
						</el-col>
						<el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
							<el-form-item>
								<el-button type="primary" size="small" @click="sysJobList"
									><el-icon><ele-Search /></el-icon>搜索</el-button
								>
								<el-button size="small" @click="resetQuery(queryRef)"
									><el-icon><ele-Refresh /></el-icon>重置</el-button
								>
								<el-button type="primary" link size="small" @click="toggleSearch">
									{{ word }}
									<el-icon v-show="showAll"><ele-ArrowUp /></el-icon>
									<el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
								</el-button>
							</el-form-item>
						</el-col>
					</el-row>
				</el-form>
				<el-row :gutter="10" class="mb8">
					<el-col :span="1.5">
						<el-button type="primary" size="small" @click="handleAdd" v-auth="'api/v1/system/sysJob/add'"
							><el-icon><ele-Plus /></el-icon>新增</el-button
						>
					</el-col>
					<el-col :span="1.5">
						<el-button type="success" size="small" :disabled="single" @click="handleUpdate(null)" v-auth="'api/v1/system/sysJob/edit'"
							><el-icon><ele-Edit /></el-icon>修改</el-button
						>
					</el-col>
					<el-col :span="1.5">
						<el-button type="danger" size="small" :disabled="multiple" @click="handleDelete(null)" v-auth="'api/v1/system/sysJob/delete'"
							><el-icon><ele-Delete /></el-icon>删除</el-button
						>
					</el-col>
				</el-row>
			</div>
			<el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
				<el-table-column type="selection" width="55" align="center" />
				<el-table-column label="任务ID" align="center" prop="jobId" min-width="100px" />
				<el-table-column label="任务名称" align="center" prop="jobName" min-width="100px" />
				<el-table-column label="任务组名" align="center" prop="jobGroup" :formatter="jobGroupFormat" min-width="100px" />
				<el-table-column label="任务方法" align="center" prop="invokeTarget" min-width="100px" />
				<el-table-column label="cron执行表达式" align="center" prop="cronExpression" min-width="100px" />
				<el-table-column label="计划执行策略" align="center" prop="misfirePolicy" :formatter="misfirePolicyFormat" min-width="100px" />
				<el-table-column label="状态" align="center">
					<template #default="scope">
						<el-switch v-model="scope.row.status" :active-value="0" :inactive-value="1" @change="handleStatusChange(scope.row)"></el-switch>
					</template>
				</el-table-column>
				<el-table-column label="操作" align="center" class-name="small-padding" min-width="200px" fixed="right">
					<template #default="scope">
						<el-button size="small" type="primary" link @click="handleView(scope.row)" v-auth="'api/v1/system/sysJob/view'"
							><el-icon><ele-View /></el-icon>详情</el-button
						>
						<el-button size="small" type="primary" link @click="handleUpdate(scope.row)" v-auth="'api/v1/system/sysJob/edit'"
							><el-icon><ele-EditPen /></el-icon>修改</el-button
						>
						<el-button size="small" type="primary" link @click="handleJobRun(scope.row)" v-auth="'api/v1/system/sysJob/run'"
							><el-icon><ele-CaretRight /></el-icon>执行一次</el-button
						>
						<el-button size="small" type="primary" link @click="handleDelete(scope.row)" v-auth="'api/v1/system/sysJob/delete'"
							><el-icon><ele-DeleteFilled /></el-icon>删除</el-button
						>
					</template>
				</el-table-column>
			</el-table>
			<pagination
				v-show="tableData.total > 0"
				:total="tableData.total"
				v-model:page="tableData.param.pageNum"
				v-model:limit="tableData.param.pageSize"
				@pagination="sysJobList"
			/>
		</el-card>
		<apiV1SystemSysJobEdit
			ref="editRef"
			:jobGroupOptions="sys_job_group"
			:misfirePolicyOptions="sys_job_policy"
			:statusOptions="sys_job_status"
			@sysJobList="sysJobList"
		></apiV1SystemSysJobEdit>
		<apiV1SystemSysJobDetail
			ref="detailRef"
			:jobGroupOptions="sys_job_group"
			:misfirePolicyOptions="sys_job_policy"
			:statusOptions="sys_job_status"
			@sysJobList="sysJobList"
		></apiV1SystemSysJobDetail>
	</div>
</template>
<script setup lang="ts">
import { ItemOptions } from '/@/api/items';
import { toRefs, reactive, onMounted, ref, defineComponent, computed, getCurrentInstance, toRaw } from 'vue';
import { ElMessageBox, ElMessage, FormInstance } from 'element-plus';
import { listSysJob, getSysJob, delSysJob, addSysJob, updateSysJob, getUserList, runJob, startJob, stopJob } from '/@/api/system/tools/job/sysJob';
import { SysJobTableColumns, SysJobInfoData, SysJobTableDataState } from '/@/views/system/sysJob/list/component/model';
import apiV1SystemSysJobEdit from '/@/views/system/sysJob/list/component/edit.vue';
import apiV1SystemSysJobDetail from '/@/views/system/sysJob/list/component/detail.vue';
defineOptions({ name: "apiV1SystemSysJobList"})
const { proxy } = <any>getCurrentInstance();
const loading = ref(false);
const queryRef = ref();
const editRef = ref();
const detailRef = ref();
// 是否显示所有搜索选项
const showAll = ref(false);
// 非单个禁用
const single = ref(true);
// 非多个禁用
const multiple = ref(true);
const word = computed(() => {
  if (showAll.value === false) {
    //对文字进行处理
    return '展开搜索';
  } else {
    return '收起搜索';
  }
});
// 字典选项数据
const { sys_job_group, sys_job_policy, sys_job_status } = proxy.useDict('sys_job_group', 'sys_job_policy', 'sys_job_status');
const state = reactive<SysJobTableDataState>({
  jobIds: [],
  tableData: {
    data: [],
    total: 0,
    loading: false,
    param: {
      pageNum: 1,
      pageSize: 10,
      jobName: undefined,
      jobGroup: undefined,
      status: undefined,
    },
  },
});
const { tableData } = toRefs(state);
// 页面加载时
onMounted(() => {
  initTableData();
});
// 初始化表格数据
const initTableData = () => {
  sysJobList();
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
  sysJobList();
};
// 获取列表数据
const sysJobList = () => {
  loading.value = true;
  listSysJob(state.tableData.param).then((res: any) => {
    let list = res.data.list ?? [];
    state.tableData.data = list;
    state.tableData.total = res.data.total;
    loading.value = false;
  });
};
const toggleSearch = () => {
  showAll.value = !showAll.value;
};
// 任务组名字典翻译
const jobGroupFormat = (row: SysJobTableColumns) => {
  return proxy.selectDictLabel(sys_job_group.value, row.jobGroup);
};
// 计划执行策略字典翻译
const misfirePolicyFormat = (row: SysJobTableColumns) => {
  return proxy.selectDictLabel(sys_job_policy.value, row.misfirePolicy);
};
// 状态字典翻译
const statusFormat = (row: SysJobTableColumns) => {
  return proxy.selectDictLabel(sys_job_status.value, row.status);
};
// 多选框选中数据
const handleSelectionChange = (selection: Array<SysJobInfoData>) => {
  state.jobIds = selection.map((item) => item.jobId);
  single.value = selection.length != 1;
  multiple.value = !selection.length;
};
const handleAdd = () => {
  editRef.value.openDialog();
};
const handleUpdate = (row: SysJobTableColumns|null) => {
  if (!row) {
    row = state.tableData.data.find((item: SysJobTableColumns) => {
      return item.jobId === state.jobIds[0];
    }) as SysJobTableColumns;
  }
  editRef.value.openDialog(toRaw(row));
};
/** 执行按钮操作 */
const handleJobRun = (row: SysJobTableColumns) => {
  const jobId = row.jobId || 0;
  ElMessageBox.confirm('是否确认立即执行一次该任务?', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(function () {
      return runJob(jobId);
    })
    .then(() => {
      ElMessage.success('执行成功');
    })
    .catch(function () {});
};
const handleDelete = (row: SysJobTableColumns|null) => {
  let msg = '你确定要删除所选数据？';
  let jobId: number[] = [];
  if (row) {
    msg = `此操作将永久删除数据，是否继续?`;
    jobId = [row.jobId];
  } else {
    jobId = state.jobIds;
  }
  if (jobId.length === 0) {
    ElMessage.error('请选择要删除的数据。');
    return;
  }
  ElMessageBox.confirm(msg, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      delSysJob(jobId).then(() => {
        ElMessage.success('删除成功');
        sysJobList();
      });
    })
    .catch(() => {});
};
const handleView = (row: SysJobTableColumns) => {
  detailRef.value.openDialog(toRaw(row));
};
// 任务状态修改
const handleStatusChange = (row: SysJobTableColumns) => {
  let text = row.status === 0 ? '启用' : '停用';
  ElMessageBox.confirm('确认要"' + text + '""' + row.jobName + '"任务吗?', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(function () {
      if (row.status === 0) {
        return startJob(row.jobId);
      } else {
        return stopJob(row.jobId);
      }
    })
    .then(() => {
      ElMessage.success(text + '成功');
    })
    .catch(function () {
      row.status = row.status === 0 ? 1 : 0;
    });
};
</script>
<style lang="scss" scoped>
.colBlock {
	display: block;
}
.colNone {
	display: none;
}
</style>
