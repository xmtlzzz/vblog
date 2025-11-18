package test

import (
	"log"
	"os"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/server"
)

// 设置IoC的配置
func LoadConfig() {
	// 导入单元测试的配置，application.toml
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	workspecDir := os.Getenv("workdir")
	if workspecDir == "" {
		workspecDir = "D:\\Desktop\\code\\Go\\vblog"
	} else {
		req.ConfigFile.Path = workspecDir + "\\etc\\application.toml"
	}
	err := ioc.ConfigIocObject(req)
	if err != nil {
		log.Fatal(err)
	}
	// 加载配置给ginserver使用
	server.DefaultConfig = req
}
