package utils

import "time"

type ResourceMetadata struct {
	// 用户ID，唯一的定位到一个用户
	Id int `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	// 创建用户时间
	CreatedAt time.Time `json:"createAt" gorm:"column:created_at;not null"`
	// 被哪个用户创建，value传入的就是创建blog的用户
	// 被哪个用户创建
	CreateBy string `json:"createBy" gorm:"column:create_by"`
	// 更新用户时间
	UpdateAt time.Time `json:"updateAt" gorm:"column:update_at"`
}

// 构造方法，用于User New方法初始化
func NewResourceMetadata() *ResourceMetadata {
	return &ResourceMetadata{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}

// 用于blog interface部分的请求参数，blogid为通用参数所以单独拿出来定义
type GetRequest struct {
	Id uint `json:"id"`
}

// 对外的构造函数
func NewGetRequest(id uint) *GetRequest {
	return &GetRequest{
		Id: id,
	}
}
