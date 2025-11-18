package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/http/gin/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	ioc_gin "github.com/infraboard/mcube/v2/ioc/config/gin"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/xmtlzzz/vblog/apps/token"
	"github.com/xmtlzzz/vblog/middleware"
)

func init() {
	ioc.Controller().Registry(&TokenAPIHandler{})
}

type TokenAPIHandler struct {
	ioc.ObjectImpl
	token token.Outer
}

func (t *TokenAPIHandler) Name() string {
	return "tokens"
}

// 重写ioc框架的Init方法实现服务注册，等于是在main中直接调用registry方法注册路由
func (t *TokenAPIHandler) Init() error {
	t.token = token.GetService()

	// 获取模块路由，带有url前缀
	r := ioc_gin.ObjectRouter(t)
	// 从ioc获取gin来实现api接口
	r.POST("/", t.IssueToken)
	r.Use(middleware.Auth)
	r.POST("/revolk", t.RevolkToken)
	r.POST("/refresh", t.RefreshToken)
	return nil
}

// 颁发token
func (t *TokenAPIHandler) IssueToken(ctx *gin.Context) {
	tk := token.NewIssueTokenRequest("", "")
	if err := ctx.BindJSON(tk); err != nil {
		response.Failed(ctx, err)
		return
	}
	// 传入httpRequest的cnotext
	ins, err := t.token.IssueToken(ctx.Request.Context(), *tk)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	// 获取cookie作用的域名信息，配合etc目录的toml文件的app表的address实现
	domain := application.Get().Domain()
	log.L().Debug().Msgf("cookie value=%v", domain)
	// 通过gin的setcookie方法得到cookie，内容就是用户携带的token
	// path表示cookie的作用路径，/表示对所有url都起作用
	// domain表示客户端访问指定域名的时候浏览器才会发送cookie
	// secure表示仅https
	ctx.SetCookie(token.CookieName, ins.AccessToken, ins.AccessTokenExpireTTL(), "/", domain, false, true)
	response.Success(ctx, ins)
}

// 注销token
func (t *TokenAPIHandler) RevolkToken(ctx *gin.Context) {
	tk := token.RevolkTokenRequest{}
	// BindJSON实际上内部通过json.Unmarshall反序列化，所以需要传入指针对象
	if err := ctx.BindJSON(&tk); err != nil {
		response.Failed(ctx, err)
		return
	}
	ins, err := t.token.RevolkToken(ctx.Request.Context(), &tk)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}

// 刷新accesstoken
func (t *TokenAPIHandler) RefreshToken(ctx *gin.Context) {
	tk := token.RefreshTokenRequest{}
	if err := ctx.BindJSON(tk); err != nil {
		response.Failed(ctx, err)
		return
	}
	ins, err := t.token.RefreshToken(ctx.Request.Context(), tk)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}
