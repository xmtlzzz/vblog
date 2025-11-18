<template>
  <!--  布局整体内容-->
  <a-layout class="layout-all">
    <!--    顶不导航-->
    <a-layout-header class="layout-header">
      <div class="header-value">VBLOG管理界面</div>
      <div
        style="margin-left: auto; margin-right: 10px; margin-top: 20px"
        @click="$router.push({ name: 'frontend_blog_list' })"
      >
        <a-button>
          <span>返回前台</span>
          <icon-import />
        </a-button>
      </div>
      <div style="margin-top: 10px; margin-right: 10px" @click="logout">
        <LoginOrLogout></LoginOrLogout>
      </div>
    </a-layout-header>
    <!--    侧边栏部分参照arco官网放到另一个layout-->
    <a-layout class="layout-sider">
      <!--    侧边栏菜单内容-->
      <!--      实现菜单栏可以收缩靠边-->
      <a-layout-sider :collapsible="true" breakpoint="xl">
        <a-menu
          @menu-item-click="handlerMenuClick"
          :selected-keys="menu_key.system_current_menu_key"
          :default-open-keys="[menu_key.system_current_menu_key]"
          :auto-open="true"
        >
          <a-sub-menu key="blog">
            <template #title>
              <span><icon-bookmark />文章管理</span>
            </template>
            <a-menu-item key="backend_blog_list">文章管理</a-menu-item>
            <a-menu-item key="backend_tag_list">标签管理</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="comment">
            <template #title>
              <span><icon-message />评论管理</span>
            </template>
            <a-menu-item key="backend_comment_list">评论管理</a-menu-item>
          </a-sub-menu>
        </a-menu>
      </a-layout-sider>
      <a-layout-content>
        <router-view></router-view>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { IsLogin, token } from '@/stores/token'
import { menu_key } from '@/stores/system'
import { LOGOUT } from '@/api/token'
import LoginOrLogout from '@/components/LoginOrLogout.vue'

const router = useRouter()
const handlerMenuClick = (key) => {
  router.push({ name: key })
  // 保存当前选中的key
  menu_key.value.system_current_menu_key = key
}

const logout = async () => {
  const token_data = {
    access_token: token.value.access_token,
    refresh_token: token.value.refresh_token,
  }
  console.log(token_data)
  await LOGOUT(token_data)
  token.value = null
  await router.push({ name: 'login' })
}
</script>

<style lang="css" scoped>
/* 将layout占据屏幕的整体高度 */
.layout-all {
  height: 100vh;
}
/* 参照arco design官网的导航栏样式 */
.layout-header {
  width: 100%;
  height: 60px;
  box-sizing: border-box;
  min-height: 60px;
  max-height: 60px;
  display: flex;
  border-bottom: 1px solid var(--color-border);
  background-color: var(--color-bg-2);
  justify-content: space-between;
}
/* header和sider都是layout的一部分，sider占据header用剩下的全部 */
.layout-sider {
  height: 100%;
}

.header-value {
  width: 30vh;
  margin-top: 10px;
  margin-left: 10px;
  font-weight: bold;
}
</style>
