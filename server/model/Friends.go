package model

// Friends 友链模型
type Friends struct {
	ID     int    `gorm:"primary_key;not null;auto_increment" json:"id"`
	Name   string `gorm:"type:varchar(20);not null;comment:友链名称" json:"name"`
	Link   string `gorm:"type:varchar(100);not null;comment:网址" json:"link"`
	Avatar string `gorm:"type:varchar(100);not null;comment:头像" json:"avatar"`
}
