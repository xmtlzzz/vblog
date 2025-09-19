# impl层
用来实现interface层定义的方法，并且对外提供调用
```go
var TokenService token.Service = &TokenServiceImpl{}
```
外部调用可以将token.Service嵌套，并且使用token.TokenService初始化否则就是nil pointer