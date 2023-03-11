package controllers

import "github.com/astaxie/beego"

type IndexController struct {

	beego.Controller  //继承beego的控制器
}

//使用注解路由

// @router / [get]
func (c *IndexController) Index() {

	//设置模板路径
	c.TplName = "home/index.html"
}


