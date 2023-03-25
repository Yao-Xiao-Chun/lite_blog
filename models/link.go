package models

import (
	"github.com/jinzhu/gorm"
)

/**
友情链接表
*/
type LiteLink struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255);not null;"`  //名称
	Url    string `gorm:"type:varchar(255);not nul;"`   //地址
	Sort   int    `gorm:"type:int;not null;default:0;"` //默认排序
	Status int    `gorm:"type:int;not null;default:1"`  //启用状态 1 启用 0禁用
}
