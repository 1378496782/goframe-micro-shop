<template>
  <div class="system-sysNotice-edit">
    <!-- 添加或修改私信对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.system-sysNotice-edit .el-dialog', '.system-sysNotice-edit .el-dialog__header']">
          <span style="font-weight: 500;font-size: 18px;color: rgb(31 34 37);"> {{
              (!formData.id || formData.id == 0 ? '发送' : '修改') + title
            }} </span>
        </div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="80px">
        <el-form-item label="标题" prop="title">
          <el-input v-model="formData.title" placeholder="请输入标题"/>
        </el-form-item>
        <!--        <el-form-item label="类型" prop="type">
                  <el-select v-model="formData.type" placeholder="请选择类型">
                    <el-option label="请选择字典生成" value=""/>
                  </el-select>
                </el-form-item>-->

        <el-form-item v-if="formData.type==2" label="接收用户">
          <el-select
              style="width: 100%;"
              v-model="formData.receiver"
              multiple
              filterable
              remote
              reserve-keyword
              placeholder="可输入要指定的用户名称搜索"
              :remote-method="remoteUserMethod"
              :loading="loading"
          >
            <el-option
                v-for="item in userListOptions"
                :key="item.id"
                :label="item.userNickname"
                :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="标签" prop="tag">
          <el-select v-model="formData.tag" placeholder="请选择标签">
            <el-option
                v-for="dict in tagOptions"
                :key="dict.value"
                :label="dict.label"
                :value="dict.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="内容">
          <gf-ueditor editorId="ueSysNoticeContent" v-model="formData.content"></gf-ueditor>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <!--          <el-radio-group v-model="formData.status">
                      <el-radio>请选择字典生成</el-radio>
                    </el-radio-group>-->
          <el-switch
              v-model="formData.status"
              style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
              :active-value="1"
              :inactive-value="0"
              active-text="正常"
              inactive-text="停用"
          />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input show-word-limit maxlength="200" v-model="formData.remark" type="textarea" placeholder="请输入备注"/>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number type="num" v-model="formData.sort"/>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="onSubmit" :disabled="loading">确 定</el-button>
          <el-button @click="onCancel">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import {reactive, onMounted, toRefs, defineComponent, ref, unref, getCurrentInstance} from 'vue';
import { ElMessage} from 'element-plus';
import {
  getSysNotice,
  addSysNotice,
  updateSysNotice, getUserList,
} from "/@/api/system/notice/sysNotice";
import GfUeditor from "/@/components/ueditor/index.vue"
import {
  SysNoticeInfoData,
  SysNoticeEditState
} from "/@/views/system/sysNotice/list/component/model"
defineOptions({ name: "NoticeMessageEdit"})
const props = defineProps({
  tagOptions: {
    type: Array,
    default: () => []
  },
})
const emit = defineEmits(["sysNoticeList"]);
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<SysNoticeEditState>({
  loading: false,
  isShowDialog: false,
  title: "",
  userListOptions: [],
  formData: {
    id: undefined,
    receiver: undefined,
    title: undefined,
    type: undefined,
    tag: undefined,
    content: undefined,
    remark: undefined,
    sort: undefined,
    status: 0,
    createdBy: undefined,
    updatedBy: undefined,
    createdAt: undefined,
    updatedAt: undefined,
    deletedAt: undefined,
  },

  // 表单校验
  rules: {
    id: [
      {required: true, message: "ID不能为空", trigger: "blur"}
    ],
    title: [
      {required: true, message: "标题不能为空", trigger: "blur"}
    ],
    type: [
      {required: true, message: "类型不能为空", trigger: "blur"}
    ],
    content: [
      {required: true, message: "内容不能为空", trigger: "blur"}
    ],
    status: [
      {required: true, message: "状态不能为空", trigger: "blur"}
    ],
  }
});
const { isShowDialog,formData,loading,userListOptions,rules,title} = toRefs(state)
onMounted(() => {
  remoteUserMethod("");
});
// 打开弹窗
const openDialog = (row?: SysNoticeInfoData) => {
  resetForm();
  if (row) {
    getSysNotice(row.id!).then((res: any) => {
      const data = res.data;
      data.type = parseInt(data.type)
      if(data.type===2&&data.receiverUser){
        const userListOptions = [...state.userListOptions,...data.receiverUser]
        let uniqueSet = new Set(userListOptions.map(item => item.id));
        state.userListOptions = userListOptions.filter((value, index, self) => {
          return uniqueSet.has(value.id) && uniqueSet.delete(value.id)
        });
      }
      data.tag = '' + data.tag
      data.status = parseInt(data.status)
      state.formData = data;
    })
  }
  state.isShowDialog = true;
};
const remoteUserMethod = (query: string) => {
  //console.log("remoteMethod", query)
  state.userListOptions = []
  getUserList(query).then((res: any) => {
    /*console.log(res)*/
    // let list:object[]
    //list=res.data
    state.userListOptions = res.data.userList
  })
  /*      if (query) {
    loading.value = true
    setTimeout(() => {
      loading.value = false
      options.value = list.value.filter((item) => {
        return item.label.toLowerCase().includes(query.toLowerCase())
      })
    }, 200)
  } else {
    options.value = []
  }*/
}
//设置类型
const setType = (type: number) => {
  state.formData.type = type
  if (type == 1) {
    state.title = "通知"
  } else if (type == 2) {
    state.title = "私信"
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
// 提交
const onSubmit = () => {
  const formWrap = unref(formRef) as any;
  if (!formWrap) return;
  formWrap.validate((valid: boolean) => {
    if (valid) {
      state.loading = true;
      if (!state.formData.id || state.formData.id === 0) {
        //添加
        addSysNotice(state.formData).then(() => {
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('sysNoticeList')
        }).finally(() => {
          state.loading = false;
        })
      } else {
        //修改
        updateSysNotice(state.formData).then(() => {
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('sysNoticeList')
        }).finally(() => {
          state.loading = false;
        })
      }
    }
  });
};
const resetForm = () => {
  state.formData = {
    receiver: undefined,
    type: undefined,
    id: undefined,
    title: undefined,
    tag: undefined,
    content: undefined,
    remark: undefined,
    sort: 0,
    status: 1,
    createdBy: undefined,
    updatedBy: undefined,
    createdAt: undefined,
    updatedAt: undefined,
    deletedAt: undefined
  }
};
//富文本编辑器内容
const setContentEditContent = (data: string) => {
  state.formData.content = data
}
defineExpose({openDialog,setType})
</script>
<style scoped>
.kv-label {
  margin-bottom: 15px;
  font-size: 14px;
}

.mini-btn i.el-icon {
  margin: unset;
}

.kv-row {
  margin-bottom: 12px;
}
</style>
