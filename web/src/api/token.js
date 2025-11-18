import { client } from './client.js'

// 初始化一个LOGIN对象，用于后续请求token颁发
export const LOGIN= (data)=>client({
  url:"/api/vblog/v1/tokens/",
  method:"POST",
  data:data
})

export const LOGOUT=(data)=>client({
  url:"/api/vblog/v1/tokens/revolk",
  method:"POST",
  data:data
})
