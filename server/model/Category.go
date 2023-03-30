package model

// Category 分类模型
// hasMany: 一个分类下可以有多篇文章
type Category struct {
	ID   uint   `gorm:"primary_key;not null;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null;comment:分类名称" json:"name"`
	//Articles []Article `gorm:"foreignKey:Cid"` // 外键

}
