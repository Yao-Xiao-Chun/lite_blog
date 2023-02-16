package models

import "github.com/jinzhu/gorm"

type LiteCrontab struct {
	gorm.Model
	Task
}

/**
	定时任务表
 */
type Task struct {
	TaskName string `gorm:"type:varchar(255);not null;"`//名称
	Descript string `gorm:"type:varchar(255);null;"`//地址
	Status int `gorm:"type:int;not null;default:0"`//启用状态 1 启用 0禁用
	Frequency string `gorm:"type:varchar(255);null;"`//运行的间隔
	CreateName string `gorm:"type:varchar(50);null;` //创建者
	TaskId string `gorm:"type:varchar(50);null;` //执行任务id
}


/**
	新创建任务

 */
func TaskAdd(data LiteCrontab)error  {

	//var cro LiteCrontab

	return db.Create(&data).Error
}

/**
	查询数据
 */
func FindTask(page int,size int) (task []LiteCrontab,err error){

	return task, db.Offset((page - 1) * 10).Limit(10).Find(&task).Error
}

/**
	统计数据
 */
func CountPage()(num int,err error)  {

	var task []LiteCrontab

	return num,db.Find(&task).Count(&num).Error
}

/**
	查询任务执行计划
 */
func FindInfoTask(id int) (task LiteCrontab,err error) {

	return task,db.Where("id = ?",id).Select("task_name,frequency,status,id,task_id").Find(&task).Error
}


/**
	更新后台数据库
 */
func UpdateTaskStatus(task LiteCrontab){

	db.Omit("created_at,created_at,descript").Save(&task)
}


/**
	删除
 */
func DeleteTask(id int){

	var task LiteCrontab

	db.Where("id = ?",id).Delete(&task)
}


/**
	获取当前状态为1的定时任务
 */
func RunTask() (task []LiteCrontab,err error){

	return task,db.Where("status = ?",1).Select("task_name,frequency,status,id,task_id").Find(&task).Error
}