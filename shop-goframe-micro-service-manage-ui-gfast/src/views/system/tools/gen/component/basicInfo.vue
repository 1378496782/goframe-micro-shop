<template>
	<div class="form-adapt-container">
		<el-card shadow="never" header="基本信息配置">
			<el-form ref="basicInfoFormRef" :model="info" :rules="rules" label-width="150px" size="default" class="mt35 mb35">
				<el-row :gutter="35">
					<el-col class="mb20" :span="12">
						<el-form-item label="数据库" prop="group">
							<el-input placeholder="请输入数据库Group" v-model="info.group" />
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="12">
						<el-form-item label="表名称" prop="tableName">
							<el-input placeholder="请输入仓库名称" v-model="info.tableName" />
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="12">
						<el-form-item label="表描述" prop="tableComment">
							<el-input placeholder="请输入" v-model="info.tableComment" />
						</el-form-item>
					</el-col>

					<el-col class="mb20" :span="12">
						<el-form-item label="实体类名称" prop="className">
							<el-input placeholder="请输入" v-model="info.className" />
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="12">
						<el-form-item label="作者" prop="functionAuthor">
							<el-input placeholder="请输入" v-model="info.functionAuthor" />
						</el-form-item>
					</el-col>

					<el-col class="mb20" :span="12">
						<el-form-item label="排序字段" prop="sortColumn">
							<el-select v-model="info.sortColumn" placeholder="请选择">
								<el-option v-for="(column, index) in info.columns" :key="index" :label="column.columnComment" :value="column.columnName"> </el-option>
							</el-select>
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="12">
						<el-form-item label="排序方式" prop="sortType">
							<el-select placeholder="请选择" v-model="info.sortType">
								<el-option label="顺序" value="asc" />
								<el-option label="倒序" value="desc" />
							</el-select>
						</el-form-item>
					</el-col>

					<el-col class="mb20" :span="24">
						<el-form-item label="备注" prop="remark">
							<el-input type="textarea" :rows="3" v-model="info.remark"></el-input>
						</el-form-item>
					</el-col>
          <el-col class="mb20" :span="4">
            <el-form-item label="是否覆盖原有文件" prop="overwrite">
              <el-checkbox v-model="info.overwrite" @change="overwriteChange"/>
            </el-form-item>
          </el-col>
					<el-col class="mb20" :span="4">
						<el-form-item prop="useSnowId">
              <template #label>
							<span>
								主键使用雪花ID
								<el-tooltip content="选择此项，数据表中主键不要使用自增" placement="top">
									<el-icon><ele-QuestionFilled /></el-icon>
								</el-tooltip>
							</span>
              </template>
							<el-checkbox v-model="info.useSnowId" />
						</el-form-item>
					</el-col>
          <el-col class="mb20" :span="4">
            <el-form-item label="导出excel" prop="overwrite">
              <el-checkbox v-model="info.excelPort" />
            </el-form-item>
          </el-col>
          <el-col class="mb20" :span="4">
            <el-form-item label="导入excel" prop="overwrite">
              <el-checkbox v-model="info.excelImp" />
            </el-form-item>
          </el-col>
					<el-col class="mb20" :span="4">
						<el-form-item label="是否有显示详情功能" prop="showDetail">
							<el-checkbox v-model="info.showDetail" />
						</el-form-item>
					</el-col>
          <el-col class="mb20" :span="24" style="border:solid 1px #e2e3e3;width: 90%;padding: 12px;" v-show="info.overwrite">
            <el-row :gutter="18">
              <el-col class="mb20" :span="3" v-for="(item,ind) in overwriteOptions" :key="ind">
                <el-form-item :label="item.name" :prop="item.key">
                  <el-checkbox v-model="info.overwriteInfo[ind]['value']" />
                </el-form-item>
              </el-col>
            </el-row>
          </el-col>
				</el-row>
			</el-form>
			<el-row>
				<el-col :span="10"></el-col>
				<el-col :span="4" class="base-footer">
					<span>
						<el-button size="default" @click="close">关 闭</el-button>
						<el-button type="primary" @click="nextTip" size="default">下一步</el-button>
					</span>
				</el-col>
				<el-col :span="10"></el-col>
			</el-row>
		</el-card>
	</div>
</template>

<script setup lang="ts">
import { inject, reactive, ref, unref } from 'vue';
import type { FormInstance, FormRules } from 'element-plus';
import type { TableDataInfo } from '/@/views/system/tools/gen/component/model';
import { ElMessage } from 'element-plus';
defineOptions({ name: "toolsBasicInfoForm"})
const emit = defineEmits(['goNext','close']);
const basicInfoFormRef = ref<FormInstance>();
const overwriteOptions = ref([
  {key:"api",name:"覆盖api"},
  {key:"controller",name:"覆盖controller"},
  {key:"dao",name:"覆盖dao"},
  {key:"dao_internal",name:"覆盖dao_internal"},
  {key:"logic",name:"覆盖logic"},
  {key:"model",name:"覆盖model"},
  {key:"model_do",name:"覆盖model_do"},
  {key:"model_entity",name:"覆盖model_entity"},
  {key:"router",name:"覆盖router"},
  {key:"router_func",name:"覆盖router_func"},
  {key:"service",name:"覆盖service"},
  {key:"sql",name:"覆盖菜单sql"},
  {key:"tsApi",name:"覆盖tsApi"},
  {key:"tsModel",name:"覆盖tsModel"},
  {key:"vue",name:"覆盖vue-list"},
  {key:"vueDetail",name:"覆盖vue-detail"},
  {key:"vueEdit",name:"覆盖vue-edit"},
])
const info = inject<TableDataInfo>('tableData') as TableDataInfo;
const nextTip = (): boolean => {
  //表单验证
  let tag = false;
  const formWrap = unref(basicInfoFormRef) as any;
  if (!formWrap) return false;
  formWrap.validate((valid: boolean) => {
    if (valid) {
      emit('goNext');
      tag = true;
    } else {
      ElMessage.error('请将表单填写完整');
    }
  });
  return tag;
};
const close = ()=>{
  emit('close')
}
const overwriteChange = (value:boolean)=>{
  info.value.overwriteInfo.map((item:any)=>{
    item.value = value
  })
}
const rules = reactive<FormRules>({
  tableName: [{ required: true, message: '请输入表名称', trigger: 'blur' }],
  tableComment: [{ required: true, message: '请输入表描述', trigger: 'blur' }],
  className: [{ required: true, message: '请输入实体类名称', trigger: 'blur' }],
  functionAuthor: [{ required: true, message: '请输入作者', trigger: 'blur' }],
  sortColumn: [{ required: true, message: '请选择排序字段', trigger: 'blur' }],
  sortType: [{ required: true, message: '请选择排序类型', trigger: 'blur' }],
});
</script>

<style scoped>
.base-footer {
	text-align: center;
}
</style>
