package dto

import (
	"mywork/models"
	"mywork/pkg/model"
)

/**
新创建任务

*/
func TaskAdd(data models.LiteCrontab) error {

	//var cro LiteCrontab

	return model.Db.Create(&data).Error
}

/**
查询数据
*/
func FindTask(page int, size int) (task []models.LiteCrontab, err error) {

	return task, model.Db.Offset((page - 1) * 10).Limit(10).Find(&task).Error
}

/**
统计数据
*/
func CountPage() (num int, err error) {

	var task []models.LiteCrontab

	return num, model.Db.Find(&task).Count(&num).Error
}

/**
查询任务执行计划
*/
func FindInfoTask(id int) (task models.LiteCrontab, err error) {

	return task, model.Db.Where("id = ?", id).Select("task_name,frequency,status,id,task_id").Find(&task).Error
}

/**
更新后台数据库
*/
func UpdateTaskStatus(task models.LiteCrontab) {

	model.Db.Omit("created_at,created_at,descript").Save(&task)
}

/**
删除
*/
func DeleteTask(id int) {

	var task models.LiteCrontab

	model.Db.Where("id = ?", id).Delete(&task)
}

/**
获取当前状态为1的定时任务
*/
func RunTask() (task []models.LiteCrontab, err error) {

	return task, model.Db.Where("status = ?", 1).Select("task_name,frequency,status,id,task_id").Find(&task).Error
}
