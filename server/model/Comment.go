package model

import "gorm.io/gorm"

// Comment 评论模型
// hasMany: 一篇文章可以有多个评论，一个用户可以有多个评论
type Comment struct {
	gorm.Model
	Content   string `gorm:"not null;" json:"comment_content"`
	Uid       int    `gorm:"type:int;comment:用户ID" json:"uid"`
	UserEmail string `gorm:"type:varchar(20);not null;" label:"邮箱" json:"userEmail"`
	ArticleID int    `gorm:"type:int;not null;comment:文章ID" json:"articleID"`
}
