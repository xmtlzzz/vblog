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
        <div class="left">left</div>
        <div class="center">
          <router-view></router-view>
        </div>
        <div class="right">right</div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script setup>
// 组件化的形式导入使用
import LoginOrLogout from '@/components/LoginOrLogout.vue'
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
}
.left{
  width: 30%;
  overflow: auto;
}
.right{
  width: 30%;
  overflow: auto;
}
.center{
  width: 100%;
  overflow: auto;

}
.center::-webkit-scrollbar {
  display: none;
  /* Chrome, Safari 和 Opera */
}
</style>
