package model

import (
	"gorm.io/gorm"
)

// User 用户模型
// todo 采用字段unique
type User struct {
	//Articles []Article `gorm:"foreignKey:Uid"` // 外键
	gorm.Model
	NickName string `gorm:"type:varchar(20);" label:"昵称" json:"nickName"`
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=1,max=40" label:"用户名"` //控制器的时候，绑定json，方便以后上下文引用和交互
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=40" label:"密码"`
	Desc     string `gorm:"type:varchar(40);"  label:"描述" json:"desc"`
	Website  string `gorm:"type:varchar(40);"  label:"个人主页" json:"site"`
	Email    string `gorm:"type:varchar(40);"  label:"邮箱" json:"email"`
	Avatar   string `gorm:"type:varchar(250);"  label:"头像" json:"avatar"`
	Role     int    `gorm:"type:int" json:"role" validate:"" label:"身份码"`
	Salt     string `gorm:"type:varchar(20);not null" json:"salt"`
}

// BeforeCreate 权限控制Hook
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	return nil
}
