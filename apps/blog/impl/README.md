# impl层
用来实现interface层定义的方法，并且对外提供调用
```go
var BlogService blog.Service = &BlogServiceImpl{}
```
外部调用可以将blog.Service嵌套，并且使用blog.BlogService初始化否则就是nil pointer