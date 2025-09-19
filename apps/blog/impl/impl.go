package impl

import (
	"context"
	"fmt"
	"strconv"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/xmtlzzz/vblog/apps/blog"
	"github.com/xmtlzzz/vblog/middleware"
	"github.com/xmtlzzz/vblog/utils"
	"gorm.io/gorm"
)

func init() {
	// 注册到ioc池子
	ioc.Controller().Registry(&BlogServiceImpl{})
}

var BlogService blog.Service = &BlogServiceImpl{}
var db *gorm.DB

type BlogServiceImpl struct {
	// 继承object接口实现
	ioc.ObjectImpl
}

func (*BlogServiceImpl) Name() string {
	return blog.AppName
}

func (b BlogServiceImpl) CreateBlog(ctx context.Context, request *blog.CreateBlogRequest) (*blog.Blog, error) {
	ins, err := blog.NewBlog(request)
	if err != nil {
		return nil, err
	}
	db = utils.NewDBConnecter()
	// 调用中间件逻辑，获取ctx中的token，这里ctx后续api层传入的是gin.request.Context获取http请求报文中的内容
	// 测试的时候通过Header的Bear xxx形式传入token即可实现简单的鉴权
	tk := middleware.GetTokenFromCtx(ctx)
	// 补充一段通过上面token找到user的查询，设置created_by字段
	if err := db.WithContext(ctx).Where("access_token = ?", &tk.AccessToken).Take(tk).Error; err != nil {
		return nil, err
	}
	fmt.Println(tk)
	// 添加关联信息
	ins.CreateBy = strconv.Itoa(tk.RefUserId)
	if err := db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (b BlogServiceImpl) QueryBlog(ctx context.Context, request *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	db = utils.NewDBConnecter()
	// 指定操作的数据表
	db = db.WithContext(ctx).Model(&blog.Blog{})
	if request.Keywords != "" {
		db.Where("title like ?", "%"+request.Keywords+"%")
	}
	if request.Stage != nil {
		db.Where("stage = ?", request.Stage)
	}
	if request.Category != "" {
		db.Where("category = ?", request.Category)
	}
	if request.CreateBy != "" {
		db.Where("create_by = ?", request.CreateBy)
	}
	// 提取k、v然后分别通过占位符的形式去传递k、v，得到k对应的value
	// 因为属性是map，转换之后需要这样操作
	for k, v := range request.Tags {
		db = db.Where(fmt.Sprintf("tags->>'$.%s' = ?", k), v)
	}
	set := blog.NewBlogSet()
	// 判断总共有多少条数据
	if err := db.Count(&set.Total).Error; err != nil {
		return nil, err
	}
	// 处理分页，按照created_at排序然后执行分页
	// 这里默认不offset，limit限制为20条
	// 防止用户不传入pagesize导致limit 0无法返回，至少返回一条
	if request.PageSize == 0 {
		request.PageSize = 1
	}
	if err := db.Order("created_at").Offset(int(request.Offset())).Limit(int(request.PageSize)).Find(&set.Items).Error; err != nil {
		return nil, err
	}
	return set, nil

}

func (b BlogServiceImpl) DescribeBlog(ctx context.Context, request *blog.DescribeBlogRequest) (*blog.Blog, error) {
	var bg = &blog.Blog{}
	db = utils.NewDBConnecter()
	if err := db.WithContext(ctx).Where("id = ?", request.Id).Take(bg).Error; err != nil {
		return nil, exception.NewBadRequest("查找的blog不存在,err: %v", err)
	}
	return bg, nil
}

func (b BlogServiceImpl) UpdateBlog(ctx context.Context, request *blog.UpdateBlogRequest) (*blog.Blog, error) {
	var bg = &blog.Blog{}
	db = utils.NewDBConnecter()
	if err := db.WithContext(ctx).Where("id = ?", request.Id).Take(bg).Error; err != nil {
		return nil, exception.NewBadRequest("查找的blog不存在,err: %v", err)
	}
	if request.Title != "" {
		bg.Title = request.Title
	}
	if request.Summary != "" {
		bg.Summary = request.Summary
	}
	if request.Content != "" {
		bg.Content = request.Content
	}
	if request.Category != "" {
		bg.Category = request.Category
	}
	if len(request.Tags) != 0 {
		// 防止tags默认为空导致panic提示nil map写入
		bg.Tags = make(map[string]string)
		for i, v := range request.Tags {
			bg.Tags[i] = v
		}
	}
	if err := db.WithContext(ctx).Where("id = ?", request.Id).Updates(bg).Error; err != nil {
		return nil, exception.NewBadRequest("更新失败查找的blog不存在,err: %v", err)
	}
	return bg, nil
}

func (b BlogServiceImpl) PublishBlog(ctx context.Context, request *blog.PublishBlogRequest) (*blog.Blog, error) {
	var bg = &blog.Blog{}
	// 发布
	db = utils.NewDBConnecter()
	bg.Stages = blog.STAGE_PUBLISHED
	if err := db.WithContext(ctx).Where("id = ?", request.Id).Updates(bg).Error; err != nil {
		return nil, exception.NewBadRequest("查找的blog不存在，公开文章失败,err: %v", err)
	}
	if err := db.WithContext(ctx).Where("id = ?", request.Id).Take(bg).Error; err != nil {
		return nil, exception.NewBadRequest("查找的blog不存在,err: %v", err)
	}
	return bg, nil
}

func (b BlogServiceImpl) DeleteBlog(ctx context.Context, request *blog.DeleteBlogRequest) error {
	var bg = &blog.Blog{}
	db = utils.NewDBConnecter()
	if err := db.WithContext(ctx).Where("id = ?", request.Id).Delete(bg).Error; err != nil {
		return exception.NewBadRequest("查找的blog不存在或已经被删除,err: %v", err)
	}
	return nil
}
