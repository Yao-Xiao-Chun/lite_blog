package models

import (
	"github.com/jinzhu/gorm"
	"mywork/pkg/model"
)

/**
日志记录表
*/
type LiteLog struct {
	gorm.Model
	Content string `gorm:"type:varchar(255);not null"`  //存放日志记录
	Level   int    `gorm:"type:int;not null;default:0"` //等级
}

/**
insert 存放日志
@param info sting 描述日志，不超过255字符
@param level int  层次等级
*/
func InsertLog(info string, level int) error {

	var log LiteLog

	log.Content = info

	log.Level = level

	return model.Db.Create(&log).Error

}

/**
查询数据
*/
func SelectLog(page int) (log []LiteLog, err error) {

	return log, model.Db.Order("created_at desc,id desc").Offset((page - 1) * 10).Limit(10).Find(&log).Error

}

/**
条数
*/
func CountLog() (num int, err error) {

	var log []LiteLog

	return num, model.Db.Order("created_at desc,id desc").Find(&log).Count(&num).Error
}

func FindLogAll() (log []LiteLog, err error) {

	return log, model.Db.Order("created_at desc,id desc").Find(&log).Error
}
