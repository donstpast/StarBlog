package model

import "gorm.io/gorm"

// Comment 评论模型
// hasMany: 一篇文章可以有多个评论，一个用户可以有多个评论
type Comment struct {
	gorm.Model
	Content   string `gorm:"not null;" json:"comment_content"`
	UserEmail string `gorm:"type:varchar(20);not null;" json:"email" label:"邮箱" json:"userEmail"`
	ArticleID int    `gorm:"type:int;not null;comment:文章ID" json:"articleID"`
}
