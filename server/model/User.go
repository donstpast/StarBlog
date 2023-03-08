package model

import (
	"gorm.io/gorm"
	"starblog/utils/errmsg"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"` //控制器的时候，绑定json，方便以后上下文引用和交互
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// services层

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	DB.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1002 用户名已被使用
	}
	return errmsg.SUCCESS //200
}

// CreateUser  新增用户
func CreateUser(data *User) int {
	err := DB.Create(data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var err error
	var users []User
	err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}
