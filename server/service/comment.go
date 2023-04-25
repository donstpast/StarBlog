package service

import (
	"starblog/model"
	"starblog/utils/errmsg"
)

// CreateComment 新增评论
func CreateComment(data *model.Comment) int {
	err := model.DB.Create(data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// ShowComments 查询评论列表
func ShowComments(pageSize int, pageNum int) ([]model.Comment, int64) {
	var comment []model.Comment
	var totalNum int64
	err := model.DB.Model(&comment).Order("Updated_At DESC").Count(&totalNum).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&comment).Error //分页通用做法
	//如果err不为空，并且gorm的ErrRecordNotFound不为空，则异常，返回nil
	if err != nil {
		return nil, 0
	}
	return comment, totalNum
}

// EditComment 编辑评论
func EditComment(id int, data *model.Comment) int {
	var comment model.Comment
	var maps = make(map[string]interface{})
	//判断评论是否存在
	err := model.DB.Select("id").Where("id = ?", id).First(&comment).Error
	if err != nil {
		return errmsg.ERROR_COMMENT_NOT_EXIST //5001 评论不存在
	}
	//判断用户是否传入空值
	if data.Content == "" {
		return errmsg.ERROR_COMMENT_IS_EMPTY //5002 评论不能为空
	} else if data.UserEmail == "" {
		return errmsg.ERROR_COMMENT_EMAIL_IS_EMPTY
	} else if data.NickName == "" {
		return errmsg.ERROR_COMMENT_NICKNAME_IS_EMPTY
	} else {
		maps["Content"] = data.Content
		maps["NickName"] = data.NickName
		maps["UserEmail"] = data.UserEmail
		err = model.DB.Model(&comment).Where("id = ?", id).Updates(maps).Error
		if err != nil {
			return errmsg.ERROR
		}
		return errmsg.SUCCESS
	}

}

// DelComment 删除评论
func DelComment(id int) int {
	var comment model.Comment
	//先判断分类是否存在
	err := model.DB.Select("id").Where("id = ?", id).First(&comment).Error
	if err != nil {
		return errmsg.ERROR_COMMENT_NOT_EXIST //5001 评论不存在
	}
	err = model.DB.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
