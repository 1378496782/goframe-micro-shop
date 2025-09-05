<template>
	<div class="system-edit-dic-container">
		<el-dialog :title="(ruleForm.dictId!==0?'修改':'添加')+'字典'" v-model="isShowDialog" width="769px">
			<el-form :model="ruleForm" ref="formRef" :rules="rules" size="default" label-width="90px">
				<el-form-item label="上级" prop="pid">
					<el-cascader v-model="ruleForm.pid" :options="dictTypeOpt" :props="typeProps" clearable />
				</el-form-item>
        <el-form-item label="字典名称" prop="dictName">
          <el-input v-model="ruleForm.dictName" placeholder="请输入字典名称" />
        </el-form-item>
        <el-form-item label="字典类型" prop="dictType">
          <el-input v-model="ruleForm.dictType" placeholder="请输入字典类型" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="ruleForm.status">
            <el-radio :value="1" >启用</el-radio>
            <el-radio :value="0" >禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="ruleForm.remark" type="textarea" placeholder="请输入内容"></el-input>
        </el-form-item>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default">{{ruleForm.dictId!==0?'修 改':'添 加'}}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import { reactive, toRefs, ref, unref, getCurrentInstance } from 'vue';
import { getType, addType, editType, optionselect } from '/@/api/system/dict/type';
import {ElMessage} from "element-plus";
interface RuleFormState {
  dictId:number;
	pid:number;
  dictName:string;
  dictType:string;
  status:number;
  remark:string;
}
interface DicState {
	isShowDialog: boolean;
	ruleForm: RuleFormState;
  rules:{}
}
defineOptions({ name: "systemEditDic"})
const {proxy} = getCurrentInstance() as any;
const emit = defineEmits(['typeList']);
const formRef = ref<HTMLElement | null>(null);
const dictTypeOpt = ref<RuleFormState>()
const typeProps = ref({
	value: 'dictId',
	label: 'dictName',
	children: 'children',
	checkStrictly:true,
	emitPath: false
})
const state = reactive<DicState>({
  isShowDialog: false,
  ruleForm: {
    dictId:0,
		pid:0,
    dictName:'',
    dictType:'',
    status:1,
    remark:''
  },
  rules: {
    dictName: [
      { required: true, message: "字典名称不能为空", trigger: "blur" }
    ],
    dictType: [
      { required: true, message: "字典类型不能为空", trigger: "blur" }
    ]
  }
});
const { isShowDialog,ruleForm,rules } = toRefs(state);
const getParent = () => {
	optionselect(true).then((res:any)=>{
		const data = res.data.dictType??[]
		dictTypeOpt.value = proxy.handleTree(data, 'dictId', 'pid', 'children', true)
  })
}
// 打开弹窗
const openDialog = (row: RuleFormState|null) => {
  resetForm();
	getParent();
  if (row){
    getType(row.dictId).then((res:any)=>{
      state.ruleForm = res.data.dictType
    })
    state.ruleForm = row;
  }
  state.isShowDialog = true;
};
defineExpose({ openDialog})
const resetForm = ()=>{
  state.ruleForm = {
    dictId:0,
		pid:0,
    dictName:'',
    dictType:'',
    status:1,
    remark:''
  }
};
// 关闭弹窗
const closeDialog = () => {
  state.isShowDialog = false;
};
// 取消
const onCancel = () => {
  closeDialog();
};
// 新增
const onSubmit = () => {
  const formWrap = unref(formRef) as any;
  if (!formWrap) return;
  formWrap.validate((valid: boolean) => {
    if (valid) {
      if(state.ruleForm.dictId!==0){
        //修改
        editType(state.ruleForm).then(()=>{
          ElMessage.success('字典类型修改成功');
          closeDialog(); // 关闭弹窗
          emit('typeList')
        })
      }else{
        //添加
        addType(state.ruleForm).then(()=>{
          ElMessage.success('字典类型添加成功');
          closeDialog(); // 关闭弹窗
          emit('typeList')
        })
      }
    }
  });
};

</script>
