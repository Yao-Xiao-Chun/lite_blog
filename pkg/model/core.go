package model

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"lite_blog/models"
	"time"
)

var (
	Db *gorm.DB
)

type MysqlSdk struct {
}

func (m MysqlSdk) GetClient() *gorm.DB {

	return Db
}

/**
伏见老贼丧天良，吾与黑猫共存亡
*/
func init() {

	logs.Info("数据库初始化！")

	var err error

	var pwd, mysqldb, sqlhost, sqluser, dbtype string

	pwd = beego.AppConfig.String("mysqlpass")

	mysqldb = beego.AppConfig.String("mysqldb")

	sqlhost = beego.AppConfig.String("mysqlurls")

	sqluser = beego.AppConfig.String("mysqluser")

	dbtype = beego.AppConfig.String("dbtype")

	Db, err = gorm.Open(dbtype, ""+sqluser+":"+pwd+"@tcp("+sqlhost+")/"+mysqldb+"?charset=utf8&parseTime=True&loc=Local") //连接数据库

	if err != nil {
		fmt.Println(err)
		panic("Mysql:连接数据库错误！请确认是否启动Mysql")
	}

	//设置一下mysql连接池
	pool := Db.DB()
	//最多连接100
	pool.SetMaxOpenConns(100)
	//空闲连接10
	pool.SetMaxIdleConns(10)
	pool.SetConnMaxLifetime(5 * time.Second)

	//初始化
	models.MigrateData(&MysqlSdk{})

	//defer db.Close()

}
