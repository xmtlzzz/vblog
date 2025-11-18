package impl_test

import (
	"context"
	"testing"

	"github.com/xmtlzzz/vblog/apps/blog"
	"github.com/xmtlzzz/vblog/apps/token"
	"github.com/xmtlzzz/vblog/middleware"
)

var ctx = context.Background()

// 从ioc池子获取对象

func BenchmarkBlogServiceImpl_QueryBlog(t *testing.B) {
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	var it = blog.NewQueryBlogRequest()
	it.Tags = map[string]string{"language": "golang"}
	ins, err := blog.GetService().QueryBlog(ctx, it)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range ins.Items {
		t.Log(v)
	}
}

func TestCreateBlog(t *testing.T) {
	t.Setenv("workdir", "D:\\Desktop\\code\\Go\\vblog")
	var it = blog.CreateBlogRequest{
		Title:    "测试5",
		Summary:  "gin and gorm1",
		Content:  "test2",
		Category: "软件开发1",
		Tags: map[string]string{
			"language": "golang",
		},
	}
	// 携带token进行中间件鉴权
	tk := token.Token{AccessToken: "123"}
	ct := context.WithValue(ctx, middleware.TokenCtxKey{}, &tk)
	// 从ioc池子获取对象，定义在interface层
	ins, err := blog.GetService().CreateBlog(ct, &it)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
func TestQueryBlog(t *testing.T) {
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	var it = blog.NewQueryBlogRequest()
	it.Tags = map[string]string{"language": "golang"}
	ins, err := blog.GetService().QueryBlog(ctx, it)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range ins.Items {
		t.Log(v)
	}
}

func TestDescribeBlog(t *testing.T) {
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	var it = blog.NewDescribeBlogRequest()
	it.Id = 7
	ins, err := blog.GetService().DescribeBlog(ctx, it)
	if err != nil {
		t.Error(err)
	}
	t.Log(ins)
}

func TestUpdateBlog(t *testing.T) {
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	var it = blog.NewUpdateBlogRequest()
	it.Id = 7
	it.Tags = map[string]string{"test": "123"}
	it.Title = "sz"
	ins, err := blog.GetService().UpdateBlog(ctx, it)
	if err != nil {
		t.Error(err)
	}
	t.Log(ins)
}

func TestPublishBlog(t *testing.T) {
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	var it = blog.NewPublishBlogRequest()
	it.Id = 7
	var status = &blog.StatusSpec{Stages: blog.STAGE_PUBLISHED}
	ins, err := blog.GetService().PublishBlog(ctx, it, status)
	if err != nil {
		t.Error(err)
	}
	t.Log(ins)
}
func TestDeleteBlog(t *testing.T) {
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	var it = blog.NewDeleteBlogRequest()
	it.Id = 7
	err := blog.GetService().DeleteBlog(ctx, it)
	if err != nil {
		t.Error(err)
	}
}
