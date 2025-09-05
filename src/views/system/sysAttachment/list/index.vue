<template>
  <div class="system-sysAttachment-container">
    <el-card shadow="hover">
        <div class="system-sysAttachment-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>
                <el-col :span="8" class="colBlock">
                  <el-form-item label="应用ID" prop="appId">
                    <el-input
                        v-model="tableData.param.appId"
                        placeholder="请输入应用ID"
                        clearable
                        @keyup.enter.native="sysAttachmentList"
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="8" class="colBlock">
                  <el-form-item label="上传驱动" prop="drive">
                    <el-select v-model="tableData.param.drive" placeholder="请选择上传驱动" clearable style="width:200px;">
                        <el-option
                            v-for="dict in sys_upload_drive"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>
                <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="sysAttachmentList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="文件原始名" prop="name">
                    <el-input
                        v-model="tableData.param.name"
                        placeholder="请输入文件原始名"
                        clearable
                        @keyup.enter.native="sysAttachmentList"
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="上传类型" prop="kind">
                    <el-select v-model="tableData.param.kind" placeholder="请选择上传类型" clearable style="width:200px;">
                        <el-option
                            v-for="dict in sys_upload_file_type"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="扩展类型" prop="mimeType">
                    <el-input
                        v-model="tableData.param.mimeType"
                        placeholder="请输入扩展类型"
                        clearable
                        @keyup.enter.native="sysAttachmentList"
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="状态" prop="status">
                      <el-switch  v-model="tableData.param.status" class="ml-2" />
                  </el-form-item>
                </el-col>
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="创建时间" prop="createdAt">
                    <el-date-picker
                        clearable  style="width: 200px"
                        v-model="tableData.param.createdAt"
                        format="YYYY-MM-DD HH:mm:ss"
                        value-format="YYYY-MM-DD HH:mm:ss"
                        type="datetimerange"
                        range-separator="至"
                        start-placeholder="开始时间"
                        end-placeholder="结束时间"
                    ></el-date-picker>
                  </el-form-item>
                </el-col>
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="sysAttachmentList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                        {{ word }}
                        <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                        <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>
              </el-row>
            </el-form>
            <el-row :gutter="10" class="mb8">
              <el-col :span="1.5">
                <el-button
                  type="primary"
                  @click="handleAdd('file')"
                ><el-icon><ele-Upload /></el-icon>上传文件</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                    color="#626aef"
                    :dark="true"
                    type="success"
                    @click="handleAdd('image')"
                ><el-icon><ele-PictureFilled /></el-icon>上传图片</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  @click="handleAdd('bigFile')"
                ><i class="iconfont icon-shangchuan"></i>上传大文件</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/system/sysAttachment/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="文件ID" align="center" prop="id"
            min-width="80px"
             />
          <el-table-column label="应用ID" align="center" prop="appId"
            min-width="100px"
             />
          <el-table-column label="上传驱动" align="center" prop="drive" :formatter="driveFormat"
            min-width="100px"
             />
          <el-table-column label="文件原始名" align="center" prop="name" show-overflow-tooltip
            min-width="260px"
             />
          <el-table-column label="上传类型" align="center" prop="kind" :formatter="kindFormat"
            min-width="100px"
             />
          <el-table-column label="本地路径" align="center" prop="path" min-width="100px">
            <template #default="scope">
              <el-image
                  v-if="scope.row.kind=='image'"
                  style="width: 60px; height: 60px"
                  :src="proxy.getUpFileUrl(scope.row.path)"
                  :zoom-rate="1.2"
                  :max-scale="7"
                  :min-scale="0.2"
                  :preview-src-list="[proxy.getUpFileUrl(scope.row.path)]"
                  :initial-index="0"
                  preview-teleported
                  hide-on-click-modal
                  fit="fill"
              />
              <el-image
                  v-else
                  style="width: 60px; height: 60px">
                <template #error>
                  <div class="image-slot">
                    {{getExt(scope.row.name)}}
                  </div>
                </template>
              </el-image>
            </template>
          </el-table-column>
          <el-table-column label="文件大小" align="center" prop="size" :formatter="formatFileSize"
            min-width="100px"
             />
          <el-table-column label="状态" align="center" prop="status"
          min-width="150px"
          >
              <template #default="scope">
                  <el-switch  v-model="scope.row.status" class="ml-2" @change="changeStatus(scope.row)"/>
              </template>
          </el-table-column>
          <el-table-column label="上传时间" align="center" prop="updatedAt"
            min-width="150px"
            >
            <template #default="scope">
                <span>{{ proxy.parseTime(scope.row.updatedAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" class-name="small-padding" min-width="160px" fixed="right">
            <template #default="scope">
              <el-button
                type="primary"
                link
                @click="handleDownload(scope.row)"
              ><el-icon><ele-Download /></el-icon>下载</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/system/sysAttachment/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="sysAttachmentList"
        />
    </el-card>
    <ApiV1SystemSysAttachmentEdit
       ref="editRef"
       @sysAttachmentList="sysAttachmentList"
    ></ApiV1SystemSysAttachmentEdit>
  </div>
</template>
<script setup lang="ts">
import {toRefs, reactive, onMounted, ref, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listSysAttachment,
    delSysAttachment,
    changeSysAttachmentStatus,
} from "/@/api/system/sysAttachment";
import {
    SysAttachmentTableColumns,
    SysAttachmentInfoData,
    SysAttachmentTableDataState,
} from "/@/views/system/sysAttachment/list/component/model"
import ApiV1SystemSysAttachmentEdit from "/@/views/system/sysAttachment/list/component/edit.vue"
defineOptions({ name: "apiV1SystemSysAttachmentList"})
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const queryRef = ref()
const editRef = ref();
// 是否显示所有搜索选项
const showAll =  ref(false)
// 非单个禁用
const single = ref(true)
// 非多个禁用
const multiple =ref(true)
const word = computed(()=>{
    if(showAll.value === false) {
        //对文字进行处理
        return "展开搜索";
    } else {
        return "收起搜索";
    }
})
// 字典选项数据
const {
    sys_upload_drive,
    sys_upload_file_type,
} = proxy.useDict(
    'sys_upload_drive',
    'sys_upload_file_type',
)
const state = reactive<SysAttachmentTableDataState>({
    ids:[],
    tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
            pageNum: 1,
            pageSize: 10,
            appId: undefined,
            drive: undefined,
            name: undefined,
            kind: undefined,
            mimeType: undefined,
            status: undefined,
            createdAt: [],
            dateRange: []
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
    sysAttachmentList()
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    sysAttachmentList()
};
// 获取列表数据
const sysAttachmentList = ()=>{
  loading.value = true
  listSysAttachment(state.tableData.param).then((res:any)=>{
    let list = res.data.list??[];
    state.tableData.data = list;
    state.tableData.total = res.data.total;
    loading.value = false
  })
};
const toggleSearch = () => {
    showAll.value = !showAll.value;
}
// 上传驱动字典翻译
const driveFormat = (row:SysAttachmentTableColumns) => {
    return proxy.selectDictLabel(sys_upload_drive.value, row.drive);
}
// 上传类型字典翻译
const kindFormat = (row:SysAttachmentTableColumns) => {
    return proxy.selectDictLabel(sys_upload_file_type.value, row.kind);
}
const changeStatus = (row:SysAttachmentTableColumns) => {
    changeSysAttachmentStatus(row.id,row.status)
        .catch(()=>{
            setTimeout(()=>{
                row.status = !row.status
            },300)
        })
}
// 多选框选中数据
const handleSelectionChange = (selection:Array<SysAttachmentInfoData>) => {
    state.ids = selection.map(item => item.id)
    single.value = selection.length!=1
    multiple.value = !selection.length
}
const handleAdd =  (upType:string)=>{
    editRef.value.openDialog(upType)
}
const handleDownload = (row: SysAttachmentTableColumns) => {
    window.open(proxy.getUpFileUrl(row.path))
};
const handleDelete = (row: SysAttachmentTableColumns|null) => {
    let msg = '你确定要删除所选数据？';
    let id:number[] = [] ;
    if(row){
    msg = `此操作将永久删除数据，是否继续?`
    id = [row.id]
    }else{
    id = state.ids
    }
    if(id.length===0){
        ElMessage.error('请选择要删除的数据。');
        return
    }
    ElMessageBox.confirm(msg, '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
    })
        .then(() => {
            delSysAttachment(id).then(()=>{
                ElMessage.success('删除成功');
                sysAttachmentList();
            })
        })
        .catch(() => {});
}
const getExt = (fileName:string):string=>{
  // 查找最后一个点的位置
  const lastDotIndex = fileName.lastIndexOf('.');
  // 如果没有找到点，或者点在第一个字符（没有扩展名）
  if (lastDotIndex === -1 || lastDotIndex === 0) {
    return '';
  }
  // 提取并返回点后面部分的后缀（不带点）
  return fileName.slice(lastDotIndex + 1);
}
const formatFileSize = (row:SysAttachmentTableColumns):string =>{
  const bytes:number = row.size;
  if (bytes === 0) return '0 Bytes';
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
  const index = Math.floor(Math.log(bytes) / Math.log(1024)); // 计算索引
  const value = (bytes / Math.pow(1024, index)).toFixed(2); // 计算具体数值并保留两位小数
  return `${value} ${sizes[index]}`;
}
</script>
<style lang="scss" scoped>
    .colBlock {
        display: block;
    }
    .colNone {
        display: none;
    }
    .ml-2{margin: 3px;}
    .image-slot {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100%;
      height: 100%;
      background: var(--el-fill-color-light);
      color: var(--el-text-color-secondary);
      font-size: 20px;
    }
</style>
