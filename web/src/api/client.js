import axios from "axios";
import { Message } from '@arco-design/web-vue';

// 初始化一个axios客户端，用于对接后端API
export const client = axios.create({
  baseURL: "",
  timeout: 5000,
})

// 通过拦截器来实现后端响应信息报错信息捕获
// interceptors就是创建一个捕获器
// response就是捕获响应报文的内容，request就是捕获请求内容
// 这里的response就是后端的response内容
client.interceptors.response.use(
  // 请求成功返回具体的值
  (value)=> {
    // 获取response的data部分也就是后端实际返回的内容
    return value.data
  },
  // 请求失败就处理告警
  (error) => {
    let msg = error.message
    try {
      // 获取报错的详细内容，比如后端bcrypto的报错信息就存储在这个路径
      msg = error.response.data.message
    }catch(error) {
      // 保持msg为原始错误信息
      console.log(error)
    }
    Message.error(msg)
  },
)
