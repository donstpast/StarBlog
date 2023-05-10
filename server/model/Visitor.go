package model

import "time"

// Visitor 访客模型
type Visitor struct {
	ID          uint   `gorm:"primary_key;not null;auto_increment" json:"id"`
	RequestPath string `gorm:"type:varchar(80);comment:受访地址" json:"requestPath"`
	ClientIP    string `gorm:"type:varchar(20);comment:访客IP地址" json:"clientIP"`
	Agent       string `gorm:"type:varchar(40);comment:访客代理" `
	Referer     string `gorm:"type:varchar(80);comment:来源地址" json:"referer"`
	Date        time.Time
}
