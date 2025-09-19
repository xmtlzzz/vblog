# 博客管理

+ 创建博客
+ 博客列表查询
+ 博客详情查询
+ 博客编辑
+ 博客发布
+ 博客撤销（删除）

## interface规划
每个interface都通过定义结构体来作为形参，后续传入结构体实例进行调用

## queryBlog方法

```go
interface.go:
QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)

model.go:
type BlogSet struct {
// 返回blog的数量
BlogNum uint
}
```
这样定义主要是为了避免后续方法的修改，想要修改返回的内容直接修改结构体即可
