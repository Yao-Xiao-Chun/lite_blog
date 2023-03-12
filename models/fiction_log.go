package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"mywork/pkg/model"
)

/**
下载小说记录
*/
type LiteFictionLog struct {
	gorm.Model
	FictionId int    `gorm:"type:int;not null;index:fiction_id"`  //关联小说表id
	Ip        string `gorm:"type:varchar(100);not null;index:ip"` //下载用户的ip地址
	SubNum    int    `gorm:"type:int;default:1;not null"`         //下载本次书籍次数
}

/**
下载小说批次表
*/
type LiteFictionOperation struct {
	gorm.Model
	FictionId   int `gorm:"type:int;not null;index:fiction_id"` //关联小说表id
	DownloadNum int `gorm:"type:int;default:0;not null"`        //改小说下载的量

}

/**
此小说的下载记录
*/
func FictionOperation(id int) (op LiteFictionOperation, err error) {

	return op, model.Db.Where("fiction_id = ?", id).Select("download_num").Find(&op).Error
}

/**
更新下载次数
*/
func UpdateOperation(data LiteFiction) {
	//查询此id是否存在
	var op LiteFictionOperation

	ids, _ := findLog(data.ID)

	logs.Info(ids)

	if ids > 0 {
		//存在，更新
		model.Db.Model(&op).UpdateColumn("download_num", gorm.Expr("download_num + ?", 1))
	} else {

		//不存在，创建
		op.FictionId = int(data.ID)

		op.DownloadNum = 1

		model.Db.Create(&op)

	}
}

/**
创建日志记录
*/
func CreateFictionLog(data LiteFiction, ip string) {

	var log LiteFictionLog

	log.FictionId = int(data.ID)

	log.SubNum = 1

	log.Ip = ip

	model.Db.Create(&log)
}

/**
查询
*/
func findLog(id uint) (num int, err error) {
	/*	var num int*/

	var op LiteFictionOperation

	return num, model.Db.Where("fiction_id = ?", id).Find(&op).Count(&num).Error
}

/**
查询小说日志统计
*/
func CountFictionLog() (num int, err error) {

	var log []LiteFictionLog

	return num, model.Db.Find(&log).Count(&num).Error
}

/**
小说日志列表
*/
func GetFictionLogList(page, size int) (data []LiteFictionLog, err error) {

	return data, model.Db.Order("created_at desc").Offset((page - 1) * size).Limit(size).Find(&data).Error
}
