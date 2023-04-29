package service

import (
	"starblog/model"
	"starblog/utils/errmsg"
)

// CheckTag 查询标签是否存在
func CheckTag(name string) (code int) {
	var tag model.Tag
	_ = model.DB.Select("id", "name").Where("name = ?", name).First(&tag).Error
	if tag.ID > 0 {
		return errmsg.ERROR_TAG_USED //4002 标签已存在
	}
	return errmsg.SUCCESS //200
}

// CreateTag 新增标签
func CreateTag(data *model.Tag) int {
	err := model.DB.Create(data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// ShowTags 查询标签列表
func ShowTags(pageSize int, pageNum int) ([]model.Tag, int64) {
	var tag []model.Tag
	var totalNum int64
	err := model.DB.Model(&tag).Preload("Articles").Count(&totalNum).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&tag).Error //分页通用做法
	//如果err不为空，并且gorm的ErrRecordNotFound不为空，则异常，返回nil
	if err != nil {
		return nil, 0
	}
	return tag, totalNum
}

// ShowTagInfo 查询单个标签信息
func ShowTagInfo(id int) (model.Tag, int) {
	var tag model.Tag
	err := model.DB.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return tag, errmsg.ERROR_TAG_NOT_EXIST
	}
	return tag, errmsg.SUCCESS
}

// EditTag 编辑标签
func EditTag(id int, data *model.Tag) int {
	var tag model.Tag
	var maps = make(map[string]interface{})
	//判断标签是否存在
	err := model.DB.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil {
		return errmsg.ERROR_TAG_NOT_EXIST //4001 标签不存在
	}
	//判断用户是否传入空值
	if data.Name == "" {
		return errmsg.ERROR_TAG_IS_EMPTY //4003 标签名不能为空
	} else {
		maps["Name"] = data.Name
		err = model.DB.Model(&tag).Where("id = ?", id).Updates(maps).Error
		if err != nil {
			return errmsg.ERROR
		}
		return errmsg.SUCCESS
	}

}

// DelTag 删除标签
func DelTag(id int) int {
	var tag model.Tag
	//先判断标签是否存在
	err := model.DB.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil {
		return errmsg.ERROR_TAG_NOT_EXIST //4001 标签不存在
	}
	err = model.DB.Where("id = ?", id).Delete(&tag).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
