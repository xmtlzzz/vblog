package user

import (
	"context"
)

type Service interface {
	AdminService
	UserService
}

// 管理员接口
type AdminService interface {
	UpdateUserStatus(context.Context, UpdateUserStatusRequest) (*User, error)
	// 查询用户信息
	DescribeUser(context.Context, DescribeUserRequest) (*User, error)
}

// 调用枚举实现查询指标
type DescribeUserRequest struct {
	// 根据枚举指定查询基于userid还是username
	DescribeBy DESCRIBE_BY `json:"describe_by"`
	// DescribeBy字段的value通过value属性携带
	Value string `json:"value"`
}

// 普通用户接口
type UserService interface {
	//+ 注册
	Registry(context.Context, *RegistryRequest) (*User, error)
	// 更新用户密码
	UpdateUserPassword(context.Context, UpdatePasswordRequest) error
	// 更新用户模板如email、电话等
	UpdateUserProfile(context.Context, UpdateProfileRequest) (*User, error)
	// 重置用户密码
	ResetUserPassword(context.Context, ResetPasswordRequest) error
	//+ 注销
	UnRegistry(context.Context, UnRegistryRequest) error
}

// 根据UserId去修改用户的Status
type UpdateUserStatusRequest struct {
	// UserId string `json:"userid"`
	Username string `json:"username"`
	Status
}

type UpdatePasswordRequest struct {
	// 用户名
	Username string `json:"username" gorm:"column:username"`
	// 老密码
	OldPassword string `json:"oldPassword" gorm:"column:old_password"`
	// 新密码
	NewPassword string `json:"newPassword" gorm:"column:new_password"`
}

type ResetPasswordRequest struct {
	// 用户名
	Username string `json:"username" gorm:"column:username" `
	// 新密码
	NewPassword string `json:"newpassword" gorm:"column:new_password" validate:"required"`
	// 验证码
	VerifyCode string `json:"verifycode" gorm:"column:verify_code"`
}

type UpdateProfileRequest struct {
	// 用username查询更加合理，用户知道用户名
	// UserId string `json:"userid"`
	Username string `json:"username"`
	Profile
}

type UnRegistryRequest struct {
	// 根据用户名去注销
	Username string `json:"username"`
}
