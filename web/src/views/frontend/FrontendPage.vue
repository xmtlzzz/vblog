<template>
  <div>
    <a-layout class="layout-all">
      <a-layout-header class="layout-header">
        <div class="header-value">VBLOG前台</div>
        <div style="margin-top: 10px; margin-right: 10px">
          <LoginOrLogout :redirectToLogin="false"></LoginOrLogout>
        </div>
      </a-layout-header>
      <a-layout-content class="layout-content">
        <div class="left">
          <FrontendBlogScrollElement v-if="isDetailPage" :scrollElement="centerRef"></FrontendBlogScrollElement>
        </div>
        <div class="center" ref="centerRef">
          <router-view></router-view>
        </div>
<!--        <div class="right">right</div>-->
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import LoginOrLogout from '@/components/LoginOrLogout.vue'
import FrontendBlogScrollElement from '@/components/FrontendBlogScrollElement.vue'

const centerRef = ref(null)
const route = useRoute()
const isDetailPage = computed(() => route.name === 'frontend_blog_detail')
</script>

<style lang="css" scoped>
.layout-all {
  height: 100vh;
}
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
.header-value {
  width: 30vh;
  margin-top: 10px;
  margin-left: 10px;
  font-weight: bold;
}
.layout-content {
  display: flex;
  justify-content: space-between;
  height: calc(100vh - 60px);
}
.left{
  width: 20%;
  overflow: auto;
  position: sticky;
  top: 0;
  height: calc(100vh - 60px);
}
.right{
  width: 20%;
  overflow: auto;
}
.center{
  width: 80%;
  padding-right: 200px;
  overflow: auto;
  height: calc(100vh - 60px);
}
.center::-webkit-scrollbar {
  display: none;
  /* Chrome, Safari 和 Opera */
}
</style>
