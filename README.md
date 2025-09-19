# vblog
engined by gin gorm and vue

## OOP基本组织
在这个项目中主要涉及到了5个对象
1. TokenApiHandler
2. TokenServiceImpl
3. BlogApiHandler
4. BlogServiceImpl
5. UserServiceImpl

![img.png](docs/oop.png)

```go
blogApi.NewBlogApiHandler(&blogImpl.BlogServiceImpl{}).Registry(ge)
tokenApi.NewBTokenAPIHandler(tokenImpl.NewTokenServiceImpl(&userImpl.UserServiceImpl{})).Registry(ge)

token/impl.go
var TokenService token.Service = &TokenServiceImpl{}

type TokenServiceImpl struct {
    UserSvc user.AdminService
}

func NewTokenServiceImpl(user user.AdminService) token.Service {
    return &TokenServiceImpl{
        UserSvc: user,
    }
}

token/api.go
type TokenAPIHandler struct {
    token token.Outer
}

func NewBTokenAPIHandler(tokenService impl.TokenServiceImpl) *TokenAPIHandler {
    return &TokenAPIHandler{
        token: &tokenService,
    }
}
```

OOP的思想实现，可以将每个对象的构造函数中定义形参，类型就是属性类型，将属性设置转移到外部调用时确定
这个就是OOP的依赖注入思想

定义构造函数的时候最好就是不要将属性定死，而是通过上面的方式去组织代码逻辑

并且在返回的时候返回接口更容易扩展

## IoC托管
![img.png](docs/ioc.png)
将5个对象托管到IoC容器中，后续互相调用和IoC容器请求即可

```go
token/interface.go
const (
	AppName = "blog"
)

// 对外的构造函数
func GetService() Service {
	// 调用ioc池子中指定名称的对象，这里就是在调用blog中的impl
	return ioc.Controller().Get(AppName).(Service)
}

token/impl.go
func init() {
// 注册到ioc池子
ioc.Controller().Registry(&TokenServiceImpl{})
}

type TokenServiceImpl struct {
// 继承mcube框架的Object接口实现，从而能够注册	
    ioc.ObjectImpl
// 调用DescribeUser查询用户
    UserSvc user.AdminService
}

// 设置注册到IoC容器的名字
func (*TokenServiceImpl) Name() string {
    return token.AppName
}

// 重写ioc框架的Init方法实现服务注册，等于是在main中直接调用registry方法注册路由
func (t *TokenAPIHandler) Init() error {
	// 通过IoC池子获取属性
    t.token = token.GetService()
    t.Registry(server.GinServer)
    return nil
}

func (t *TokenAPIHandler) Registry(ge *gin.Engine) {
    server := ge.Group("/vblog/api/v1/tokens")
    server.POST("/", t.IssueToken)
    server.POST("/revolk", t.RevolkToken)
    server.POST("/refresh", t.RefreshToken)
}

// 具体调用
ins, err := token.GetService().RevolkToken(ctx, it)
```
后续其他模块中如果想要调用impl的方法直接token.GetService即可，因为此时已经将对象实例注册到了Ioc池子中
