package models

import (
	"github.com/jinzhu/gorm"
)

type LiteBase struct {
	gorm.Model
	Content string `gorm:"type:text;null"`     //存放数据
	Type    string `gorm:"type:int;default:0"` //存放数据
}
