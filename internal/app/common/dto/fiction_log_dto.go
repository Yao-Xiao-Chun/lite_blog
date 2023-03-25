package dto

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"mywork/models"
	"mywork/pkg/model"
)

/**
此小说的下载记录
*/
func FictionOperation(id int) (op models.LiteFictionOperation, err error) {

	return op, model.Db.Where("fiction_id = ?", id).Select("download_num").Find(&op).Error
}

/**
更新下载次数
*/
func UpdateOperation(data models.LiteFiction) {
	//查询此id是否存在
	var op models.LiteFictionOperation

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
func CreateFictionLog(data models.LiteFiction, ip string) {

	var log models.LiteFictionLog

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

	var op models.LiteFictionOperation

	return num, model.Db.Where("fiction_id = ?", id).Find(&op).Count(&num).Error
}

/**
查询小说日志统计
*/
func CountFictionLog() (num int, err error) {

	var log []models.LiteFictionLog

	return num, model.Db.Find(&log).Count(&num).Error
}

/**
小说日志列表
*/
func GetFictionLogList(page, size int) (data []models.LiteFictionLog, err error) {

	return data, model.Db.Order("created_at desc").Offset((page - 1) * size).Limit(size).Find(&data).Error
}
