<template>
  <div>
    <el-tabs v-model="tabsActive" @tab-change="handleTabChange">
      <el-tab-pane :name="1">
        <template v-slot:label>
          <span style="font-size: 18px;font-weight: 500;">通知</span>
          <el-badge type="warning" :value="count.notify" style="margin-left: 10px"/>
        </template>

        <div class="layout-navbars-breadcrumb-user-news">
          <div class="content-box">
            <template v-if="noticeList.length > 0">
              <div class="content-box-item" v-for="(v, k) in noticeList" :key="k">
                <div @click="handleRead(v)">
                  <div>{{ v.title }}</div>
                  <div class="content-box-msg" v-html="v.content">
                  </div>
                  <div class="msg-box-attr">
                    <span class="content-box-time">{{ v.createdAt }}</span>
                    <span>
                    <el-tag type="success" effect="plain" v-if="v.isRead==true">已读</el-tag>
                    <el-tag type="danger" effect="plain" v-else>未读</el-tag>
                    </span>
                  </div>
                </div>
              </div>
            </template>
            <el-empty :description="$t('message.user.newDesc')" v-else></el-empty>
          </div>
          <div class="foot-box" v-if="noticeList.length > 0">
            <el-button @click="hendleShowMore()" size="small">查看更多</el-button>
          </div>
        </div>

      </el-tab-pane>


      <el-tab-pane :name="2">
        <template v-slot:label>
          <span style="font-size: 18px;font-weight: 500;"> 私信</span>
          <el-badge type="danger" :value="count.notice" style="margin-left: 10px"/>
        </template>

        <div class="layout-navbars-breadcrumb-user-news">
          <div class="content-box">
            <template v-if="noticeList.length > 0">
              <div class="content-box-item" v-for="(v, k) in noticeList" :key="k">
                <div @click="handleRead(v)">
                  <div>{{ v.title }}</div>
                  <div class="content-box-msg" v-html="v.content">

                  </div>
                  <div class="msg-box-attr">
                    <span class="content-box-time">{{ v.createdAt }}</span>
                    <span>
                    <el-tag type="success" effect="plain" v-if="v.isRead==true">已读</el-tag>
                    <el-tag type="danger" effect="plain" v-else>未读</el-tag>
                    </span>
                  </div>
                </div>
              </div>
            </template>
            <el-empty :description="$t('message.user.newDesc')" v-else></el-empty>
          </div>
          <div class="foot-box" v-if="noticeList.length > 0">
            <el-button @click="hendleShowMore()" size="small">查看更多</el-button>
          </div>
        </div>
      </el-tab-pane>
    </el-tabs>

    <!--查看更多弹出框-->
    <el-dialog v-model="isShowDialog" width="1200px" :close-on-click-modal="false" :destroy-on-close="true">
      <el-tabs  v-model="activeInfo" type="border-card">
        <el-tab-pane :name="1">
          <template v-slot:label>
            <span style="font-weight: 600;font-size: 18px;color: #6c6c6c ">通知</span>
          </template>
        </el-tab-pane>
        <el-tab-pane label="Config" :name="2">
          <template v-slot:label>
            <span style="font-weight: 600;font-size: 18px;color: #6c6c6c ">私信</span>
          </template>
        </el-tab-pane>
      </el-tabs>

    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import {reactive, toRefs, onMounted} from 'vue';
import {
  listShowNotice,
   unReadCount,
} from "/@/api/system/notice/sysNotice";
import {ElMessage, ElMessageBox} from "element-plus";
import {readNotice} from "/@/api/system/notice/sysNoticeRead";
import {SysNoticeInfoData} from "/@/views/system/sysNotice/list/component/model";
import {useRouter} from "vue-router";
defineOptions({ name: "layoutBreadcrumbUserNews"})

onMounted(()=>{
  getUnReadCount()
  getData(1)
})

