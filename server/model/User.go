package model

import (
	"gorm.io/gorm"
	"starblog/utils/errmsg"
)

// User 用户模型
type User struct {
	Universal
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"` //控制器的时候，绑定json，方便以后上下文引用和交互
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 用控制器控制模型的增删改查
// services层

// CreateUser  新增用户
func CreateUser(data *User) int {
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}
