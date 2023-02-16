package models

import "github.com/jinzhu/gorm"

// 推送任务执行表
type LiteTask struct {

	gorm.Model

	TaskContent string `gorm:"type:varchar(255);not null;index:task_content;"` //推送内容

	TaskStatus int `gorm:"type:int;not null;default:0;"` //是否开启推送

	Num int `gorm:"type:int;not null;default:0;"`//推送次数

}


