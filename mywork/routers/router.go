package routers

import (
	"mywork/controllers"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {

	//注册错误处理
	beego.ErrorController(&controllers.ErrorController{})
	
	beego.Include(&controllers.IndexController{})//前台首页

	beego.Include(&controllers.AdminIndexController{})//后台首页

	beego.Include(&controllers.AdminUserController{})//后台登录用户

	beego.Include(&controllers.AdminTimeController{})//时间线控制器

	beego.Include(&controllers.AdminArticleController{})//文章控制器

	beego.Include(&controllers.AdminTagController{})//Tag控制器

	beego.Include(&controllers.AdminMenuController{})//menu控制器

	beego.Include(&controllers.AdminReviewController{})//留言控制器

	beego.Handler("/captcha/*.png", captcha.Server(120, 38)) //设置验证码


}


