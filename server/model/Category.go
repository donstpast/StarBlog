package model

import "gorm.io/gorm"

// Category 分类模型
type Category struct {
	Universal
	gorm.Model
	Name     string    `gorm:"type:varchar(20);not null;comment:分类名称" json:"name"`
	Articles []Article `gorm:"foreignKey:Cid"` // 外键

}
