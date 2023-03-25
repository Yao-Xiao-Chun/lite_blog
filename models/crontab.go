package models

import (
	"github.com/jinzhu/gorm"
)

type LiteCrontab struct {
	gorm.Model
	Task
}

/**
定时任务表
*/
type Task struct {
	TaskName   string `gorm:"type:varchar(255);not null;"` //名称
	Descript   string `gorm:"type:varchar(255);null;"`     //地址
	Status     int    `gorm:"type:int;not null;default:0"` //启用状态 1 启用 0禁用
	Frequency  string `gorm:"type:varchar(255);null;"`     //运行的间隔
	CreateName string `gorm:"type:varchar(50);null;`       //创建者
	TaskId     string `gorm:"type:varchar(50);null;`       //执行任务id
}
