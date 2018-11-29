package controllers

import (
	"mywork/models"
	"time"
	"github.com/astaxie/beego/logs"
)

type AdminIndexController struct {

	AdminBaseController
}

/**
	注解路由 后台首页
 */


// @router /admin [get] 后台首页
func (this *AdminIndexController)AdminIndex(){
	this.SetLogs("我被请求到了，看看日志是否存在呢！")
	this.TplName = "admin/index.html"
}

// @router /main [get] 后台 main
func (this *AdminIndexController)AdminMain(){

	this.getTimeLines()

	this.countMessage()

	this.TplName = "admin/main.html"
}


// @router /admin/login [get] 后台 登录页面
func (this *AdminIndexController)AdminLogin(){

	this.TplName = "admin/login.html"
}

// @router /admin/addtime [get] 后台 获取时间先路由
func (this *AdminIndexController)AdminGettime(){

	this.Data["key"] = this.GetUUID()

	this.TplName = "admin/news/addtime.html"
}




/**
	查询 后台的时间线
 */

 func (this *AdminIndexController)getTimeLines(){

	 line, err := models.GetAdminTimeLine() //多维结构体

	 arr := make(map[int]map[string]interface{},10)

	 if err != nil {

	 	arr[0]["code"] = "没有数据"

	 	arr[0]["content"] = "没有数据"

	 }else{

		 //变量赋值
		 for key,val := range line{

			 data := make(map[string]interface{},4) // 每次使用都要初始化一次

			 data["code"] = val.Title

			 data["content"] = val.Content

			 data["id"] = val.ID

			 data["token"] = val.Token

			 arr[key] = data

		 }
	 }

	 num,_:= models.GetHomeCountTimeLine()

	 this.Data["count"] = num

	 this.Data["line"] = arr

 }


 /**
 	基本资料
  */
// @router /admin/user/message [get] 后台 获取用户资料
func (this *AdminIndexController) GetUserMessage(){

	//获取当前登录用户的id

	id := this.User.ID

	user,err := models.FindUser(int(id));

	if err != nil{

	}else{

		this.Data["userinfo"] = map[string]interface{}{
			"title":user.Nikename,
			"status":user.Status,
			"is_admin":user.Is_admin,
			"head_img":user.Head_img,
			"account":user.Account,
			"email":user.Email,
			"password":user.Password,
			"id":user.ID,
		}

	}

	this.TplName = "admin/user/edit.html"
}



/**
	统计今天新增的留言数量
	后台
 */
 func (this *AdminIndexController) countMessage(){

 	//获取今天的时间
 	timeStr := time.Now().Format("2006-01-02 00:00:00")

 	count,_:= models.GetWhereReviewCount(timeStr)

 	this.Data["reviewCount"] = count

 }


 /**
 	后台关于
  */
// @router /admin/baseseting [get] 后台 获取用户资料
func (this *AdminIndexController) SetAbort(){


	this.Data["abort"],_ = models.GetAbort()

	this.TplName = "admin/abort/index.html"
}

/**
	后台关于 数据处理
 */
// @router /admin/baseseting [post] 后台 获取用户资料
func (this *AdminIndexController) AbortFormData() {

	data := this.GetString("content") //获取数据

	logs.Info(data)

	if data == ""{

		this.Data["json"] = map[string]interface{}{
			"code":1003,
			"errmsg":"数据丢失",
		}
	}else{
		models.UpdateBase(data)

		this.Data["json"] = map[string]interface{}{
			"code":0,
			"errmsg":"更新成功",
		}
	}

	this.ServeJSON()
}


/**
	后台公告
 */
// @router /admin/baseplacard [get] 后台 获取用户资料
 func (this *AdminIndexController) SetPlacard(){

 	this.Data["abort"],_ = models.GetPlacard()

 	this.TplName = "admin/placard/index.html"
 }

/**
   后台公告 数据处理
*/
// @router /admin/baseplacard [post] 后台 获取用户资料
func (this *AdminIndexController) PlacardFormData()  {

	data := this.GetString("content") //获取数据

	logs.Info(data)

	if data == ""{

		this.Data["json"] = map[string]interface{}{
			"code":1003,
			"errmsg":"数据丢失",
		}
	}else{
		models.UpdatePlacard(data)

		this.Data["json"] = map[string]interface{}{
			"code":0,
			"errmsg":"更新成功",
		}
	}

	this.ServeJSON()
}



/**
	获取日志记录
 */
// @router /admin/log/index [get]
func (this *AdminIndexController) GetLogList(){

	this.Data["num"],_ = models.CountLog()

	this.TplName = "admin/log/index.html"
}

/**
	日志数据获取
 */

// @router /admin/log/page/?:key [get]
func (this *AdminIndexController) GetLogPage(){

	var page int

	this.Ctx.Input.Bind(&page,"page")

	var res []models.LiteLog

	res = make([]models.LiteLog,0)

	if page == 0{

		res,_ = models.SelectLog(1)
	}else{

		res,_ = models.SelectLog(page)
	}


	this.Data["json"] = map[string]interface{}{
		"code":"0",
		"data":res,
	}

	this.ServeJSON()
}