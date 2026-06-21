<template>
	<div class="system-dic-container">
		<el-row :gutter="10" style="width: 100%;">
			<el-col :span="5">
				<el-card shadow="hover">
					<el-aside style="width: 100%;">
						<el-scrollbar>
							<div class="act-btn">
								<el-button size="default" type="success" class="ml10 btm8" @click="onOpenAddDic">
									<el-icon>
										<ele-FolderAdd />
									</el-icon>
									新增
								</el-button>
								<el-button size="default" type="primary" class="ml10 btm8" @click="onOpenEditDic" :disabled="selectedNode===null">
									<el-icon>
										<ele-Edit />
									</el-icon>
									编辑
								</el-button>
								<el-button size="default" type="danger" class="ml10 btm8" @click="onRowDel(null)" :disabled="checkedNodes.length===0">
									<el-icon>
										<ele-Delete />
									</el-icon>
									删除
								</el-button>
								<el-button size="default" type="warning" class="ml10 btm8" @click="expandedOn">
									<el-icon>
										<ele-DCaret />
									</el-icon>
									{{expanded?'收起':'展开'}}
								</el-button>
							</div>
							<el-divider />
							<el-input :prefix-icon="search" v-model="filterText" placeholder="请输入字典名称" clearable style="margin-bottom: 8px;"/>
							<el-tree
								ref="treeRef"
								class="filter-tree"
								:data="dictTypeData"
								:props="treeProps"
								:filter-node-method="filterNode"
								@node-click="handleNodeClick"
								show-checkbox
								@check-change="handleCheckChange"
								node-key="dictId"
                :check-on-click-leaf="false"
							/>
						</el-scrollbar>
					</el-aside>
				</el-card>
			</el-col>
			<el-col :span="19">
				<apiV1SystemDictDataList ref="dataViewRef"></apiV1SystemDictDataList>
			</el-col>
		</el-row>
		<EditDic ref="editDicRef" @typeList="getTypeData"/>
	</div>
</template>

<script setup lang="ts">
import {   onMounted, ref, getCurrentInstance, watch } from 'vue';
import { ElMessageBox, ElMessage,  ElTree } from 'element-plus';
import EditDic from '/@/views/system/dict/component/editDic.vue';
import { deleteType,  optionselect } from '/@/api/system/dict/type';
import { Search } from '@element-plus/icons-vue'
import apiV1SystemDictDataList from '/@/views/system/dict/dataList.vue'
// 定义接口来定义对象的类型
interface TableDataRow {
	dictId:number;
	pid:number;
	dictName: string;
	dictType: string;
	status: number;
	remark:string;
	createdAt:string;
}
defineOptions({ name: "apiV1SystemDictTypeList"})
const search = Search
const filterText = ref('');
const treeRef = ref()
const {proxy} = getCurrentInstance() as any;
const editDicRef = ref();
const dictTypeData = ref([])
const selectedNode = ref<TableDataRow|null>(null)
const checkedNodes = ref<TableDataRow[]>([])
const expanded = ref<boolean>(false)
const dataViewRef = ref()
const treeProps = ref({
	value: 'dictId',
	label: 'dictName',
	children: 'children',
	checkStrictly:true,
	emitPath: false
})
const getTypeData = () => {
	optionselect(true).then((res:any)=>{
		const data = res.data.dictType??[]
		dictTypeData.value = proxy.handleTree(data, 'dictId', 'pid', 'children', true)
	})
}
// 初始化表格数据
const initTableData = () => {
	getTypeData()
};
// 打开新增字典弹窗
const onOpenAddDic = () => {
	editDicRef.value.openDialog();
};
// 打开修改字典弹窗
const onOpenEditDic = () => {
	const row = selectedNode.value;
	editDicRef.value.openDialog(row);
};
// 删除字典
const onRowDel = (row: TableDataRow|null) => {
	let msg = '你确定要删除所选数据？';
	let ids:number[] = [] ;
	if(row){
		msg = `此操作将永久删除用户：“${row.dictName}”，是否继续?`
		ids = [row.dictId]
	}else{
		checkedNodes.value.forEach(item=>{
			ids.push(item.dictId);
		})
	}
	if(ids.length===0){
		ElMessage.error('请选择要删除的数据。');
		return
	}
	ElMessageBox.confirm(msg, '提示', {
		confirmButtonText: '确认',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			deleteType(ids).then(()=>{
				ElMessage.success('删除成功');
				getTypeData()
				checkedNodes.value = []
			})
		})
		.catch(() => {});
};
// 页面加载时
onMounted(() => {
	initTableData();
});
const filterNode = (value: string, data:any) => {
	if (!value) return true;
	return data.dictName.includes(value)
};
// 节点单击事件
const handleNodeClick = (data:TableDataRow) => {
	selectedNode.value = data
	dataViewRef.value.initTableData(data.dictId,data.dictType)
};

const handleCheckChange = (
	data: TableDataRow,
	checked: boolean,
) => {
	if(checked){
		checkedNodes.value.push(data)
	}else{
		checkedNodes.value = checkedNodes.value.filter((item:TableDataRow) => {
			if(item.dictId!==data.dictId){
				return true
			}
			return false
		})
	}
}
const expandedOn = ()=>{
	expanded.value = !expanded.value
	let treeList = dictTypeData.value as Array<TableDataRow>;
	for (let i = 0; i < treeList.length; i++) {
		treeRef.value.store.nodesMap[treeList[i].dictId].expanded = expanded.value;
	}
}
watch(filterText, (val) => {
	treeRef.value!.filter(val)
});

</script>
<style scoped lang="scss">
.btm8{margin-bottom: 8px;}
.filter-tree{
	height: calc(100vh - 290px);
	overflow-y: auto;
}
</style>
