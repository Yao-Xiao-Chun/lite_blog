package models

import (
	"github.com/jinzhu/gorm"
)

type LiteArticleTag struct {
	gorm.Model
	Aid         int    `gorm:"not null;type:int;index:aid"` //文章id
	Tid         string `gorm:"not null;type:varchar(255)"`  //标签id
	Create_name string `gorm:"type:varchar(255);not null;"` //创建时间
	Uid         int    `gorm:"type:int;not null"`           //用户

}
