package controllers

import (
	"mywork/models"
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego/logs"

)

/**
	时间线处理控制器；
 */
type AdminTimeController struct {

	AdminBaseController
}

// @router /admin/timepost/:key [post] Time新增处理方法
func (this *AdminTimeController) AddTime(){
	//获取key
	token := this.Ctx.Input.Param(":key")

	//查询token是否存在
	line,err := models.QueryToken(token)

	//判断是否查找到记录
	if !gorm.IsRecordNotFoundError(err) {

		this.Data["json"] = map[string]interface{}{
			"code":2,
			"msg":"token创建失败",
		}
	}else{

		//获取标题
		title := this.CheckMustKey("title","没有获取到正确的标题")

		content := this.GetString("content")

		status := this.CheckMustKey("status","请选择启用状态")

		logs.Info(status)

		line.Title = title

		line.Content = content

		line.Uid = int(this.User.ID)//用户id

		line.Token = token

		line.Status = status

		models.CreateTimeLine(line)

		this.Data["json"] = map[string]interface{}{
			"code":0,
			"msg":"创建成功",
		}

	}



	this.ServeJSON()
}


/**
	删除 时间线操作
 */

// @router /admin/deltime/:key [get] Time新增处理方法
func (this *AdminTimeController) DelTime(){

	tid := this.GetString("id")

	token := this.GetString("token")

	if (tid == "" || token == ""){

		this.Data["json"] = map[string]interface{}{
			"code":1002,
			"msg":"数据丢失",

		}

	}else{

		_,err := models.SetDelTimes(tid,token)

		if err == nil{

			this.Data["json"] = map[string]interface{}{
				"code":0,
				"msg":"删除成功",

			}
		}else{

			this.Data["json"] = map[string]interface{}{
				"code":1002,

				"msg":"删除失败",

			}

		}

	}

	this.ServeJSON()
}


/**
	修改时间线页面
 */
// @router /admin/timeinfo/?:id [get]
func (this *AdminTimeController) GetTimeInfo()  {

	var tid int
    this.Ctx.Input.Bind(&tid,"id")

	var token string

	this.Ctx.Input.Bind(&token,"token")

	line,_ :=models.GetTileLineFind(tid,token)

	data := map[string]interface{}{

		"title":line.Title,
		"content":line.Content,
		"id":line.ID,
		"token":line.Token,
		"status":line.Status,

	}
	this.Data["lines"] = data

	this.TplName = "admin/news/time.html"

}


/**
	修改提交表单
 */

 //@router /admin/formtime/:key [post]
 func (this *AdminTimeController) SetTimeInfo(){
	 //获取key
	 token := this.Ctx.Input.Param(":key")

	 //查询token是否存在
	 line,err := models.QueryToken(token)

	 //判断是否查找到记录
	 if !gorm.IsRecordNotFoundError(err) {
		 //获取标题
		 title := this.CheckMustKey("title","没有获取到正确的标题")

		 content := this.GetString("content")

		 id := this.GetString("id")

		 status := this.CheckMustKey("status","请选择启用状态")

		 line.Status =status

		 line.Content =content

		 line.Title = title

		 models.SetTimeInfo(id,line)

		 this.Data["json"] = map[string]interface{}{
			 "code":0,
			 "msg":"修改成功",
		 }

	 }else{
		 this.Data["json"] = map[string]interface{}{
			 "code":2,
			 "msg":"数据不存在，禁止操作",
		 }
	 }

	 this.ServeJSON()
 }

