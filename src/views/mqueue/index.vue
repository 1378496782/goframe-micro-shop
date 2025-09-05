<template>
  <div>
    <el-card>
      <el-alert type="error" title="消息队列测试，先生产一条消息再订阅，先订阅不存在的topic会失败" :closable="false"></el-alert>
      <el-row :gutter="20" style="margin-top: 20px;">
        <el-col :span="10">
          <div class="my-title">生产者</div>
          <el-form :label-width="80">
            <el-form-item label="topic">
              <el-input v-model="demoForm.topic" />
            </el-form-item>
            <el-form-item label="topic">
              <el-input v-model="demoForm.body" type="textarea" :rows="6" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" size="default" @click="handleProduce">生产消息</el-button>
            </el-form-item>
          </el-form>
        </el-col>
        <el-col :span="14">
          <div class="my-title">消费者</div>
          <el-form :label-width="80">
            <el-form-item label="topic">
              <el-input v-model="demoForm.topic" :disabled="wsReady" />
            </el-form-item>
            <el-form-item label="channel">
              <el-input v-model="demoForm.channel" :disabled="wsReady" />
            </el-form-item>
            <el-form-item>
              <el-button size="default" type="primary" @click="handleStartSubscribe" :disabled="wsReady">订阅消费</el-button>
              <el-button size="default" type="primary" @click="handleUnSubscribe" :disabled="!wsReady">取消订阅</el-button>
            </el-form-item>
            <el-form-item label="消费结果">
              <el-input :model-value="messageResult" type="textarea" :rows="12"></el-input>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script setup lang="ts">

import {reactive, ref} from "vue";
import {ElMessageBox} from "element-plus";
import request from "/@/utils/request";

const demoForm = reactive({
  topic: "producer_topic_test",
  body: "测试消息队列内容",
  channel: "channel1"
})
const consumerForm = reactive({
  topic:"producer_topic_test",
  channel:"channel1"
})

function handleProduce(){
  request({
    url: '/api/v1/mqueue/demo/produce',
    method: 'post',
    data: {topic:demoForm.topic,body:demoForm.body}
  }).then((res:any)=>{
    window.console.log(res)
  })
}
const webSocketURL = import.meta.env.VITE_WEBSOCKET_URL
const messageResult = ref("")
let ws:WebSocket|null = null
const wsReady = ref(false)
function handleStartSubscribe(){
  if(ws) return
  const url = webSocketURL+"api/v1/mqueue/demo/subscribe?topic="+demoForm.topic+"&channel="+demoForm.channel;
  ws = new WebSocket(url);
  try {
    // ws连接成功
    ws.onopen = function () {
      messageResult.value += "WebSocket Server [" + url +"] 连接成功！\r\n";
      wsReady.value = true
    };
    // ws连接关闭
    ws.onclose = function () {
      if (ws) {
        ws.close();
        ws = null;
        wsReady.value = false
      }
      messageResult.value += "WebSocket Server [" + url +"] 连接被服务器关闭！\r\n";
    };
    // ws连接错误
    ws.onerror = function () {
      if (ws) {
        ws.close();
        ws = null;
        wsReady.value = false
      }
      messageResult.value += "WebSocket Server [" + url +"] 连接错误！\r\n";
    };
    // ws数据返回处理
    ws.onmessage = function (result) {
      window.console.log(result)
      messageResult.value += " > " + result.data+"\r\n";
    };
  } catch (e:any) {
    ElMessageBox.alert(e.message)
  }
}

function handleUnSubscribe(){
  if(ws) {
    ws.close()
    ws = null;
    wsReady.value = false
    messageResult.value += "已手动取消订阅！\r\n";
  }
}
</script>

<style scoped lang="scss">
.my-title{
  font-size: 16px; font-weight: bold;line-height: 2;
  color: #1890ff;
  text-align: center;

}
</style>
