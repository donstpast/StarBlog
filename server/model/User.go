package model

import (
	"gorm.io/gorm"
)

// User 用户模型
// todo 采用字段unique
type User struct {
	gorm.Model
	Articles []Article `gorm:"foreignKey:Uid"`
	Username string    `gorm:"type:varchar(20);not null" json:"username"` //控制器的时候，绑定json，方便以后上下文引用和交互
	Password string    `gorm:"type:varchar(20);not null" json:"password"`
	Role     int       `gorm:"type:int" json:"role"`
	Salt     string    `gorm:"type:varchar(20);not null" json:"salt"`
}
