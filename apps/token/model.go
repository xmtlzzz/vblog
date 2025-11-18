package token

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

// token是表结构
type Token struct {
	// 作为数据表的主键
	Id int `json:"id" gorm:"id;primaryKey"`
	// 用户id
	RefUserId int `json:"ref_user_id" gorm:"column:ref_user_id;unique;index"`
	// 颁发的Access token，token要唯一
	AccessToken string `json:"access_token" gorm:"column:access_token;unique;index"`
	// 颁发时间，不能为nil
	IssueAt time.Time `json:"issue_at" gorm:"column:issue_at"`
	// 指针，因为token第一次颁发没有过期可以为nil
	AccessTokenExpireAt *time.Time `json:"access_token_expire_at" gorm:"column:access_token_expire_at"`
	// 刷新token，解决user session问题模拟长连接，tokn要唯一
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token;unique;index"`
	// 同上因为第一次颁发可以不过期
	RefreshTokenExpireAt *time.Time `json:"refresh_token_expire_at" gorm:"column:refresh_token_expire_at"`
	// 不作为字段插入数据库，但是希望可以进行关联查询：通过RefUserId找到UserName并返回
	RefUserName string `json:"ref_user_name" gorm:"-"`
}

func GenNewToken(RefUserID int) *Token {
	aTokenExpireTime := time.Now().AddDate(0, 0, 1)
	rTokenExpireTime := time.Now().AddDate(0, 0, 7)
	return &Token{
		RefUserId:            RefUserID,
		AccessToken:          uuid.NewString(),
		IssueAt:              time.Now(),
		AccessTokenExpireAt:  &aTokenExpireTime,
		RefreshToken:         uuid.NewString(),
		RefreshTokenExpireAt: &rTokenExpireTime,
	}
}

func (t *RefreshTokenRequest) Validate() error {
	valid := validator.New()
	return valid.Struct(t)
}

// 有效期是否在过期时间之前
func (t *Token) IsAccessTkExpired() error {
	if time.Now().Before(*t.AccessTokenExpireAt) {
		return nil
	}
	return exception.NewBadRequest("AccessToken已过期")
}

func (t *Token) AccessTokenExpireTTL() int {
	// 如果过期时间为空，那么cookie持续时间为0
	if t.AccessTokenExpireAt == nil {
		return 0
	}
	// 如果不为空，那么就根据现在到过期时间还有多久返回
	return int(t.AccessTokenExpireAt.Sub(time.Now()).Seconds())
}

func (t *Token) IsRefreshTkExpired() error {
	if time.Now().Before(*t.RefreshTokenExpireAt) {
		return nil
	}
	return exception.NewBadRequest("RefreshToken已过期")
}

// 链式编程，实现RefUserName字段的设置
func (t *Token) SetRefUserName(refusername string) *Token {
	t.RefUserName = refusername
	return t
}

func (in *IssueTokenRequest) Validate() error {
	valid := validator.New()
	return valid.Struct(in)
}

func (in *Token) String() string {
	return pretty.ToJSON(in)
}

func (t *Token) TableName() string {
	return "tokens"
}
