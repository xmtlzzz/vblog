package blog

import (
	"context"
	"strings"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/xmtlzzz/vblog/utils"
)

const (
	AppName = "blog"
)

func GetService() Service {
	// 调用ioc池子中指定名称的对象，这里就是在调用blog中的impl
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	//+ 创建博客
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	//+ 博客列表查询，返回结构体让方法返回不需要变动，变结构体属性
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	//+ 博客详情查询
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	//+ 博客编辑
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	//+ 博客发布
	PublishBlog(context.Context, *PublishBlogRequest) (*Blog, error)
	//+ 博客撤销（删除）
	DeleteBlog(context.Context, *DeleteBlogRequest) error
}

type DeleteBlogRequest struct {
	utils.GetRequest
}

// 发布博客根据blogid
type PublishBlogRequest struct {
	utils.GetRequest
}

// 修改blog的属性根据blogid匹配
type UpdateBlogRequest struct {
	utils.GetRequest
	// 可以修改的属性
	CreateBlogRequest
}

// 根据blogid找到blog查看详细信息
type DescribeBlogRequest struct {
	utils.GetRequest
}

// 设置form tag让后续url解析表单的时候属性对应
type QueryBlogRequest struct {
	// 分页查找参数导入
	utils.PageRequest
	// 状态过滤参数，管理员nil（都可以查），普通用户只能公开STAGE_PUBLISHED
	Stage *STAGE `json:"stage" form:"stage"`
	// 文章标题模糊匹配
	Keywords string `json:"keywords" form:"keywords"`
	// 根据分类查
	Category string `json:"category" form:"category"`
	// 查看指定用户的文章，作者使用
	CreateBy string `json:"create_by" form:"create_by"`
	// 根据tag匹配
	Tags map[string]string `json:"tags" form:"-"`
}

// 解析Tags属性，按照K1=v1这样的格式输出
func (r *QueryBlogRequest) SetTag(tag string) {
	// ，分割多个tag kv对参数
	kvItem := strings.Split(tag, ",")
	for i := range kvItem {
		kvString := kvItem[i]
		// 使用=分割
		kv := strings.Split(kvString, "=")
		if len(kv) > 1 {
			r.Tags[kv[0]] = strings.Join(kv[1:], "=")
		}
	}
}
