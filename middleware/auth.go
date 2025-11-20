package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xmtlzzz/vblog/apps/token"
	"github.com/xmtlzzz/vblog/apps/token/impl"
	"github.com/xmtlzzz/vblog/response"
)

// 自定义中间件返回token以及cookie
func Auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	// 基于bear token的形式，得到存储bear和token的切片
	tkList := strings.Split(authHeader, " ")
	accessToken := ""
	if len(tkList) == 2 {
		accessToken = tkList[1]
	}

	// 如果accessToken为空就提取cookie作为accesstoken
	if accessToken == "" {
		// 提取gin.Context对象携带的cookie信息
		ck, err := c.Cookie(token.CookieName)
		if err != nil {
			response.Failed(c, err)
			//  add
			c.Abort()
			return
		} else {
			accessToken = ck
		}
	}

	tk, err := impl.TokenService.ValidateToken(c.Request.Context(), token.NewValidateTokenRequest(accessToken))
	if err != nil {
		response.Failed(c, err)
		// context对象直接退出
		c.Abort()
		return
	}
	// 声明新的context对象携带token信息，用于后续其他模块调用验证token
	// 这里key直接使用字符串不行，所以自定义一个结构体
	ctx := context.WithValue(c.Request.Context(), TokenCtxKey{}, tk)
	// 携带ctx信息
	c.Request = c.Request.WithContext(ctx)
	// 鉴权成功后继续后续处理
	c.Next()
	// 继续后续处理链
	c.Next()
}

type TokenCtxKey struct {
	context.Context
}

func GetTokenFromCtx(ctx context.Context) *token.Token {
	// 后续外部模块调用GetTokenFromCtx方法就获取到了http头部中携带的token信息，
	return ctx.Value(TokenCtxKey{}).(*token.Token)
}
