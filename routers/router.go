package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"mywork/internal/app/admin"
	"mywork/internal/app/common"
)

func init() {

	//注册错误处理
	beego.ErrorController(&common.ErrorController{})

	beego.Include(&common.IndexController{}) //前台首页

	beego.Include(&admin.AdminIndexController{}) //后台首页

	beego.Include(&admin.AdminUserController{}) //后台登录用户

	beego.Include(&admin.AdminTimeController{}) //时间线控制器

	beego.Include(&admin.ArticleController{}) //文章控制器

	beego.Include(&admin.AdminTagController{}) //Tag控制器

	beego.Include(&admin.AdminMenuController{}) //menu控制器

	beego.Include(&admin.ReviewController{}) //留言控制器

	beego.Include(&admin.AdminLinkController{}) //友情链接

	beego.Include(&common.UploadController{}) //大文件上传demo

	beego.Include(&common.CronTabController{}) //定时任务控制器

	beego.Handler("/captcha/*.png", captcha.Server(120, 38)) //设置验证码

	beego.Router("/ws/index", &admin.PushSocketController{}) //websocket控制器

	beego.Router("/admin/webpush/index", &admin.AdminTaskController{}, "get:TaskIndex") //推送任务列表首页

	beego.Router("admin/fiction/index", &admin.FileController{}, "get:FileIndex") //小说列表页

	beego.Router("admin/fiction/delete/?:key", &admin.FileController{}, "get:SetFictionStatus") //禁止小说下载

	beego.Router("admin/fiction/index/page/?:key", &admin.FileController{}, "get:FilePage") //小说列表页

	beego.Router("/fiction", &common.IndexController{}, "get:GetHomeFiction") //前台小说列表页

	beego.Router("/fiction/page/?:key", &common.IndexController{}, "get:HomeFictionPage") //前台小说列表页

	beego.Router("/fiction/download/?:key", &common.IndexController{}, "get:HomeFictionDownload") //前台小说下载

	//小说日志路由模块 admin
	beego.Router("admin/fiction/log", &admin.FileController{}, "get:FictionLog") //小说日志访问路由

	beego.Router("/admin/fiction/log/page/?:key", &admin.FileController{}, "get:FictionLogPage") //小说日志列表

	beego.Router("/admin/fiction/banned/?:key", &admin.FileController{}, "get:FictionBanned") //ip黑名单

}
