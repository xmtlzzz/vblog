package token

import (
	"context"

	"github.com/infraboard/mcube/v2/ioc"
)

const (
	AppName = "token"
)

// token模块调用实现ioc资源池对象获取
func GetService() Service {
	// 断言实现
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	Outer
	Inner
}

// 1. 外部
type Outer interface {
	// + 登录、颁发令牌
	IssueToken(context.Context, IssueTokenRequest) (*Token, error)

	// + 退出、撤销令牌
	RevolkToken(context.Context, RevolkTokenRequest) (*Token, error)

	// refreshtoken有效期内刷新accesstoken
	RefreshToken(context.Context, RefreshTokenRequest) (*Token, error)
}

// 2. 内部
type Inner interface {
	// + 令牌校验方法
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

// 产生token的必要信息，不需要写入数据库
type IssueTokenRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	// 记住我，token过期可能是1天，浏览器在期间是否记住token信息
	RememberMe bool `json:"remember_me"`
}

func NewIssueTokenRequest(username, password string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
		Password: password,
	}
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func NewRefreshTokenRequest(refreshtoken string) *RefreshTokenRequest {
	return &RefreshTokenRequest{
		RefreshToken: refreshtoken,
	}
}

type RevolkTokenRequest struct {
	// 取消token的时候需要传入一对token来实现校验，避免用户得到AccessToken就直接注销
	AccessToken string `json:"access_token"`
	// Refreshtoken一般不返回
	RefreshToken string `json:"refresh_token"`
}

func NewRevolkTokenRequest(accesstoken, refreshtoken string) *RevolkTokenRequest {
	return &RevolkTokenRequest{
		AccessToken:  accesstoken,
		RefreshToken: refreshtoken,
	}
}

// 正常情况下用户只会达到一个AccessToken
type ValidateTokenRequest struct {
	AccessToken string `json:"access_token"`
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: accessToken,
	}
}
