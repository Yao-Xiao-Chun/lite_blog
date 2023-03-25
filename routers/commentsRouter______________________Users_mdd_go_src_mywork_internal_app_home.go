package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           "/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "IndexAbout",
			Router:           "/about",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "GetHomePageArticle",
			Router:           "/article/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "ArticleClick",
			Router:           "/article/click/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "GetArticleInfo",
			Router:           "/article/info/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "CommitArticle",
			Router:           "/blog/pages/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "TypeArticle",
			Router:           "/category/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "IndexDetails",
			Router:           "/details",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "DownFile",
			Router:           "/download/file/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "IndexMessage",
			Router:           "/message",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "SteReview",
			Router:           "/message/review",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "HomePageReview",
			Router:           "/message/review/page/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "IndexTime",
			Router:           "/time",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"] = append(beego.GlobalControllerRouter["lite_blog/internal/app/home:IndexController"],
		beego.ControllerComments{
			Method:           "GetTimePage",
			Router:           "/time/page/?:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
