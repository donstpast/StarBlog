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
	if users.Username == "" || users.Password == "" {
		return errmsg.ERROR_USERNAME_OR_PASSWORD_IS_EMPTY
	}
	err := DB.Select("id").Where("username = ?", name).First(&users).Error
	if err != nil {
		return errmsg.ERROR
	}
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

// ShowUsers 查询用户列表
func ShowUsers(pageSize int, pageNum int) []User {
	var users []User
	result := DB.Limit(pageSize).Find(&users).Offset((pageNum - 1) * pageSize) //分页通用做法
	//如果err不为空，并且gorm的ErrRecordNotFound不为空，则异常，返回nil
	//&& result.Error != gorm.ErrRecordNotFound会报错，暂时不加，之后解决
	//todo 解决result.Error != gorm.ErrRecordNotFound问题
	if result.Error != nil {
		return nil
	}
	return users
}
