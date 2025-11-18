import { useStorage } from '@vueuse/core'

// 这里通过useStorage去调用localStorage存储请求后端返回的2个token信息，以及颁发给哪个用户的token
export const token = useStorage('token',
  {"access_token":"","ref_user_name":"","refresh_token":""},
  localStorage,{mergeDefaults:true}
)

// 用来校验用户是否登录
export const IsLogin= ()=>{
  console.log(token.value)
  console.log(token.value === '')
  return token.value.access_token!==undefined || token.value.ref_user_name!==""
}
