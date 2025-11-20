package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/ioc"
	ioc_gin "github.com/infraboard/mcube/v2/ioc/config/gin"
)

// 注册
func init() {
	ioc.Config().Registry(&Web{})
}

// 添加mcube ioc的基本实现
type Web struct {
	ioc.ObjectImpl
}

func (w *Web) Name() string {
	return "web"
}

//go:embed dist/assets/*
var assets embed.FS

//go:embed dist/index.html
var webDir embed.FS

func (w Web) Init() error {
	// 获取根路由
	rr := ioc_gin.RootRouter()

	// 从 embed.FS 中获取 dist/assets 目录的子文件系统，用于提供静态资源访问
	staticFS, err := fs.Sub(assets, "dist/assets")
	if err != nil {
		return err
	}

	// 将assets目录的内容直接映射到x.x.x.x/assets目录，url映射
	rr.StaticFS("/assets", http.FS(staticFS))

	// 当get/根路径的时候调用自定义的Index方法返回dist/index.html的内容
	rr.GET("/", w.Index)
	// 所有非 API 路由回退到前端入口
	rr.NoRoute(w.Index)
	return nil
}

//// RedirectIndex 重定向
//func (h *Web) RedirectIndex(c *gin.Context) {
//	c.Redirect(http.StatusFound, "/")
//}

// Index 实现主页内容返回
func (h *Web) Index(c *gin.Context) {
	c.Header("Content-Type", "text/html;charset=utf-8")
	index, err := webDir.ReadFile("dist/index.html")
	if err != nil {
		panic(err)
	}
	c.String(200, string(index))
}
