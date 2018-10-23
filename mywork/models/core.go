package models

import (
	_"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

var(
	db *gorm.DB
)
/**
	伏见老贼丧天良，吾与黑猫共存亡
 */
func init()  {

	logs.Info("数据库初始化！")

	var err error

	var pwd string

	pwd = beego.AppConfig.String("mysqlpass")

	db, err = gorm.Open("mysql", "root:"+pwd+"@/lite_blog?charset=utf8&parseTime=True&loc=Local") //连接数据库

	if err != nil{

		panic("Mysql:连接数据库错误！请确认是否启动Mysql")
	}

	//同步表结构 做数据迁移的时候使用
	auotoData()


	var count int

	errs := db.Model(&LiteOauthUser{}).Count(&count).Error //判断是否存在数据

	if errs == nil && count == 0{
		//新增一条数据
		db.Create(&LiteOauthUser{
			Uid:1,
			Nikename:"王大锤",
			Types:1,
			Account:"949656336@qq.com",
			Password:"1234@abcd",
			Is_admin:1,
			Status:1,

		})
	}


	//defer db.Close()

}

func auotoData(){
	//同步表结构
	db.AutoMigrate(&LiteOauthUser{}) //如果不存在就会生成一个新表

	db.AutoMigrate(&LiteAdminMenu{}) //创建一张新的菜单表

	db.AutoMigrate(&LiteTimeLine{}) //创建一张时间线表

	db.AutoMigrate(&LiteArticleTag{}) //创建一张文章管理表

	db.AutoMigrate(&LiteTag{}) //创建一张标签表

	db.AutoMigrate(&LiteArticle{}) //创建一张文章

	db.AutoMigrate(&LiteReview{}) //创建一张留言表

}
