<template>
  <div>
    <a-list
      class="a-list-action"
      :bordered="true"
      :data="data.items"
      :pagination-props="{ defaultPageSize: Blog_Query.page_size, total: data.total }"
      @page-change="PageNumChanged"
      @page-size-change="PageSizeChanged"
    >
      <template #item="{ item }">
        <a-list-item
          class="list-demo-item"
          action-layout="vertical"
          @click="$router.push({ name: 'frontend_blog_detail', query: { id: item.id } })"
        >
          <template #actions>
            <span><icon-heart />83</span>
            <span><icon-star />{{ item.id }}</span>
            <span><icon-message />Reply</span>
          </template>
          <template #extra>
            <div className="image-area">
              <img
                alt="arco-design"
                width="200"
                height="120"
                src="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/1f61854a849a076318ed527c8fca1bbf.png~tplv-uwbnlip3yd-webp.webp"
              />
            </div>
          </template>
          <a-list-item-meta :title="item.title" :description="item.summary"> </a-list-item-meta>
        </a-list-item>
      </template>
    </a-list>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'
import { FrontendBlogQueryRequest } from '@/api/blog.js'

const Blog_Query = reactive({
  // keywords: '',
  // category: '',
  page_size: 5,
  page_num: 1,
})

// blog请求逻辑
// data.items就是请求后端返回的所有列表，然后交给a-list去渲染
const data = ref({
  items: [],
  total: 0,
})

const BlogQueryLoading = ref(false)
const FrontendBlogQuery = async () => {
  try {
    BlogQueryLoading.value = true
    data.value = await FrontendBlogQueryRequest(Blog_Query)
  } finally {
    BlogQueryLoading.value = false
  }
}

const PageNumChanged = (current) => {
  Blog_Query.page_num = current
  FrontendBlogQuery(Blog_Query)
}
const PageSizeChanged = (pageSize) => {
  Blog_Query.page_size = pageSize
  console.log(pageSize)
  // 防止pagesize的数量超过查询数据总rows，pagenum定位到空内容
  Blog_Query.page_num = 1
  FrontendBlogQuery(Blog_Query)
}

onMounted(() => {
  FrontendBlogQuery()
})
</script>

<style lang="css" scoped></style>
