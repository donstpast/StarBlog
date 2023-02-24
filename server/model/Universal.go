package model

import (
	"time"
)

// Universal 通用模型
type Universal struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id" mapstructure:"id"`
	CreatedAt time.Time `json:"created_at" mapstructure:"-"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"-"`
}
