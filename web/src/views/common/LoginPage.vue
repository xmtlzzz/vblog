<template>
  <div class="content">
    <div class="login-container">
      <h2 class="login-title">欢迎来到VBLOG登录页面</h2>
      <!--      绑定响应式对象、以及提交事件-->
      <a-form class="login-form" :model="form" @submit="handlerSubmit">
        <a-form-item
          style="padding-top: 20px"
          field="username"
          tooltip="请输入用户名"
          :rules="{ required: true, message: `请输入用户名` }"
          hide-label
          hide-asterisk
        >
          <a-input v-model="form.username" placeholder="Please Enter Your Username">
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item
          field="password"
          tooltip="请输入密码"
          :rules="{ required: true, message: `请输入密码` }"
          hide-label
          hide-asterisk
        >
          <a-input-password v-model="form.password" placeholder="Please Enter Password" allow-clear>
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item>
          <!--          boolean表示的是类型-->
          <a-checkbox
            style="margin-left: auto"
            v-model="form.remember_me"
            value="boolean"
            @change="isRemember"
            >记住</a-checkbox
          >
        </a-form-item>
        <a-form-item hide-asterisk hide-label>
          <a-button :loading="submitLoading" style="width: 100%" type="primary" html-type="submit">
            登录
          </a-button>
        </a-form-item>
      </a-form>
      <a-form style="  width: 80%;margin-top: 10px">
        <a-form-item hide-asterisk hide-label>
          <a-button @click="handlerRevolk" :loading="revolkLoading" style="width: 100%" type="primary" html-type="submit">
            注销登录信息
          </a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'
import { LOGIN } from '@/api/token'
import { token } from '@/stores/token'
import { login_info } from '@/stores/system.js'
import { useRouter } from 'vue-router'
import { getDecryptedStorage, setEncryptedStorage } from '@/utils/encode_decode.js'
import { Notification } from '@arco-design/web-vue'

// form响应式对象，属性和vblog中颁发token的属性相同
const form = reactive({
  username: '',
  password: '',
  remember_me: false,
})

const submitLoading = ref(false)
const router = useRouter()

// 修改表单提交处理函数
const handlerSubmit = async (data) => {
  // 检查是否有验证错误
  if (data && data.errors === undefined) {
    try {
      if (form.remember_me) {
        setEncryptedStorage('userPassword', form.password)
        setEncryptedStorage('userName', form.username)
        login_info.value.remember_me = form.remember_me
        login_info.value.username =localStorage.getItem("userName")
        login_info.value.password =localStorage.getItem("userPassword")
      }
      submitLoading.value = true
      // 得到client.js捕获器return的正确response结果
      const respData = await LOGIN(data.values)
      console.log('登录成功:', respData)
      // 将后端返回的token实际内容存储到localStorage持久化，后续其他模块、页面可以重复调用token实现校验
      token.value.access_token = respData.access_token
      token.value.refresh_token = respData.refresh_token
      token.value.ref_user_name = respData.ref_user_name
      router.push({ name: 'backend_blog' })
    } catch (error) {
      // 只输出错误信息，而不是整个错误对象
      console.error('登录失败:', error)
      // 可以添加用户友好的错误提示
    } finally {
      submitLoading.value = false
    }
  } else if (data && data.errors) {
    console.log('表单验证失败:', data.errors)
  }
}

const revolkLoading = ref(false)
const handlerRevolk = () => {
  localStorage.removeItem('userName')
  localStorage.removeItem('userPassword')
  localStorage.removeItem('login_info')
  Notification.success({content:"登录状态信息已经清除，请刷新页面重新输入用户名密码",position:'topLeft'})
}

onMounted(() => {
  if (login_info.value.remember_me) {
    form.username = getDecryptedStorage('userName')
    form.password = String(getDecryptedStorage('userPassword'))
  }
})

const isRemember = () => {}
</script>

<style lang="css" scoped>
.content {
  width: 100%;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-container {
  width: 460px;
  height: 360px;
  flex-direction: column;
  display: flex;
  align-items: center;
  background: rgba(255, 255, 255, 0.3);
  /* 增加透明度 */
  border-radius: 15px;
  backdrop-filter: blur(15px);
  /* 增加模糊程度 */
  box-shadow: 0 4px 30px rgba(0, 0, 0, 0.2);
}

.login-title {
  padding-top: 20px;
  font-size: 22px;
  font-weight: bold;
  color: var(--color-neutral-8);
}

.login-form {
  width: 80%;
  height: 60%;
}
</style>
