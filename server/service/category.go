package service

import (
	"starblog/model"
	"starblog/utils/errmsg"
)

// CheckCategory 查询分类名是否存在
func CheckCategory(name string) (code int) {
	var cate model.Category
	_ = model.DB.Select("id", "name").Where("name = ?", name).First(&cate).Error
	if cate.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED //3002 分类已存在
	}
	return errmsg.SUCCESS //200
}

// CreateCategory 新增分类
func CreateCategory(data *model.Category) int {
	err := model.DB.Create(data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// ShowCategories 查询分类列表
func ShowCategories(pageSize int, pageNum int) ([]model.Category, int64) {
	var cate []model.Category
	var totalNum int64
	err := model.DB.Model(&cate).Count(&totalNum).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error //分页通用做法
	//如果err不为空，并且gorm的ErrRecordNotFound不为空，则异常，返回nil
	if err != nil {
		return nil, 0
	}
	return cate, totalNum
}

// ShowCategoryInfo 查询单个分类信息
func ShowCategoryInfo(id int) (model.Category, int) {
	var category model.Category
	err := model.DB.Where("id = ?", id).First(&category).Error
	if err != nil {
		return category, errmsg.ERROR_CATEGORY_NOT_EXIST
	}
	return category, errmsg.SUCCESS
}

// EditCategory 编辑分类
func EditCategory(id int, data *model.Category) int {
	var cate model.Category
	var maps = make(map[string]interface{})
	//判断分类是否存在
	err := model.DB.Select("id").Where("id = ?", id).First(&cate).Error
	if err != nil {
		return errmsg.ERROR_CATEGORY_NOT_EXIST //3001 分类不存在
	}
	//判断用户是否传入空值
	if data.Name == "" {
		return errmsg.ERROR_CATEGORY_IS_EMPTY //3003 分类名不能为空
	} else {
		maps["Name"] = data.Name
		err = model.DB.Model(&cate).Where("id = ?", id).Updates(maps).Error
		if err != nil {
			return errmsg.ERROR
		}
		return errmsg.SUCCESS
	}

}

// DelCategory 删除分类
func DelCategory(id int) int {
	var cate model.Category
	//先判断分类是否存在
	err := model.DB.Select("id").Where("id = ?", id).First(&cate).Error
	if err != nil {
		return errmsg.ERROR_CATEGORY_NOT_EXIST //3001 分类不存在
	}
	err = model.DB.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
