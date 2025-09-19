package impl

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/xmtlzzz/vblog/apps/token"
	"github.com/xmtlzzz/vblog/apps/user"
	"github.com/xmtlzzz/vblog/apps/user/impl"
	"github.com/xmtlzzz/vblog/utils"
	"gorm.io/gorm"
)

func init() {
	// 注册到ioc池子
	ioc.Controller().Registry(&TokenServiceImpl{})
}

var TokenService token.Service = &TokenServiceImpl{}
var db *gorm.DB

type TokenServiceImpl struct {
	ioc.ObjectImpl
	// 调用DescribeUser查询用户
	UserSvc user.AdminService
}

func (*TokenServiceImpl) Name() string {
	return token.AppName
}

// 属性依赖通过外部传入，并且返回接口更加容易扩展
func NewTokenServiceImpl(user user.AdminService) *TokenServiceImpl {
	return &TokenServiceImpl{
		UserSvc: user,
	}
}

func (t *TokenServiceImpl) IssueToken(ctx context.Context, request token.IssueTokenRequest) (*token.Token, error) {
	if err := request.Validate(); err != nil {
		return nil, exception.NewBadRequest("%v", err)
	}
	// 1. 找到用户
	// 这里一定要初始化属性，否则就是nil pointer，UserService实现了user.AdminService接口
	t.UserSvc = impl.UserService
	ins, err := t.UserSvc.DescribeUser(ctx, user.DescribeUserRequest{
		DescribeBy: user.Describe_By_UserName,
		Value:      request.Username,
	})
	fmt.Println(ins)
	if err != nil {
		return nil, exception.NewBadRequest("%v", err)
	}
	// 2. 对比密码，传入的密码是申请token的密码，去和查询得到的hash比对
	if err := ins.CheckPassword([]byte(request.Password)); err != nil {
		return nil, exception.NewBadRequest("密码不对应")
	}
	// 3. 颁发token，使用查询到的User实例的UserId填充token表，两表逻辑对应
	// 链式编程实现不能多个返回值必须返回*Token类型，设置RefUserName，两表逻辑对应
	tk := token.GenNewToken(ins.Id).SetRefUserName(ins.Username)
	db = utils.NewDBConnecter()
	if err := db.WithContext(ctx).Create(tk).Error; err != nil {
		return nil, exception.NewBadRequest("key冲突，用户token已存在")
	}
	return tk, nil
}

// 撤销token
func (t *TokenServiceImpl) RevolkToken(ctx context.Context, request token.RevolkTokenRequest) (*token.Token, error) {
	var tk = &token.Token{}
	db = utils.NewDBConnecter()
	if err := db.WithContext(ctx).Where("access_token = ? and refresh_token = ?", request.AccessToken, request.RefreshToken).Delete(tk).Error; err != nil {
		return nil, exception.NewBadRequest("传入的AccessToken和RefreshToken错误")
	}
	return tk, nil
}

func (t *TokenServiceImpl) ValidateToken(ctx context.Context, request *token.ValidateTokenRequest) (*token.Token, error) {
	// 1.确定token是否是我们颁发的
	// 根据传入的AccessToken查表有的话就是颁发的
	var tk = &token.Token{}
	if err := utils.NewDBConnecter().WithContext(ctx).Where("access_token = ?", request.AccessToken).Take(tk).Error; err != nil {
		return nil, exception.NewBadRequest("携带的token非本地颁发（不存在）,err: %v", err)
	}
	// 2.确定token是否在有效期内
	if err := tk.IsAccessTkExpired(); err != nil {
		return nil, err
	}
	return tk, nil
}

// 判断refreshtoken是否有效，有效则调用刷新token
func (t *TokenServiceImpl) RefreshToken(ctx context.Context, request token.RefreshTokenRequest) (*token.Token, error) {
	var tk = &token.Token{}
	db = utils.NewDBConnecter()
	if err := request.Validate(); err != nil {
		return nil, exception.NewBadRequest("RefreshToken未指定，具体报错为:%v", err)
	}
	if err := db.WithContext(ctx).Where("refresh_token = ?", request.RefreshToken).Take(tk).Error; err != nil {
		return nil, exception.NewBadRequest("RefreshToken不存在，提供正确的token查询，具体报错为:%v", err)
	}
	if err := tk.IsRefreshTkExpired(); err != nil {
		return nil, exception.NewBadRequest("RefreshToken已经过期，请重新申请token，具体报错为:%v", err)
	}
	// 若refreshtoken在有效期内，则刷新accesstoken
	tk.AccessToken = uuid.NewString()
	if err := db.WithContext(ctx).Where("ref_user_id = ?", tk.RefUserId).Updates(tk).Error; err != nil {
		return nil, err
	}
	return tk, nil
}
