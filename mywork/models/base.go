package models

import "github.com/jinzhu/gorm"

type LiteBase struct {
	gorm.Model
	Content string `gorm:"type:text;null"` //存放数据
	Type string `gorm:"type:int;default:0"` //存放数据
}

/**
	创建数据
 */
 func UpdateBase(content string) {

 	var base LiteBase
 	base.Content = content
 	base.Type = "1"
 	db.Save(&base)
 }


 /**
 	获取数据
  */
func GetAbort() (list LiteBase,err error) {

	return list,db.Where("type = ?",1).Order("created_at desc").Limit(1).Find(&list).Error
}