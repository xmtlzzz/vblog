package main

import (
	"log"

	_ "github.com/xmtlzzz/vblog/apps"
	"github.com/xmtlzzz/vblog/server"
	"github.com/xmtlzzz/vblog/test"
)

func main() {
	// 配置解析导入
	test.LoadConfig()
	if err := server.GinServer.Run(); err != nil {
		log.Println(err)
	}
}
