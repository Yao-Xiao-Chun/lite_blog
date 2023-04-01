package cmd

import (
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
	"lite_blog/models"
	"lite_blog/pkg/search"
	"lite_blog/routers"
	"strconv"
	"strings"
	"time"
)

// Execute 初始化执行命令
func Execute() {

	routers.HomeApi()

	initSearch()

	tip()

}

func tip() {
	initSetSession()

	setTemplate()

	gob.Register(models.LiteOauthUser{})

	beego.Run(beego.AppConfig.String("httpport"))
}

func setTemplate() {

	//创建新的模板函数
	err := beego.AddFuncMap("equrl", func(x, y string) bool {

		//去除获取url中的 /
		x1 := strings.Trim(x, "/")

		y1 := strings.Trim(y, "/")

		if strings.Compare(x1, y1) == 0 {

			return true
		} else {

			return false
		}

	})
	if err != nil {
		return
	}

	//拼接参数
	err = beego.AddFuncMap("urljoin", func(x string, y int) string {

		str := strconv.Itoa(y)
		return x + str
	})
	if err != nil {
		return
	}

	/**
	判读一个参数是否存在另一个集合里面
	*/
	err = beego.AddFuncMap("in_array", func(x uint, sli []string) bool {

		str := strconv.Itoa(int(x))

		for _, val := range sli {

			if val == str {

				return true
				break
			}
		}

		return false
	})
	if err != nil {
		return
	}

}

/**
初始化session
*/
func initSetSession() {

	beego.BConfig.WebConfig.Session.SessionOn = true

	//设置 cookies 的名字，Session 默认是保存在用户的浏览器 cookies 里面的，默认名是 beegosessionID，配置文件对应的参数名是
	beego.BConfig.WebConfig.Session.SessionName = "liteblog"

	beego.BConfig.WebConfig.Session.SessionProvider = "file"

	beego.BConfig.WebConfig.Session.SessionProviderConfig = "docs/session"
}

//初始化搜索引擎
func initSearch() {

	host := beego.AppConfig.String("searchhost")
	port := beego.AppConfig.String("searchport")
	key := beego.AppConfig.String("searchkey")
	log.Infoln(fmt.Sprintf("搜索引擎host:%v,端口：%v,master_key:%v", host, port, key))

	ser := NewMeiliSearch(host, port, key)

	search.SearchSDK = ser.OpenClient() //创建全局使用变量

	fmt.Println("init search success...")

	//开启搜索引擎监控
	go openElasticsearchPing()

	search.AutoSearchData()
}

// 监听搜索引擎服务是否存活，不影响主线程
func openElasticsearchPing() {

	defer func() {
		if r := recover(); r != nil {

			log.Infoln("捕获错误准备重连搜索引擎...")
			go tryReconnecting() //重新连接搜索引擎
		}
	}()

	for {
		log.Info("ping search...")
		resp, err := search.SearchSDK.Health()

		if err != nil {
			_ = fmt.Errorf("搜索引擎错误！%v", err)
		}
		log.Infoln(fmt.Sprintf("搜索引擎运行状态...%v", resp.Status))

		time.Sleep(time.Second * 2)

	}

}

// 重新连接搜索引擎
// 每10s重新连接一场
func tryReconnecting() {
	host := beego.AppConfig.String("searchhost")
	port := beego.AppConfig.String("searchport")
	key := beego.AppConfig.String("searchkey")

	log.Infoln(fmt.Sprintf("搜索引擎host:%v,端口：%v,master_key:%v", host, port, key))

	ser := NewMeiliSearch(host, port, key)

	for {
		if tmp := ser.OpenClient().IsHealthy(); tmp {
			//重新连接成功
			search.SearchSDK = ser.OpenClient()
			log.Infoln("重新连接成功success.....")
			go openElasticsearchPing()
			break
		}
		//重连失败
		log.Info(fmt.Sprintf("重新连接失败10s后重新尝试...."))

		time.Sleep(time.Second * 10)
	}

}
