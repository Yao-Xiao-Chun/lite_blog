package controllers

import (
	"strconv"
	"mywork/models"
)

/**
	友情链接控制器
 */
type AdminLinkController struct {
	AdminBaseController
}

// @router /admin/link/index [get] 友情首页
func (this *AdminLinkController) Index()  {

	this.Data["num"],_ = models.GetLinkTotal()

	this.TplName = "admin/link/index.html"
}


// @router /admin/link/add [get] 友情新增
func (this *AdminLinkController) LinkAdd()  {

	this.TplName = "admin/link/add.html"
}

// @router /admin/link/add [post] 友情新增
func (this *AdminLinkController) LinkAddForm()  {
	linkName := this.GetString("link_name")

    linkUrl := this.GetString("link_url")

	linkStatus := this.GetString("link_status")

	linkSort := this.GetString("link_sort")

	if linkName == ""|| linkUrl == "" || linkSort == "" || linkStatus == ""{

		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"msg":"参数不完整",
		}
	}else{

		var link models.LiteLink

		link.Name = linkName
		link.Url = linkUrl
		link.Sort,_ = strconv.Atoi(linkSort)
		link.Status,_ = strconv.Atoi(linkStatus)

		err := models.AddLinkData(link)

		if err == nil{

			this.ReadLog("账号:"+this.User.Nikename+" 操作：添加友链:'"+linkName+"',状态：成功",2)	//写入操作日志
		}

		this.Data["json"] = map[string]interface{}{
			"code":"0",
			"msg":"创建成功",
		}

	}

	this.ServeJSON()
}


/**
	获取前台展示数据
 */
// @router /admin/link/page/?:key [get]
 func (this *AdminLinkController) GetLinkIndex(){
	 var page int

	 this.Ctx.Input.Bind(&page,"page")

	 var res []models.LiteLink

	 res = make([]models.LiteLink,0)

	 if page == 0{

		 res,_ = models.SelectLink(1)
	 }else{

		 res,_ = models.SelectLink(page)
	 }


	 this.Data["json"] = map[string]interface{}{
		 "code":"0",
		 "data":res,
	 }

	 this.ServeJSON()
 }


// @router /admin/link/edit/?:key [get] 编辑页面
func (this *AdminLinkController) LinkInfo(){

	var id int

	this.Ctx.Input.Bind(&id,"id")

	this.Data["link"],_ = models.FindLinkInfo(id)

	this.TplName = "admin/link/edit.html"
}

// @router /admin/link/edit [post] 编辑保存
func (this *AdminLinkController) LinkFormData()  {

	linkName := this.GetString("link_name")

	linkUrl := this.GetString("link_url")

	linkStatus := this.GetString("link_status")

	linkSort := this.GetString("link_sort")

	id := this.GetString("id")

	if linkName == ""|| linkUrl == "" || linkSort == "" || linkStatus == "" || id == ""{

		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"msg":"参数不完整",
		}
	}else{

		ids,_ := strconv.Atoi(id)

		var link models.LiteLink

		link.Name = linkName
		link.Url = linkUrl
		link.Sort,_ = strconv.Atoi(linkSort)
		link.Status,_ = strconv.Atoi(linkStatus)
		link.ID = uint(ids)

		err := models.SaveLink(link)

		if err == nil{

			this.ReadLog("账号:"+this.User.Nikename+" 操作：修改友链:'"+linkName+"',状态：成功",2)	//写入操作日志
		}

		this.Data["json"] = map[string]interface{}{
			"code":"0",
			"msg":"修改成功",
		}

	}

	this.ServeJSON()
}


/**
	删除
 */
// @router /admin/link/delete/?:key [get] 编辑保存
func (this *AdminLinkController) LinkDel()  {

	var id int

	this.Ctx.Input.Bind(&id,"id")

	if id == 0{

		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"msg":"参数不完整",
		}
	}else{

		err := models.DeleteLink(id)

		if err == nil{

			this.ReadLog("账号:"+this.User.Nikename+" 操作：删除友链:'"+strconv.Itoa(id)+"',状态：成功",2)	//写入操作日志
		}

		this.Data["json"] = map[string]interface{}{
			"code":"0",
			"msg":"删除成功",
		}
	}

	this.ServeJSON()
}