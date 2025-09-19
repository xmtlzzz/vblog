# apps模块
主要是存放后端业务逻辑的部分

## 单元测试
这里blog、token、user子模块通过interface定义方法以及结构体模型

最终都通过单测的形式去构建数据表
```go
var db *gorm.DB

func TestTokenTable(t *testing.T) {
	// 进程内环境变量设置，toml配置文件路径
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	db = utils.NewDBConnecter()
	if err := db.AutoMigrate(&token.Token{}); err != nil {
		t.Fatal(err)
	}
}
```
总体就是设置toml配置文件的路径，将项目根路径作为环境变量在单测读取，因为只需要读取一个就这样写了

## registry
用于针对apps中的模块进行ioc的统一依赖导入

因为每个对象中都通过init方法实现了到ioc资源池的注册，那么只需要通过匿名导入的形式去执行init函数即可
