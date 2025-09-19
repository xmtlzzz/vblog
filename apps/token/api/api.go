package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/http/gin/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/xmtlzzz/vblog/apps/token"
	"github.com/xmtlzzz/vblog/server"
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
	t.Registry(server.GinServer)
	return nil
}

//func NewBTokenAPIHandler(tokenService *impl.TokenServiceImpl) *TokenAPIHandler {
//	return &TokenAPIHandler{
//		token: tokenService,
//	}
//}

func (t *TokenAPIHandler) Registry(ge *gin.Engine) {
	server := ge.Group("/vblog/api/v1/tokens")
	server.POST("/", t.IssueToken)
	server.POST("/revolk", t.RevolkToken)
	server.POST("/refresh", t.RefreshToken)
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
	response.Success(ctx, ins)
}

// 注销token
func (t *TokenAPIHandler) RevolkToken(ctx *gin.Context) {
	tk := token.RevolkTokenRequest{}
	if err := ctx.BindJSON(tk); err != nil {
		response.Failed(ctx, err)
		return
	}
	ins, err := t.token.RevolkToken(ctx.Request.Context(), tk)
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
