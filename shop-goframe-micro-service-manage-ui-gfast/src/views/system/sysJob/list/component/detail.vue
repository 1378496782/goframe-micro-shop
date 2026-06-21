<template>
  <!-- 定时任务详情抽屉 -->
  <div class="system-sysJob-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>定时任务详情</h4>
      </template>
      <el-form ref="formRef" :model="formData" label-width="120px">
        <el-row>
          <el-col :span="12">
            <el-form-item label="任务ID">{{ formData.jobId }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="任务名称">{{ formData.jobName }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="参数">{{ formData.jobParams }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="任务组名">{{ proxy.getOptionValue(formData.jobGroup, jobGroupOptions,'value','label') }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="任务方法">{{ formData.invokeTarget }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="cron执行表达式">{{ formData.cronExpression }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="计划执行策略">{{ proxy.getOptionValue(formData.misfirePolicy, misfirePolicyOptions,'value','label') }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="状态">{{ proxy.getOptionValue(formData.status, statusOptions,'value','label') }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="创建者">{{ formData.createdUser?.userNickname }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="更新者">{{ formData.updatedUser?.userNickname }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注信息">{{ formData.remark }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="创建时间">{{ proxy.parseTime(formData.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <el-divider />
      <h3>任务执行日志</h3>
      <el-button
          class="btn-del"
          type="danger"
          size="small"
          @click="handleDelete"
      ><el-icon><ele-Delete /></el-icon>清空日志</el-button>
      <el-divider />
      <el-table v-loading="logList.loading" :data="logList.data" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center" />
        <el-table-column label="日志ID" align="center" prop="id"
                         min-width="100px"
        />
        <el-table-column label="执行方法" align="center" prop="targetName"
                         min-width="150px"
        />
        <el-table-column label="执行日期" align="center" prop="createdAt" />
        <el-table-column label="执行结果" align="center" prop="result"/>
      </el-table>
      <pagination
          v-show="logList.total>0"
          :total="logList.total"
          v-model:page="logList.param.pageNum"
          v-model:limit="logList.param.pageSize"
          @pagination="getLogList(formData.invokeTarget)"
      />
    </el-drawer>
  </div>
</template>
<script setup lang="ts">
import { reactive, toRefs, defineComponent,ref,unref,getCurrentInstance,computed } from 'vue';
import {ElMessageBox, ElMessage, FormInstance,UploadProps} from 'element-plus';
import {
  listSysJob,
  getSysJob,
  delSysJob,
  addSysJob,
  updateSysJob,
  getUserList, getJobLogs, delSysJobLog,
} from "/@/api/system/tools/job/sysJob";
import {
  SysJobTableColumns,
  SysJobInfoData,
  SysJobTableDataState,
  SysJobEditState, SysJobLog, SysJobLogData,
} from "/@/views/system/sysJob/list/component/model"
defineOptions({ name: "apiV1SystemSysJobDetail"})
const props = defineProps({
  jobGroupOptions:{
    type:Array,
    default:()=>[]
  },
  misfirePolicyOptions:{
    type:Array,
    default:()=>[]
  },
  statusOptions:{
    type:Array,
    default:()=>[]
  },
})
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
// 非多个禁用
const multiple =ref(true)
const logList = reactive<SysJobLog>({
  logIds:[],
  data: [],
  loading: false,
  param: {
    pageNum: 1,
    pageSize: 10,
    targetName: undefined
  },
  total: 0
})
const state = reactive<SysJobEditState>({
  loading:false,
  isShowDialog: false,
  formData: {
    jobId: undefined,
    jobName: undefined,
    jobParams: undefined,
    jobGroup: undefined,
    invokeTarget: undefined,
    cronExpression: undefined,
    misfirePolicy: false ,
    concurrent: undefined,
    status: false ,
    createdBy: undefined,
    updatedBy: undefined,
    remark: undefined,
    createdAt: undefined,
    updatedAt: undefined,
    createdUser:{userNickname:''},
    updatedUser:{userNickname:''}
  },
  // 表单校验
  rules: {}
});
const { isShowDialog,formData} = toRefs(state)
// 打开弹窗
const openDialog = (row?: SysJobInfoData) => {
resetForm();
if(row) {
getSysJob(row.jobId!).then((res:any)=>{
  state.formData = res.data;
})
logList.loading=true
getLogList(row.invokeTarget!);
}
state.isShowDialog = true;
};
defineExpose({openDialog})
const getLogList = (invokeTarget:string)=>{
  logList.param.targetName=invokeTarget
  getJobLogs(logList.param).then((res:any)=>{
    logList.data = res.data.list||[];
    logList.total=res.data.total;
    logList.loading=false;
  })
};
// 关闭弹窗
const closeDialog = () => {
  state.isShowDialog = false;
};
// 取消
const onCancel = () => {
  closeDialog();
};
const resetForm = ()=>{
  state.formData = {
    jobId: undefined,
    jobName: undefined,
    jobParams: undefined,
    jobGroup: undefined,
    invokeTarget: undefined,
    cronExpression: undefined,
    misfirePolicy: false ,
    concurrent: undefined,
    status: false ,
    createdBy: undefined,
    updatedBy: undefined,
    remark: undefined,
    createdAt: undefined,
    updatedAt: undefined,
    createdUser:{userNickname:''},
    updatedUser:{userNickname:''}
  }
};
// 多选框选中数据
const handleSelectionChange = (selection:Array<SysJobLogData>) => {
  logList.logIds = selection.map(item => item.id)
  multiple.value = !selection.length
};
const handleDelete = () => {
  ElMessageBox.confirm('你确定要清空日志？', '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        delSysJobLog(logList.param.targetName!).then(()=>{
          ElMessage.success('删除成功');
          getLogList(logList.param.targetName!);
        })
      })
      .catch(() => {});
}
</script>
<style scoped>
  .system-sysJob-detail :deep(.el-form-item--large .el-form-item__label){
    font-weight: bolder;
  }
  .pic-block{
    margin-right: 8px;
  }
  .file-block{
    width: 100%;
    border: 1px solid var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
    margin-bottom: 5px;
    padding: 3px 6px;
  }
  .ml-2{margin-right: 5px;}
  .btn-del{
    margin-top: 20px;
  }
</style>
