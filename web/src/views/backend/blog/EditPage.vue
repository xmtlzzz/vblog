<template>
  <div class="edit-page-container">
    <div>
      <!--    三元表达式实现基于idEdit判断是编辑还是创建-->
      <a-page-header
        :title="isEdit ? '编辑文章' : '创建文章'"
        @back="$router.push({ name: 'backend_blog_list' })"
        style="width: 100%"
      >
        <template #breadcrumb>
          <a-breadcrumb>
            <a-breadcrumb-item><icon-home /></a-breadcrumb-item>
            <a-breadcrumb-item>{{ isEdit ? 'EditBlog' : 'CreateBlog' }}</a-breadcrumb-item>
          </a-breadcrumb>
        </template>
        <!--      提交新建文章按钮-->
        <!--      跳转到编辑页面-->
        <template #extra>
          <a-button
            type="primary"
            @click="
              $router.push({
                name: 'frontend_blog_detail',
                query: { id: current_blog_id.current_blog_id },
              })
            "
            style="margin-top: 0px; margin-right: 10px"
          >
            <template #icon>
              <icon-save />
            </template>
            跳转预览页
          </a-button>
          <a-button
            class="Submit_Blog"
            type="primary"
            @click="handlerEditBlog(BlogData)"
            style="margin-top: 0px"
            :loading="EditBlogLoading"
          >
            <template #icon>
              <icon-save />
            </template>
            保存
          </a-button>
        </template>
      </a-page-header>
    </div>
    <div class="form-container">
      <a-form :model="BlogData" layout="vertical" size="large" class="content">
        <a-form-item :rules="{ required: true, message: '请输入文章标题' }" label="文章标题">
          <a-input v-model="BlogData.title" placeholder="请输入文章标题" />
        </a-form-item>
        <a-form-item :rules="{ required: true, message: '请选择文章分类' }" label="文章分类">
          <!--        <a-input placeholder="请输入文章分类" />-->
          <a-select placeholder="请选择文章分类" v-model="BlogData.category">
            <a-option>软件开发</a-option>
            <a-option>系统运维</a-option>
            <a-option>网络工程</a-option>
          </a-select>
        </a-form-item>
        <a-form-item :field="BlogData.summary" label="文章汇总内容">
          <a-textarea
            v-model="BlogData.summary"
            placeholder="在此处编写文章的简要信息"
            allow-clear
            auto-size
          />
        </a-form-item>
        <a-form-item label="文章正文内容">
          <div class="editor_mapper">
            <!--        需要v-model绑定才可以渲染内容-->
            <MdEditor v-model="BlogData.content" @onUploadImg="handleUploadImage" />
          </div>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { useRoute } from 'vue-router'
import { BlogCreateRequest, BlogUpdateRequest, BlogDescribeRequest } from '@/api/blog.js'
import { ImageUploadRequest } from '@/api/upload.js'
import { onMounted, ref } from 'vue'
import router from '@/router/index.js'
import { Notification } from '@arco-design/web-vue'
import axios from 'axios'
import { current_blog_id } from '@/stores/system.js'

// 通过前端路由携带的query参数去判断用户是编辑文章还是新建文章
const route = useRoute()
// 三元表达式，id不为空就是true，反之false
const isEdit = ref(route.query.id ? true : false)
const BlogData = ref({
  title: '',
  summary: '',
  content: '',
  category: '',
  // tags: [],
})

const EditBlogLoading = ref(false)
const handlerEditBlog = async (data) => {
  if (!isEdit.value) {
    try {
      EditBlogLoading.value = true
      isEdit.value = true
      const respData = await BlogCreateRequest(data)
      // 调用Notification实现消息提示
      // 实现页面切换，当保存之后获取到blog的id然后因为有了id变为编辑文章页面，url中也携带id保证刷新还是编辑页面
      router.replace({ name: 'backend_blog_edit', query: { id: respData.id } })
    } catch (err) {
      if (err === undefined) {
        Notification.success({ content: '文章新建成功！' })
      } else {
        // Notification.error({ content: `文章新建失败: ${err.message || err}` })
        Notification.error({
          content: `文章新建失败，请检查必选项是否都填写，如若填写，请联系管理员`,
        })
      }
    } finally {
      EditBlogLoading.value = false
    }
  } else {
    try {
      EditBlogLoading.value = true
      isEdit.value = true
      await BlogUpdateRequest(data, route.query.id)
      router.replace({ name: 'backend_blog_edit', query: { id: route.query.id } })
    } catch (err) {
      if (err === undefined) {
        Notification.success({ content: '文章新建成功！' })
      } else {
        // Notification.error({ content: `文章新建失败: ${err.message || err}` })
        Notification.error({
          content: `文章新建失败，请检查必选项是否都填写，如若填写，请联系管理员`,
        })
      }
    } finally {
      EditBlogLoading.value = false
    }
  }
}

// 当BlogList组件内点击编辑之后，请求后端API指定id的blog然后赋值给BlogData即可实现blog内容显示
onMounted(async () => {
  if (isEdit.value) {
    const BlogDescribeData = await BlogDescribeRequest(route.query.id)
    BlogData.value.title = BlogDescribeData.title
    BlogData.value.summary = BlogDescribeData.summary
    BlogData.value.content = BlogDescribeData.content
    BlogData.value.category = BlogDescribeData.category
  }
})

// 处理图片上传,by claude
const handleUploadImage = async (files, callback) => {
  try {
    console.log('开始上传图片，文件数量:', files.length)

    // files 是一个数组，包含所有要上传的文件
    const uploadPromises = files.map(async (file, index) => {
      try {
        console.log(`上传文件 ${index + 1}:`, file.name)

        const response = await ImageUploadRequest(file)

        console.log(`文件 ${index + 1} 上传响应:`, response)

        // mcube 的 response.Success 返回格式: { code: 0, data: { url, filename, size } }
        // 前端 axios 拦截器返回 value.data，所以这里直接访问 response.url
        const imageUrl = response.url

        if (!imageUrl) {
          console.error('无法从响应中提取 URL，完整响应:', JSON.stringify(response, null, 2))
          Notification.error({ content: `文件 ${file.name} 上传失败：无法获取图片URL` })
          return ''
        }

        console.log(`文件 ${index + 1} 上传成功，URL:`, imageUrl)
        return imageUrl
      } catch (error) {
        console.error(`文件 ${index + 1} 上传失败:`, error)
        console.error('错误详情:', error.response || error.message || error)
        Notification.error({ content: `图片上传失败: ${error.message || error}` })
        return ''
      }
    })

    // 等待所有图片上传完成
    const urls = await Promise.all(uploadPromises)
    console.log('所有图片上传完成，有效 URL 数量:', urls.filter((u) => u).length)

    // 调用回调函数，传入上传后的图片 URL 数组
    callback(urls)
  } catch (error) {
    console.error('图片上传异常:', error)
    Notification.error({ content: `图片上传异常: ${error.message || error}` })
  }
}
</script>

<style lang="css" scoped>
/* 覆盖arco design的原有样式，更改a-form的位置 */
:deep(.arco-form-item-label-col) {
  margin-left: 40px;
}
:deep(.arco-form-item-wrapper-col) {
  margin-left: 30px;
}
.editor_mapper {
  width: 100%;
}
.form-container {
  width: 97%;
}
/* 这里是为了得到EditPage的实际高度，60px即使上面的layout-header */
.edit-page-container {
  height: calc(100vh - 60px);
}
::-webkit-scrollbar {
  overflow: hidden;
}
</style>
