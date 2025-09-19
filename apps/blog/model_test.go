package blog_test

import (
	"testing"

	"github.com/xmtlzzz/vblog/apps/blog"
	"github.com/xmtlzzz/vblog/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestBlogTable(t *testing.T) {
	// 进程内环境变量设置，toml配置文件路径
	t.Setenv("workdir", "C:\\Users\\Administrator\\Desktop\\code\\Go\\vblog")
	db = utils.NewDBConnecter()
	if err := db.AutoMigrate(&blog.Blog{}); err != nil {
		t.Fatal(err)
	}
}
