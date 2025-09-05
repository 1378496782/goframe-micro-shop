<template>
  <div class="system-sysJob-edit">
    <!-- 添加或修改定时任务对话框 -->
    <el-dialog v-model="isShowDialog" width="769px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.system-sysJob-edit .el-dialog', '.system-sysJob-edit .el-dialog__header']">{{(!formData.jobId || formData.jobId==0?'添加':'修改')+'定时任务'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" size="default" label-width="120px">
        <el-form-item label="任务名称" prop="jobName">
          <el-input v-model="formData.jobName" placeholder="请输入任务名称" />
        </el-form-item>
        <el-form-item prop="jobParams">
          <template #label>
							<span>
								执行参数
								<el-tooltip content="以|分隔多个参数 示例：param1|param1；参数说明：支持字符串，布尔类型，长整型，浮点型，整型" placement="top">
									<el-icon><ele-QuestionFilled /></el-icon>
								</el-tooltip>
							</span>
          </template>
          <el-input v-model="formData.jobParams" placeholder="请输入参数" />
        </el-form-item>
        <el-form-item label="任务组名" prop="jobGroup">
          <el-select v-model="formData.jobGroup" placeholder="请选择任务组名" >
            <el-option
              v-for="dict in jobGroupOptions"
              :key="dict.value"
              :label="dict.label"
                  :value="dict.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="invokeTarget">
          <template #label>
							<span>
								任务方法
								<el-tooltip content="示例：demo1" placement="top">
									<el-icon><ele-QuestionFilled /></el-icon>
								</el-tooltip>
							</span>
          </template>
          <el-input v-model="formData.invokeTarget" placeholder="请输入任务方法" />
        </el-form-item>
        <el-form-item label="cron执行表达式" prop="cronExpression">
          <el-input v-model="formData.cronExpression" placeholder="请输入cron执行表达式" />
        </el-form-item>
        <ul class="cronExpression-rule">
          <li><el-icon><ele-WarningFilled /></el-icon> */5 * * * * ? : 每隔5秒执行一次</li>
          <li><el-icon><ele-WarningFilled /></el-icon> 20 */1 * * * ? : 每隔1分钟执行一次</li>
          <li><el-icon><ele-WarningFilled /></el-icon> 30 0 23 * * ? : 每天23点执行一次</li>
          <li><el-icon><ele-WarningFilled /></el-icon> 0 0 1 * * ? : 每天凌晨1点执行一次</li>
          <li><el-icon><ele-WarningFilled /></el-icon> 0 0 1 1 * ? : 每月1号凌晨1点执行一次</li>
        </ul>
        <el-form-item label="计划执行策略" prop="misfirePolicy">
          <el-radio-group v-model="formData.misfirePolicy">
            <el-radio
              v-for="dict in misfirePolicyOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio
              v-for="dict in statusOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注信息" prop="remark">
          <el-input v-model="formData.remark" placeholder="请输入备注信息" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="onSubmit">确 定</el-button>
          <el-button @click="onCancel">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import { reactive, toRefs, defineComponent,ref,unref,getCurrentInstance } from 'vue';
import {ElMessageBox, ElMessage, FormInstance,UploadProps} from 'element-plus';
import {
  listSysJob,
  getSysJob,
  delSysJob,
  addSysJob,
  updateSysJob,
  getUserList,
} from "/@/api/system/tools/job/sysJob";
import {
  SysJobTableColumns,
  SysJobInfoData,
  SysJobTableDataState,
  SysJobEditState,
} from "/@/views/system/sysJob/list/component/model"
defineOptions({ name: "apiV1SystemSysJobEdit"})
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
const emit = defineEmits(['sysJobList'])
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
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
  rules: {
    cronExpression: [
      { required: true, message: "cron执行表达式不能为空", trigger: "blur" }
    ],
    jobName : [
        { required: true, message: "任务名称不能为空", trigger: "blur" }
    ],
    invokeTarget : [
        { required: true, message: "任务方法不能为空", trigger: "blur" }
    ],
    status : [
        { required: true, message: "状态不能为空", trigger: "blur" }
    ],
  }
});
const {isShowDialog,formData,rules} = toRefs(state);
// 打开弹窗
const openDialog = (row?: SysJobInfoData) => {
  resetForm();
  if(row) {
    getSysJob(row.jobId!).then((res:any)=>{
      const data = res.data;
      data.jobGroup = ''+data.jobGroup
      data.misfirePolicy = ''+data.misfirePolicy
      data.status = ''+data.status
      state.formData = data;
  })
}
  state.isShowDialog = true;
};
defineExpose({openDialog})
// 关闭弹窗
const closeDialog = () => {
  state.isShowDialog = false;
};
// 取消
const onCancel = () => {
  closeDialog();
};
// 提交
const onSubmit = () => {
  const formWrap = unref(formRef) as any;
  if (!formWrap) return;
  formWrap.validate((valid: boolean) => {
    if (valid) {
      state.loading = true;
      if(!state.formData.jobId || state.formData.jobId===0){
        //添加
      addSysJob(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('sysJobList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateSysJob(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('sysJobList')
        }).finally(()=>{
          state.loading = false;
        })
      }
    }
  });
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
</script>
<style scoped lang="scss">
.cronExpression-rule{
  list-style: none;
  width: 60%;
  margin: 20px auto;
  li{
    margin-bottom: 15px;
    font-weight: bolder;
  }
}
</style>
