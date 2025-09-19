package blog

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/xmtlzzz/vblog/utils"
)

// 优化interface定义的QueryBlog方法的返回值
type BlogSet struct {
	// 返回blog的数量
	Total int64 `json:"total"`
	// blog切片
	Items []*Blog `json:"items"`
}

func (b *Blog) String() string {
	return pretty.ToJSON(b)
}

// 构造方法
func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type Blog struct {
	utils.ResourceMetadata
	CreateBlogRequest
	Status
}

// 对外构造方法
func NewBlog(b *CreateBlogRequest) (*Blog, error) {
	if err := b.Valiate(); err != nil {
		return nil, err
	}
	return &Blog{
		ResourceMetadata:  *utils.NewResourceMetadata(),
		CreateBlogRequest: *b,
	}, nil
}

type CreateBlogRequest struct {
	// 标题
	Title string `json:"title" gorm:"column:title;type:varchar(255)" validate:"required"`
	// 摘要，内容较大使用text存储
	Summary string `json:"summary" gorm:"column:summary;type:text" validate:"required"`
	// 内容
	Content string `json:"content" gorm:"column:content;type:text" validate:"required"`
	// 分类
	Category string `json:"category" gorm:"column:category;type:varchar(255)" validate:"required"`
	// 标签，gorm中想要直接存储map类型的结构需要序列化serializer实现json序列化
	Tags map[string]string `json:"tags" gorm:"column:tags;serializer:json"`
}

// 构造方法
func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageRequest: *utils.NewPageRequest(),
		Tags:        map[string]string{},
	}
}

func NewDescribeBlogRequest() *DescribeBlogRequest {
	return &DescribeBlogRequest{}
}

func NewUpdateBlogRequest() *UpdateBlogRequest {
	return &UpdateBlogRequest{}
}

func NewPublishBlogRequest() *PublishBlogRequest {
	return &PublishBlogRequest{}
}

func NewDeleteBlogRequest() *DeleteBlogRequest {
	return &DeleteBlogRequest{}
}

type Status struct {
	StatusSpec
	// 状态何时改变，所以传指针可以是nil也就是没变化
	ChangeAt *time.Time `json:"change_at" gorm:"column:change_at"`
}

// 因为正常情况下用户只需要更新blog的状态，时间是自动更新了所以这里单独定义
type StatusSpec struct {
	// 自定义类型枚举实现文章状态，后续可以通过状态找文章，所以加索引
	Stages STAGE `json:"stages" gorm:"column:stages;type:tinyint;index"`
}

func (c *CreateBlogRequest) Valiate() error {
	valid := validator.New()
	if err := valid.Struct(c); err != nil {
		return err
	}
	return nil
}
