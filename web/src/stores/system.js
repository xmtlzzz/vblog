import { useStorage } from '@vueuse/core'

// 这里通过useStorage去调用localStorage存储请求后端返回的2个token信息，以及颁发给哪个用户的token
export const menu_key = useStorage('menu',
  {"system_current_menu_key":""},
  localStorage,{mergeDefaults:true}
)


export const login_info = useStorage('login_info',
  {"username":"","password":""},
  localStorage,{mergeDefaults:true}
)
