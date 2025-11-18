package server

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/ioc"
	ioc_gin "github.com/infraboard/mcube/v2/ioc/config/gin"
)

var GinServer = gin.Default()

func init() {
	ioc.Config().Registry(&StaticFileConfig{})
}

type StaticFileConfig struct {
	ioc.ObjectImpl
}

func (s *StaticFileConfig) Name() string {
	return "static_file"
}

// Init 初始化静态文件服务
func (s *StaticFileConfig) Init() error {
	// 获取 Gin 引擎
	engine := ioc_gin.RootRouter()

	// 配置静态文件服务，将 /uploads 路径映射到 uploads 目录
	engine.Static("/uploads", "./uploads")

	return nil
}
