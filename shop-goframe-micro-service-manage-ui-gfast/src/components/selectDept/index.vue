<template>
  <el-cascader
      filterable
      :options="deptData"
      :props="{ checkStrictly: true,emitPath: false, value: 'deptId', label: 'deptName',multiple: multiple }"
      placeholder="请选择部门"
      clearable
      class="w100"
      v-model="deptIds"
  >
    <template #default="{ node, data }">
      <span>{{ data.deptName }}</span>
      <span v-if="!node.isLeaf"> ({{ data.children.length }}) </span>
    </template>
  </el-cascader>
</template>

<script setup lang="ts">
import {computed, onMounted, ref} from "vue";
import {getDeptTree} from "/@/api/system/user";

defineOptions({ name: "selectDept"})
const props = defineProps({
  multiple:{
    type:Boolean,
    default:false
  },
  modelValue:{
    type:[Array,Number],
    default:()=>[]
  }
})
const deptIds = computed({
  get: () => {
    if(props.multiple){
      return props.modelValue as number[]
    }else{
      return props.modelValue instanceof Array ? props.modelValue[0] : props.modelValue
    }
  },
  set: (value) => {
    emit('update:modelValue', value)
  }
})
const emit = defineEmits(['update:modelValue'])
const deptData = ref([])
onMounted(()=>{
  getDeptTree().then((res)=>{
    deptData.value = res.data.deps
  })
})
</script>

<style scoped lang="scss">

</style>
