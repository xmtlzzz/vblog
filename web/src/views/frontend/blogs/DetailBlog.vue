<template>
  <div class="detail-container">
    <!-- 添加一个锚点偏移元素，高度等于导航栏高度 -->
    <div class="anchor-offset"></div>
<!--    modelValue绑定的内容需要是字符串否则md-editor组件内部逻辑过不去-->
    <div class="markdown-content">
      <MdPreview editorId="blog-preview" :modelValue="blog_data.content" showCodeRowNumber  :codeFoldable="false" />
    </div>
    <!-- 在这里可以添加其他内容 -->
    <div class="additional-content">
      <h3>评论区</h3>
      <p>这里是评论内容区域</p>
      <!-- 你可以在这里添加评论组件、推荐文章等内容 -->
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { MdPreview } from 'md-editor-v3'
// preview.css相比style.css少了编辑器那部分样式
import 'md-editor-v3/lib/preview.css'
import { useRoute } from 'vue-router'
import { BlogDescribeRequest } from '@/api/blog.js'

const route = useRoute()

const blog_data = ref({
  title: '',
  // // summary: '',
  content: '',
  // // category: '',
  // // tags: [],
})

const QueryLoading = ref(false)
const QueryThisBlog = async () => {
  try {
    QueryLoading.value = true
    const query_data = await BlogDescribeRequest(route.query.id)
    blog_data.value.content = query_data.content
    blog_data.value.title=query_data.title
  } finally {
    QueryLoading.value = false
  }
}

onMounted(() => {
  QueryThisBlog()
})
</script>

<style lang="scss" scoped>
.detail-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.markdown-content {
  flex: 1;
}

.additional-content {
  /* 设置一些样式以便于识别 */
  padding: 20px;
  background-color: #f5f5f5;
  margin-top: 100px; /* 调整间距，不要紧贴底部 */
  margin-bottom: 50px; /* 添加底部间距 */
}
</style>
