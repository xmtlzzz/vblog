# utils模块
这个模块主要是编写了一些公共内容
```go
package utils
type ResourceMetadata struct {
	// 用户ID，唯一的定位到一个用户
	UserId string `json:"userid" gorm:"column:user_id;type:varchar(255);unique;index"`
	// 创建用户时间
	CreatedAt time.Time `json:"createAt";gorm:"column:created_at;not null"`
	// 被哪个用户创建
	CreateBy string `json:"createBy";gorm:"column:create_by"`
	// 更新用户时间
	UpdateAt time.Time `json:"updateAt";gorm:"column:update_at"`
}
```
User管理部分的元数据信息