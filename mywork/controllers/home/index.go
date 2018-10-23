package controllers

import "github.com/astaxie/beego"

type IndexController struct {

	beego.Controller  //继承beego的控制器
}

//使用注解路由

// @router / [get]
func (this *IndexController) Index() {

	//设置模板路径
	this.TplName = "home/index.html"
}


