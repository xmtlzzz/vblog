package user

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/xmtlzzz/vblog/utils"
	"golang.org/x/crypto/bcrypt"
)

// 用户的Profile信息

type User struct {
	// 调用utils包内定义的元数据
	utils.ResourceMetadata
	// 继承注册基本属性
	RegistryRequest
}

// 调用mcube封装的json格式化
func (u *User) String() string {
	return pretty.ToJSON(u)
}

// User类型的构造函数
func New(in *RegistryRequest) (*User, error) {
	if err := in.Validate(); err != nil {
		// 调用mcube封装的逻辑
		return nil, exception.NewBadRequest("参数校验错误: %s", err)
	}
	return &User{
		ResourceMetadata: *utils.NewResourceMetadata(),
		RegistryRequest:  *in,
	}, nil
}

func (in *RegistryRequest) CheckPassword(password []byte) error {
	// 这里传入的password是明文password
	return bcrypt.CompareHashAndPassword([]byte(in.Password), password)
}

func (in *RegistryRequest) Validate() error {
	valid := validator.New()
	return valid.Struct(in)
}

func (in *ResetPasswordRequest) Validate() error {
	valid := validator.New()
	return valid.Struct(in)
}

type Status struct {
	//	用户什么时候处于禁用状态
	BlockAT time.Time `json:"blockAT" gorm:"column:block_at"`
	// 禁用的原因
	BlockReason string `json:"blockReason" gorm:"column:block_reason"`
}

type Profile struct {
	// 头像
	Avatar string `json:"avatar" gorm:"column:avatar;type:varchar(100)"`
	// 用户昵称
	NickName string `json:"nickname" gorm:"column:nick_name;type:varchar(255)"`
	// 邮箱
	Email string `json:"email" gorm:"column:email;type:varchar(255)"`
}

// 用户注册模型，因为*User继承RegistryRequest所以把User该有的都放在这里，避免反复调用
type RegistryRequest struct {
	// 用户名
	Username string `json:"username" gorm:"column:username;type:varchar(255);unique;index" validate:"required"`
	// 密码
	Password string `json:"password" gorm:"column:password;type:varchar(255)"  validate:"required"`
	// 用户基本模板
	Profile
	// 用户状态
	Status
}

func NewRegistryRequest() *RegistryRequest {
	return &RegistryRequest{}
}

func (*User) TableName() string {
	return "users"
}
