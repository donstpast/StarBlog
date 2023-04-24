package model

// Tag 分类模型
// hasMany: 一个分类下可以有多篇文章
type Tag struct {
	Articles []*Article `gorm:"many2many:article_tags;"`
	ID       uint       `gorm:"primary_key;not null;auto_increment" json:"id"`
	Name     string     `gorm:"type:varchar(20);not null;comment:分类名称" json:"name"`
}
