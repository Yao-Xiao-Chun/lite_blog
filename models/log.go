package models

import (
	"github.com/jinzhu/gorm"
)

/**
日志记录表
*/
type LiteLog struct {
	gorm.Model
	Content string `gorm:"type:varchar(255);not null"`  //存放日志记录
	Level   int    `gorm:"type:int;not null;default:0"` //等级
}
