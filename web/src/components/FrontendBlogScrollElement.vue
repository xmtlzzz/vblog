<template>
  <div>
    <MdCatalog v-if="showCatalog" editorId="blog-preview" :scrollElement="scrollElement" />
  </div>
</template>

<script setup>
import { MdCatalog } from 'md-editor-v3'
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

const props = defineProps({
  scrollElement: {
    type: Object,
    default: null,
  },
})

const scrollElement = ref(props.scrollElement || document.documentElement)
const route = useRoute()

// 监听路由变化，当不在详情页时隐藏目录
const showCatalog = ref(false)
watch(
  () => route.name,
  (newRouteName) => {
    // 只在详情页显示目录
    showCatalog.value = newRouteName === 'frontend_blog_detail'
  },
  { immediate: true },
)
</script>

<style lang="scss" scoped></style>
