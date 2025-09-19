package test

import (
	"fmt"
	"log"
	"os"

	"github.com/infraboard/mcube/v2/ioc"
)

func LoadConfig() {
	// 导入单元测试的配置，application.toml
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	workspecDir := os.Getenv("workdir")
	if workspecDir == "" {
		workspecDir = "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog"
	} else {
		req.ConfigFile.Path = workspecDir + "\\etc\\application.toml"
	}
	fmt.Println(req.ConfigFile.Path)
	err := ioc.ConfigIocObject(req)
	if err != nil {
		log.Fatal(err)
	}

}
