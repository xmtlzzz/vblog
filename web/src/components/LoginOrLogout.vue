<template>
  <div style="margin-top: 10px; margin-right: 10px">
    <a-button v-if="isLogin" @click="LoginOrOut">
      <span style="margin-right: 10px">退出登录</span>
      <icon-import />
    </a-button>
    <a-button v-else @click="router.push({ name: 'login' })">
      <span style="margin-right: 10px">登录</span>
      <icon-import />
    </a-button>
  </div>
</template>

<script setup>
import { Notification } from '@arco-design/web-vue'
import { IsLogin, token } from '@/stores/token.js'
import { login_info } from '@/stores/system.js'
import { LOGOUT } from '@/api/token.js'
import { useRouter } from 'vue-router'
import { computed } from 'vue'

// 定义一个props属性，只有当redirectToLogin为true的时候才会跳转到login界面
// 后台退出登录是要的，前台不需要
const props = defineProps({
  redirectToLogin: {
    type: Boolean,
    default: true,
  },
})
const router = useRouter()

// store/token.js中的IsLogin不是响应式的，所以点击退出登录会执行LoginOrOut函数
// 但是按钮不会刷新，导致一直是退出登录的逻辑
const isLogin = computed(() => {
  return token.value && token.value.access_token && token.value.access_token !== ''
})

const LoginOrOut = async () => {
  if (IsLogin()) {
    const token_data = {
      access_token: token.value.access_token,
      refresh_token: token.value.refresh_token,
    }
    await LOGOUT(token_data)
    // 退出登录全局气泡提示，用了arco design的组件
    Notification.success({content:'Account Is Already Logout!', position: "topLeft"})
    token.value = { access_token: '' , refresh_token: '' ,ref_user_name: ''}
    if (!login_info.value.remember_me) {
      login_info.value = {username: '', password: '',remember_me: false}
    }
    if (props.redirectToLogin) {
      router.push({ name: 'login' })
    } else {
      router.push({ name: 'frontend_blog_list' })
    }
  }
}

</script>

<style lang="scss" scoped></style>
