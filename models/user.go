package models

import (
	_ "github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

type LiteOauthUser struct {
	gorm.Model
	//Id int `gorm:"PRIMARY_KEY;not null"` //主键id
	Uid           int    `gorm:"not null;DEFAULT :0"` //关联本站的id
	Types         int    `gorm:"not null;default :1"` //类型 1：QQ  2：新浪微博 3：豆瓣 4：人人 5：开心网'
	Nikename      string `gorm:"not null;defaul:"`    //第三方昵称
	Head_img      string `gorm:"not null;defaul:"`    //头像
	Openid        string `gorm:"not null;defaul:"`    //第三方用户id
	Access_token  string `gorm:"not null;defaul:0"`   //绑定时间
	Create_time   int    `gorm:"not null;defaul:0"`   //最后登录时间
	Last_login_ip string `gorm:"not null;defaul:"`    //最后登录ip
	Login_times   int    `gorm:"not null;defaul:0"`   //登录次数
	Status        int    `gorm:"not null;defaul:1"`   //状态
	Email         string `gorm:"not null;defaul:"`    //邮箱
	Is_admin      int    `gorm:"not null;defaul:0"`   //是否是admin
	Account       string `gorm:"not null"`            //账户
	Password      string `gorm:"not null"`            //密码
}
