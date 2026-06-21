<template>
	<div class="app-container">
		<el-card shadow="hover">
			<div class="mb15">
				<el-row>
					<el-col :span="4"></el-col>
					<el-col :span="16">
						<el-steps :active="active" align-center simple>
							<el-step title="1.基本信息" @click="changeSteps(1)"
								><template #icon
									><el-icon><ele-Comment /></el-icon></template
							></el-step>
							<el-step title="2.生成信息" @click="changeSteps(2)"
								><template #icon
									><el-icon><ele-Document /></el-icon></template
							></el-step>
							<el-step title="3.字段信息" @click="changeSteps(3)"
								><template #icon
									><el-icon><ele-Expand /></el-icon></template
							></el-step>
						</el-steps>
					</el-col>
					<el-col :span="4"></el-col>
				</el-row>
			</div>
			<toolsBasicInfoForm ref="toolsBasicInfoFormRef" v-show="active === 1" @goNext="goNext" @close="close"></toolsBasicInfoForm>
			<genInfoForm ref="genInfoFormRef" v-show="active === 2" @goNext="goNext" @close="close"></genInfoForm>
      <genTableColumns ref="tableColumnsRef" v-show="active === 3"></genTableColumns>
			<div class="tools-footer" v-show="active === 3">
				<el-button size="default" @click="close">关 闭</el-button>
				<el-button type="primary" @click="onSubmit" size="default">保 存</el-button>
			</div>
		</el-card>
	</div>
</template>

<script setup lang="ts">
import {onBeforeMount, ref, provide, getCurrentInstance} from 'vue';
import { useRoute } from 'vue-router';
import toolsBasicInfoForm from '/@/views/system/tools/gen/component/basicInfo.vue';
import { TableColumns, TableDataInfo } from '/@/views/system/tools/gen/component/model';
import { ElMessage } from 'element-plus/es';
import {getGenTable, saveEdit} from '/@/api/system/tools/gen';
import genInfoForm from '/@/views/system/tools/gen/component/genInfoForm.vue';
import genTableColumns from '/@/views/system/tools/gen/component/tableColumns.vue';
defineOptions({ name: "apiV1SystemToolsGenColumnEdit"})
const { proxy } = getCurrentInstance() as any;
const route = useRoute();
const toolsBasicInfoFormRef = ref();
const genInfoFormRef = ref();
const tableColumnsRef = ref()
const active = ref(1);

const changeSteps = (args: number) => {
  if (args > active.value) {
    switch (args) {
      case 2:
        //验证基本信息表单
        if (toolsBasicInfoFormRef.value.nextTip && toolsBasicInfoFormRef.value.nextTip()) {
          active.value = args;
        }
        break;
      case 3:
        //验证生成信息表单
        if (active.value==1){
          //验证基本信息表单
          if (toolsBasicInfoFormRef.value.nextTip&&toolsBasicInfoFormRef.value.nextTip()) {
            active.value = args;
          }
        }else{
          //验证生成信息表单
          if (genInfoFormRef.value.nextTip&&genInfoFormRef.value.nextTip()) {
            active.value = args;
          }
        }
        break;
    }
  } else {
    active.value = args;
  }
};
const defaultOverwriteInfo = [
  {key:"api",value:false},
  {key:"controller",value:false},
  {key:"dao",value:false},
  {key:"dao_internal",value:false},
  {key:"logic",value:false},
  {key:"model",value:false},
  {key:"model_do",value:false},
  {key:"model_entity",value:false},
  {key:"router",value:false},
  {key:"router_func",value:false},
  {key:"service",value:false},
  {key:"sql",value:false},
  {key:"tsApi",value:false},
  {key:"tsModel",value:false},
  {key:"vue",value:false},
  {key:"vueDetail",value:false},
  {key:"vueEdit",value:false},
]
// 表详细信息
const info = ref(<TableDataInfo>{overwriteInfo:defaultOverwriteInfo});
provide('tableData', info);
const mergeArraysByKey = (arrays:Array<Array<object>>, key:string) :Array<any> => {
  const merged:any={}
  // 遍历所有数组，后面的数组会覆盖前面的数组
  arrays.forEach(array => {
    array.forEach((obj:any) => {
      // 使用key的值作为键，将对象存储在merged对象中
      merged[obj[key]] = obj;
    });
  });
  // 将merged对象转换为数组
  return Object.values(merged);
}
onBeforeMount(() => {
  const tableId = route.query?.tableId;
  if (tableId) {
    // 获取表详细信息
    getGenTable(tableId).then((res) => {
      let columnsTmp:TableColumns[]= [];
      columnsTmp = res.data.list.filter((item: TableColumns) => {
        return !['updated_at', 'deleted_at'].includes(item.columnName);
      });
      const data: TableDataInfo = res.data.info as TableDataInfo;
      data.columns = columnsTmp;
      data.overwriteInfo = data.overwriteInfo?mergeArraysByKey([defaultOverwriteInfo,data.overwriteInfo],'key'):defaultOverwriteInfo
      info.value = data;
    });
  } else {
    ElMessage.error('参数错误');
  }
});
// 表单组件验证
const formRulesValidate = (pageRef: string, sonRef: string) => {
  return new Promise((resolve,reject) => {
    proxy.$refs[pageRef].$refs[sonRef].validate((valid: boolean) => {
      if (valid) {
        resolve(valid)
      }else{
        reject('请将表单填写完整')
      }
    });
  });
};
const onSubmit = () => {
  Promise.all([
    formRulesValidate('toolsBasicInfoFormRef', 'basicInfoFormRef'),
    formRulesValidate('genInfoFormRef', 'genInfoFormRef'),
  ]).then(() => {
    saveEdit(info.value).then(()=>{
      ElMessage.success('保存成功');
      close();
    })
  }).catch(e=>{
    ElMessage.error(e.toString());
  });
};
const goNext = () => {
  if (active.value > 2) {
    active.value = 1;
  } else {
    active.value++;
  }
};
//取消返回
const close = ()=>{
  proxy.mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 1, ...route }));
}
</script>

<style scoped>
.tools-footer {
	margin: 30px auto;
	text-align: center;
}
</style>
