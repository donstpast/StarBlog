package service

import (
	"fmt"
	"gorm.io/gorm"
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

// UpdateCheckUser 查询要更新的用户是否存在
func UpdateCheckUser(id int, name string) (code int) {
	var users model.User
	model.DB.Select("id", "username").Where("username = ?", name).First(&users)
	if users.ID == uint(id) {
		return errmsg.SUCCESS
	} else {
		if users.ID > 0 {
			return errmsg.ERROR_USERNAME_USED //1002 用户名已被使用
		}
		return errmsg.SUCCESS //200
	}

}

// CreateUser  新增用户
func CreateUser(data *model.User) int {
	data.Password, data.Salt = crypt.PwScrypt(data.Password)
	err := model.DB.Create(data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// ShowUsers 查询用户列表
func ShowUsers(userName string, pageSize int, pageNum int) ([]model.User, int64) {
	var err error
	var users []model.User
	var totalNum int64
	if userName == "" {
		err = model.DB.Model(&users).Count(&totalNum).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users).Error //分页通用做法
	} else {
		err = model.DB.Model(&users).Where("username LIKE ?", fmt.Sprintf("%%%s%%", userName)).Count(&totalNum).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users).Error
		//model.DB.Model(&users).Where("username LIKE ?", fmt.Sprintf("%%%s%%", userName)).Count(&totalNum)
	}
	//如果err不为空，并且gorm的ErrRecordNotFound不为空，则异常，返回nil
	//&& result.Error != gorm.ErrRecordNotFound会报错，暂时不加，之后解决
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, totalNum
}

// ShowUserInfo 查询单个用户信息
func ShowUserInfo(id int) (model.User, int) {
	var user model.User
	err := model.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	return user, errmsg.SUCCESS
}

// ShowUserID 查询用户ID
func ShowUserID(username string) (uint, int) {
	var user model.User
	err := model.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return 0, errmsg.ERROR_USER_NOT_EXIST
	}
	return user.ID, errmsg.SUCCESS
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
	if data == nil {
		return errmsg.ERROR
	} else if data.Username == "" {
		return errmsg.ERROR_USERNAME_OR_PASSWORD_IS_EMPTY
	} else if data.Role == 0 {
		return errmsg.ERROE_ROLE_WRONG
	} else if data.Password == "" {
		maps["username"] = data.Username
		maps["role"] = data.Role
	} else {
		maps["username"] = data.Username
		maps["role"] = data.Role
		maps["password"], maps["salt"] = crypt.PwScrypt(data.Password)
	}
	err = model.DB.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
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

// CheckLogin 登录验证
func CheckLogin(username string, password string) (model.User, int) {
	var user model.User
	model.DB.Where("username = ?", username).First(&user)
	//验证密码是否正确
	isEqual := crypt.VerifyPassword(password, user.Password, user.Salt)
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if !isEqual {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NOT_ADMIN
	}
	return user, errmsg.SUCCESS
}