const router = useRouter();
const state = reactive({
  type1Num: 0,
  type2Num: 0,
  // notifyList: [],
  noticeList: [],
  count: {
    notify: 0,
    notice: 0
  },
  tabsActive: 1,
  activeInfo: 1,
  isShowDialog: false,
  barName: "通知"
});
const { tabsActive, activeInfo, isShowDialog, count, noticeList} = toRefs(state)
/** 改变tab*/
const handleTabChange = (tabName: number) => {
  if (tabName === 1) {
    state.barName = "通知"
  } else {
    state.barName = "公告"
  }
  getData(tabName)
};
//获取维度信息
const getUnReadCount = () => {
  unReadCount().then((res: any) => {
    if (res.data != null) {
      state.count.notice = res.data.noticeCount || 0
      state.count.notify = res.data.notifyCount || 0
    }
  })
}
//获取数据
const getData = (barName: number | undefined) => {
  /*    let notifyParam = {
        pageNum: 1,
        pageSize: 5,
        type: barName,
      }
      listSysNotice(notifyParam).then((res: any) => {
        console.log("listSysNotice",res)
        state.notifyList = res.data.list || []
      })*/
  let noticeParam = {
    pageNum: 1,
    pageSize: 10,
    type: barName,
    isTrim:true
  }
  listShowNotice(noticeParam).then((res: any) => {
    state.noticeList = res.data.list || []
  })
};
const readAllItem = () => {

}
// 全部已读点击
const onAllReadClick = () => {
  state.noticeList = [];
};
// 前往通知中心点击
const onGoToGiteeClick = () => {
  /*window.open('https://gitee.com/tiger1103/gfast');*/
  state.isShowDialog = true
};
const hendleClear = (type: string) => {
  ElMessageBox.confirm("是否清除全部" + type + "!", '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {/*
        delSysNotice(id).then(() => {
          ElMessage.success('删除成功');
        })*/
        let typeIndex = (type == "通知") ? 1 : 2
        let query = {
          type: typeIndex,
        }
        /*  clearNews(query).then((res: any) => {
            console.log(res)
            ElMessage.success('清空成功');
          })*/
      })
      .catch(() => {
      });

};
const hendleAllread = (type: string) => {

  ElMessageBox.confirm("是否将全部" + type + "标记为已读!", '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {/*
        delSysNotice(id).then(() => {
          ElMessage.success('删除成功');
        })*/
      })
      .catch(() => {
      });
};
const hendleShowMore = () => {
 // console.log(emit)
  router.push("/system/sysNotice/show")
};


const handleRead = (item: SysNoticeInfoData) => {
  // console.log("handleRead", item)
  let query = {
    noticeId: item.id
  }
  readNotice(query).then(() => {
    //  console.log(res)
    getData(item.type)
    ElMessage.success("已读");
  })
  getUnReadCount()
}
</script>

<style scoped lang="scss">
.layout-navbars-breadcrumb-user-news {
  height: calc(100vh - 150px);
  overflow-y: auto;
  .content-box {
    font-size: 13px;

    .content-box-item {
      padding-top: 12px;
      border-bottom: 1px solid var(--el-border-color-lighter);

      &:last-of-type {
        padding-bottom: 12px;
      }

      .content-box-msg {
        color: var(--el-text-color-secondary);
        margin-top: 5px;


      }

      .content-box-time {
        color: var(--el-text-color-secondary);
      }
    }
  }

  .foot-box {
    height: 35px;
    color: var(--el-color-primary);
    font-size: 13px;
    cursor: pointer;
    opacity: 0.8;
    display: flex;
    align-items: center;
    justify-content: center;
    /* border-top: 1px solid var(--el-border-color-lighter);*/
    margin-top: 10px;

    &:hover {
      opacity: 1;
    }
  }

  :deep(.el-empty__description p) {
    font-size: 13px;
  }
}

:deep(.el-tabs__active-bar) {
  padding-left: 100px;
  padding-right: 100px;
}

:deep(.el-tabs__item) {
  width: 200px;
}
.msg-box-attr{
  padding-bottom: 5px;
  flex-direction: row;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
