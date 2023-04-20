package service

import (
	"fmt"
	"gorm.io/gorm"
	"starblog/model"
	"starblog/utils/errmsg"
)

// CreateArticle 新增文章
func CreateArticle(data *model.Article) int {
	err := model.DB.Preload("Category").Preload("User").Create(data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// todo 显示文章总数

// ShowArticles 查询文章列表
func ShowArticles(title string, pageSize int, pageNum int) ([]model.Article, int, int64) {
	var err error
	var arti []model.Article
	var totalNum int64
	if title == "" {
		err = model.DB.Preload("Category").Preload("User").Model(&arti).Count(&totalNum).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arti).Error //分页通用做法
	} else {
		err = model.DB.Preload("Category").Preload("User").Model(&arti).Where("title LIKE ?", fmt.Sprintf("%%%s%%", title)).Count(&totalNum).Offset((pageNum - 1) * pageSize).Find(&arti).Error
	}
	//如果err不为空，并且gorm的ErrRecordNotFound不为空，则异常，返回nil
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return arti, errmsg.SUCCESS, totalNum
}

// ShowCategoryArticles 显示分类下所有文章
func ShowCategoryArticles(id int, pageSize int, pageNum int) ([]model.Article, int, int64) {
	var cateArti []model.Article
	var totalNum int64
	err := model.DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArti).Count(&totalNum).Error
	if err != nil {
		return nil, errmsg.ERROR_CATEGORY_NOT_EXIST, 0
	}
	return cateArti, errmsg.SUCCESS, totalNum
}

// ShowSingleArticle 查看单个文章内容
func ShowSingleArticle(id int) (model.Article, int) {
	var arti model.Article
	err := model.DB.Preload("Category").Where("id = ?", id).First(&arti).Error
	if err != nil {
		return arti, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return arti, errmsg.SUCCESS
}

// EditArticle 编辑文章
func EditArticle(id int, data *model.Article) int {
	var arti model.Article
	var maps = make(map[string]interface{})
	//判断文章是否存在
	err := model.DB.Preload("Category").Preload("User").Select("id").Where("id = ?", id).First(&data).Error
	if err != nil {
		return errmsg.ERROR_ARTICLE_NOT_EXIST //2001 文章不存在
	}
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["uid"] = data.Uid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = model.DB.Model(&arti).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

// DelArticle 删除文章
func DelArticle(id int) int {
	var arti model.Article
	//先判断分类是否存在
	err := model.DB.Select("id").Where("id = ?", id).First(&arti).Error
	if err != nil {
		return errmsg.ERROR_ARTICLE_NOT_EXIST //2001 文章不存在
	}
	err = model.DB.Where("id = ?", id).Delete(&arti).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
