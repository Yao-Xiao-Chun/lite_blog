package models

import (
	"github.com/jinzhu/gorm"

	_"github.com/astaxie/beego/logs"

)

type LiteOauthUser struct {
	gorm.Model
	//Id int `gorm:"PRIMARY_KEY;not null"` //主键id
	Uid int `gorm:"not null;DEFAULT :0"` //关联本站的id
	Types int `gorm:"not null;default :1"` //类型 1：QQ  2：新浪微博 3：豆瓣 4：人人 5：开心网'
	Nikename string `gorm:"not null;defaul:"`//第三方昵称
	Head_img string `gorm:"not null;defaul:"`//头像
	Openid string    `gorm:"not null;defaul:"`//第三方用户id
	Access_token string `gorm:"not null;defaul:0"`//绑定时间
	Create_time int	`gorm:"not null;defaul:0"`//最后登录时间
	Last_login_ip string `gorm:"not null;defaul:"`//最后登录ip
	Login_times int	`gorm:"not null;defaul:0"`//登录次数
	Status int	`gorm:"not null;defaul:1"`//状态
	Email string `gorm:"not null;defaul:"`//邮箱
	Is_admin int `gorm:"not null;defaul:0"`//是否是admin
	Account string `gorm:"not null"` //账户
	Password string `gorm:"not null"`//密码
}

/**
	查询登录的账户是否存在
 */
func QueryAccountAndPwd(account string,pwd string) (user LiteOauthUser,err error)   {

	//logs.Info(db.Where("account = ? and password = ?",account,pwd).Take(&user)) 注意Take里面穿的参数，不能出啊挪了
	return user,db.Where("account = ? and password = ? and status = 1 and is_admin = 1",account,pwd).Take(&user).Error

}

/**
	查询所有的用户数据
 */
func QueryUserList(page int,limit int) (user []LiteOauthUser,err error) {

	if limit == 0{

		limit = 10;
	}

	var list []LiteOauthUser

	return list,db.Order("created_at desc,id desc").Offset((page-1)*limit).Limit(limit).Find(&list).Error
}


/**
	获取当前所有的id
 */
func GetUserNum() (num int,err error) {

	var count int

	var user []LiteOauthUser

	return count,db.Order("id desc").Find(&user).Count(&count).Error
}


/**
	查询当前账户是否存在
 */
 func GetIsAccount(account string)(user LiteOauthUser,err error){

	 return user,db.Where("account = ?",account).Limit(1).Take(&user).Error
 }


 /**
 	创建新的用户
  */
  func CreateUser(user *LiteOauthUser){

  	db.Save(&user)
  }

  /**
  	 更新本机ip地址
   */
   func UpdateIP(ip string,id int){

   		var user LiteOauthUser

		db.Model(&user).Where("id = ?",id).Update("last_login_ip",ip).Limit(1)
   }

   /**
   		删除用户
    */
    func DelUser(id int)(user LiteOauthUser,err error){

    	return user,db.Where("id = ?",id).Delete(&user).Error
	}


	/**
		查询当前用户的数据详情
	 */
	 func FindUser(id int)(user LiteOauthUser,err error){

	 	return user, db.Where("id = ?",id).Limit(1).First(&user).Error
	 }


	 /**
	 	修改用户信息
	  */
func EditUser(id uint,user LiteOauthUser)  {

	db.Save(&user)

}