<template>
  <div>
    <gf-ueditor v-if="show" editorId="demoEdit01" v-model="content"></gf-ueditor>
    <h3>同步获取编辑器内容如下：</h3>
    <div v-html="content"></div>
  </div>
</template>

<script lang="ts">
import {defineComponent,ref,onActivated,onDeactivated} from "vue";
import GfUeditor from "/@/components/ueditor/index.vue"
export default defineComponent({
  name:'demoUeditor',
  components:{GfUeditor},
  setup(){
    const show=ref(true)
    const content = ref(`<P>这是有一个测试内容</P>
    <p><img src="https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2022-06-08/timg.jpg"/></p>
    <p style="text-indent: 2em;">在软件工程中，有些主题和写程序并没有直接的关联，但它们为你提供了工具和基础设施 支持，使得软件对每个人都变得更易用。这些主题包括：
文档：通过附带的 rustdoc 生成库文档给用户。
测试：为库创建测试套件，确保库准确地实现了你想要的功能。
基准测试（benchmark）：对功能进行基准测试，保证其运行速度足够快。</p>
    <p style="text-indent: 2em;">Rust 是一门注重安全（safety）、速度（speed）和并发（concurrency）的现代系统编程语言。Rust 通过内存安全来实现以上目标，但不使用垃圾回收机制（garbage collection, GC）。
《通过例子学 Rust》（Rust By Example, RBE）内容由一系列可运行的实例组成，通过这些例子阐明了各种 Rust 的概念和基本库。想获取这些例子外的更多内容，不要忘了安装 Rust 到本地并查阅官方标准库文档。另外为了满足您的好奇心，您还可以查阅本网站的源代码。</p>`)
    const setEditContent = (data:string)=>{
      content.value = data
    }
    onActivated(()=>{
      if(!show.value){
        show.value=true
      }
    })
    onDeactivated(()=>{
      show.value=false
    })
    return {
      show,
      content,
      setEditContent
    }
  }
})
</script>
