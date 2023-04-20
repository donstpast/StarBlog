package model

import "gorm.io/gorm"

// Article 文章模型
// 目前逻辑：
// belongTo: 一个文章 属于 一个分类
// belongTo: 一个文章 属于 一个用户
type Article struct {
	Category Category `gorm:"foreignKey:Cid;references:ID"`
	User     User     `gorm:"foreignKey:Uid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null;comment:文章标题" json:"title"`
	Cid     int    `gorm:"type:int;not null;comment:分类 ID" json:"cid"`
	Uid     int    `gorm:"type:int;not null;comment:用户_ID" json:"uid"`
	Desc    string `gorm:"type:varchar(200);comment:文章描述" json:"desc"`
	Content string `gorm:"type:longtext;comment:文章内容" json:"content"`
	Img     string `gorm:"type:varchar(100);comment:封面图片地址" json:"img"`
}
