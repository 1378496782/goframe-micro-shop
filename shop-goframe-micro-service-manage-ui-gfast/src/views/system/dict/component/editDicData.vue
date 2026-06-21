<template>
	<div class="system-edit-dic-container">
		<el-dialog :title="(ruleForm.dictCode!==0?'修改':'添加')+'字典'" v-model="isShowDialog" width="769px">
			<el-form :model="ruleForm" ref="formRef" :rules="rules" size="default" label-width="90px">
        <el-form-item label="字典类型" prop="dictType">
					<el-cascader v-model="ruleForm.dictType" :options="dictTypeOpt" :props="typeProps" clearable :show-all-levels="false"/>
        </el-form-item>
        <el-form-item label="数据标签" prop="dictLabel">
          <el-input v-model="ruleForm.dictLabel" placeholder="请输入数据标签" />
        </el-form-item>
        <el-form-item label="数据键值" prop="dictValue">
          <el-input v-model="ruleForm.dictValue" placeholder="请输入数据键值" />
        </el-form-item>
        <el-form-item label="显示排序" prop="dictSort">
          <el-input-number v-model="ruleForm.dictSort" controls-position="right" :min="0" />
        </el-form-item>

        <el-form-item label="系统默认">
          <el-switch v-model="ruleForm.isDefault"
                     active-text="是"
                     inactive-text="否"
                     :active-value="1"
                     :inactive-value="0"
          ></el-switch>
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
					<el-button type="primary" @click="onSubmit" size="default">{{ruleForm.dictCode!==0?'修 改':'添 加'}}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import { reactive, toRefs, ref, unref, getCurrentInstance } from 'vue';
import { getData,addData,editData } from '/@/api/system/dict/data';
import {ElMessage} from "element-plus";
import { optionselect } from '/@/api/system/dict/type';
interface RuleFormState {
  dictCode: number;
  dictLabel: string;
  dictValue: string;
  dictSort:number;
  isDefault:number;
  status:number;
  remark:string;
  dictType:string;
}
interface DicState {
	isShowDialog: boolean;
	ruleForm: RuleFormState;
  rules:{}
}
defineOptions({ name: "systemEditDicData"})
const {proxy} = getCurrentInstance() as any;
const prop = defineProps({
  dictType:{
    type:String,
    default:''
  }
})
const emit = defineEmits(['dataList']);
const formRef = ref<HTMLElement | null>(null);
const dictTypeOpt = ref<RuleFormState>()
const typeProps = ref({
	value: 'dictType',
	label: 'dictName',
	children: 'children',
	emitPath: false,
	checkStrictly:true
})
const state = reactive<DicState>({
  isShowDialog: false,
  ruleForm: {
    dictCode: 0,
    dictLabel: '',
    dictValue: '',
    dictSort: 0,
    isDefault:0,
    status: 1,
    remark: '',
    dictType:prop.dictType
  },
  rules: {
		dictType:[
			{ required: true, message: "请选择字典类型", trigger: "change" }
		],
    dictLabel: [
      { required: true, message: "数据标签不能为空", trigger: "blur" }
    ],
    dictValue: [
      { required: true, message: "数据键值不能为空", trigger: "blur" }
    ],
    dictSort: [
      { required: true, message: "数据顺序不能为空", trigger: "blur" }
    ]
  }
});
const { isShowDialog,ruleForm,rules } = toRefs(state);
const getTypeOpt = () => {
	optionselect(true).then((res:any)=>{
		const data = res.data.dictType??[]
		dictTypeOpt.value = proxy.handleTree(data, 'dictId', 'pid', 'children', true)
	})
}
// 打开弹窗
const openDialog = (row: RuleFormState|null) => {
	getTypeOpt()
  resetForm()
  if (row){
    getData(row.dictCode).then((res:any)=>{
      state.ruleForm = res.data.dict
    })
    state.ruleForm = row;
  }
  state.isShowDialog = true;
};
defineExpose({ openDialog})
const resetForm = ()=>{
  state.ruleForm = {
    dictCode: 0,
    dictLabel: '',
    dictValue: '',
    dictSort: 0,
    isDefault:0,
    status: 1,
    remark: '',
    dictType:prop.dictType
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
      if(state.ruleForm.dictCode!==0){
        //修改
        editData(state.ruleForm).then(()=>{
          ElMessage.success('字典数据修改成功');
          closeDialog(); // 关闭弹窗
          emit('dataList')
        })
      }else{
        //添加
        addData(state.ruleForm).then(()=>{
          ElMessage.success('字典数据添加成功');
          closeDialog(); // 关闭弹窗
          emit('dataList')
        })
      }
    }
  });
};
</script>
