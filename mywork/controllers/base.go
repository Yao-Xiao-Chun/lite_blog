package controllers

import (
	"github.com/astaxie/beego"
	"crypto/md5"
	"io"
	"fmt"
	"github.com/satori/go.uuid"
	"mywork/syserror"
	"time"
	"github.com/tabalt/ipquery"
	"strings"
)

/**
	基础控制器
 */
type BaseController struct {

	beego.Controller


}

func (this *BaseController) Prepare(){

	this.Data["Path"] = this.Ctx.Request.RequestURI  //获取当前中的url



}


func (this *BaseController)Abort500(err error){

	this.Data["error"] = err

	this.Abort("500")
}

/**
	生成随机TOKen
 */
func (this *BaseController)SetToken()  {


}

/**
	生成加密密码
 */
 func (this *BaseController)SetMd5Pwd(pwd string) interface{}{

	 //假设用户名abc，密码123456
	 h := md5.New()

	 io.WriteString(h, pwd)

	 //pwmd5等于e10adc3949ba59abbe56e057f20f883e
	 pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	 return pwmd5
 }


 /**
 	生成 uuid
  */
func (this *BaseController) GetUUID() interface{} {

	uuidStr,err := uuid.NewV4()

	if err != nil{

		this.Abort500(syserror.New("uuid获取错误",err))
	}

	return uuidStr
}


/**
	生成格式化时间
 */

 func (this *BaseController) Date(times string) string{

	 datetime := time.Now().Format(times)

	 return datetime
 }

 /**
 	设置全局的图片上传类型
  */
  func (this *BaseController) GetUploadTypeImage()(arr map[string]string){

	  return map[string]string{
	  	"0":"jpg",
	  	"1":"gif",
	  	"2":"jpeg",
	  	"3":"png",
	  }
  }


func JsonFormat(retcode int, retmsg string, retdata interface{}, stime time.Time) (json map[string]interface{}) {
	cost := time.Now().Sub(stime).Seconds()
	if retcode == 1 {
		json = map[string]interface{}{
			"code": retcode,
			"data": retdata,
			"desc": retmsg,
			"cost": cost,
		}
	} else {
		json = map[string]interface{}{
			"code": retcode,
			"desc": retmsg,
			"cost": cost,
		}
	}

	return json
}


/**
	获取ip省份
 */
 func (this *BaseController) GetAddress(ip string) string{

 	if ip == ""{

 		return "未知省份"
	}

	 df := "conf/testdata/test_10000.txt"

	 err := ipquery.Load(df)
	 if err != nil {

		 fmt.Println(err)
	 }

	 dt, err := ipquery.Find(ip)

	 if err != nil {

		return "未知地址"

	 } else {

	 	ips := strings.Split(string(dt),"	")

	 	ips = append(ips[:3],ips[4:]...) //移除第三个切片

	 	ipStr := strings.Join(ips,"-")

	 	return ipStr
	 }
 }