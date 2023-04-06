package model

import (
	"gorm.io/gorm"
)

// User 用户模型
// todo 采用字段unique
type User struct {
	gorm.Model
	Articles []Article `gorm:"foreignKey:Uid"`
	Username string    `gorm:"type:varchar(20);not null" json:"username" validate:"required" label:"用户名"` //控制器的时候，绑定json，方便以后上下文引用和交互
	Password string    `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=40" label:"密码"`
	Role     int       `gorm:"type:int" json:"role" validate:"" label:"身份码"`
	Salt     string    `gorm:"type:varchar(20);not null" json:"salt"`
}

// BeforeCreate 权限控制Hook
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Role = 2
	return nil
}
