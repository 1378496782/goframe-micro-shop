<template>
	<el-card shadow="never" header="生成信息配置">
		<el-form ref="genInfoFormRef" :model="info" :rules="rules" label-width="150px">
			<el-row :gutter="35">
				<el-col class="mb20" :span="12">
					<el-form-item prop="tplCategory">
						<template #label>
							<span>生成模板</span>
						</template>
						<el-select v-model="info.tplCategory">
							<el-option label="单表（增删改查）" value="crud" />
							<el-option label="树表（增删改查）" value="tree" />
						</el-select>
					</el-form-item>
				</el-col>

				<el-col class="mb20" :span="12">
					<el-form-item prop="packageName">
						<template #label>
							<span>
								生成包路径
								<el-tooltip content="生成在哪个包下，例如 internal/app/system，必须internal/app开头，不可/结尾" placement="top">
									<el-icon><ele-QuestionFilled /></el-icon>
								</el-tooltip>
							</span>
						</template>
						<el-input v-model="info.packageName" @input="setModuleName"/>
					</el-form-item>
				</el-col>

				<el-col class="mb20" :span="12">
					<el-form-item prop="businessName">
						<template #label>
							<span>
								生成业务名
								<el-tooltip content="可理解为功能英文名，例如 user" placement="top">
									<el-icon><ele-QuestionFilled /></el-icon>
								</el-tooltip>
							</span>
						</template>
						<el-input v-model="info.businessName" />
					</el-form-item>
				</el-col>

				<el-col class="mb20" :span="12">
					<el-form-item prop="functionName">
						<template #label>
							<span>
								生成功能名
								<el-tooltip content="用作类描述，例如 用户" placement="top">
									<el-icon><ele-QuestionFilled /></el-icon>
								</el-tooltip>
							</span>
						</template>
						<el-input v-model="info.functionName" />
					</el-form-item>
				</el-col>
        <el-col class="mb20" :span="12" >
          <el-form-item label="上级菜单">
            <el-cascader
                :options="menuData"
                :props="{ label: 'title',value: 'id',checkStrictly: true,emitPath: false }"
                placeholder="请选择上级菜单,默认为根目录"
                clearable
                class="w100"
                v-model="info.menuPid"
            >
              <template #default="{ node, data }">
                <span>{{ data.title }}</span>
                <span v-if="!node.isLeaf"> ({{ data.children.length }}) </span>
              </template>
            </el-cascader>
          </el-form-item>
        </el-col>
			</el-row>

			<div v-show="info.tplCategory == 'tree'">
				<h4 class="form-header">其他信息</h4>
				<el-divider />
				<el-row :gutter="35">
					<el-col class="mb20" :span="8">
						<el-form-item>
							<template #label>
								<span>
									树编码字段
									<el-tooltip content="树显示的编码字段名， 如：dept_id" placement="top">
										<el-icon><ele-QuestionFilled /></el-icon>
									</el-tooltip>
								</span>
							</template>
							<el-select v-model="info.treeCode" placeholder="请选择">
								<el-option
									v-for="column in info.columns"
									:key="column.htmlField"
									:label="column.columnName + '：' + column.columnComment"
									:value="column.htmlField"
								></el-option>
							</el-select>
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="8">
						<el-form-item>
							<template #label>
								<span>
									树父编码字段
									<el-tooltip content="树显示的父编码字段名， 如：parent_Id" placement="top">
										<el-icon><ele-QuestionFilled /></el-icon>
									</el-tooltip>
								</span>
							</template>
							<el-select v-model="info.treeParentCode" placeholder="请选择">
								<el-option
									v-for="column in info.columns"
									:key="column.htmlField"
									:label="column.columnName + '：' + column.columnComment"
									:value="column.htmlField"
								></el-option>
							</el-select>
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="8">
						<el-form-item>
							<template #label>
								<span>
									树名称字段
									<el-tooltip content="树节点的显示名称字段名， 如：dept_name" placement="top">
										<el-icon><ele-QuestionFilled /></el-icon>
									</el-tooltip>
								</span>
							</template>
							<el-select v-model="info.treeName" placeholder="请选择">
								<el-option
									v-for="column in info.columns"
									:key="column.htmlField"
									:label="column.columnName + '：' + column.columnComment"
									:value="column.htmlField"
								></el-option>
							</el-select>
						</el-form-item>
					</el-col>
          <el-col class="mb20" :span="8">
            <el-form-item>
              <template #label>
								<span>
									使用虚拟化树表
									<el-tooltip content="如果该表数据量较大，请选择使用虚拟表，否则请勿选择，因为虚拟表会增加代码复杂度，后续如果修改所生成的页面会增加工作量" placement="top">
										<el-icon><ele-QuestionFilled /></el-icon>
									</el-tooltip>
								</span>
              </template>
              <el-checkbox v-model="info.useVirtual" />
            </el-form-item>
          </el-col>
				</el-row>
			</div>
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
</template>

<script setup lang="ts">
import {defineComponent, getCurrentInstance, inject, onMounted, reactive, ref, unref} from 'vue';
import type { FormInstance, FormRules } from 'element-plus';
import { TableDataInfo } from '/@/views/system/tools/gen/component/model';
import {ElMessage} from "element-plus";
import {Ref, UnwrapRef} from "@vue/reactivity";
import {getMenuParams} from "/@/api/system/menu";
defineOptions({ name: "genInfoForm"})
const emit = defineEmits(['goNext','close']);
const genInfoFormRef = ref<FormInstance>();
const info = inject<Ref<UnwrapRef<TableDataInfo>>>('tableData') as Ref<UnwrapRef<TableDataInfo>>;
const rules = reactive<FormRules>({
  tplCategory: [{ required: true, message: '请选择生成模板', trigger: 'blur' }],
  packageName: [{ required: true, message: '请输入生成包路径', trigger: 'blur' }],
  moduleName: [{ required: true, message: '请输入生成模块名', trigger: 'blur' }],
  businessName: [{ required: true, message: '请输入生成业务名', trigger: 'blur' }],
  functionName: [{ required: true, message: '请输入生成功能名', trigger: 'blur' }],
});
const { proxy } = <any>getCurrentInstance();
const menuData = ref<any[]>([]);
onMounted(()=>{
  getMenuParams().then((res:any)=>{
    const menu = { id: 0, title: '主类目', children: [] };
    menu.children = proxy.handleTree(res.data.menus, "id","pid");
    menuData.value=new Array(menu) as any;
  });
})
const nextTip = ()=>{
  //表单验证
  let tag = false;
  const formWrap = unref(genInfoFormRef) as any;
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
}
const close = ()=>{
  emit('close')
}
const getLastSubstring = (str:string):string => {
  let lastIndex = str.lastIndexOf('/');
  if (lastIndex !== -1) {
    return str.substring(lastIndex + 1);
  } else {
    return str;
  }
}
const setModuleName = (value:string)=>{
  info.value.moduleName = getLastSubstring(value)
}
defineExpose({nextTip})
</script>

<style scoped>
.form-header {
}
</style>
