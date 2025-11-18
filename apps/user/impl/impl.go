package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/xmtlzzz/vblog/apps/user"
	"github.com/xmtlzzz/vblog/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var UserService user.Service = &UserServiceImpl{}
var db *gorm.DB

type UserServiceImpl struct {
}

func (u *UserServiceImpl) UpdateUserStatus(ctx context.Context, request user.UpdateUserStatusRequest) (*user.User, error) {
	var ru = &user.User{}
	db = utils.NewDBConnecter()
	if err := db.Table("users").WithContext(ctx).Where("username = ?", request.Username).Updates(&request).Error; err != nil {
		return nil, exception.NewBadRequest("找不到用户无法更新状态，错误内容为: %v", err)
	}
	if err := db.WithContext(ctx).Where("username = ?", request.Username).Take(&ru).Error; err != nil {
		return nil, exception.NewBadRequest("找不到用户，错误内容为: %v", err)
	}
	return ru, nil
}

func (u *UserServiceImpl) DescribeUser(ctx context.Context, request user.DescribeUserRequest) (*user.User, error) {
	var ru = user.User{}
	db = utils.NewDBConnecter()
	switch request.DescribeBy {
	case user.Describe_By_UserId:
		db.WithContext(ctx).Where("user_id = ?", request.Value).Take(&ru)
	case user.Describe_By_UserName:
		db.WithContext(ctx).Where("username = ?", request.Value).Take(&ru)
	}
	return &ru, nil
}

func (u *UserServiceImpl) Registry(ctx context.Context, request *user.RegistryRequest) (*user.User, error) {
	ins, err := user.New(request)
	if err != nil {
		return nil, err
	}
	t, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	ins.UpdateAt = t
	ins.BlockAT = t
	// bcrypt实现hash
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	ins.Password = string(hashPassword)
	db = utils.NewDBConnecter()
	if err := db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (u *UserServiceImpl) UpdateUserPassword(ctx context.Context, request user.UpdatePasswordRequest) error {
	var ru = &user.User{}
	db = utils.NewDBConnecter()
	if err := db.WithContext(ctx).Where("username = ?", request.Username).Take(ru).Error; err != nil {
		return exception.NewBadRequest("找不到用户无法更新密码，错误内容为: %v", err)
	}
	fmt.Println(request.OldPassword)
	fmt.Println(ru.Password)
	if err := ru.CheckPassword([]byte(request.OldPassword)); err != nil {
		return exception.NewBadRequest("oldPassword错误，具体错误内容为: %v", err)
	}
	newPass, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return exception.NewBadRequest("newpassword加密失败，错误内容为: %v", err)
	}
	ru.Password = string(newPass)
	if err := db.Table("users").WithContext(ctx).Where("username = ?", request.Username).Updates(&ru).Error; err != nil {
		return exception.NewBadRequest("用户存在但无法更新密码，错误内容为: %v", err)
	}
	return nil
}

func (u *UserServiceImpl) UpdateUserProfile(ctx context.Context, request user.UpdateProfileRequest) (*user.User, error) {
	var ru = &user.User{}
	db = utils.NewDBConnecter()
	if err := db.Table("users").WithContext(ctx).Where("username = ?", request.Username).Updates(&request).Error; err != nil {
		return nil, exception.NewBadRequest("找不到用户无法更新Profile，错误内容为: %v", err)
	}
	if err := db.WithContext(ctx).Where("username = ?", request.Username).Take(&ru).Error; err != nil {
		return nil, exception.NewBadRequest("找不到用户，错误内容为: %v", err)
	}
	return ru, nil
}

func (u *UserServiceImpl) ResetUserPassword(ctx context.Context, request user.ResetPasswordRequest) error {
	var ru = &user.User{}
	if err := request.Validate(); err != nil {
		return exception.NewBadRequest("newpassword为空，无法重置密码，错误内容为: %v", err)
	}
	newPass, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return exception.NewBadRequest("newpassword加密失败，错误内容为: %v", err)
	}
	ru.Password = string(newPass)
	db = utils.NewDBConnecter()
	if err := db.WithContext(ctx).Where("username = ?", request.Username).Updates(&ru).Error; err != nil {
		return exception.NewBadRequest("用户不存在无法重置密码，错误内容为: %v", err)
	}
	return nil
}

func (u *UserServiceImpl) UnRegistry(ctx context.Context, request user.UnRegistryRequest) error {
	var ru = &user.User{}
	db = utils.NewDBConnecter()
	if err := db.WithContext(ctx).Where("username = ?", request.Username).Delete(ru).Error; err != nil {
		return exception.NewBadRequest("需要注销额度用户不存在，错误内容为: %v", err)
	}
	return nil
}
