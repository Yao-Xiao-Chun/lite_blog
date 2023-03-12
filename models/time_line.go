package models

import (
	_ "github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"mywork/pkg/model"
)

/**
时间线model
*/
type LiteTimeLine struct {
	gorm.Model

	Content string `gorm:"not null;type:text;"`

	Status string `gorm:"type:int;default:1"`

	Uid int `gorm:"type:int(6);not null"` //创建用户

	Token string `gorm:"type:varchar(250);not null;unique_index"` //token,临牌，禁止重复，唯一索引

	Title string `gorm:"null;type:text;"` //标题

}

/**
创建新的时间线
*/

func CreateTimeLine(line LiteTimeLine) {

	model.Db.Save(&line)

}

/**
查询token是否存在
*/
func QueryToken(key string) (line LiteTimeLine, err error) {

	return line, model.Db.Where("token = ?", key).Take(&line).Error

}

/**
前台 获取所有的时间线
*/
func GetHomeTimeLine() (line []LiteTimeLine, err error) {

	var nx2 []LiteTimeLine
	//启用状态
	return nx2, model.Db.Where("status = 1").Order("id desc").Limit(10).Find(&nx2).Error
}

/**
前台获取时间线的所有条数
*/

func GetHomeCountTimeLine() (num int, err error) {

	var count int

	var lite []LiteTimeLine

	return count, model.Db.Where("status = 1").Order("id desc").Find(&lite).Count(&count).Error
}

/**
后台获取所有
*/

func GetAdminTimeLine() (line []LiteTimeLine, err error) {

	var nx2 []LiteTimeLine
	//启用状态
	return nx2, model.Db.Order("id desc").Limit(10).Find(&nx2).Error
}

/**
删除时间线
*/

func SetDelTimes(id string, token string) (line LiteTimeLine, err error) {

	return line, model.Db.Where("id = ? and token = ?", id, token).Delete(&line).Error

}

/**
后台获取一条数据 用户修改
*/
func GetTileLineFind(id int, token string) (line LiteTimeLine, err error) {

	return line, model.Db.Where("id = ? and token = ?", id, token).First(&line).Error
}

/**
修改数据
*/
func SetTimeInfo(id string, line LiteTimeLine) bool {

	model.Db.Save(&line)

	return true
}

/**
前台时间线分页方法
*/
func GetPageTimeLine(id int, limit int) (line []LiteTimeLine, err error) {

	if limit == 0 {

		limit = 10
	}

	var nx2 []LiteTimeLine
	//启用状态
	return nx2, model.Db.Where("status = 1 and id < ?", id).Order("id desc").Limit(limit).Find(&nx2).Error

}
