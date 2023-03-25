package service

import (
	"starblog/model"
	"starblog/utils/crypt"
	"starblog/utils/errmsg"
)

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var users model.User
	model.DB.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1002 用户名已被使用
	}
	return errmsg.SUCCESS //200
}

// CreateUser  新增用户
func CreateUser(data *model.User) int {
	data.Password = crypt.PwScrypt(data.Password)
	err := model.DB.Create(data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// ShowUsers 查询用户列表
func ShowUsers(pageSize int, pageNum int) []model.User {
	var users []model.User
	err := model.DB.Limit(pageSize).Find(&users).Offset((pageNum - 1) * pageSize).Error //分页通用做法
	//如果err不为空，并且gorm的ErrRecordNotFound不为空，则异常，返回nil
	//&& result.Error != gorm.ErrRecordNotFound会报错，暂时不加，之后解决
	if err != nil {
		return nil
	}
	return users
}

// EditUser 编辑用户
func EditUser(id int, data *model.User) int {
	var user model.User
	var maps = make(map[string]interface{})
	//判断用户是否存在
	err := model.DB.Select("id").Where("id = ?", id).First(&data).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	//判断用户是否传入空值
	if data.Username == "" {
		return errmsg.ERROR_USERNAME_OR_PASSWORD_IS_EMPTY
	} else {
		maps["username"] = data.Username
		err = model.DB.Model(&user).Where("id = ?", id).Updates(maps).Error
		if err != nil {
			return errmsg.ERROR
		}
		return errmsg.SUCCESS
	}

}

// DelUser 删除用户
func DelUser(id int) int {
	var user model.User
	//先判断用户是否存在
	err := model.DB.Select("deleted_at").Where("id = ?", id).First(&user).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	err = model.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
