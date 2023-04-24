package service

import (
	"starblog/model"
	"starblog/utils/errmsg"
)

// CreateFriends 新增友链
func CreateFriends(data *model.Friends) int {
	err := model.DB.Create(data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// ShowFriends 查询友链列表
func ShowFriends(pageSize int, pageNum int) ([]model.Friends, int64) {
	var friends []model.Friends
	var totalNum int64
	err := model.DB.Model(&friends).Count(&totalNum).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&friends).Error //分页通用做法
	//如果err不为空，并且gorm的ErrRecordNotFound不为空，则异常，返回nil
	if err != nil {
		return nil, 0
	}
	return friends, totalNum
}

// EditFriends 编辑友链
func EditFriends(id int, data *model.Friends) int {
	var friends model.Friends
	var maps = make(map[string]interface{})
	//判断友链是否存在
	err := model.DB.Select("id").Where("id = ?", id).First(&friends).Error
	if err != nil {
		return errmsg.ERROR_FRIENDS_NOT_EXIST //6001 友链不存在
	}
	//判断用户是否传入空值
	if data.Link == "" {
		return errmsg.ERROR_FRIENDS_LINKS_IS_EMPTY //6002 友链为空
	} else {
		maps["name"] = data.Name
		maps["link"] = data.Link
		maps["avatar"] = data.Avatar
		err = model.DB.Model(&friends).Where("id = ?", id).Updates(maps).Error
		if err != nil {
			return errmsg.ERROR
		}
		return errmsg.SUCCESS
	}

}

// DelFriends 删除友链
func DelFriends(id int) int {
	var friends model.Friends
	//先判断友链是否存在
	err := model.DB.Select("id").Where("id = ?", id).First(&friends).Error
	if err != nil {
		return errmsg.ERROR_FRIENDS_NOT_EXIST //6001 友链不存在
	}
	err = model.DB.Where("id = ?", id).Delete(&friends).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
