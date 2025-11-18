package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/ioc"
	ioc_gin "github.com/infraboard/mcube/v2/ioc/config/gin"
	"github.com/xmtlzzz/vblog/apps/blog"
	"github.com/xmtlzzz/vblog/middleware"
	"github.com/xmtlzzz/vblog/response"
	"github.com/xmtlzzz/vblog/utils"
)

func init() {
	ioc.Api().Registry(&BlogApiHandler{})
}

type BlogApiHandler struct {
	ioc.ObjectImpl
	blog blog.Service
}

func (b *BlogApiHandler) Name() string {
	return "blogs"
}

// 重写ioc框架的Init方法实现服务注册，等于是在main中直接调用registry方法注册路由
func (b *BlogApiHandler) Init() error {
	b.blog = blog.GetService()
	// url路径一般为/api/mcube_service/v1/hello_module/
	r := ioc_gin.ObjectRouter(b)
	r.GET("/frontend_query", b.FrontendQueryBlog)
	r.Use(middleware.Auth)
	r.POST("/create", b.CreateBlog)
	r.GET("/query", b.QueryBlog)
	r.GET("/describe/:id", b.DescribeBlog)
	r.PUT("/update/:id", b.UpdateBlog)
	r.POST("/publish/:id", b.PublishBlog)
	r.DELETE("/delete/:id", b.DeleteBlog)
	return nil
}

func (b *BlogApiHandler) CreateBlog(ctx *gin.Context) {
	bg := &blog.CreateBlogRequest{}
	if err := ctx.BindJSON(bg); err != nil {
		response.Failed(ctx, err)
		return
	}
	ins, err := b.blog.CreateBlog(ctx.Request.Context(), bg)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}

// + 博客列表查询，返回结构体让方法返回不需要变动，变结构体属性
func (b *BlogApiHandler) QueryBlog(ctx *gin.Context) {
	bg := &blog.QueryBlogRequest{}
	if err := ctx.BindQuery(bg); err != nil {
		response.Failed(ctx, err)
		return
	}
	// 从url中获取kv对参数，指定tag=xxx
	// 通过SetTag方法实现自定义Tags属性输出
	bg.SetTag(ctx.Query("tags"))
	ins, err := b.blog.QueryBlog(ctx.Request.Context(), bg)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}

// + 博客详情查询
func (b *BlogApiHandler) DescribeBlog(ctx *gin.Context) {
	// 解析路径的id去定位blog
	bg := &blog.DescribeBlogRequest{}
	idInt, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	bg.Id = uint(idInt)
	ins, err := b.blog.DescribeBlog(ctx.Request.Context(), bg)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}

// + 博客编辑
// param获取id，body通过json传递数据
func (b *BlogApiHandler) UpdateBlog(ctx *gin.Context) {
	bg := &blog.UpdateBlogRequest{}
	idInt, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	bg.Id = uint(idInt)
	// 解析body携带的json信息
	if err := ctx.BindJSON(bg); err != nil {
		response.Failed(ctx, err)
		return
	}
	ins, err := b.blog.UpdateBlog(ctx.Request.Context(), bg)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}

// + 博客发布
func (b *BlogApiHandler) PublishBlog(ctx *gin.Context) {
	bg := &blog.PublishBlogRequest{}
	status := &blog.StatusSpec{}
	if err := ctx.BindJSON(status); err != nil {
		response.Failed(ctx, err)
		return
	}
	idInt, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	bg.Id = uint(idInt)
	ins, err := b.blog.PublishBlog(ctx.Request.Context(), bg, status)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}

// + 博客撤销（删除）
func (b *BlogApiHandler) DeleteBlog(ctx *gin.Context) {
	bg := &blog.DeleteBlogRequest{}
	idInt, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	bg.Id = uint(idInt)
	err = b.blog.DeleteBlog(ctx.Request.Context(), bg)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, "删除成功")
}

func (b *BlogApiHandler) FrontendQueryBlog(ctx *gin.Context) {
	pg := utils.PageRequest{}
	if err := ctx.BindQuery(&pg); err != nil {
		response.Failed(ctx, err)
		return
	}
	ins, err := b.blog.FrontendQueryBlog(ctx.Request.Context(), pg)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}
