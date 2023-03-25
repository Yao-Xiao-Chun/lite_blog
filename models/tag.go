package models

import (
	"github.com/jinzhu/gorm"
)

type LiteTag struct {
	gorm.Model
	Tag_name  string `gorm:"not null;type:varchar(255)"` //标签名称
	Is_status int    `gorm:"not null;default:0"`
}
