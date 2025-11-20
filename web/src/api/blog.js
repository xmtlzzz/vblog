import { client } from '@/api/client.js'


export const BlogQueryRequest=(params)=>client({
  url:"/api/vblog/v1/blogs/query",
  method:"GET",
  params,
})


export const BlogDeleteRequest=(id)=>client({
  // ${}实现可选参数拼接
  url:`/api/vblog/v1/blogs/delete/${id}`,
  method:"DELETE",
  data: id,
})


export const BlogCreateRequest=(blog_info)=>client({
  url:'/api/vblog/v1/blogs/create',
  method:"POST",
  data: blog_info,
})

export const BlogUpdateRequest=(blog_info,id)=>client({
  url:`/api/vblog/v1/blogs/update/${id}`,
  method:"PUT",
  data: blog_info,
})


export const BlogDescribeRequest=(id)=>client({
  url:`/api/vblog/v1/blogs/describe/${id}`,
  method:"GET",
  data:id,
})

export const BlogPublishRequest=(id,data)=>client({
  url:`/api/vblog/v1/blogs/publish/${id}`,
  method:"POST",
  data: data,
})


export const FrontendBlogQueryRequest=(params)=>client({
  url:"/api/vblog/v1/blogs/frontend_query",
  method:"GET",
  params
})
