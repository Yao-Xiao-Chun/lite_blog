package common

import (
	"github.com/astaxie/beego/logs"
	"lite_blog/syserror"
)

/**
错误处理代码
*/
type ErrorController struct {
	BaseController
}

/**
404
是否是ajax 请求 json{"code:","message:"}
*/
func (c *ErrorController) Error404() {

	c.TplName = "error/404.html"
	//是否是ajax
	if c.IsAjax() {

		//同时设置 请求头
		c.Ctx.Output.Status = 200

		c.Data["json"] = map[string]interface{}{
			"code":    1002,
			"message": "错误请求，ajax",
		}
		c.ServeJSON()

	}

}

/**
500 错误
*/
func (c *ErrorController) Error500() {

	c.TplName = "error/500.html"

	err, ok := c.Data["error"].(error)

	if !ok {

		err = syserror.New("未知错误", nil)
	}

	rsErr, ok := err.(syserror.Error)

	if !ok {

		rsErr = syserror.New(err.Error(), nil) //获取的是go自带的
	}

	if rsErr.ReasonError() != nil {

		logs.Info(rsErr.Error(), rsErr.ReasonError()) //判断是否为空，不为空，则输出错误信息
	}

	if c.IsAjax() {

		c.jsonErr(rsErr)

	} else {

		c.Data["content"] = rsErr.Error()
	}
}

func (c *ErrorController) jsonErr(rsErr syserror.Error) {

	//同时设置 请求头
	c.Ctx.Output.Status = 200

	c.Data["json"] = map[string]interface{}{
		"code":    rsErr.Code(),  //key z值
		"message": rsErr.Error(), //下标信息
	}
	c.ServeJSON()
}
