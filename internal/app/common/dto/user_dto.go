package dto

import (
	"github.com/astaxie/beego/logs"
	"mywork/models"
	"mywork/pkg/model"

	_ "github.com/astaxie/beego/logs"
)

/**
查询登录的账户是否存在
*/
func QueryAccountAndPwd(account string, pwd string) (user models.LiteOauthUser, err error) {

	logs.Info(model.Db.Debug().Where("account = ? and password = ?", account, pwd).Take(&user)) //注意Take里面穿的参数，不能出啊挪了
	return user, model.Db.Where("account = ? and password = ? and status = 1 and is_admin = 1", account, pwd).Take(&user).Error

}

/**
查询所有的用户数据
*/
func QueryUserList(page int, limit int) (user []models.LiteOauthUser, err error) {

	if limit == 0 {

		limit = 10
	}

	var list []models.LiteOauthUser

	return list, model.Db.Order("created_at desc,id desc").Offset((page - 1) * limit).Limit(limit).Find(&list).Error
}

/**
获取当前所有的id
*/
func GetUserNum() (num int, err error) {

	var count int

	var user []models.LiteOauthUser

	return count, model.Db.Order("id desc").Find(&user).Count(&count).Error
}

/**
查询当前账户是否存在
*/
func GetIsAccount(account string) (user models.LiteOauthUser, err error) {

	return user, model.Db.Where("account = ?", account).Limit(1).Take(&user).Error
}

/**
创建新的用户
*/
func CreateUser(user *models.LiteOauthUser) {

	model.Db.Save(&user)
}

/**
更新本机ip地址
*/
func UpdateIP(ip string, id int) {

	var user models.LiteOauthUser

	model.Db.Model(&user).Where("id = ?", id).Update("last_login_ip", ip).Limit(1)
}

/**
删除用户
*/
func DelUser(id int) (user models.LiteOauthUser, err error) {

	return user, model.Db.Where("id = ?", id).Delete(&user).Error
}

/**
查询当前用户的数据详情
*/
func FindUser(id int) (user models.LiteOauthUser, err error) {

	return user, model.Db.Where("id = ?", id).Limit(1).First(&user).Error
}

/**
修改用户信息
*/
func EditUser(id uint, user models.LiteOauthUser) {

	model.Db.Save(&user)

}
