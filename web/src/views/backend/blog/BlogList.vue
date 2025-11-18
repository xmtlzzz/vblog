<template>
  <div class="page_list_div">
    <div>
      <a-page-header title="VBLOG文章列表" @back="handlerBack" style="width: 100%">
        <template #breadcrumb>
          <a-breadcrumb>
            <a-breadcrumb-item><icon-home /></a-breadcrumb-item>
            <a-breadcrumb-item>BlogList</a-breadcrumb-item>
          </a-breadcrumb>
        </template>
        <!--      提交新建文章按钮-->
        <!--      跳转到编辑页面-->
        <template #extra>
          <a-button
            class="Submit_Blog"
            type="primary"
            @click="$router.push({ name: 'backend_blog_edit' })"
            style="margin-top: 0px"
          >
            <template #icon>
              <icon-plus />
            </template>
            新建文章
          </a-button>
        </template>
      </a-page-header>
    </div>
    <div class="input_value">
      <div>
        <a-select
          :style="{ width: '320px' }"
          placeholder="Please select ..."
          v-model="Blog_Query.category"
          @change="() => QueryBlogData(Blog_Query)"
        >
          <a-option>软件开发</a-option>
          <a-option>系统运维</a-option>
          <a-option>网络工程</a-option>
        </a-select>
      </div>
      <div style="padding-right: 10px">
        <a-input-search
          :style="{ width: '400px' }"
          placeholder="输入搜索内容 按下回车键进行搜索"
          search-button
          v-model="Blog_Query.keywords"
          @change="() => QueryBlogData(Blog_Query)"
        >
          <template #button-icon>
            <icon-search />
          </template>
          <template #button-default> Search </template>
        </a-input-search>
      </div>
    </div>
    <div style="padding-left: 8px; margin-top: 13px">
      <!--      table绑定的data相当于是传递一个响应式变量，然后根据这个对象的属性去渲染-->
      <!--      不加.items会导致内容无限加载-->
      <a-table
        :data="data.items"
        :loading="BlogQueryLoading"
        style="width: 99%"
        :pagination="false"
      >
        <template #columns>
          <a-table-column title="编号" data-index="id"></a-table-column>
          <a-table-column title="标题" data-index="title"></a-table-column>
          <a-table-column title="文章状态" data-index="stages"></a-table-column>
          <a-table-column title="最后更新于" data-index="updateAt">
            <!--            自定义时间格式化，针对后端发来的内容进行格式化-->
            <!--            默认情况下js会将后端传来的时间内容更新为utc时间-->
            <template #cell="{ record }">
              <!--              {{ formatTime(record.updateAt) }}-->
              {{ dayjs(record.updateAt).format('YYYY-MM-DD HH:mm:ss') }}
            </template>
          </a-table-column>
          <a-table-column title="分类" data-index="category"></a-table-column>
          <a-table-column title="操作" align="center">
            <template #cell="{ record }">
              <a-space>
                <!--                primary表示和主题色相同，danger表示红色状态样式-->
                <!--                跳转到编辑页面，query属性表示传递的url参数类似gin中的:id-->
                <a-button
                  type="primary"
                  @click="$router.push({ name: 'backend_blog_edit', query: { id: record.id } })"
                >
                  <template #icon>
                    <icon-edit />
                  </template>
                  编辑
                </a-button>
                <a-button type="primary" @click="handlerPublishBlog(record.id)">
                  <template #icon>
                    <icon-launch />
                  </template>
                  发布
                </a-button>
                <!--                气泡提示框实现二次确认，点击ok事件才去执行删除而不是让button去实现-->
                <a-popconfirm
                  :content="`确定要删除${record.title}吗?`"
                  @ok="DeleteBlogData(record.id)"
                  type="error"
                >
                  <!--                这里需要注意，因为是批量渲染的button按钮，但是删除加载是针对选择的那行内容进行的-->
                  <!--                所以需要将BlogDeleteLoading和某一行的id进行比较，默认看到为false，然后点击函数运行之后为true执行加载-->
                  <a-button status="danger" :loading="BlogDeleteLoading === record.id">
                    <template #icon>
                      <icon-delete />
                    </template>
                    删除
                  </a-button>
                </a-popconfirm>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>
      <div>
        <!--        调用后端返回的total字段总数，表示查询的总rows-->
        <a-pagination
          :total="data.total"
          :current="Blog_Query.page_num"
          :page-size="Blog_Query.page_size"
          @change="PageNumChanged"
          @page-size-change="PageSizeChanged"
          :page-size-options="[2, 10, 20, 30, 50]"
          show-total
          show-jumper
          show-page-size
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'
import { BlogQueryRequest, BlogDeleteRequest,BlogPublishRequest } from '@/api/blog'
import dayjs from 'dayjs'

const handlerBack = () => {
  console.log('test back')
}

const Blog_Query = reactive({
  keywords: '',
  category: '',
  page_size: 10,
  page_num: 1,
})

// blog请求逻辑
const data = ref({})
const BlogQueryLoading = ref(false)
const QueryBlogData = async (value) => {
  try {
    BlogQueryLoading.value = true
    data.value = await BlogQueryRequest(value)
  } finally {
    BlogQueryLoading.value = false
  }
}

// blog删除逻辑，不能直接设置为true或false，会导致所有行都统一loading状态
const BlogDeleteLoading = ref(0)
const DeleteBlogData = async (id) => {
  try {
    // 配合template中的逻辑实现
    BlogDeleteLoading.value = id
    await BlogDeleteRequest(id)
    QueryBlogData(Blog_Query)
  } finally {
    BlogDeleteLoading.value = 0
  }
}

onMounted(() => {
  QueryBlogData(Blog_Query)
})

// 分页实现逻辑
const PageNumChanged = (current) => {
  Blog_Query.page_num = current
  QueryBlogData(Blog_Query)
}
const PageSizeChanged = (pageSize) => {
  Blog_Query.page_size = pageSize
  // 防止pagesize的数量超过查询数据总rows，pagenum定位到空内容
  Blog_Query.page_num = 1
  QueryBlogData(Blog_Query)
}

// 对接后端blog文章发布功能
const publishLoading = ref(false)
const handlerPublishBlog = async (id) =>{
  try {
    publishLoading.value = true
    await BlogPublishRequest(id,{"stages":"已发布"})
    QueryBlogData(Blog_Query)
  }finally{
    publishLoading.value = false
  }
}
</script>

<style lang="css" scoped>
.page_list_div {
  padding-left: 10px;
}

:deep(.arco-page-header-wrapper) {
  padding-left: 8px;
  /* 让pageheader中的新建文章按钮位置对齐 */
  padding-right: 0px;
}

.input_value {
  display: flex;
  padding-left: 8px;
  justify-content: space-between;
}

/* 覆盖table原有的样式，主要是为了无论标题内容多长，每行元素的内容的width都相同 */
/* !important是无论什么情况下都会优先生效。权重最高 */
:deep(.arco-table-th),
:deep(.arco-table-td) {
  width: 10% !important;
}

.Submit_Blog {
  width: 100px;
  margin-top: 50px;
  margin-bottom: 10px;
  margin-left: auto;
  margin-right: 10px;
}
</style>
