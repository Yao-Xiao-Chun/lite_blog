package dto

import (
	_ "github.com/astaxie/beego/logs"
	"lite_blog/models"
	"lite_blog/pkg/model"
)

/**
时间线model
*/

/**
创建新的时间线
*/

func CreateTimeLine(line models.LiteTimeLine) {

	model.Db.Save(&line)

}

/**
查询token是否存在
*/
func QueryToken(key string) (line models.LiteTimeLine, err error) {

	return line, model.Db.Where("token = ?", key).Take(&line).Error

}

/**
前台 获取所有的时间线
*/
func GetHomeTimeLine() (line []models.LiteTimeLine, err error) {

	var nx2 []models.LiteTimeLine
	//启用状态
	return nx2, model.Db.Where("status = 1").Order("id desc").Limit(10).Find(&nx2).Error
}

/**
前台获取时间线的所有条数
*/

func GetHomeCountTimeLine() (num int, err error) {

	var count int

	var lite []models.LiteTimeLine

	return count, model.Db.Where("status = 1").Order("id desc").Find(&lite).Count(&count).Error
}

/**
后台获取所有
*/

func GetAdminTimeLine() (line []models.LiteTimeLine, err error) {

	var nx2 []models.LiteTimeLine
	//启用状态
	return nx2, model.Db.Order("id desc").Limit(10).Find(&nx2).Error
}

/**
删除时间线
*/

func SetDelTimes(id string, token string) (line models.LiteTimeLine, err error) {

	return line, model.Db.Where("id = ? and token = ?", id, token).Delete(&line).Error

}

/**
后台获取一条数据 用户修改
*/
func GetTileLineFind(id int, token string) (line models.LiteTimeLine, err error) {

	return line, model.Db.Where("id = ? and token = ?", id, token).First(&line).Error
}

/**
修改数据
*/
func SetTimeInfo(id string, line models.LiteTimeLine) bool {

	model.Db.Save(&line)

	return true
}

/**
前台时间线分页方法
*/
func GetPageTimeLine(id int, limit int) (line []models.LiteTimeLine, err error) {

	if limit == 0 {

		limit = 10
	}

	var nx2 []models.LiteTimeLine
	//启用状态
	return nx2, model.Db.Where("status = 1 and id < ?", id).Order("id desc").Limit(limit).Find(&nx2).Error

}
