package service

import (
	"starblog/model"
	"starblog/utils/errmsg"
)

// CreateArticle 新增文章
func CreateArticle(data *model.Article) int {
	err := model.DB.Create(data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// ShowArticles 查询文章列表
// todo 显示文章总数
func ShowArticles(pageSize int, pageNum int) []model.Article {
	var arti []model.Article
	err := model.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arti).Error //分页通用做法
	//如果err不为空，并且gorm的ErrRecordNotFound不为空，则异常，返回nil
	if err != nil {
		return nil
	}
	return arti
}

//todo 显示分类下所有文章

// todo 查看单个文章内容

// EditArticle 编辑文章
func EditArticle(id int, data *model.Article) int {
	var arti model.Article
	var maps = make(map[string]interface{})
	//判断文章是否存在
	err := model.DB.Select("id").Where("id = ?", id).First(&data).Error
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
