package models

import "github.com/jinzhu/gorm"

/**
	websocket执行的记录表 推送表，超级管理登录可以发送推送，其他用户可以及时的查看 阅读后，访日已读表格
 */
type TaskLog struct {
	gorm.Model
	UserID int `gorm:"type:int;not null;index"`
	TaskType  int `gorm:"type:int;not null;default:0"`//推送类型
	TaskID int	`gorm:"type:int;null"`//执行任务id
	TaskRead int `gorm:"type:int;not null;default:0;"`//是否查看
	Status int `gorm:"type:int;not null;default:1;"`//是否重新推送 0：是 1，否
}


/**

 */