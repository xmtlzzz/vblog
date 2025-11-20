import { useStorage } from '@vueuse/core'

// 这里通过useStorage去调用localStorage存储请求后端返回的2个token信息，以及颁发给哪个用户的token
export const menu_key = useStorage('menu',
  {"system_current_menu_key":""},
  localStorage,{mergeDefaults:true}
)


export const login_info = useStorage('login_info',
  {"username":"","password":"","remember_me":false},
  localStorage,{mergeDefaults:true}
)

export const current_blog_id = useStorage('blog_id',
  {"current_blog_id":""},
  localStorage,{mergeDefaults:true}
)
