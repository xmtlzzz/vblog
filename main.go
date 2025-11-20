package main

import (
	"context"
	"log"

	"github.com/infraboard/mcube/v2/ioc/server"
	_ "github.com/xmtlzzz/vblog/apps"
	"github.com/xmtlzzz/vblog/test"
	//_ "github.com/xmtlzzz/vblog/web"
)

func main() {
	// 配置解析导入
	test.LoadConfig()
	//if err := server.GinServer.Run(); err != nil {
	//	log.Println(err)
	//}
	if err := server.Run(context.Background()); err != nil {
		log.Println(err)
	}
}
